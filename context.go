package httpkit

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: r,
	}
}

type R map[string]any

func (c *Context) JSON(code int, obj any) error {
	c.Writer.WriteHeader(code)
	c.Writer.Header().Set("Content-type", "application/json")
	return json.NewEncoder(c.Writer).Encode(obj)
}

func (c *Context) Error(err error) {
	if httpError, ok := err.(*HTTPError); ok {
		c.JSON(httpError.Code, httpError)
		return
	}

	c.JSON(http.StatusInternalServerError, &HTTPError{
		Code:    http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
	})
}

func (c *Context) Query(key string) string {
	values := c.Request.URL.Query()
	return values.Get(key)
}

func (c *Context) Param(name string) string {
	return c.Request.PathValue(name)
}

func getTyped[T any](c *Context, key any) (result T) {
	if val := c.Get(key); val != nil {
		result, _ = val.(T)
	}
	return
}

func (c *Context) Get(key any) any {
	return c.Request.Context().Value(key)
}

func (c *Context) GetString(key any) string {
	return getTyped[string](c, key)
}

func (c *Context) GetInt(key any) int {
	return getTyped[int](c, key)
}

func (c *Context) GetBool(key any) bool {
	return getTyped[bool](c, key)
}

func (c *Context) GetFloat64(key any) float64 {
	return getTyped[float64](c, key)
}

func (c *Context) GetFloat32(key any) float32 {
	return getTyped[float32](c, key)
}

func (c *Context) GetInt64(key any) int64 {
	return getTyped[int64](c, key)
}

func (c *Context) GetInt32(key any) int32 {
	return getTyped[int32](c, key)
}

func (c *Context) GetUint(key any) uint {
	return getTyped[uint](c, key)
}

func (c *Context) GetUint64(key any) uint64 {
	return getTyped[uint64](c, key)
}

func (c *Context) GetUint32(key any) uint32 {
	return getTyped[uint32](c, key)
}

func (c *Context) GetTime(key any) time.Time {
	return getTyped[time.Time](c, key)
}

func (c *Context) GetDuration(key any) time.Duration {
	return getTyped[time.Duration](c, key)
}

func (c *Context) GetStringSlice(key any) []string {
	return getTyped[[]string](c, key)
}

func (c *Context) GetStringMap(key any) map[string]any {
	return getTyped[map[string]any](c, key)
}

func (c *Context) BindJSON(obj any) error {
	contentType := c.Request.Header.Get("Content-Type")
	if !strings.HasPrefix(strings.ToLower(contentType), "application/json") {
		return errors.New("content-type is not application/json")
	}

	if c.Request.Body == nil {
		return errors.New("empty request body")
	}

	if err := json.NewDecoder(c.Request.Body).Decode(obj); err != nil {
		return &HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Invalid json" + err.Error(),
		}
	}

	return nil
}
