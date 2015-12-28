package application

import (
	"fmt"
	"github.com/anton-107/learn-go-gateway/middleware"
	"github.com/anton-107/learn-go-gateway/request"
	"github.com/anton-107/learn-go-gateway/response"
	"net/http"
)

type Application struct {
	Middleware middlleware.Collection
}

func (app *Application) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	app.handle(request.NewRequest(req), response.NewResponse(res), 0)
	//	fmt.Fprintf(res, "TEST")
}

func (app *Application) handle(req *request.Request, res *response.Response, mwIndex int) {

	defer func() {
		if err := recover(); err != nil {
			panic(err)
			//			fmt.Errorf("err: %v", err)
			http.Error(res.Writer, http.StatusText(500), 500)
		}
	}()

	fmt.Println("Calling middleware %d", mwIndex)
	middleware := app.Middleware.Get(mwIndex)
	middleware.Handle(req, res, func() {
		if app.Middleware.Size() > mwIndex+1 {
			mwIndex += 1
			app.handle(req, res, mwIndex)
		}
	})
}

func (app *Application) Listen(port int) {
	address := fmt.Sprintf(":%d", port)
	fmt.Println("Listening to address %s", address)
	http.ListenAndServe(address, app)

}
