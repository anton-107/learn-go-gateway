package request

import (
	"net/http"
)

type Request struct {
	HTTP *http.Request
}

func (req *Request) AddHeader(name string, value string) {
	req.HTTP.Header.Add(name, value)
}

func (req *Request) SetHeader(name string, value string) {
	req.HTTP.Header.Set(name, value)
}

func (req *Request) GetHeader(name string) string {
	return req.HTTP.Header.Get(name)
}

func NewRequest(httpRequest *http.Request) *Request {
	r := Request{
		HTTP: httpRequest,
	}
	return &r
}
