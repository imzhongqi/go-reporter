package dingtalk

import (
	"context"
	"log"
)

func ExampleNewReporter() {
	reporter := NewReporter(
		"https://oapi.dingtalk.com/robot/send?access_token=xxx",
		Secret("xxx"),
	)
	if err := reporter.Report(context.Background(), NewMessage(
		NewText("我就是我, @XXX 是不一样的烟火"),
		AtAll(),
	)); err != nil {
		log.Fatalf("report the message error: %s", err)
	}
	// Output:
}

func ExampleNewMarkdown() {
	reporter := NewReporter(
		"https://oapi.dingtalk.com/robot/send?access_token=xxx",
		Secret("xxx"),
	)
	if err := reporter.Report(context.Background(), NewMessage(
		NewMarkdown("title", `# title
> references
`),
	)); err != nil {
		log.Fatalf("report the message error: %s", err)
	}
	// Output:
}
