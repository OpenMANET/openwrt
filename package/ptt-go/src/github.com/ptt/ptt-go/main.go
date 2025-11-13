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

	"github.com/digineo/go-uci"
	"github.com/gordonklaus/portaudio"
	"github.com/gvalkov/golang-evdev"
	"github.com/hraban/opus"
	"golang.org/x/net/ipv4"
)

/********* defaults *********/
const (
	sampleRate        = 48000
	channels          = 1
	frameSize         = 960
	targetBitrate     = 12000
	encoderComplexity = 3
	packetLossPerc    = 10
	defaultKey        = "any"
	defaultIface      = "br-ahwlan" // ← use bridge by default; override in UCI if needed
	defaultG          = "224.0.0.1"
	defaultPort       = 5007
)

var (
	// codec/network
	encoder *opus.Encoder
	decoder *opus.Decoder
	udpSendConn *net.UDPConn
	udpRecvConn *net.UDPConn
	localIP              string
	playbackBuffer       = make(chan []float32, 2)
	beepBufferStart      = make([]float32, frameSize)
	beepBufferStop       = make([]float32, frameSize)
	broadcastStream      *portaudio.Stream
	broadcasting         bool
	recordMutex          sync.Mutex

	// config from UCI (with fallbacks)
	ifaceName  = defaultIface
	mcastAddr  = defaultG
	mcastPort  = defaultPort
	pttKey     = defaultKey
)

/********* helpers: UCI *********/
func loadConfig() {
	tree := uci.NewTree("/etc/config")
	if v, ok := tree.Get("pttradio", "main", "iface"); ok && len(v) > 0 && v[0] != "" {
		ifaceName = v[0]
	}
	if v, ok := tree.Get("pttradio", "main", "mcast_addr"); ok && len(v) > 0 && v[0] != "" {
		mcastAddr = v[0]
	}
	if v, ok := tree.Get("pttradio", "main", "mcast_port"); ok && len(v) > 0 {
		if p, err := strconv.Atoi(v[0]); err == nil {
			mcastPort = p
		}
	}
	if v, ok := tree.Get("pttradio", "main", "ptt_key"); ok && len(v) > 0 && v[0] != "" {
		pttKey = v[0]
	}
}

/********* helpers: net *********/
func getIfaceIPv4(name string) (string, *net.Interface, error) {
	ifi, err := net.InterfaceByName(name)
	if err != nil {
		return "", nil, err
	}
	addrs, err := ifi.Addrs()
	if err != nil {
		return "", nil, err
	}
	for _, a := range addrs {
		if ipn, ok := a.(*net.IPNet); ok && ipn.IP.To4() != nil {
			return ipn.IP.String(), ifi, nil
		}
	}
	return "", ifi, fmt.Errorf("no IPv4 on iface %s", name)
}

func joinMulticastGroup(iface *net.Interface, conn *net.UDPConn, group net.IP) error {
	p := ipv4.NewPacketConn(conn)
	return p.JoinGroup(iface, &net.UDPAddr{IP: group})
}

/********* app *********/
func main() {
	loadConfig()

	var err error
	encoder, err = opus.NewEncoder(sampleRate, channels, opus.AppVoIP)
	check(err)
	check(encoder.SetBitrate(targetBitrate))
	check(encoder.SetComplexity(encoderComplexity))
	check(encoder.SetInBandFEC(true))
	check(encoder.SetPacketLossPerc(packetLossPerc))
	check(encoder.SetDTX(false))

	decoder, err = opus.NewDecoder(sampleRate, channels)
	check(err)

	check(portaudio.Initialize())
	defer portaudio.Terminate()

	// playback stream
	device := getDeviceByIndex(1)
	params := portaudio.StreamParameters{
		Output: portaudio.StreamDeviceParameters{
			Device:   device,
			Channels: channels,
		},
		SampleRate:      float64(sampleRate),
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
	check(playbackStream.Start())
	defer playbackStream.Stop()
	defer playbackStream.Close()

	// mic stream (opened, not started)
	broadcastStream, err = portaudio.OpenDefaultStream(channels, 0, float64(sampleRate), frameSize, func(in []float32) {
		pcm := make([]int16, len(in))
		for i, v := range in {
			pcm[i] = int16(v * 32767)
		}
		buf := make([]byte, 4000)
		if n, err := encoder.Encode(pcm, buf); err == nil {
			_, _ = udpSendConn.Write(buf[:n])
		}
	})
	check(err)
	defer broadcastStream.Close()

	// beeps
	for i := range beepBufferStart {
		beepBufferStart[i] = float32(math.Sin(2*math.Pi*1000*float64(i)/float64(sampleRate))) * 0.2
		beepBufferStop[i] = float32(math.Sin(2*math.Pi*600*float64(i)/float64(sampleRate))) * 0.2
	}

	// networking: bind send to iface IP; listen on :port and join group on iface
	ifIP, ifi, err := getIfaceIPv4(ifaceName)
	check(err)
	localIP = ifIP

	// sender bound to iface IP so traffic egresses that iface
	dst := &net.UDPAddr{IP: net.ParseIP(mcastAddr), Port: mcastPort}
	src := &net.UDPAddr{IP: net.ParseIP(ifIP), Port: 0}
	udpSendConn, err = net.DialUDP("udp4", src, dst)
	check(err)

	// receiver on all, then join group on iface
	udpRecvConn, err = net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4zero, Port: mcastPort})
	check(err)
	_ = udpRecvConn.SetReadBuffer(65535)
	check(joinMulticastGroup(ifi, udpRecvConn, net.ParseIP(mcastAddr)))

	go receiveLoop()

	// PTT input (kept as-is for now)
	ptt := findPTTDevice()
	fmt.Println("🎙️ Listening for PTT on:", ptt.Name)
	go monitorPTT(ptt)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("Exiting.")
}

func receiveLoop() {
	buf := make([]byte, 1500)
	for {
		n, src, err := udpRecvConn.ReadFromUDP(buf)
		if err != nil {
			log.Println("Recv error:", err)
			continue
		}
		// (optional) drop self
		if src.IP.IsLoopback() || src.IP.String() == localIP {
			// pass if testing loopback; otherwise skip
		}

		frame := make([]byte, n)
		copy(frame, buf[:n])

		pcm := make([]int16, frameSize)
		n, err := decoder.Decode(frame, pcm)
		if err != nil {
			continue
		}
		if isBroadcasting() {
			continue
		}
		out := make([]float32, n)
		for i := 0; i < n; i++ {
			out[i] = float32(pcm[i]) / 32768
		}
		select {
		case playbackBuffer <- out:
		default:
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
		if ev.Type != evdev.EV_KEY {
			continue
		}
		match := false
		if pttKey == "any" {
			match = true
		} else if kc, err := strconv.Atoi(pttKey); err == nil && ev.Code == uint16(kc) {
			match = true
		}
		if !match {
			continue
		}

		switch ev.Value {
		case 1:
			beginTransmission()
		case 0:
			endTransmission()
		}
	}
}

func drainPlaybackBuffer() {
	for {
		select {
		case <-playbackBuffer:
		default:
			return
		}
	}
}

func isBroadcasting() bool {
	recordMutex.Lock()
	defer recordMutex.Unlock()
	return broadcasting
}

func beginTransmission() {
	recordMutex.Lock()
	if broadcasting {
		recordMutex.Unlock()
		return
	}
	broadcasting = true
	recordMutex.Unlock()

	drainPlaybackBuffer()
	playbackBuffer <- beepBufferStart
	time.Sleep(200 * time.Millisecond)
	if err := broadcastStream.Start(); err != nil {
		log.Printf("start mic: %v", err)
		recordMutex.Lock()
		broadcasting = false
		recordMutex.Unlock()
	}
}

func endTransmission() {
	recordMutex.Lock()
	if !broadcasting {
		recordMutex.Unlock()
		return
	}
	recordMutex.Unlock()

	if err := broadcastStream.Stop(); err != nil {
		log.Printf("stop mic: %v", err)
	}
	drainPlaybackBuffer()
	playbackBuffer <- beepBufferStop

	recordMutex.Lock()
	broadcasting = false
	recordMutex.Unlock()
}

func getDeviceByIndex(index int) *portaudio.DeviceInfo {
	devs, err := portaudio.Devices()
	check(err)
	if len(devs) <= index {
		log.Fatalf("Device index %d not found; only %d devices available", index, len(devs))
	}
	return devs[index]
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

func check(err error) { if err != nil { log.Fatal(err) } }
