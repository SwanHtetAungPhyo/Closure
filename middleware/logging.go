package middleware

import (
	"log"
	"time"

	"github.com/SwanHtetAungPhyo/closure/closure"
	"github.com/valyala/fasthttp"
)

const (
	LOG    = "LOG"
	green  = "\033[32m"
	reset  = "\033[0m"
	red    = "\033[31m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	magenta= "\033[35m"
	cyan   = "\033[36m"
	white  = "\033[37m"
	black  = "\033[30m"
	gray   = "\033[90m"
	lightRed   = "\033[91m"
	lightGreen = "\033[92m"
	lightYellow = "\033[93m"
	lightBlue  = "\033[94m"
	lightMagenta = "\033[95m"
	lightCyan   = "\033[96m"
	lightWhite  = "\033[97m"
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