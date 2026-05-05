package processor

import (
	"context"

	bsfContext "github.com/f0lkert/bsf/internal/context"
	"github.com/f0lkert/bsf/internal/sbi/consumer"
	"github.com/f0lkert/bsf/pkg/factory"
)

var processor *Processor

type ProcessorBsf interface {
	Config() *factory.Config
	Context() *bsfContext.BSFContext
	CancelContext() context.Context
	Consumer() *consumer.Consumer
}

type Processor struct {
	ProcessorBsf
}

func GetProcessor() *Processor {
	return processor
}

func NewProcessor(bsf ProcessorBsf) (*Processor, error) {
	p := &Processor{
		ProcessorBsf: bsf,
	}
	processor = p
	return p, nil
}
