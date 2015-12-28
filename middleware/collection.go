package middlleware

import (
	"github.com/anton-107/learn-go-gateway/request"
	"github.com/anton-107/learn-go-gateway/response"
)

type Middleware interface {
	Handle(req *request.Request, res *response.Response, next func())
}

type Collection struct {
	middlewareList []Middleware
}

func (col *Collection) Add(mw Middleware) *Collection {
	col.middlewareList = append(col.middlewareList, mw)
	return col
}

func (col *Collection) Get(index int) Middleware {
	return col.middlewareList[index]
}

func (col *Collection) Size() int {
	return len(col.middlewareList)
}
