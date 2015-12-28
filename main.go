package main

import (
	"github.com/anton-107/learn-go-gateway/application"
	"github.com/anton-107/learn-go-gateway/middleware/identifier"
)

func main() {
	app := application.Application{}
	app.Middleware.Add(identifier.NewRequestIdentifier())
	app.Middleware.Add(identifier.NewResponseIdentifier())
	app.Listen(3020)
}
