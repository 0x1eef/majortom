package control

import (
	"errors"
)

const (
	Version = "0.3.0"
)

var (
	ErrUseAfterFree = errors.New("context has been freed")
	ErrNullPtr      = errors.New("null pointer")
)

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
