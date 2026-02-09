package handlers

import (
	"groupie-tracker/services"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	var err error
	tmpl, err = template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
}

func HomeHandler(rw http.ResponseWriter, rq *http.Request) {
	//* Fetch artists from API
	artists, err := services.FetchArtist()
	if err != nil {
		//! Log the Error
		log.Printf("Error fetching artists: %v", err)

		//! Send Error to User
		http.Error(rw, "Failed to load artists", http.StatusInternalServerError)
		return
	}

	// Execute template with artists data
	err = tmpl.Execute(rw, artists)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(rw, "Failed to render page", http.StatusInternalServerError)
	}
}
