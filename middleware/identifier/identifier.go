package identifier

import (
	"fmt"
	"github.com/anton-107/learn-go-gateway/request"
	"github.com/anton-107/learn-go-gateway/response"
	"github.com/pborman/uuid"
)

const idHeader string = "Request-Id"

type RequestIdentifier struct{}

func (mw *RequestIdentifier) Handle(req *request.Request, res *response.Response, next func()) {
	fmt.Println("Request identifier is called")
	req.SetHeader(idHeader, uuid.NewRandom().String())
	next()
}

type ResponseIdentifier struct{}

func (mw *ResponseIdentifier) Handle(req *request.Request, res *response.Response, next func()) {
	fmt.Println("Response identifier is called")
	res.AddHeader(idHeader, req.GetHeader(idHeader))
	next()
}

func NewRequestIdentifier() *RequestIdentifier {
	r := RequestIdentifier{}
	return &r
}

func NewResponseIdentifier() *ResponseIdentifier {
	r := ResponseIdentifier{}
	return &r
}
