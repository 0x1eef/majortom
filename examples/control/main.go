package main

import (
	"fmt"
	"github.com/0x1eef/majortom/control"
	"os"
)

func main() {
	file, err := os.CreateTemp("", "test")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	defer os.Remove(file.Name())

	ctx, err := control.NewContext(control.Namespace("user"))
	if err != nil {
		panic(err)
	}

	features, err := ctx.FeatureNames()
	if err != nil {
		panic(err)
	}

	for _, name := range features {
		fmt.Printf("feature: %s\n", name)
	}

	feature, target := "mprotect", file.Name()
	if err := ctx.Enable(feature, target); err != nil {
		panic(err)
	}
	if err := ctx.Disable(feature, target); err != nil {
		panic(err)
	}
	if err := ctx.Sysdef(feature, target); err != nil {
		panic(err)
	}

	if status, err := ctx.Status(feature, target); err != nil {
		panic(err)
	} else {
		fmt.Printf("The mprotect feature has the status: %s\n", status)
	}
}
