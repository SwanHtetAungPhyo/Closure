package middleware

import (
	"github.com/SwanHtetAungPhyo/swan_lib"

	"github.com/valyala/fasthttp"
	"github.com/SwanHtetAungPhyo/closure/closure"
)

func JWTMiddleware(secret string) *closure.Middleware {
	manager := swan_lib.NewJWTMiddleware(secret)
	return &closure.Middleware{
		Name: "JWT",
		Handler: func(next closure.Handler) closure.Handler{
			return func(ctx *fasthttp.RequestCtx) any{
				manager.FastAuthorize( func(ctx *fasthttp.RequestCtx) {
					next(ctx)
				})(ctx)
				return nil
			}
		},
	}
}