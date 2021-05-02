package wecom

import (
	"context"
	"log"
)

func ExampleNewReporter() {
	reporter := NewReporter("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=18da439d-bda9-43cd-a76a-f3b04b730d60")
	if err := reporter.Report(context.Background(), NewMessage(NewText("hello,world"))); err != nil {
		log.Fatalf("report the message error: %s", err)
	}
	// Output:
}
