package dingtalk

type Content interface {
	Type() string
	Body() interface{}
}

type text struct {
	Content string `json:"content"`
}

func NewText(content string) Content {
	return &text{
		Content: content,
	}
}

func (t *text) Type() string {
	return "text"
}

func (t *text) Body() interface{} {
	return t
}

type markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func NewMarkdown(title string, text string) Content {
	return &markdown{
		Title: title,
		Text:  text,
	}
}

func (m *markdown) Type() string {
	return "markdown"
}

func (m *markdown) Body() interface{} {
	return m
}

type link struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

func NewLink(title, text, MessageUrl string, picUrl ...string) Content {
	l := &link{
		Title:      title,
		Text:       text,
		MessageUrl: MessageUrl,
	}
	if len(picUrl) > 0 && picUrl[0] != "" {
		l.PicUrl = picUrl[0]
	}
	return l
}

func (l *link) Type() string {
	return "link"
}

func (l *link) Body() interface{} {
	return l
}
