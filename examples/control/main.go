package main

import (
	"fmt"
	"github.com/0x1eef/majortom/control"
)

func main() {
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
