package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/home", handleHome)
	http.HandleFunc("/console", handleConsole)

	fmt.Println("Listening on port 8080...")
	server.ListenAndServe()
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "index", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/home.html", "templates/library.html", "templates/lyrics.html", "templates/search.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "home", nil)
}

func handleConsole(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/home.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "home", nil)
}
