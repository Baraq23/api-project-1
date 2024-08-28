package main

import (
	"log"
	"net/http"

	"groupie-tracker/handlers"
)

func main() {
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.HomeHandler)

	log.Println("http://localhost:8008")
	s := http.Server{
		Addr: "127.0.0.1:8008",
	}
	s.ListenAndServe()
}
