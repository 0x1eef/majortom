## About

The majortom module provides Go bindings for the libhbsdcontrol
library for the [HardenedBSD](https://git.hardenedbsd.org/hardenedbsd/hardenedbsd)
operating system. The library provides an interface that can enable, disable,
restore and query feature states for a given file.

## Background

The control package can enable or disable security features
that are managed by the [HardenedBSD](https://hardenedbsd.org)
kernel on a per-file basis. The package is not pure Go and
also requires C code to be compiled. The dependency on C largely
exists because HardenedBSD does not implement its own system calls
since they could conflict with FreeBSD.

Since HardenedBSD does not provide system calls that can enable or
disable feature state that leaves the primary interface as the
C libraries that HardenedBSD does provide. In this case, that interface is
[libhbsdcontrol](https://git.hardenedbsd.org/hardenedbsd/hardenebsd).

## Examples

#### Features

The following example demonstrates how to create an instance of
`control.Context` and then how to query all feature names:

```go
package main

import (
	"fmt"
	"github.com/0x1eef/majortom/control"
)

func main() {
	ns := control.Namespace("system")
	ctx, err := control.NewContext(ns)
	if err != nil {
		panic(err)
	}
	defer ctx.Free()

	features, err := ctx.FeatureNames()
	if err != nil {
		panic(err)
	}
	for _, name := range features {
		fmt.Printf("feature: %s\n", name)
	}
}
```

#### Settings

The next example shows how to enable, disable, and restore the system default
settings for a given file and feature:

```go
package main

import (
	"fmt"
	"github.com/0x1eef/majortom/control"
)

func main() {
	ns := control.Namespace("system")
	ctx, err := control.NewContext(ns)
	if err != nil {
		panic(err)
	}
	defer ctx.Free()

	feature, target := "mprotect", "/usr/bin/mdo"
	if err := ctx.Enable(feature, target); err != nil {
		panic(err)
	}
	if err := ctx.Disable(feature, target); err != nil {
		panic(err)
	}
	if err := ctx.Sysdef(feature, target); err != nil {
		panic(err)
	}
}
```

#### Status

The last example demonstrates how to query the status of a feature
for a given file:

```go
package main

import (
	"fmt"
	"github.com/0x1eef/majortom/control"
)

func main() {
	ns := control.Namespace("system")
	ctx, err := control.NewContext(ns)
	if err != nil {
		panic(err)
	}
	defer ctx.Free()

	feature, target := "mprotect", "/usr/bin/mdo"
	status, err := ctx.Status(feature, target)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The %s feature is %s\n", feature, status)
}
```

#### Concurrency

The control package expects that each instance of `control.Context`
is not shared across goroutines, otherwise the behavior is undefined
and it could lead to program crashes. In other words, create one context
per goroutine. The following example spawns three goroutines that
correctly create one context per goroutine, and most important,
the context is not shared between goroutines:

```go
package main

import (
	"fmt"
	"github.com/0x1eef/majortom/control"
)

func worker() {
	ns := control.Namespace("system")
	ctx, err := control.NewContext(ns)
	if err != nil {
		panic(err)
	}
	defer ctx.Free()
}

func main() {
	for i := 0; i < 3; i++ {
		go worker()
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

## Motivation

The main motivation behind this project was to learn more about Go.
I was already familiar with the libhbsdcontrol interface, and since I
stood to learn a lot in the process, I decided to create Go bindings
for libhbsdcontrol. Over time the library has grown and improved as
my Go skills also improved, and I ended up using it in a new utility
I wrote named [control(8)](https://github.com/0x1eef/control).

## License

[BSD Zero Clause](https://choosealicense.com/licenses/0bsd/)
<br>
See [LICENSE](./LICENSE)
