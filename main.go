package main

import (
	"github.com/SwanHtetAungPhyo/closure/closure"
	logging "github.com/SwanHtetAungPhyo/closure/log"
	"github.com/SwanHtetAungPhyo/closure/middleware"
	"github.com/valyala/fasthttp"
)

type user struct{
	Name string `json:"name"`
}

func main(){
	app := closure.New()
	// jwtMiddleware := middleware.JWTMiddleware("your-secret-key")
	cors := middleware.NewCORSMiddleware().
		AllowOrigins([]string{"*"}).
		AllowMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}).
		AllowHeaders([]string{"Content-Type", "Authorization", "X-Requested-With"})
	
	app.ApplyMiddleware(*cors.ToMiddleware())
	// app.ApplyMiddleware((*jwtMiddleware))
	app.Cluster("/api", func(api *closure.Cluster) {
		// cacheMiddleware := middleware.CacheMiddleware("cache-key")
		// api.AddMiddleware(*jwtMiddleware)
		// api.AddMiddleware(*cacheMiddleware)
		api.Get("/", GetHandler)
		api.Post("/user",PostHandler) 
	})

	app.Cluster("/api2",func(api2 *closure.Cluster) {
		api2.Get("/", GetHandler)
		api2.Post("/user",PostHandler)	
	})


	app.Start(":3500")
}

func GetHandler(ctx *fasthttp.RequestCtx) any{
	user := user{
		Name:  "Swan",
	}
	logging.Logger.Println("Get Handler is called")
	closure.JSONfiy(ctx,fasthttp.StatusAccepted,"user created", user.Name )
	return nil 
}

func PostHandler(ctx *fasthttp.RequestCtx) interface{} {
	return `{"status": "User created"}`
}