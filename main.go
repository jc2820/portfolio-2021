package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type pagedata struct {
	Title   string
	Content interface{}
}

var templates *template.Template

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	data := pagedata{"About Me", ""}
	templates.ExecuteTemplate(w, "about.html", data)
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	data := pagedata{"Selected Projects", AllProjects}
	templates.ExecuteTemplate(w, "projects.html", data)
}

func stackHandler(w http.ResponseWriter, r *http.Request) {
	data := pagedata{"Tech Stack", Stack}
	templates.ExecuteTemplate(w, "stack.html", data)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./favicon.ico")
}

func robotsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./robots.txt")
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
	mux.HandleFunc("/stack", stackHandler)
	mux.HandleFunc("/favicon.ico", faviconHandler)
	mux.HandleFunc("/robots.txt", robotsHandler)

	fmt.Printf("Listening on %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
