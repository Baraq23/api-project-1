package handlers

import (
	"groupie-tracker/errorhandlers"
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorhandlers.Server404Page(w, r)
		log.Println("Page not found: ", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		errorhandlers.Server404Page(w, r)
		log.Println("Error parsing index.html:", err)
		return
	}

	err = tmpl.Execute(w, nil)

}
