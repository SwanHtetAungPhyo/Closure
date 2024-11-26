package closure

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/valyala/fasthttp"
)

const (
	green = "\033[32m"
	reset = "\033[0m"
)

var statement = fmt.Sprintf("%s[CLOSURE] %s", green, reset)

type App struct {
	router    *Router
	middleware []Middleware
	Logger    *log.Logger
}

func New() *App {
	return &App{
		router:    NewRouter(),
		middleware: []Middleware{},
		Logger:    log.New(os.Stdout, statement, log.LstdFlags),
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
	a.Logger.Printf("Server is listening at the address %shttp://localhost:%s", green, addr)
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
		a.Logger.Printf("Received shutdown signal: %v", stop)
	case err := <-serverError:
		a.Logger.Fatalf("Server error %s", err.Error())
		return
	}
	a.Logger.Println("Shutting down the server .....")
	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 5*time.Second)
	defer shutdownCancel()

	if err := server.ShutdownWithContext(shutdownCtx); err != nil {
		a.Logger.Fatalf("Error occurred during shutdown %s", err.Error())
	} else {
		a.Logger.Fatal("Server is successfully shut down")
	}
}
