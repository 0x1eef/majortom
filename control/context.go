package control

/*
	#cgo LDFLAGS: -lhbsdcontrol
	#include "control.h"
*/
import "C"

import (
	"errors"
	"unsafe"
)

var (
	ErrUseAfterFree = errors.New("context has been freed")
	ErrNullPtr			= errors.New("null pointer")
)

type Context struct {
	namespace string
	flags     uint64
	ptr       *C.struct__hbsdctrl_ctx
}

func NewContext(opts ...Option) (Context, error) {
	ctx := Context{namespace: "system", flags: 0}
	for _, set := range opts {
		set(&ctx)
	}
	flags, ns := C.hbsdctrl_flag_t(ctx.flags), C.CString(ctx.namespace)
	if ctx.ptr = C.hbsdctrl_ctx_new(flags, ns); ctx.ptr == nil {
		return Context{}, ErrNullPtr
	} else {
		return ctx, nil
	}
}

func (ctx *Context) FeatureNames() ([]string, error) {
	if ctx.ptr == nil {
		return []string{}, ErrUseAfterFree
	} else {
		names := []string{}
		cary := C.hbsdctrl_ctx_all_feature_names(ctx.ptr)
		if cary == nil {
			return names, ErrNullPtr
		} else {
			defer C.hbsdctrl_ctx_free_feature_names(cary)
			names = gostrings(cary)
			return names, nil
		}
	}
}

func (ctx *Context) Status(feature, path string) (string, error) {
	if ctx.ptr == nil {
		return "", ErrUseAfterFree
	} else {
		cStatus, cFeature, cPath := C.CString(""), C.CString(feature), C.CString(path)
		cPtr := (**C.char)(unsafe.Pointer(&cStatus))
		result := C.feature_status(ctx.ptr, cFeature, cPath, cPtr)
		if result == 0 {
			return C.GoString(cStatus), nil
		} else {
			return "", handle(result)
		}
	}
}

func (ctx *Context) IsEnabled(feature, path string) (bool, error) {
	if ctx.ptr == nil {
		return false, ErrUseAfterFree
	} else {
		if status, err := ctx.Status(feature, path); err != nil {
			return false, err
		} else {
			return status == "enabled", err
		}
	}
}

func (ctx *Context) IsDisabled(feature, path string) (bool, error) {
	if ctx.ptr == nil {
		return false, ErrUseAfterFree
	} else {
		if status, err := ctx.Status(feature, path); err != nil {
			return false, err
		} else {
			return status == "disabled", err
		}
	}
}

func (ctx *Context) IsSysdef(feature, path string) (bool, error) {
	if ctx.ptr == nil {
		return false, ErrUseAfterFree
	} else {
		if status, err := ctx.Status(feature, path); err != nil {
			return false, err
		} else {
			return status == "sysdef", err
		}
	}
}

func (ctx *Context) Enable(feature, path string) error {
	if ctx.ptr == nil {
		return ErrUseAfterFree
	}
	result := C.enable_feature(ctx.ptr, C.CString(feature), C.CString(path))
	return handle(result)
}

func (ctx *Context) Disable(feature, path string) error {
	if ctx.ptr == nil {
		return ErrUseAfterFree
	}
	result := C.disable_feature(ctx.ptr, C.CString(feature), C.CString(path))
	return handle(result)
}

func (ctx *Context) Sysdef(feature, path string) error {
	if ctx.ptr == nil {
		return ErrUseAfterFree
	}
	result := C.sysdef_feature(ctx.ptr, C.CString(feature), C.CString(path))
	return handle(result)
}

func (ctx *Context) Free() {
	ptr := (**C.struct__hbsdctrl_ctx)(unsafe.Pointer(&ctx.ptr))
	C.hbsdctrl_ctx_free(ptr)
	ctx.ptr = nil
}
