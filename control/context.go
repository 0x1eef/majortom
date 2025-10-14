package control

/*
	#cgo LDFLAGS: -lhbsdcontrol
	#include "control.h"
*/
import "C"

import (
	"unsafe"
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
	cFlags, cNamespace := C.hbsdctrl_flag_t(ctx.flags), C.CString(ctx.namespace)
	defer free(cNamespace)
	if ctx.ptr = C.hbsdctrl_ctx_new(cFlags, cNamespace); ctx.ptr == nil {
		return Context{}, ErrNullPtr
	} else {
		return ctx, nil
	}
}

func (ctx *Context) FeatureNames() ([]string, error) {
	if ctx.ptr == nil {
		return []string{}, ErrUseAfterFree
	} else {
		cNames := C.hbsdctrl_ctx_all_feature_names(ctx.ptr)
		if cNames == nil {
			return []string{}, ErrNullPtr
		} else {
			defer C.hbsdctrl_ctx_free_feature_names(cNames)
			return gostrings(cNames), nil
		}
	}
}

func (ctx *Context) Status(feature, path string) (string, error) {
	if ctx.ptr == nil {
		return "", ErrUseAfterFree
	} else {
		cStatus, cFeature, cPath := C.CString(""), C.CString(feature), C.CString(path)
		defer free(cStatus, cFeature, cPath)
		cPtr := (**C.char)(unsafe.Pointer(&cStatus))
		result := C.feature_status(ctx.ptr, cFeature, cPath, cPtr)
		if result == 0 {
			defer free(cStatus)
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
	} else {
		cFeature, cPath := C.CString(feature), C.CString(path)
		defer free(cFeature, cPath)
		result := C.enable_feature(ctx.ptr, cFeature, cPath)
		return handle(result)
	}
}

func (ctx *Context) Disable(feature, path string) error {
	if ctx.ptr == nil {
		return ErrUseAfterFree
	} else {
		cFeature, cPath := C.CString(feature), C.CString(path)
		defer free(cFeature, cPath)
		result := C.disable_feature(ctx.ptr, cFeature, cPath)
		return handle(result)
	}
}

func (ctx *Context) Sysdef(feature, path string) error {
	if ctx.ptr == nil {
		return ErrUseAfterFree
	} else {
		cFeature, cPath := C.CString(feature), C.CString(path)
		defer free(cFeature, cPath)
		result := C.sysdef_feature(ctx.ptr, cFeature, cPath)
		return handle(result)
	}
}

func (ctx *Context) Namespace() string {
	return ctx.namespace
}

func (ctx *Context) Free() {
	ptr := (**C.struct__hbsdctrl_ctx)(unsafe.Pointer(&ctx.ptr))
	C.hbsdctrl_ctx_free(ptr)
	ctx.ptr = nil
}
