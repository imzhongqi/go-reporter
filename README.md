## Reporter

提供一系列的告警，例如钉钉，企业微信(*TODO*)等。

**钉钉告警使用**

```go
package main

import (
	"context"
	"log"
	
	"github.com/imzhongqi/go-reporter"
	"github.com/imzhongqi/go-reporter/dingtalk"
)

func main() {
	rp := dingtalk.NewReporter(
		"https://oapi.dingtalk.com/robot/send?access_token=xxx",
		"xxxxx",
	)
	if err := rp.Report(context.Background(), dingtalk.NewMessage(
		dingtalk.NewText("我就是我, @XXX 是不一样的烟火"),
		dingtalk.AtAll(),
	)); err != nil {
		log.Fatalf("report the message error: %s", err)
	}
}
```