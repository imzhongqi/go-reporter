package dingtalk

import "github.com/imzhongqi/go-reporter"

type atOptions struct {
	IsAtAll   bool     `json:"isAtAll"`
	AtUserIds []string `json:"atUserIds"`
	AtMobiles []string `json:"atMobiles"`
}

type message struct {
	At      *atOptions
	Content Content
}

func (m *message) Metadata() map[string]string {
	return make(map[string]string)
}

func (m *message) Body() interface{} {
	msg := make(map[string]interface{})
	if m.At != nil {
		msg["at"] = m.At
	}
	msg["msgtype"] = m.Content.Type()
	msg[m.Content.Type()] = m.Content.Body()
	return msg
}

type Options struct {
	at *atOptions
}

type Option func(o *Options)

func lazyInitAtOpts(o *Options) {
	if o.at == nil {
		o.at = new(atOptions)
	}
}

func AtAll() Option {
	return func(o *Options) {
		lazyInitAtOpts(o)
		o.at.IsAtAll = true
	}
}

func AtUserIds(userIds ...string) Option {
	return func(o *Options) {
		lazyInitAtOpts(o)
		o.at.AtUserIds = userIds
	}
}

func AtMobiles(mobiles ...string) Option {
	return func(o *Options) {
		lazyInitAtOpts(o)
		o.at.AtMobiles = mobiles
	}
}

func NewMessage(content Content, opts ...Option) reporter.Message {
	options := &Options{}
	for _, o := range opts {
		o(options)
	}
	return &message{
		At:      options.at,
		Content: content,
	}
}
