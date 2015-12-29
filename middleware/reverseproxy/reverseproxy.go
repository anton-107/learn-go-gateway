package reverseproxy

import (
	"fmt"
	"github.com/anton-107/learn-go-gateway/request"
	"github.com/anton-107/learn-go-gateway/response"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

type ReverseProxy struct{}

func (mw *ReverseProxy) Handle(req *request.Request, res *response.Response, next func()) {
	if req.UpstreamService == nil {
		res.SetStatus(http.StatusNotFound)

		next()
		return
	}

	fmt.Printf("Proxying request to host: %s", req.UpstreamService.Address)
	fmt.Printf("Proxying request to port: %s", req.UpstreamService.Port)

	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   req.UpstreamService.Address + ":" + strconv.Itoa(req.UpstreamService.Port),
	})

	proxy.ServeHTTP(res.Writer, req.HTTP)
	next()
}

func NewReverseProxy() *ReverseProxy {
	return &ReverseProxy{}
}
