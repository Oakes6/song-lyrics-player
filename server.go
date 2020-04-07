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

   fmt.Println("Listening on port 8080...")
	server.ListenAndServe()
}

func handleHome(w http.ResponseWriter, r *http.Request) {
   t, err := template.ParseFiles("index.html")
	if err != nil {
      panic(err)
	}
   t.ExecuteTemplate(w, "index", nil)
}
