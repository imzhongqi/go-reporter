// dingtalk document on here:
// https://developers.dingtalk.com/document/app/custom-robot-access

package dingtalk

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/imzhongqi/go-reporter"

	jsoniter "github.com/json-iterator/go"
)

type dingTalk struct {
	cli      *http.Client
	endpoint string
	secret   []byte
}

func NewReporter(endpoint string, secret string) reporter.Reporter {
	return &dingTalk{
		cli:      &http.Client{},
		endpoint: endpoint,
		secret:   []byte(secret),
	}
}

func (d *dingTalk) newRequest(ctx context.Context, message reporter.Message) (req *http.Request, err error) {
	buf := &bytes.Buffer{}
	if err = jsoniter.NewEncoder(buf).Encode(message.Body()); err != nil {
		return
	}
	if req, err = http.NewRequestWithContext(ctx, http.MethodPost, d.endpoint, buf); err != nil {
		return
	}

	// set headers
	req.Header.Set("Content-Type", "application/json")

	// sign the request
	if len(d.secret) == 0 {
		return
	}
	timestamp := time.Now().Unix() * 1e3
	sign := hmacSha256Sign(timestamp, d.secret)

	// set params
	params := req.URL.Query()
	params.Set("timestamp", strconv.FormatInt(timestamp, 10))
	params.Set("sign", sign)
	req.URL.RawQuery = params.Encode()
	return
}

type Response struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (e *Response) Error() string {
	return fmt.Sprintf("dingtalk: errcode=%d, errmsg=%s", e.ErrCode, e.ErrMsg)
}

func (d *dingTalk) Report(ctx context.Context, message reporter.Message) (err error) {
	req, err := d.newRequest(ctx, message)
	if err != nil {
		return err
	}

	resp, err := d.cli.Do(req)
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
	return
}

func hmacSha256Sign(timestamp int64, secret []byte) string {
	mac := hmac.New(sha256.New, secret)
	fmt.Fprintf(mac, "%d\n%s", timestamp, secret)
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
