package control

import (
	"testing"
)

func TestDefaultNamespace(t *testing.T) {
	ctx := New()
	defer ctx.Free()
	if ctx.namespace != "system" {
		t.Fatalf("The default namespace should be 'system' but got '%s'", ctx.namespace)
	}
}

func TestUserNamespace(t *testing.T) {
	ctx := New(Namespace("user"))
	defer ctx.Free()
	if ctx.namespace != "user" {
		t.Fatalf("The namespace should have been 'user' but got '%s'", ctx.namespace)
	}
}

func TestFeatureNames(t *testing.T) {
	ctx := New(Namespace("user"))
	defer ctx.Free()
	if names, err := ctx.FeatureNames(); err != nil {
		t.Fatalf("The FeatureNames method has an error: %s", err)
	} else {
		if len(names) == 0 {
			t.Fatalf("The FeatureNames method has zero features")
		}
	}
}
