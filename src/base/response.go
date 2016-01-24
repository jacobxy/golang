package base

import (
	"net/http"
)

type Response struct {
	httpResp *http.Response
	depth    uint32
}

func NewResponse(httpResp *http.Response, depth uint32) {
	return &Response{httpResp: httpResp, depth: depth}
}

func (resp *httpResp) HttpResp() *http.Response {
	return resp.httpResp
}

func (resp *httpResp) Depth() uint32 {
	return resp.depth
}
