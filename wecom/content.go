package wecom

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

type Content interface {
	Type() string
	Body() interface{}
}

type text struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list,omitempty"`
	MentionedMobileList []string `json:"mentioned_mobile_list,omitempty"`
}

type TextOption func(t *text)

func MentionedList(l ...string) TextOption {
	return func(t *text) {
		t.MentionedList = l
	}
}

func MentionedMobileList(l ...string) TextOption {
	return func(t *text) {
		t.MentionedMobileList = l
	}
}

func NewText(content string, opts ...TextOption) Content {
	t := &text{
		Content: content,
	}
	for _, o := range opts {
		o(t)
	}
	return t
}

func (t *text) Type() string {
	return "text"
}

func (t *text) Body() interface{} {
	return t
}

type markdown struct {
	Content string `json:"content"`
}

func NewMarkdown(c string) Content {
	return &markdown{
		Content: c,
	}
}

func (m *markdown) Type() string {
	return "markdown"
}

func (m *markdown) Body() interface{} {
	return m
}

type image struct {
	Base64 string `json:"base64"` // 图片内容的base64编码
	Md5    string `json:"md5"`    // 图片内容（base64编码前）的md5值
}

func NewImage(b64, md5 string) Content {
	return &image{
		Base64: b64,
		Md5:    md5,
	}
}

func NewImageFromBytes(b []byte) Content {
	h := md5.New()
	h.Write(b)
	return &image{
		Base64: base64.StdEncoding.EncodeToString(b),
		Md5:    hex.EncodeToString(h.Sum(nil)),
	}
}

func (i *image) Type() string {
	return "image"
}

func (i *image) Body() interface{} {
	return i
}

type news struct {
	Articles []*Article
}

func (n *news) Type() string {
	return "news"
}

func (n *news) Body() interface{} {
	return n
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PicUrl      string `json:"picurl"`
}

func NewNews(articles []*Article) Content {
	return &news{
		Articles: articles,
	}
}

type file struct {
	MediaId string `json:"media_id"`
}

func NewFile(mediaId string) Content {
	return &file{
		MediaId: mediaId,
	}
}

func (f file) Type() string {
	return "file"
}

func (f *file) Body() interface{} {
	return f
}
