// Package httpkit provides a lightweight, extensible wrapper around Go's standard net/http package.
// It is NOT an HTTP router.
//
// httpkit is designed to help you build HTTP applications using the standard library,
// while making it easier to:
//   - Attach middleware in a clean, composable way.
//   - Register plugins that can be accessed in request handlers.
//   - Share state and utilities across requests via plugins.
//   - Extend net/http without introducing a new routing layer.
//
// # Overview
//
// At its core, httpkit provides:
//   - An App type that wraps an http.ServeMux.
//   - Middleware and Plugin interfaces for reusable request processing.
//   - A Context type for request/response handling and plugin access.
//
// Quick Start
//
//	package main
//
//	import (
//		"fmt"
//		"net/http"
//		"github.com/yourusername/httpkit"
//	)
//
//	type HelloPlugin struct{}
//
//	func (p HelloPlugin) Name() string { return "hello" }
//	func (p HelloPlugin) Middleware() httpkit.Middleware {
//		return func(next httpkit.HandlerFunc) httpkit.HandlerFunc {
//			return func(ctx *httpkit.Context) {
//				fmt.Println("HelloPlugin: before handler")
//				next(ctx)
//			}
//		}
//	}
//
//	func main() {
//		app := httpkit.New()
//		app.RegisterPlugin(HelloPlugin{})
//
//		app.Handle("/", func(ctx *httpkit.Context) {
//			ctx.JSON(http.StatusOK, httpkit.H{"message": "Hello, World!"})
//		})
//
//		app.Run(":8080")
//	}
//
// # Plugin Scope
//
// Plugins can be registered globally via App.RegisterPlugin (available to all routes)
// or passed per-route when calling App.Handle.
//
// See the README for more advanced usage and plugin examples.
package httpkit

import (
	"fmt"
	"net/http"
	"os"
)

// Middleware represents a function that wraps a HandlerFunc to process
// HTTP requests before and/or after calling the next handler.
// Middleware can be used for logging, authentication, metrics, etc.
type Middleware func(HandlerFunc) HandlerFunc

// HandlerFunc defines the signature for request handlers in httpkit.
// It receives a custom Context instead of the standard http.ResponseWriter and *http.Request,
// making it easier to access plugins and utility methods.
type HandlerFunc func(ctx *Context)

// App represents an HTTP application built on top of Go's standard http.ServeMux.
// It supports global and per-route plugins and middleware for flexible request processing.
type App struct {
	mux     *http.ServeMux
	plugins []Plugin
}

// New creates and returns a new App instance with an empty plugin list and a fresh ServeMux.
func New() *App {
	return &App{
		mux:     http.NewServeMux(),
		plugins: make([]Plugin, 0),
	}
}

// Handle registers a route with the given pattern and handler.
// You can optionally provide per-route plugins that will wrap the handler in addition
// to any globally registered plugins.
//
// Example:
//
//	app.Handle("/hello", func(ctx *httpkit.Context) {
//	    ctx.JSON(http.StatusOK, httpkit.H{"message": "Hello"})
//	}, MyPlugin{})
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

// RegisterPlugin adds a plugin to the global plugin list.
// Global plugins apply to all registered routes unless overridden by per-route plugins.
func (app *App) RegisterPlugin(plugin Plugin) {
	app.plugins = append(app.plugins, plugin)
}

// Run starts the HTTP server on the provided address.
// If no address is provided, it defaults to the PORT environment variable or ":8080".
func (app *App) Run(addr ...string) {
	address := resolveAddress(addr)
	fmt.Printf("Server is running on port %s\n", address)
	http.ListenAndServe(address, app.mux)
}

// resolveAddress determines the listening address for the server.
// If no argument is provided, it checks the PORT environment variable, falling back to ":8080".
// If one argument is provided, it uses that directly.
// Providing more than one argument causes a panic.
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
