package middleware

import (
	"fmt"
	"sync"

	"github.com/SwanHtetAungPhyo/closure/closure"
	logging "github.com/SwanHtetAungPhyo/closure/log"
	"github.com/valyala/fasthttp"
)

var statement = fmt.Sprintf("%s[CLOCACHE] %s", cyan, reset)

func CacheMiddleware(key string) *closure.Middleware {
	clocache := make(map[string]any)
	var mu sync.RWMutex

	return &closure.Middleware{
		Name: "CloCache",
		Handler: func(next closure.Handler) closure.Handler {
			return func(ctx *fasthttp.RequestCtx) interface{} {
				cacheKey := ctx.URI().String() + key

				mu.RLock()
				defer mu.RUnlock()
				if data, found := clocache[cacheKey]; found {
					logging.Logger.Printf("Cache hit with this Key: %s", cacheKey)
					ctx.SetBody(data.([]byte))
				} else {
					logging.Logger.Printf("Cache miss for: %s\n", cacheKey)
					var responseBody []byte
					ctx.SetBody([]byte{})

	
					next(ctx)
					responseBody = ctx.Response.Body()

					mu.Lock()
					defer mu.Unlock()
					clocache[cacheKey] = responseBody
					ctx.SetBody(responseBody)
				}
				return nil
			}
		},
	}
}
