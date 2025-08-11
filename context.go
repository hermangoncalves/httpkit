package httpkit

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Context represents the context of an HTTP request in the httpkit framework.
// It wraps Go's standard http.ResponseWriter and *http.Request,
// and adds a plugin store for sharing data or functionality between middleware and handlers.
type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Plugins map[string]any
}

// H is a shorthand for map[string]any.
// It is commonly used to construct JSON responses or generic key-value maps.
//
// Example:
//
//	ctx.JSON(http.StatusOK, httpkit.H{"message": "Hello, world"})
type H map[string]any

// JSON writes the given object as a JSON response with the specified HTTP status code.
//
// Example:
//
//	ctx.JSON(http.StatusOK, httpkit.H{"status": "success"})
func (ctx *Context) JSON(code int, obj any) {
	ctx.Writer.WriteHeader(code)
	ctx.Writer.Header().Set("Content-type", "application/json")
	json.NewEncoder(ctx.Writer).Encode(obj)
}

// Query retrieves the value of a query string parameter from the request URL.
// If the parameter is not present, it returns an empty string.
//
// Example:
//
//	// GET /search?q=golang
//	query := ctx.Query("q") // returns "golang"
func (ctx *Context) Query(key string) string {
	values := ctx.Request.URL.Query()
	return values.Get(key)
}

// Param retrieves the value of a path parameter from the request URL.
// Path parameters require Go 1.22+ with the new ServeMux pattern syntax.
//
// Example:
//
//	// Registered pattern: /users/{id}
//	// Request: GET /users/42
//	id := ctx.Param("id") // returns "42"
func (ctx *Context) Param(name string) string {
	return ctx.Request.PathValue(name)
}

func getTyped[T any](ctx *Context, key any) (result T) {
	if val := ctx.Get(key); val != nil {
		result, _ = val.(T)
	}
	return
}

func (ctx *Context) Get(key any) any {
	return ctx.Request.Context().Value(key)
}

func (ctx *Context) GetString(key any) string {
	return getTyped[string](ctx, key)
}

func (ctx *Context) GetInt(key any) int {
	return getTyped[int](ctx, key)
}

func (ctx *Context) GetBool(key any) bool {
	return getTyped[bool](ctx, key)
}

func (ctx *Context) GetFloat64(key any) float64 {
	return getTyped[float64](ctx, key)
}

func (ctx *Context) GetFloat32(key any) float32 {
	return getTyped[float32](ctx, key)
}

func (ctx *Context) GetInt64(key any) int64 {
	return getTyped[int64](ctx, key)
}

func (ctx *Context) GetInt32(key any) int32 {
	return getTyped[int32](ctx, key)
}

func (ctx *Context) GetUint(key any) uint {
	return getTyped[uint](ctx, key)
}

func (ctx *Context) GetUint64(key any) uint64 {
	return getTyped[uint64](ctx, key)
}

func (ctx *Context) GetUint32(key any) uint32 {
	return getTyped[uint32](ctx, key)
}

func (ctx *Context) GetTime(key any) time.Time {
	return getTyped[time.Time](ctx, key)
}

func (ctx *Context) GetDuration(key any) time.Duration {
	return getTyped[time.Duration](ctx, key)
}

func (ctx *Context) GetStringSlice(key any) []string {
	return getTyped[[]string](ctx, key)
}

func (ctx *Context) GetStringMap(key any) map[string]any {
	return getTyped[map[string]any](ctx, key)
}

// DecodeJSON reads and decodes a JSON request body into the provided destination object.
//
// Example:
//
//	var data struct {
//	    Name string `json:"name"`
//	}
//	if err := ctx.DecodeJSON(&data); err != nil {
//	    ctx.JSON(http.StatusBadRequest, httpkit.H{"error": err.Error()})
//	    return
//	}
func (ctx *Context) DecodeJSON(obj any) error {
	contentType := ctx.Request.Header.Get("Content-Type")
	if !strings.HasPrefix(strings.ToLower(contentType), "application/json") {
		return errors.New("content-type is not application/json")
	}

	if ctx.Request.Body == nil {
		return errors.New("empty request body")
	}

	if err := json.NewDecoder(ctx.Request.Body).Decode(obj); err != nil {
		return err
	}

	return nil
}

// GetPlugin retrieves a registered plugin of the specified type from the context's plugin store.
// It returns the plugin instance if found and correctly typed, otherwise it returns the zero value of T.
//
// Example:
//
//	logger, err := httpkit.GetPlugin[*LoggerPlugin](ctx, "logger")
//	if err != nil {
//	    // Handle missing or wrong type plugin.
//	    log.Println(err)
//	    return
//	}
//	logger.Info("Processing request")
func GetPlugin[T Plugin](ctx *Context, pluginName string) (T, error) {
	plugin, ok := ctx.Plugins[pluginName].(T)
	if !ok {
		var zero T
		return zero, fmt.Errorf("plugin %q not found or wrong type", pluginName)
	}
	return plugin, nil
}
