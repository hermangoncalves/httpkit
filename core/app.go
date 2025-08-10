package core

import (
	"fmt"
	"net/http"
	"os"
)

type Middleware func(HandlerFunc) HandlerFunc
type HandlerFunc func(ctx *Context)

type App struct {
	mux         *http.ServeMux
	middlewares []Middleware
}

func New() *App {
	return &App{
		mux:         http.NewServeMux(),
		middlewares: make([]Middleware, 0),
	}
}

func (app *App) Use(middleware Middleware) {
	app.middlewares = append(app.middlewares, middleware)
}

func (app *App) Handle(pattern string, handler HandlerFunc) {
	for i := len(app.middlewares) - 1; i >= 0; i-- {
		handler = app.middlewares[i](handler)
	}

	app.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		ctx := &Context{
			Writer:  w,
			Request: r,
			Plugins: make(map[string]any),
		}
		handler(ctx)
	})
}

func (app *App) Run(addr ...string) {
	address := resolveAddress(addr)
	fmt.Printf("Server is running on port %s\n", address)
	http.ListenAndServe(address, app.mux)
}

func resolveAddress(addr []string) string {
	switch len(addr) {
	case 0:
		if port := os.Getenv("PORT"); port != "" {
			return ":" + port
		}
		return ":8080"
	case 1:
		return addr[0]
	default:
		panic("too many parameters")
	}
}
