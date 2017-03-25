package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	authLogin := BasicAuth{http.HandlerFunc(Login)}
	authCitizenID := BasicAuth{http.HandlerFunc(TodoIndex)}

	router.HandleFunc("/Signup", SignUpHandler)

	s := router.PathPrefix("/Citizen").Subrouter()

	s.Handle("/", authLogin)
	s.Handle("/Name/{Name}", authCitizenID)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	signup := SignupBody{}
	decoder := json.NewDecoder(r.Body)

	decoder.Decode(signup)
	Signup(signup)
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Authorized")
	w.Header().Add("Status-Code", string(http.StatusOK))
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
}
