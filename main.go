package main

import (
	"fmt"
	"html/template"
	"net/http"

	"education-htmx-server/handler"
	"education-htmx-server/middleware"
	// "education-htmx-server/partials"
	// "github.com/a-h/templ"
)

type Templates struct {
	templates *template.Template
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Data struct {
	Count int
}

func main() {
	// setup
	router := http.NewServeMux()
	count := 0
	views := newTemplate()

	// setup static file serving for static file like hmtx
	static := "./static"
	fileServer := http.FileServer(http.Dir(static))

	// file server handler
	router.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// pages handler
	// router.Handle("/", templ.Handler(partials.Home()))
	// router.Handle("/board", templ.Handler(partials.Board()))
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		views.templates.ExecuteTemplate(w, "index", Data{Count: count})
		// increase amount of calls
		count++
	})

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
