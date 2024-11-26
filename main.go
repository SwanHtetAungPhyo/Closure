package main

import (
	"github.com/SwanHtetAungPhyo/closure/closure"
	"github.com/SwanHtetAungPhyo/closure/middleware"
	"github.com/valyala/fasthttp"
)

type user struct{
	Name string `json:"name"`
}

func main(){
	app := closure.New()
	jwtMiddleware := middleware.JWTMiddleware("your-secret-key")


	cors := middleware.NewCORSMiddleware().
		AllowOrigins([]string{"*"}).
		AllowMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}).
		AllowHeaders([]string{"Content-Type", "Authorization", "X-Requested-With"})
	
	app.ApplyMiddleware(*cors.ToMiddleware())

	app.Cluster("/api", func(api *closure.Cluster) {
		app.ApplyMiddleware(*jwtMiddleware)
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
	closure.JSONfiy(ctx,fasthttp.StatusAccepted,"user created", user.Name )
	return nil 
}

func PostHandler(ctx *fasthttp.RequestCtx) interface{} {
	return `{"status": "User created"}`
}