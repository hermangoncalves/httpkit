package middleware

import (
	"fmt"
	"net/http"

	"github.com/hermangoncalves/httpkit"
)

func Recover(next httpkit.HandlerFunc) httpkit.HandlerFunc {
	return func(ctx *httpkit.Context) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered panic: %v", r)
				http.Error(ctx.Writer, "Internal Server error", http.StatusInternalServerError)
			}
		}()
		next(ctx)
	}
}
