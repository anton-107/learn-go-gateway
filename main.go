package main

import (
	"github.com/anton-107/learn-go-gateway/application"
	"github.com/anton-107/learn-go-gateway/middleware/identifier"
	"github.com/anton-107/learn-go-gateway/middleware/reverseproxy"
	"github.com/anton-107/learn-go-gateway/middleware/servicefinder"
)

func main() {
	app := application.Application{}
	app.Middleware.Add(identifier.NewRequestIdentifier())
	app.Middleware.Add(servicefinder.NewServiceFinder())
	app.Middleware.Add(identifier.NewResponseIdentifier())
	app.Middleware.Add(reverseproxy.NewReverseProxy())
	app.Listen(3020)
}
