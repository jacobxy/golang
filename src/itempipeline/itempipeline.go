package itempipeliner

import (
	"base"
)

type ItemPipeline interface {
	Send(item base.Item) []error
	FailFast() bool
	SetFailFast(failFast bool)
	Count() []uint64
	ProcessingNunber() uint64
	Summary() string
}
