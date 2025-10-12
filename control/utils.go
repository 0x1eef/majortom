package control

import (
	"C"
	"errors"
	"syscall"
	"unsafe"
)

func gostrings(cary **C.char) []string {
	var strings []string
	ptr := uintptr(unsafe.Pointer(cary))
	offset := unsafe.Sizeof((*C.char)(nil))
	for {
		p := (**C.char)(unsafe.Pointer(ptr))
		if *p == nil {
			break
		}
		strings = append(strings, C.GoString(*p))
		ptr += offset
	}
	return strings
}

func handle(result C.int) error {
	if result == 0 {
		return nil
	} else if result == -1 {
		return errors.New("an unknown error happened")
	} else {
		return syscall.Errno(result)
	}
}
