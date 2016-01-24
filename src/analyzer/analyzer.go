package analyzer

import (
	"base"
)

type Analyzer interface {
	Id() uint32
	Analyzer(
		respParsers []ParseResponse,
		resp base.Response) ([]base.Data, []error)
}

type ParseResponse func(httpResp *http.Response, respDepth uint32) ([]base.Data, []error)

type AnalyzerPool interface {
	Take() (Analyzer, error)
	Return(analyzer Analyzer) error
	Total() uint32
	Used() uint32
}
