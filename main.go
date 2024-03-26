package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"htmx-server/templates"

	"github.com/a-h/templ"
)

func timerHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()

	currentSeconds := currentTime.Second()

	fmt.Fprintf(w, "Current seconds: %s", strconv.Itoa(currentSeconds))
}

func main() {
	// setup static file serving for vendor file like hmtx
	vendor := "./static"
	fileServer := http.FileServer(http.Dir(vendor))

	// setup server and handler
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.Handle("/", templ.Handler(templates.Page()))

	// timer
	http.HandleFunc("/timer", timerHandler)

	fmt.Printf("Run Server and Serving on http://localhost:8080\n")
	http.ListenAndServe("localhost:8080", nil)
}
