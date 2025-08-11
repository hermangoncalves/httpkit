// Package httpkit provides a lightweight wrapper around Go's standard net/http package.
// It is NOT an HTTP router.
//
// It provides:
//
//   - A lightweight App for registering routes and middlewares.
//   - A Context type with helpers for JSON, query/path parameters,
//     and plugin storage.
//   - Support for user-defined middlewares and plugins.
//
// Quick start:
//
//	app := httpkit.New()
//	app.Use(middleware.Logging)
//	app.Handle("/", func(ctx *httpkit.Context) {
//	    ctx.JSON(200, httpkit.H{"hello": "world"})
//	})
//	app.Run() // Uses :8080 or $PORT by default
//
// Typical usage is: create an App, register middlewares/handlers,
// and run the server.
package httpkit

import (
	"fmt"
	"net/http"
	"os"
)

// Middleware defines a function that wraps a HandlerFunc.
// A middleware receives the next handler in the chain and returns
// a new handler, typically performing logic before/after invoking
// the original one.
type Middleware func(HandlerFunc) HandlerFunc

// HandlerFunc is the signature used for httpkit handlers.
// It receives a *Context, which provides utilities for writing
// responses, reading query/path parameters, and accessing plugins.
type HandlerFunc func(ctx *Context)

// App is the main server type of the framework.
// It contains an http.ServeMux for routing and a slice of middlewares
// that are applied to handlers.
type App struct {
	mux     *http.ServeMux
	plugins []Plugin
}

// New creates and returns a new App instance with an initialized
// ServeMux and an empty middleware chain.
func New() *App {
	return &App{
		mux:     http.NewServeMux(),
		plugins: make([]Plugin, 0),
	}
}

// Handle registers a handler for the given pattern.
// The currently registered middlewares are applied to the handler
// before it is registered in the internal ServeMux.
//
// The pattern follows the rules of net/http ServeMux.
func (app *App) Handle(pattern string, handler HandlerFunc, plugin ...Plugin) {
	plugins := append(app.plugins, plugin...)
	for i := len(plugins) - 1; i >= 0; i-- {
		handler = plugins[i].Middleware()(handler)
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

// RegisterPlugin registers a Plugin by adding its Middleware to the
// application's middleware chain. Plugins must implement the Plugin
// interface (Name, Middleware).
func (app *App) RegisterPlugin(plugin Plugin) {
	app.plugins = append(app.plugins, plugin)
}

// Run starts the HTTP server and blocks the current goroutine.
// It accepts an optional address parameter (e.g. ":8080" or
// "0.0.0.0:8080"). If no address is provided, it uses the PORT
// environment variable or ":8080" by default.
func (app *App) Run(addr ...string) {
	address := resolveAddress(addr)
	fmt.Printf("Server is running on port %s\n", address)
	http.ListenAndServe(address, app.mux)
}

// resolveAddress resolves the address the server will listen on
// from the given slice of arguments. If no arguments are provided,
// it checks the PORT environment variable and finally defaults
// to ":8080".
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
