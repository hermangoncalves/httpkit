package httpkit

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Writer   http.ResponseWriter
	Resquest *http.Request
}

type H map[string]any

func (ctx *Context) JSON(code int, obj any) {
	ctx.Writer.WriteHeader(code)
	ctx.Writer.Header().Set("Content-type", "application/json")
	json.NewEncoder(ctx.Writer).Encode(obj)
}
