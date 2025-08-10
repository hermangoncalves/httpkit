package middleware

import (
	"fmt"
	"net/http"

	"github.com/hermangoncalves/httpkit/core"
)

func Recover(next core.HandlerFunc) core.HandlerFunc {
	return func(ctx *core.Context) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered panic: %v", r)
				http.Error(ctx.Writer, "Internal Server error", http.StatusInternalServerError)
			}
		}()
		next(ctx)
	}
}
