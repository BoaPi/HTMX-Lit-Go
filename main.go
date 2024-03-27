package main

import (
	"fmt"
	"net/http"

	"education-htmx-server/handler"
	"education-htmx-server/middleware"
	"education-htmx-server/templates"

	"github.com/a-h/templ"
)

func main() {
	// setup router
	router := http.NewServeMux()

	// setup static file serving for static file like hmtx
	static := "./static"
	fileServer := http.FileServer(http.Dir(static))

	// file server handler
	router.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// pages handler
	router.Handle("/", templ.Handler(templates.Home()))
	router.Handle("/board", templ.Handler(templates.Board()))

	// partial handler
	router.HandleFunc("GET /timer", handler.Timer)

	// server settings
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.Logger(router),
	}

	// start server
	fmt.Printf("Run Server and Serving on http://localhost:8080\n")
	server.ListenAndServe()
}
