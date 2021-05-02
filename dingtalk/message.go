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

type MessageOptions struct {
	at *atOptions
}

type MessageOption func(o *MessageOptions)

func lazyInitAtOpts(o *MessageOptions) {
	if o.at == nil {
		o.at = new(atOptions)
	}
}

func AtAll() MessageOption {
	return func(o *MessageOptions) {
		lazyInitAtOpts(o)
		o.at.IsAtAll = true
	}
}

func AtUserIds(userIds ...string) MessageOption {
	return func(o *MessageOptions) {
		lazyInitAtOpts(o)
		o.at.AtUserIds = userIds
	}
}

func AtMobiles(mobiles ...string) MessageOption {
	return func(o *MessageOptions) {
		lazyInitAtOpts(o)
		o.at.AtMobiles = mobiles
	}
}

func NewMessage(content Content, opts ...MessageOption) reporter.Message {
	options := &MessageOptions{}
	for _, o := range opts {
		o(options)
	}
	return &message{
		At:      options.at,
		Content: content,
	}
}
