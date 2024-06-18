package syncx

import "time"

type Option func(*option)

type option struct {
	timeout time.Duration
}

func makeOpt(opts []Option) option {
	var opt option
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

func WithTimeout(d time.Duration) Option {
	return func(o *option) {
		o.timeout = d
	}
}

var globalTimeout time.Duration

func SetGlobalTimeout(d time.Duration) {
	globalTimeout = d
}
