package control

/*
	#include <stdlib.h>
*/
import "C"

import (
	"errors"
	"syscall"
	"unsafe"
)

func gostrings(cstr **C.char) []string {
	var strings []string
	addr := uintptr(unsafe.Pointer(cstr))
	offset := unsafe.Sizeof((*C.char)(nil))
	for {
		ptr := (**C.char)(unsafe.Pointer(addr))
		if *ptr == nil {
			break
		}
		strings = append(strings, C.GoString(*ptr))
		addr += offset
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

func free(objects ...*C.char) {
	for _, object := range objects {
		if object != nil {
			C.free(unsafe.Pointer(object))
		}
	}
}
