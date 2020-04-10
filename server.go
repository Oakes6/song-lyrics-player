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

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/playlists", handlePlaylists)

	fmt.Println("Listening on port 8080...")
	server.ListenAndServe()
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("home.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "home", nil)
}

func handlePlaylists(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("playlists.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "playlists.html", nil)
}
