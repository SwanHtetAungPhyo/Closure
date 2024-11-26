package middleware

import (
	"log"
	"time"

	"github.com/SwanHtetAungPhyo/closure/closure"
	"github.com/valyala/fasthttp"
)

const (
	LOG   = "LOG"
	green = "\033[32m"
	reset = "\033[0m"
)

var Loggging = func(logger *log.Logger) *closure.Middleware {
	return &closure.Middleware{
		Name: LOG,
		Handler: func(next closure.Handler) closure.Handler {
			return func(ctx *fasthttp.RequestCtx) any {
				logger.Printf("%sRequest received: %s at %v%s\n", green, string(ctx.Path()), time.Now().Local(), reset)
				return next(ctx)
			}
		},
	}
}