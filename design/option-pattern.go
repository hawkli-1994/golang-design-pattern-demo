package design

import (
	"time"
)

// 选项模式

const (
	defaultTimeout = 10
	defaultCaching = false
)

type options struct {
	timeout  time.Duration
	cacheing bool
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithTimeout(t time.Duration) Option {
	return optionFunc(func(o *options) {
		o.timeout = t
	})
}

func WithCaching(cache bool) Option {
	return optionFunc(func(o *options) {
		o.cacheing = cache
	})
}

type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

func NewConnect(addr string, opts ...Option) (*Connection, error) {
	optIns := options{
		timeout:  defaultTimeout,
		cacheing: defaultCaching,
	}
	for _, o := range opts {
		o.apply(&optIns)
	}

	return &Connection{
		addr:    addr,
		cache:   optIns.cacheing,
		timeout: optIns.timeout,
	}, nil
}
