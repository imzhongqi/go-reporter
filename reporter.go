package reporter

import (
	"context"
	"math/rand"
	"time"
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

type reporterPool struct {
	rand.Source
	reporterList []Reporter
}

func NewReporterPool(reporterList ...Reporter) Reporter {
	return &reporterPool{
		Source:       rand.NewSource(time.Now().UnixNano()),
		reporterList: reporterList,
	}
}

func (r *reporterPool) Report(ctx context.Context, message Message) error {
	reporterNum := len(r.reporterList)
	if reporterNum == 0 {
		return nil
	}
	if reporter := r.reporterList[int(r.Int63())%reporterNum]; reporter != nil {
		return reporter.Report(ctx, message)
	}
	return nil
}
