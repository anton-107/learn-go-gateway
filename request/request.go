package request

import (
	"github.com/hashicorp/consul/api"
	"net/http"
)

type Request struct {
	HTTP            *http.Request
	MatchedServices []*api.AgentService
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

func (req *Request) GetPath() string {
	return req.HTTP.RequestURI
}

func (req *Request) AddMatchedService(serviceDef *api.AgentService) {
	req.MatchedServices = append(req.MatchedServices, serviceDef)
}

func NewRequest(httpRequest *http.Request) *Request {
	r := Request{
		HTTP: httpRequest,
	}
	return &r
}
