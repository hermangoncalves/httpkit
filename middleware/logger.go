package middleware

import (
	"log"
	"time"

	"github.com/hermangoncalves/httpkit"
)

func Logging(next httpkit.HandlerFunc) httpkit.HandlerFunc {
	return func(ctx *httpkit.Context) {
		start := time.Now()

		next(ctx) // call next handler

		duration := time.Since(start)
		log.Printf(
			"%s %s %s",
			ctx.Request.Method,
			ctx.Request.URL.Path,
			duration,
		)
	}
}
