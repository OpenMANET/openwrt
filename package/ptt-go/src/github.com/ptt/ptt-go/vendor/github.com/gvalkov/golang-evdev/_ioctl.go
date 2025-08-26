package evdev

// Bits and pieces from asm-generic/ioctl.h
import (
	"syscall"
	"unsafe"
)

const (
	_IOC_NONE  = 0x0
	_IOC_WRITE = 0x1
	_IOC_READ  = 0x2

	_IOC_NRBITS   = 8
	_IOC_TYPEBITS = 8
	_IOC_SIZEBITS = 14
	_IOC_DIRBITS  = 2
	_IOC_NRSHIFT  = 0

	_IOC_TYPESHIFT = _IOC_NRSHIFT + _IOC_NRBITS
	_IOC_SIZESHIFT = _IOC_TYPESHIFT + _IOC_TYPEBITS
	_IOC_DIRSHIFT  = _IOC_SIZESHIFT + _IOC_SIZEBITS
)

func _IOC(dir int, t int, nr int, size int) int {
	return (dir << _IOC_DIRSHIFT) | (t << _IOC_TYPESHIFT) |
		(nr << _IOC_NRSHIFT) | (size << _IOC_SIZESHIFT)
}

func _IOR(t int, nr int, size int) int {
	return _IOC(_IOC_READ, t, nr, size)
}

func ioctl(fd uintptr, name int, data unsafe.Pointer) syscall.Errno {
	_, _, err := syscall.RawSyscall(syscall.SYS_IOCTL, fd, uintptr(name), uintptr(data))
	return err
}

// input.h ioctls
var _EVIOCGID = _IOR('E', 0x02, SizeofInputId) // 8 <- sizeof(struct input_id)
var _EVIOCGNAME = _IOC(_IOC_READ, 'E', 0x06, _MAX_NAME_SIZE)
var _EVIOCGPHYS = _IOC(_IOC_READ, 'E', 0x07, _MAX_NAME_SIZE)

func _EVIOCGABS(ev int) int {
	return _IOR('E', 0x40+ev, 24) // 24 <= sizeof(struct input_absinfo)
}

func _EVIOCGBIT(ev, len int) int {
	return _IOC(_IOC_READ, 'E', 0x20+ev, len)
}
