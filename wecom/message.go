package wecom

import "github.com/imzhongqi/go-reporter"

type message struct {
	Content
}

func (m *message) Metadata() map[string]string {
	return make(map[string]string)
}

func (m *message) Body() interface{} {
	msg := make(map[string]interface{})
	msg["msgtype"] = m.Content.Type()
	msg[m.Type()] = m.Content.Body()
	return msg
}

func NewMessage(content Content) reporter.Message {
	return &message{
		Content: content,
	}
}
