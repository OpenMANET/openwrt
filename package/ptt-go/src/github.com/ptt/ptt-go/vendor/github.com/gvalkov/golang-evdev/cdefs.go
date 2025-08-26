package evdev

/*
 #include <linux/input.h>
 static int _EVIOCGNAME(int len) {return EVIOCGNAME(len);}
 static int _EVIOCGPHYS(int len) {return EVIOCGPHYS(len);}
 static int _EVIOCGUNIQ(int len) {return EVIOCGUNIQ(len);}
 static int _EVIOCGPROP(int len) {return EVIOCGPROP(len);}

 static int _EVIOCGKEY(int len) {return EVIOCGKEY(len);}
 static int _EVIOCGLED(int len) {return EVIOCGLED(len);}
 static int _EVIOCGSND(int len) {return EVIOCGSND(len);}
 static int _EVIOCGSW(int len)  {return EVIOCGSW(len);}

 static int _EVIOCGBIT(int ev, int len) {return EVIOCGBIT(ev, len);}
 static int _EVIOCGABS(int abs)    {return EVIOCGABS(abs);}
 static int _EVIOCSABS(int abs)    {return EVIOCSABS(abs);}
*/
import "C"
import "syscall"
import "unsafe"

type _InputEvent C.struct_input_event
type _InputAbsinfo C.struct_input_absinfo
type _InputId C.struct_input_id
type _InputKeymapEntry C.struct_input_keymap_entry

const (
	sizeofInputAbsinfo     = C.sizeof_struct_input_absinfo
	sizeofInputId          = C.sizeof_struct_input_id
	sizeofInputKeymapEntry = C.sizeof_struct_input_keymap_entry
)

const MAX_NAME_SIZE = 256

const (
	EVIOCGID      = C.EVIOCGID      // get device ID
	EVIOCGVERSION = C.EVIOCGVERSION // get driver version
	EVIOCGREP     = C.EVIOCGREP     // get repeat settings
	EVIOCSREP     = C.EVIOCSREP     // set repeat settings

	EVIOCGKEYCODE    = C.EVIOCGKEYCODE    // get keycode
	EVIOCGKEYCODE_V2 = C.EVIOCGKEYCODE_V2 // get keycode

	EVIOCSKEYCODE    = C.EVIOCSKEYCODE    // set keycode
	EVIOCSKEYCODE_V2 = C.EVIOCSKEYCODE_V2 // set keycode

	EVIOCSFF      = C.EVIOCSFF      // send a force effect to a force feedback device
	EVIOCRMFF     = C.EVIOCRMFF     // erase a force effect
	EVIOCGEFFECTS = C.EVIOCGEFFECTS // report number of effects playable at the same time

	EVIOCGRAB     = C.EVIOCGRAB     // grab/release device
	EVIOCSCLOCKID = C.EVIOCSCLOCKID // set clockid to be used for timestamps
)

var EVIOCGNAME = C._EVIOCGNAME(MAX_NAME_SIZE) // get device name
var EVIOCGPHYS = C._EVIOCGPHYS(MAX_NAME_SIZE) // get physical location
var EVIOCGUNIQ = C._EVIOCGUNIQ(MAX_NAME_SIZE) // get unique identifier
var EVIOCGPROP = C._EVIOCGPROP(MAX_NAME_SIZE) // get device properties

var EVIOCGKEY = C._EVIOCGKEY(MAX_NAME_SIZE) // get global key state
var EVIOCGLED = C._EVIOCGLED(MAX_NAME_SIZE) // get all LEDs
var EVIOCGSND = C._EVIOCGSND(MAX_NAME_SIZE) // get all sounds status
var EVIOCGSW = C._EVIOCGSW(MAX_NAME_SIZE)   // get all switch states

func EVIOCGBIT(ev, l int) int { return int(C._EVIOCGBIT(C.int(ev), C.int(l))) } // get event bits
func EVIOCGABS(abs int) int   { return int(C._EVIOCGABS(C.int(abs))) }          // get abs bits
func EVIOCSABS(abs int) int   { return int(C._EVIOCSABS(C.int(abs))) }          // set abs bits

func ioctl(fd uintptr, name uintptr, data unsafe.Pointer) syscall.Errno {
	_, _, err := syscall.RawSyscall(syscall.SYS_IOCTL, fd, name, uintptr(data))
	return err
}
