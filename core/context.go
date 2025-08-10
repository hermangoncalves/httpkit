package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Plugins map[string]any
}

type H map[string]any

func (ctx *Context) JSON(code int, obj any) {
	ctx.Writer.WriteHeader(code)
	ctx.Writer.Header().Set("Content-type", "application/json")
	json.NewEncoder(ctx.Writer).Encode(obj)
}

func (ctx *Context) Query(key string) string {
	values := ctx.Request.URL.Query()
	return values.Get(key)
}

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

func GetPlugin[T any](ctx *Context, pluginName string) T {
	plugin, ok := ctx.Plugins[pluginName].(T)
	if !ok {
		fmt.Println("Not found")
	}
	return plugin
}
