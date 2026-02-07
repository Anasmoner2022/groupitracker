package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe("localhost:8080", nil)
}

func homeHandler(rw http.ResponseWriter, rq *http.Request) {
	rw.Write([]byte("HelloWorld"))
}
