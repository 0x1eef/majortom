## About

The majortom module provides Go bindings for the libhbsdcontrol
library for the [HardenedbSD](https://git.hardenedbsd.org/hardenedbsd/hardenedbsd)
operating system.

## Motivation

### Lyrics

[David Bowie: Space Odyessy](https://www.youtube.com/watch?v=9_M3uw29U1U)

> Ground control to Major Tom <br>
> Ground control to Major Tom <br>
> Take your protein pills and put your helmet on (ten) <br>
> Ground control to Major Tom (nine, eight, seven, six) <br>
> Commencing countdown, engines on (five, four, three, two) <br>
> Check ignition, and may God's love be with you (one, lift off) <br>

[David Bowie: Ashes to Ashes](https://www.youtube.com/watch?v=wBDD37Xv6A0)

> Do you remember a guy that's been <br>
> In such an early song? <br>
> I've heard a rumour from Ground Control <br>
> Oh no, don't say it's true! <br>
> Ashes to ashes and funk to funky
> We know Major Tom's a junkie
> Strung out in heaven's high
> Hitting an all-time low

## Examples

### control

The control package can enable or disable security features
that are managed by the [HardenedBSD](https://hardenedbsd.org)
kernel on a per-file basis. Unlike other packages this one
happens to not be pure Go and requires C code to be compiled.
That's because HardenedBSD does not implement its own system calls
because they could conflict with FreeBSD, and HardenedBSD regularly
synchronizes updates from FreeBSD.

Given that context, HardenedBSD does not provide system calls
that can enable or disable feature state, and that leaves the
primary interface as the C libraries that HardenedBSD does
provide. In this case, that interface is
[libhbsdcontrol](https://git.hardenedbsd.org/hardenedbsd/hardenebsd).

The following example queries a list of feature names, and then proceeds
to enable, disable and restore the system default for the "mprotect"
feature. As a final step, we query the status of the "mprotect" feature.
Each method in the example is scoped to the `/usr/bin/mdo` binary:

```go
package main

import (
	"fmt"
	"github.com/0x1eef/bsd/control"
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
```

## Install

The install process is more or less straight forward

    go get github.com/0x1eef/majortom

## Sources

* [github.com/@0x1eef](https://github.com/0x1eef/majortom#readme)
* [gitlab.com/@0x1eef](https://gitlab.com/0x1eef/majortom#about)
* [hardenedbsd.org/@0x1eef](https://git.HardenedBSD.org/0x1eef/majortom#about)

## License

[BSD Zero Clause](https://choosealicense.com/licenses/0bsd/)
<br>
See [LICENSE](./LICENSE)