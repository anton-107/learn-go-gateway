package reverseproxy

import (
	"fmt"
	"github.com/anton-107/learn-go-gateway/request"
	"github.com/anton-107/learn-go-gateway/response"
	"net/http/httputil"
	"net/url"
)

type ReverseProxy struct{}

func (mw *ReverseProxy) Handle(req *request.Request, res *response.Response, next func()) {
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:3001",
	})
	fmt.Println("Proxying request")
	proxy.ServeHTTP(res.Writer, req.HTTP)
	next()
}

func NewReverseProxy() *ReverseProxy {
	return &ReverseProxy{}
}
