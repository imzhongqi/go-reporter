// wecom document: https://work.weixin.qq.com/api/doc/90000/90136/91770

package wecom

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/imzhongqi/go-reporter"
	jsoniter "github.com/json-iterator/go"
)

type wecom struct {
	cli *http.Client

	endpoint string
}

func NewReporter(endpoint string, opts ...Option) reporter.Reporter {
	opt := newOptions(opts...)

	return &wecom{
		cli:      opt.cli,
		endpoint: endpoint,
	}
}

func (w *wecom) newRequest(ctx context.Context, message reporter.Message) (req *http.Request, err error) {
	buf := &bytes.Buffer{}
	if err = jsoniter.NewEncoder(buf).Encode(message.Body()); err != nil {
		return
	}
	if req, err = http.NewRequestWithContext(ctx, http.MethodPost, w.endpoint, buf); err != nil {
		return
	}

	// set headers
	req.Header.Set("Content-Type", "application/json")
	return
}

type Response struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (e *Response) Error() string {
	return fmt.Sprintf("dingtalk: errcode=%d, errmsg=%s", e.ErrCode, e.ErrMsg)
}

func (w *wecom) Report(ctx context.Context, message reporter.Message) (err error) {
	req, err := w.newRequest(ctx, message)
	if err != nil {
		return err
	}

	resp, err := w.cli.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var reply Response
	if err = jsoniter.NewDecoder(resp.Body).Decode(&reply); err != nil {
		return
	}

	if reply.ErrCode != 0 {
		return &reply
	}

	return nil
}
