package closure

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	logging "github.com/SwanHtetAungPhyo/closure/log"
	"github.com/valyala/fasthttp"
)

var 	green  = "\033[32m"
type App struct {
	router    *Router
	middleware []Middleware

}

func New() *App {
	return &App{
		router:    NewRouter(),
		middleware: []Middleware{},
	}
}

func (a *App) ApplyMiddleware(mw ...Middleware) *App {
	a.middleware = append(a.middleware, mw...)
	return a
}

func (a *App) Cluster(prefix string, block func(*Cluster)) *App {
	cluster := NewCluster(prefix, a.router, a.middleware...)
	block(cluster)
	return a
}

func (a *App) Start(addr string) {
	logging.Logger.Printf("Server is listening at the address %shttp://localhost:%s", green, addr)
	shutDownChannel := make(chan os.Signal, 1)
	signal.Notify(shutDownChannel, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server := &fasthttp.Server{
		Handler:     a.router.ServeHTTP,
		ReadTimeout: 3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	serverError := make(chan error, 1)
	go func() {
		server.ListenAndServe(addr)
	}()

	select {
	case stop := <-shutDownChannel:
		logging.Logger.Printf("Received shutdown signal: %v", stop)
	case err := <-serverError:
		logging.Logger.Fatalf("Server error %s", err.Error())
		return
	}
	logging.Logger.Println("Shutting down the server .....")
	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 2*time.Second)
	defer shutdownCancel()

	if err := server.ShutdownWithContext(shutdownCtx); err != nil {
		logging.Logger.Fatalf("Error occurred during shutdown %s", err.Error())
	} else {
		logging.Logger.Fatal("Server is successfully shut down")
	}
}
