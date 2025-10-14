package control

import (
	"errors"
	"testing"
)

func TestDefaultNamespace(t *testing.T) {
	if ctx, err := NewContext(); err != nil {
		t.Fatalf("NewContext failure: %v", err)
	} else {
		defer ctx.Free()
		if ctx.namespace != "system" {
			t.Fatalf("The default namespace should be 'system' but got '%s'", ctx.namespace)
		}
	}
}

func TestUserNamespace(t *testing.T) {
	if ctx, err := NewContext(Namespace("user")); err != nil {
		t.Fatalf("NewContext failure: %v", err)
	} else {
		defer ctx.Free()
		if ctx.namespace != "user" {
			t.Fatalf("The namespace should have been 'user' but got '%s'", ctx.namespace)
		}
	}
}

func TestFeatureNames(t *testing.T) {
	if ctx, err := NewContext(Namespace("user")); err != nil {
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
	if ctx, err := NewContext(Namespace("user")); err != nil {
		t.Fatalf("NewContext failure: %v", err)
	} else {
		ctx.Free()
		if _, err := ctx.FeatureNames(); err != nil {
			if !errors.Is(err, ErrUseAfterFree) {
				t.Fatalf("The FeatureNames method returned an unexpected error: %s", err)
			}
		} else {
			t.Fatalf("The FeatureNames method should have returned an error, but did not")
		}
	}
}
