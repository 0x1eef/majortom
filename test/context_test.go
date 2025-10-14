package test

import (
	"errors"
	"testing"

	"github.com/0x1eef/majortom/control"
)

func TestDefaultNamespace(t *testing.T) {
	if ctx, err := control.NewContext(); err != nil {
		t.Fatalf("NewContext failure: %v", err)
	} else {
		defer ctx.Free()
		if ctx.Namespace() != "system" {
			t.Fatalf("The default namespace should be 'system' but got '%s'", ctx.Namespace())
		}
	}
}

func TestUserNamespace(t *testing.T) {
	if ctx, err := control.NewContext(control.Namespace("user")); err != nil {
		t.Fatalf("NewContext failure: %v", err)
	} else {
		defer ctx.Free()
		if ctx.Namespace() != "user" {
			t.Fatalf("The namespace should have been 'user' but got '%s'", ctx.Namespace())
		}
	}
}

func TestFeatureNames(t *testing.T) {
	if ctx, err := control.NewContext(control.Namespace("user")); err != nil {
		t.Fatalf("NewContext failure: %v", err)
	} else {
		defer ctx.Free()
		if names, err := ctx.FeatureNames(); err != nil {
			t.Fatalf("The FeatureNames method has an error: %s", err)
		} else {
			if len(names) == 0 {
				t.Fatalf("The FeatureNames method has zero features")
			}
		}
	}
}

func TestUseAfterFree(t *testing.T) {
	if ctx, err := control.NewContext(control.Namespace("user")); err != nil {
		t.Fatalf("NewContext failure: %v", err)
	} else {
		ctx.Free()
		if _, err := ctx.FeatureNames(); err != nil {
			if !errors.Is(err, control.ErrUseAfterFree) {
				t.Fatalf("The FeatureNames method returned an unexpected error: %s", err)
			}
		} else {
			t.Fatalf("The FeatureNames method should have returned an error, but did not")
		}
	}
}
