package dingtalk

import "net/http"

type Options struct {
	cli    *http.Client
	secret []byte
}

type Option func(o *Options)

func newOptions(opts ...Option) *Options {
	option := &Options{
		cli: http.DefaultClient,
	}
	for _, o := range opts {
		o(option)
	}
	return option
}

func Client(c *http.Client) Option {
	return func(o *Options) {
		o.cli = c
	}
}

func Secret(s string) Option {
	return func(o *Options) {
		if s != "" {
			o.secret = []byte(s)
		}
	}
}
