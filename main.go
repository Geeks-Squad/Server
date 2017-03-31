package main

import (
	"log"
	"net/http"

	"os"

	handler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/vyasgiridhar/server/database"
	"github.com/vyasgiridhar/server/handlers"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	authCandidateID := database.BasicAuth{Next: http.HandlerFunc(handlers.GetCandidate)}
	authCandidateName := database.BasicAuth{Next: http.HandlerFunc(handlers.GetCandidateName)}
	authCandidateSkill := database.BasicAuth{Next: http.HandlerFunc(handlers.GetSkill)}
	authRegistration := database.BasicAuth{Next: http.HandlerFunc(handlers.GetCandidateRegistration)}

	router.HandleFunc("/Signup", handlers.SignUpHandler)
	router.HandleFunc("/Login", handlers.LoginHandler)
	//Candidate SubRouter

	citizenRouter := router.PathPrefix("/candidate").Subrouter()
	citizenRouter.Handle("/id/{id}", authCandidateID).Methods("GET")
	citizenRouter.Handle("/name/{name}", authCandidateName).Methods("GET")

	//Skill SubRouter
	citizenSkill := router.PathPrefix("/Skill").Subrouter()
	citizenSkill.Handle("/cid/{cid}", authCandidateSkill).Methods("GET")

	//Registration SubRouter
	registrationRouter := router.PathPrefix("/registration").Subrouter()
	registrationRouter.HandleFunc("/training/{training}", handlers.GetTrainingCandidates).Methods("GET")
	registrationRouter.Handle("/candidate/{id}", authRegistration).Methods("GET")

	//Training SubRouter
	trainingRouter := router.PathPrefix("/training").Subrouter()
	trainingRouter.HandleFunc("/id/{id}", handlers.GetTraining).Methods("GET")
	trainingRouter.HandleFunc("/centre/{id}", handlers.GetCentreTraining).Methods("GET")
	trainingRouter.HandleFunc("/jobs/{trainingid}", handlers.GetJobsTraining).Methods("GET")

	//Jobs SubRouter
	jobRouter := router.PathPrefix("/jobs").Subrouter()
	jobRouter.HandleFunc("training/{jobid}", handlers.GetTrainingForJob).Methods("GET")

	//Industry SubRouter
	industryRouter := router.PathPrefix("/industry").Subrouter()
	industryRouter.HandleFunc("/name/{name}", handlers.GetIndustry).Methods("GET")
	industryRouter.HandleFunc("/jobs/{industry}", handlers.GetIndustryJobs).Methods("GET")

	//Training Centre SubRouter
	tcentreRouter := router.PathPrefix("/tcentre").Subrouter()
	tcentreRouter.HandleFunc("/id/{id}", handlers.GetTCentres).Methods("GET")

	//Query SubRouter
	queryRouter := router.PathPrefix("/query").Subrouter()
	queryRouter.HandleFunc("/submit", handlers.SubmitQuery).Methods("POST")
	queryRouter.HandleFunc("/get", handlers.GetQueries).Methods("GET")

	//Notif subRouter
	notifRouter := router.PathPrefix("/notif").Subrouter()
	notifRouter.HandleFunc("/send", handlers.AddNotif).Methods("POST")
	notifRouter.HandleFunc("/get", handlers.GetNotifs).Methods("GET")
	notifRouter.HandleFunc("/sig/{lastPull", handlers.GetNotifsSig).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080",
		handler.LoggingHandler(os.Stdout, router)))
}
