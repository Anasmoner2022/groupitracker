package main

import (
	"groupie-tracker/handlers"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handlers.HomeHandler)
	http.ListenAndServe("localhost:8080", nil)
}
