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
		log.Printf("Error on Parsing template: %v", err)
		return
	}

}

func HomeHandler(rw http.ResponseWriter, rq *http.Request) {

	//* Fetch artist from API
	artist, err := services.FetchArtist()
	if err != nil {
		//! Log the Error
		log.Printf("Error fetching artists: %v", err)

		//! Send Error to User
		http.Error(rw, "Failed to load artists", http.StatusInternalServerError)

		return
	}

	// if len(artist) > 0 {
	// 	fmt.Fprintf(rw, "First artist: %s\n", artist[0].Name)
	// 	fmt.Fprintf(rw, "Members: %v\n", artist[0].Members)
	// } else {
	// 	fmt.Fprintf(rw, "No Artist found")
	// }

	tmpl.Execute(rw, artist)
}
