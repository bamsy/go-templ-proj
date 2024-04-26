package router

import (
	"log"
	"net/http"
	"time"

	"github.com/bamsy23/go-templ-proj/model"
	"golang.org/x/net/context"
)

// Wrapper around a http.ServeMux to serve up middleware
type App struct {
	mux        *http.ServeMux
	middleware []func(http.HandlerFunc) http.HandlerFunc
}

func NewApp() *App {
	return &App{
		mux:        http.NewServeMux(),
		middleware: []func(http.HandlerFunc) http.HandlerFunc{},
	}
}

func (a *App) Use(middleware ...func(http.HandlerFunc) http.HandlerFunc) {
	a.middleware = append(a.middleware, middleware...)
}

func (a *App) HandleFunc(pattern string, handler http.HandlerFunc) {
	wrappedHandler := a.applyMiddleware(handler)
	a.mux.HandleFunc(pattern, wrappedHandler)
}

func (a *App) applyMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	for i := len(a.middleware) - 1; i >= 0; i-- {
		handler = a.middleware[i](handler)
	}
	return handler
}

func (a *App) Handle(pattern string, handler http.Handler) {
	a.mux.Handle(pattern, handler)
}

func (a *App) Start(addr string) error {
	log.Println("Starting server on", addr)
	return http.ListenAndServe(addr, a.mux)
}

// withUser adds a user to the request context.
func WithUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user", model.User{
			Email: "test@gmail.com",
		})
		newReq := r.WithContext(ctx)
		next(w, newReq)
	}
}

// withLogging adds a logging middleware to the app.
func WithLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// Call the next handler
		next(w, r)

		duration := time.Since(start)
		log.Printf("%s %s %v", r.Method, r.URL.Path, duration)
	}
}
