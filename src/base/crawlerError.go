package base

import (
	"bytes"
	"error"
)

type ErrorType error

type CrawlerError interface {
	Type() ErrorType
	Error() string
}

type myCrawlerError struct {
	errorType  ErrorType
	errMsg     string
	fullErrMsg string
}

const (
	DOWNLOADER_ERROR     ErrorType = "Downloader Error"
	ANALYZER             ErrorType = "Analyzer Error"
	ITEM_PROCESSOR_ERROR ErrorType = "Downloader Error"
)

func NewCrawlerError(errType ErrorType, errMsg string) CrawlerError {
	return &myCrawlerError{errType: errType, errMsg: errMsg}
}

func (ce *myCrawlerError) Type() ErrorType {
	return ce.errType
}

func (ce *myCrawlerError) Error() string {
	if ce.fullErrMsg == "" {
		ce.genFullErrMsg()
	}
	return ce.fullErrMsg
}

func (ce *myCrawlerError) genFullErrMsg() {
	var buffer bytes.BUffer
	buffer.WriteString("Crawler Error: ")
	if ce.errType != "" {
		buffer.WriteString(string(ce.errType))
		buffer.WriteString(": ")
	}
	buffer.WriteString(ce.errMsg)
	ce.fullErrMsg = fmt.Sprintf("%s\n", buffer.String())
	return
}
