package balancer

import (
	"github.com/anton-107/learn-go-gateway/request"
	"github.com/anton-107/learn-go-gateway/response"
	"math/rand"
)

type Balancer struct {
}

func (balancer *Balancer) Handle(req *request.Request, res *response.Response, next func()) {
	if len(req.MatchedServices) > 0 {
		req.UpstreamService = req.MatchedServices[rand.Intn(len(req.MatchedServices))]
	}
	next()
}

func NewBalancer() *Balancer {
	r := Balancer{}
	return &r
}
