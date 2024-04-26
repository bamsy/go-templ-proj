package main

import (
	"embed"
	"net/http"

	"github.com/bamsy23/go-templ-proj/handler"
	"github.com/bamsy23/go-templ-proj/router"
)

//go:embed static
var FS embed.FS

func main() {
	app := router.NewApp()

	// Handlers
	userHandler := handler.UserHandler{}

	// Register Middleware
	app.Use(router.WithLogging)
	app.Use(router.WithUser)

	//Routes
	app.HandleFunc("GET /user", userHandler.HandleUserShow)
	app.Handle("GET /public/", http.StripPrefix("/public/", http.FileServer(http.FS(FS))))

	// Start the server
	app.Start(":8080")
}
