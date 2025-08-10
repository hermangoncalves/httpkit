package middleware

import (
	"log"
	"time"

	"github.com/hermangoncalves/httpkit/core"
)

func Logging(next core.HandlerFunc) core.HandlerFunc {
	return func(ctx *core.Context) {
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
