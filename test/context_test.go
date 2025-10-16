package test

import (
	"os"
	"testing"

	"github.com/0x1eef/majortom/control"
)

func TestDefaultNamespace(t *testing.T) {
	ctx, err := control.NewContext()
	AssertNil(t, err)
	defer ctx.Free()
	AssertEqual(t, "system", ctx.Namespace())
}

func TestUserNamespace(t *testing.T) {
	ctx, err := control.NewContext(control.Namespace("user"))
	AssertNil(t, err)
	defer ctx.Free()
	AssertEqual(t, "user", ctx.Namespace())
}

func TestFeatureNames(t *testing.T) {
	ctx, err := control.NewContext(control.Namespace("user"))
	AssertNil(t, err)
	defer ctx.Free()
	names, err := ctx.FeatureNames()
	AssertNil(t, err)
	AssertEqual(t, false, len(names) == 0)
}

func TestStatus(t *testing.T) {
	file, err := os.CreateTemp("", "test")
	AssertNil(t, err)
	defer file.Close()
	defer os.Remove(file.Name())

	ctx, err := control.NewContext(control.Namespace("user"))
	AssertNil(t, err)
	defer ctx.Free()

	status, err := ctx.Status("mprotect", file.Name())
	AssertNil(t, err)
	AssertEqual(t, "sysdef", status)
}

func TestUseAfterFree(t *testing.T) {
	ctx, err := control.NewContext(control.Namespace("user"))
	AssertNil(t, err)
	ctx.Free()
	_, err = ctx.FeatureNames()
	AssertEqual(t, err, control.ErrUseAfterFree)
}
