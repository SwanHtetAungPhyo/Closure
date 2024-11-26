package closure

import "github.com/valyala/fasthttp"
 
type Handler func(ctx *fasthttp.RequestCtx) interface{}

type Router struct{
	routes map[string]map[string]Handler
}

func NewRouter() *Router{
	return &Router{
		routes: make(map[string]map[string]Handler),
	}
}


func (r *Router) Register(method, path string, handler Handler){
	if r.routes[method] == nil{
		r.routes[method] = make(map[string]Handler)
	}
	r.routes[method][path]= handler
}

func (r *Router) ServeHTTP(ctx *fasthttp.RequestCtx) {
	methodRoutes := r.routes[string(ctx.Method())]
	if methodRoutes == nil {
		JSONfiy(ctx,fasthttp.StatusMethodNotAllowed, "Method Not allowed", nil)
		return
	}

	handler, exists := methodRoutes[string(ctx.Path())]
	if !exists {
		JSONfiy(ctx,fasthttp.StatusNotFound, "Not found", nil)
		return
	}

	result := handler(ctx)
	if result != nil {
		JSONfiy(ctx,fasthttp.StatusAccepted, "success", result)
	}
}