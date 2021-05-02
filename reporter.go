package reporter

import (
	"context"
)

var DefaultReporter Reporter = &nopReporter{}

type Message interface {
	Metadata() map[string]string
	Body() interface{}
}

type Reporter interface {
	Report(ctx context.Context, message Message) error
}

type nopReporter struct{}

func (n nopReporter) Report(ctx context.Context, message Message) error { return nil }

func Report(ctx context.Context, message Message) error {
	return DefaultReporter.Report(ctx, message)
}
