package main

import (
	"log"
	"net/http"
	"main.go/api"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/auth", api.AuthHandler).Methods("POST")
	r.HandleFunc("/user", api.TokenMiddleware(api.UserHandler)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
