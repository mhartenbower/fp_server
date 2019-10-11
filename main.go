package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	InitConnection()
	r := mux.NewRouter()
	r.HandleFunc("/secret", HomeHandler)
	r.HandleFunc("/secret/{secretID}", HomeHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
