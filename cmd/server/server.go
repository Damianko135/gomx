package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

// Preload templates once for better performance
var templates = template.Must(template.ParseGlob("web/templates/*.html"))

// Handler is the main entry point for the server
func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		serveTemplate(w, "index.html")
	case "/projects":
		serveTemplate(w, "projects.html")
	default:
		http.NotFound(w, r)
	}
}

func StartServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      http.HandlerFunc(Handler),
		ReadTimeout:  10 * time.Second,	
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting server on port " + port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}

// serveTemplate loads the correct HTML file
func serveTemplate(w http.ResponseWriter, name string) {
	err := templates.ExecuteTemplate(w, name, nil)
	if err != nil {
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
	}
}
