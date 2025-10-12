package main

import (
	"fmt"
	"github.com/0x1eef/majortom/control"
)

func main() {
	ctx := control.New(control.Namespace("system"))
	if features, err := ctx.FeatureNames(); err != nil {
		panic(err)
	} else {
		for _, name := range features {
			fmt.Printf("feature: %s\n", name)
		}
		if err := ctx.Enable("mprotect", "/usr/bin/mdo"); err != nil {
			panic(err)
		}
		if err := ctx.Disable("mprotect", "/usr/bin/mdo"); err != nil {
			panic(err)
		}
		if err := ctx.Sysdef("mprotect", "/usr/bin/mdo"); err != nil {
			panic(err)
		}
		if status, err := ctx.Status("mprotect", "/usr/bin/mdo"); err != nil {
			panic(err)
		} else {
			fmt.Printf("The mprotect feature has the status: %s\n", status)
		}
	}
}
