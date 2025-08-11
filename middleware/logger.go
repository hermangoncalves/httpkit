package middleware

import (
	"log"
	"time"

	"github.com/hermangoncalves/httpkit"
)

func Logging(next httpkit.HandlerFunc) httpkit.HandlerFunc {
	return func(ctx *httpkit.Context) {
		start := time.Now()

		next(ctx)

		duration := time.Since(start)
		log.Printf(
			"%s %s %s",
			ctx.Request.Method,
			ctx.Request.URL.Path,
			duration,
		)
	}
}
