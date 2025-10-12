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

type Context struct {
	namespace string
	flags     uint64
	ptr       *C.struct__hbsdctrl_ctx
}

func New(opts ...Option) Context {
	ctx := Context{namespace: "system", flags: 0}
	for _, set := range opts {
		set(&ctx)
	}
	flags, ns := C.hbsdctrl_flag_t(ctx.flags), C.CString(ctx.namespace)
	ctx.ptr = C.hbsdctrl_ctx_new(flags, ns)
	return ctx
}

func (ctx *Context) FeatureNames() ([]string, error) {
	names := []string{}
	cary := C.hbsdctrl_ctx_all_feature_names(ctx.ptr)
	if cary == nil {
		return names, errors.New("null pointer")
	}
	defer C.hbsdctrl_ctx_free_feature_names(cary)
	names = gostrings(cary)
	return names, nil
}

func (ctx *Context) Status(feature, path string) (string, error) {
	cStatus, cFeature, cPath := C.CString(""), C.CString(feature), C.CString(path)
	cPtr := (**C.char)(unsafe.Pointer(&cStatus))
	result := C.feature_status(ctx.ptr, cFeature, cPath, cPtr)
	if result == 0 {
		return C.GoString(cStatus), nil
	} else {
		return "", handle(result)
	}
}

func (ctx *Context) IsEnabled(feature, path string) (bool, error) {
	if status, err := ctx.Status(feature, path); err != nil {
		return false, err
	} else {
		return status == "enabled", err
	}
}

func (ctx *Context) IsDisabled(feature, path string) (bool, error) {
	if status, err := ctx.Status(feature, path); err != nil {
		return false, err
	} else {
		return status == "disabled", err
	}
}

func (ctx *Context) IsSysdef(feature, path string) (bool, error) {
	if status, err := ctx.Status(feature, path); err != nil {
		return false, err
	} else {
		return status == "sysdef", err
	}
}

func (ctx *Context) Enable(feature, path string) error {
	result := C.enable_feature(ctx.ptr, C.CString(feature), C.CString(path))
	return handle(result)
}

func (ctx *Context) Disable(feature, path string) error {
	result := C.disable_feature(ctx.ptr, C.CString(feature), C.CString(path))
	return handle(result)
}

func (ctx *Context) Sysdef(feature, path string) error {
	result := C.sysdef_feature(ctx.ptr, C.CString(feature), C.CString(path))
	return handle(result)
}
