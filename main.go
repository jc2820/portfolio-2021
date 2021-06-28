package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var templates *template.Template

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "about.html", nil)
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "projects.html", nil)
}

func techHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "tech.html", nil)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "2820"
	}

	fs := http.FileServer(http.Dir("static"))
	templates = template.Must(template.ParseGlob("templates/*.html"))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", aboutHandler)
	mux.HandleFunc("/projects", projectsHandler)
	mux.HandleFunc("/technology", techHandler)

	fmt.Printf("Listening on %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
