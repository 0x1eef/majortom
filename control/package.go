package control

type Option func(c *Context)

func Namespace(ns string) Option {
	return func(c *Context) {
		c.namespace = ns
	}
}

func SetFlags(flags uint64) Option {
	return func(c *Context) {
		c.flags = flags
	}
}
