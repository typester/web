package web

import (
	"github.com/gorilla/mux"
	"net/http"
)

type App struct {
	Router *mux.Router
	Conf   map[string]interface{}
}

type Context struct {
	App     *App
	Request *http.Request
	Stash   map[string]interface{}
	http.ResponseWriter
}

type Handler func(c *Context)

func NewApp() *App {
	app := &App{}
	app.Router = mux.NewRouter()
	app.Conf = map[string]interface{}{}
	return app
}

func (app *App) Run(addr string) error {
	return http.ListenAndServe(addr, app.Router)
}

func (app *App) Handle(path string, f Handler) *mux.Route {
	return app.Router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		context := &Context{app, r, map[string]interface{}{}, w}
		f(context)
	})
}
