package response

import "net/http"

type Response struct {
	Writer http.ResponseWriter
}

func (res *Response) AddHeader(name string, value string) {
	res.Writer.Header().Set(name, value)
}

func (res *Response) SetStatus(status int) {
	res.Writer.WriteHeader(status)
}

func NewResponse(writer http.ResponseWriter) *Response {
	r := &Response{
		Writer: writer,
	}
	return r
}
