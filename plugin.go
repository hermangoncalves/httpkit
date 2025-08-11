package httpkit

// Plugin defines the interface that any httpkit plugin must implement.
// A Plugin provides a name and a middleware function that can be registered with an App.
type Plugin interface {
	// Name returns the unique name of the plugin.
	Name() string

	// Middleware returns the Middleware function that the plugin uses.
	// This middleware will be executed as part of the request handling chain.
	Middleware() Middleware
}
