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
	flags, namespace := C.hbsdctrl_flag_t(ctx.flags), C.CString(ctx.namespace)
	defer free(namespace)
	if ctx.ptr = C.hbsdctrl_ctx_new(flags, namespace); ctx.ptr == nil {
		return Context{}, ErrNullPtr
	} else {
		return ctx, nil
	}
}

func (ctx *Context) FeatureNames() ([]string, error) {
	if ctx.ptr == nil {
		return []string{}, ErrUseAfterFree
	}
	names := C.hbsdctrl_ctx_all_feature_names(ctx.ptr)
	if names == nil {
		return []string{}, ErrNullPtr
	} else {
		defer C.hbsdctrl_ctx_free_feature_names(names)
		return gostrings(names), nil
	}
}

func (ctx *Context) Status(feature, path string) (string, error) {
	if ctx.ptr == nil {
		return "", ErrUseAfterFree
	}
	var cstatus, cfeature, cpath *C.char
	var ptr **C.char
	cfeature, cpath = C.CString(feature), C.CString(path)
	ptr = (**C.char)(unsafe.Pointer(&cstatus))
	defer free(cfeature, cpath)
	if result := C.feature_status(ctx.ptr, cfeature, cpath, ptr); result != 0 {
		return "", handle(result)
	} else {
		return C.GoString(cstatus), nil
	}
}

func (ctx *Context) IsEnabled(feature, path string) (bool, error) {
	if ctx.ptr == nil {
		return false, ErrUseAfterFree
	}
	if status, err := ctx.Status(feature, path); err != nil {
		return false, err
	} else {
		return status == "enabled", err
	}
}

func (ctx *Context) IsDisabled(feature, path string) (bool, error) {
	if ctx.ptr == nil {
		return false, ErrUseAfterFree
	}
	if status, err := ctx.Status(feature, path); err != nil {
		return false, err
	} else {
		return status == "disabled", err
	}
}

func (ctx *Context) IsSysdef(feature, path string) (bool, error) {
	if ctx.ptr == nil {
		return false, ErrUseAfterFree
	}
	if status, err := ctx.Status(feature, path); err != nil {
		return false, err
	} else {
		return status == "sysdef", err
	}
}

func (ctx *Context) Enable(feature, path string) error {
	if ctx.ptr == nil {
		return ErrUseAfterFree
	}
	cfeature, cpath := C.CString(feature), C.CString(path)
	defer free(cfeature, cpath)
	result := C.enable_feature(ctx.ptr, cfeature, cpath)
	return handle(result)
}

func (ctx *Context) Disable(feature, path string) error {
	if ctx.ptr == nil {
		return ErrUseAfterFree
	}
	cfeature, cpath := C.CString(feature), C.CString(path)
	defer free(cfeature, cpath)
	result := C.disable_feature(ctx.ptr, cfeature, cpath)
	return handle(result)
}

func (ctx *Context) Sysdef(feature, path string) error {
	if ctx.ptr == nil {
		return ErrUseAfterFree
	}
	cfeature, cpath := C.CString(feature), C.CString(path)
	defer free(cfeature, cpath)
	result := C.sysdef_feature(ctx.ptr, cfeature, cpath)
	return handle(result)
}

func (ctx *Context) Free() {
	ptr := (**C.struct__hbsdctrl_ctx)(unsafe.Pointer(&ctx.ptr))
	C.hbsdctrl_ctx_free(ptr)
	ctx.ptr = nil
}

func (ctx *Context) Namespace() string {
	return ctx.namespace
}
