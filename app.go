package httpkit

import (
	"fmt"
	"net/http"
	"os"
)

type HandlerFunc func(ctx *Context)

type App struct {
	mux *http.ServeMux
}

func New() *App {
	return &App{
		mux: http.NewServeMux(),
	}
}

func (app *App) Handle(pattern string, handler HandlerFunc) {
	app.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		ctx := &Context{
			Writer:   w,
			Resquest: r,
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
