package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vyasgiridhar/server/database"
	"github.com/vyasgiridhar/server/handlers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	authLogin := database.BasicAuth{http.HandlerFunc(handlers.Login)}
	authCitizenID := database.BasicAuth{http.HandlerFunc(handlers.GetCitizen)}
	authAddress := database.BasicAuth{http.HandlerFunc(handlers.GetCitizen)}

	router.HandleFunc("/Signup", handlers.SignUpHandler)

	s := router.PathPrefix("/Citizen").Subrouter()

	s.Handle("/Login", authLogin)
	s.Handle("/Address/{Address}", authAddress)
	s.Handle("/{ID}", authCitizenID)

	log.Fatal(http.ListenAndServe(":8080", router))
}
