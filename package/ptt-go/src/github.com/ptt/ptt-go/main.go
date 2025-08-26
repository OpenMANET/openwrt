package main

import (
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/gordonklaus/portaudio"
	"github.com/gvalkov/golang-evdev"
	"github.com/hraban/opus"
	"golang.org/x/net/ipv4"
)

const (
	sampleRate = 48000
	channels   = 1
	frameSize  = 960 // 20ms at 48kHz

	// pttKey can now be a string, allowing for "any"
	// Example: "164" for a specific key, or "any" for any keypress
	pttKey = "any" // Changed to string

	mcastAddr = "224.0.0.1"
	mcastPort = 5007

	logPackets = false
)

var (
	broadcasting    bool
	recordMutex     sync.Mutex
	encoder         *opus.Encoder
	decoder         *opus.Decoder
	udpSendConn     *net.UDPConn
	udpRecvConn     *net.UDPConn
	localIP         string
	skipOwnPackets  = false
	playbackBuffer  = make(chan []float32, 2)
	beepBufferStart = make([]float32, frameSize)
	beepBufferStop  = make([]float32, frameSize)

	// Global variable for the microphone stream
	broadcastStream *portaudio.Stream
)

func main() {
	var err error

	encoder, err = opus.NewEncoder(sampleRate, channels, opus.AppVoIP)
	check(err)
	err = encoder.SetBitrate(6000)
	check(err)
	err = encoder.SetComplexity(3)
	check(err)

	decoder, err = opus.NewDecoder(sampleRate, channels)
	check(err)

	check(portaudio.Initialize())
	defer portaudio.Terminate()

	// --- This is the main playback stream for hearing others and beeps ---
	device := getDeviceByIndex(1)
	params := portaudio.StreamParameters{
		Output: portaudio.StreamDeviceParameters{
			Device:   device,
			Channels: channels,
		},
		SampleRate:      sampleRate,
		FramesPerBuffer: frameSize,
	}
	playbackStream, err := portaudio.OpenStream(params, func(_, out []float32) {
		select {
		case data := <-playbackBuffer:
			copy(out, data)
		default:
			for i := range out {
				out[i] = 0
			}
		}
	})
	check(err)
	check(playbackStream.Start()) // Start playback stream immediately
	defer playbackStream.Close()
	defer playbackStream.Stop()

	// --- Open the broadcast stream (microphone) here, but DON'T start it ---
	broadcastStream, err = portaudio.OpenDefaultStream(channels, 0, sampleRate, frameSize, func(in []float32) {
		// This is the microphone callback logic
		pcm := make([]int16, len(in))
		for i, v := range in {
			pcm[i] = int16(v * 32767)
		}
		buf := make([]byte, 4000)
		n, err := encoder.Encode(pcm, buf)
		if err == nil {
			_, err = udpSendConn.Write(buf[:n])
			if logPackets && err == nil {
				log.Printf("📤 Sent %d bytes", n)
			}
		}
	})
	check(err)
	defer broadcastStream.Close() // Ensure it's closed on exit

	// --- Generate beep sounds ---
	for i := range beepBufferStart {
		beepBufferStart[i] = float32(math.Sin(2*math.Pi*1000*float64(i)/sampleRate)) * 0.2
		beepBufferStop[i] = float32(math.Sin(2*math.Pi*600*float64(i)/sampleRate)) * 0.2
	}

	// --- Networking Setup ---
	localIP = getLocalIP()
	addr := net.UDPAddr{IP: net.ParseIP(mcastAddr), Port: mcastPort}
	udpSendConn, err = net.DialUDP("udp", nil, &addr)
	check(err)

	iface, err := net.InterfaceByName("wlan0")
	check(err)

	listenAddr := &net.UDPAddr{IP: net.IPv4zero, Port: mcastPort}
	udpRecvConn, err = net.ListenUDP("udp", listenAddr)
	check(err)
	udpRecvConn.SetReadBuffer(65535)

	err = joinMulticastGroup(iface, udpRecvConn, net.ParseIP(mcastAddr))
	check(err)

	go receiveLoop()

	ptt := findPTTDevice()
	fmt.Println("🎙️ Listening for PTT on:", ptt.Name)
	go monitorPTT(ptt)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("Exiting.")
}

func joinMulticastGroup(iface *net.Interface, conn *net.UDPConn, group net.IP) error {
	p := ipv4.NewPacketConn(conn)
	return p.JoinGroup(iface, &net.UDPAddr{IP: group})
}

func receiveLoop() {
	buf := make([]byte, 1500)
	for {
		n, src, err := udpRecvConn.ReadFromUDP(buf)
		if err != nil {
			log.Println("Recv error:", err)
			continue
		}
		if skipOwnPackets && src.IP.String() == localIP {
			continue
		}
		if logPackets {
			log.Printf("⬇️ Received %d bytes from %s", n, src.IP)
		}

		frame := make([]byte, n)
		copy(frame, buf[:n])

		pcm := make([]int16, frameSize)
		_, err = decoder.Decode(frame, pcm)
		if err != nil {
			continue
		}

		out := make([]float32, frameSize)
		for i := range pcm {
			out[i] = float32(pcm[i]) / 32768
		}

		select {
		case playbackBuffer <- out:
			// Log if the buffer is starting to fill up
			if len(playbackBuffer) > 2 {
				log.Printf("🎧 Playback buffer depth: %d", len(playbackBuffer))
			}
		default:
			// This will happen if the buffer is full and a packet has to be dropped
			log.Println("⚠️ Playback buffer full! Dropping packet.")
		}
	}
}
func monitorPTT(dev *evdev.InputDevice) {
	for {
		ev, err := dev.ReadOne()
		if err != nil {
			continue
		}

		isPTTKeyPress := false
		if ev.Type == evdev.EV_KEY && ev.Value == 1 { // Key press event
			if pttKey == "any" {
				isPTTKeyPress = true
			} else {
				pttKeyCode, parseErr := strconv.Atoi(pttKey)
				if parseErr == nil {
					// Cast pttKeyCode to uint16 for comparison
					if ev.Code == uint16(pttKeyCode) { 
						isPTTKeyPress = true
					}
				}
			}
		}

		if isPTTKeyPress {
			recordMutex.Lock()
			isBroadcasting := broadcasting
			recordMutex.Unlock()

			if !isBroadcasting {
				log.Println("📢 Broadcasting started")

				// Play the start beep locally
				playbackBuffer <- beepBufferStart
				// Give the beep time to play before grabbing the mic
				time.Sleep(200 * time.Millisecond)

				// Start the pre-opened microphone stream
				err := broadcastStream.Start()
				if err != nil {
					log.Printf("❌ Error starting broadcast stream: %v", err)
				}
				recordMutex.Lock()
				broadcasting = true
				recordMutex.Unlock()

			} else {
				log.Println("🛑 Broadcasting stopped")
				// Stop the microphone stream
				err := broadcastStream.Stop()
				if err != nil {
					log.Printf("❌ Error stopping broadcast stream: %v", err)
				}

				// Play the stop beep locally
				playbackBuffer <- beepBufferStop

				recordMutex.Lock()
				broadcasting = false
				recordMutex.Unlock()
			}
		}
	}
}

func getDeviceByIndex(index int) *portaudio.DeviceInfo {
	devices, err := portaudio.Devices()
	check(err)
	if len(devices) <= index {
		log.Fatalf("Device index %d not found; only %d devices available", index, len(devices))
	}
	return devices[index]
}

func findPTTDevice() *evdev.InputDevice {
	devs, err := evdev.ListInputDevices()
	check(err)
	for _, d := range devs {
		if d.Name == "Generic AB13X USB Audio" {
			return d
		}
	}
	log.Fatal("PTT device not found")
	return nil
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}