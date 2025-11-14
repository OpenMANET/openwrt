package evdev

import (
	"fmt"
	"syscall"
	"unsafe"
)

type InputEvent struct {
	Time  syscall.Timeval // time in seconds since epoch at which event occurred
	Type  uint16          // event type - one of ecodes.EV_*
	Code  uint16          // event code related to the event type
	Value int32           // event value related to the event type
}

// Get a useful description for an input event. Example:
//   event at 1347905437.435795, code 01, type 02, val 02
func (ev *InputEvent) String() string {
	return fmt.Sprintf("event at %d.%d, code %02d, type %02d, val %02d",
		ev.Time.Sec, ev.Time.Usec, ev.Code, ev.Type, ev.Value)
}

var eventsize = int(unsafe.Sizeof(InputEvent{}))

type KeyEventState uint8

const (
	KeyUp   KeyEventState = 0x0
	KeyDown KeyEventState = 0x1
	KeyHold KeyEventState = 0x2
)

// KeyEvents are used to describe state changes of keyboards, buttons,
// or other key-like devices.
type KeyEvent struct {
	Event    *InputEvent
	Scancode uint16
	Keycode  uint16
	State    KeyEventState
}

func (kev *KeyEvent) New(ev *InputEvent) {
	kev.Event = ev
	kev.Keycode = 0 // :todo
	kev.Scancode = ev.Code

	switch ev.Value {
	case 0:
		kev.State = KeyUp
	case 2:
		kev.State = KeyHold
	case 1:
		kev.State = KeyDown
	}
}

func NewKeyEvent(ev *InputEvent) *KeyEvent {
	kev := &KeyEvent{}
	kev.New(ev)
	return kev
}

func (ev *KeyEvent) String() string {
	state := "unknown"

	switch ev.State {
	case KeyUp:
		state = "up"
	case KeyHold:
		state = "hold"
	case KeyDown:
		state = "down"
	}

	return fmt.Sprintf("key event at %d.%d, %d (%d), (%s)",
		ev.Event.Time.Sec, ev.Event.Time.Usec,
		ev.Scancode, ev.Event.Code, state)
}

// RelEvents are used to describe relative axis value changes,
// e.g. moving the mouse 5 units to the left.
type RelEvent struct {
	Event *InputEvent
}

func (rev *RelEvent) New(ev *InputEvent) {
	rev.Event = ev
}

func NewRelEvent(ev *InputEvent) *RelEvent {
	rev := &RelEvent{}
	rev.New(ev)
	return rev
}

func (ev *RelEvent) String() string {
	return fmt.Sprintf("relative axis event at %d.%d, %s",
		ev.Event.Time.Sec, ev.Event.Time.Usec,
		REL[int(ev.Event.Code)])
}

// TODO: Make this work

var EventFactory map[uint16]interface{} = make(map[uint16]interface{})

func init() {
	EventFactory[uint16(EV_KEY)] = NewKeyEvent
	EventFactory[uint16(EV_REL)] = NewRelEvent
}
