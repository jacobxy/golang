package base

func (req *Request) Valid() bool {
	return req.httpReq != nil && req.httpReq.URL != nil
}

func (resp *Response) Valid() bool {
	return resp.httpResp != nil && resp.httpResp.Body != nil
}

func (item Item) Valid() bool {
	return item != nil
}
