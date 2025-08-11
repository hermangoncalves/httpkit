package httpkit

// Plugin represents an httpkit extension that can modify request handling.
//
// A Plugin must have a unique Name (used for retrieval from Context)
// and provide a Middleware function that wraps request processing.
//
// Example:
//
//	type LoggerPlugin struct {}
//
//	func (l *LoggerPlugin) Name() string {
//	    return "logger"
//	}
//
//	func (l *LoggerPlugin) Middleware() httpkit.Middleware {
//	    return func(next httpkit.HandlerFunc) httpkit.HandlerFunc {
//	        return func(ctx *httpkit.Context) {
//	            log.Println("Request received")
//	            next(ctx)
//	        }
//	    }
//	}
type Plugin interface {
	// Name returns the unique identifier for the plugin.
	// This name is used to store and retrieve the plugin from Context.Plugins.
	Name() string
	// Middleware returns a Middleware function that wraps the provided HandlerFunc.
	// It can perform actions before and/or after the handler executes, such as logging,
	// authentication, or modifying the request/response.
	Middleware() Middleware
}
