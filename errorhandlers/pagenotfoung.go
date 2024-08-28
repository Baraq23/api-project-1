package errorhandlers

import (
	"html/template"
	"log"
	"net/http"
)

func Server404Page(w http.ResponseWriter, r *http.Request) {
	
	tmpl, err := template.ParseFiles("templates/errorpage.html")
	if err != nil {
		log.Println("Error parsing errorpage.html")
		return
	}

	err = tmpl.Execute(w, nil)

	if err != nil {
		log.Println("Error executing errorpage.html template.")
		return
	}
}