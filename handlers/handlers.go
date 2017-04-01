package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/vyasgiridhar/server/database"
	"github.com/vyasgiridhar/server/models"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	authorizationArray := r.Header["Authorization"]

	authorization := strings.TrimSpace(authorizationArray[0])

	userpass := strings.Split(authorization, ":")
	if len(userpass) != 2 {
		fmt.Println("Y")
		unauthorized(w)
		return
	}
	if database.CheckLogin(userpass[0], userpass[1]) {
		fmt.Fprint(w, "1")
	} else {
		unauthorized(w)
	}
}

func unauthorized(w http.ResponseWriter) {
	w.Header().Set("Authorization", "No")
	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}

func GetCandidate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idS := vars["id"]

	id, err := strconv.ParseInt(idS, 10, 64)
	fmt.Println(idS)
	if err != nil {
		fmt.Fprint(w, "Invalid Request")
		w.Header().Set("Status-Code", string(400))
		return
	}
	database.GetCandidateID(id, &w)
}
func GetCandidateName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	database.GetCandidateName(name, &w)
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idS := vars["id"]

	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		fmt.Fprint(w, "Invalid Request")
		w.Header().Set("Status-Code", string(400))
		return
	}

	database.GetCandidateStatus(id, &w)
}

func GetSkill(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	id, err := strconv.ParseInt(vars["cid"], 32, 8)
	if err != nil {
		fmt.Fprint(w, "Invalid Request")
		w.Header().Set("Status-Code", string(400))
		return
	}
	database.GetCandidateDSkill(id, &w)
}

func GetCandidateRegistration(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 32, 8)
	if err != nil {
		fmt.Fprint(w, "Invalid Request")
		w.Header().Set("Status-Code", string(400))
		return
	}
	database.GetCandidateRegistration(id, &w)
}

func GetTrainingCandidates(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idS := vars["training"]

	database.GetCandidateTraining(idS, &w)
}

func GetTraining(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idS := vars["id"]

	id, err := strconv.ParseInt(idS, 32, 8)
	if err != nil {
		fmt.Fprint(w, "Invalid Request")
		w.Header().Set("Status-Code", string(400))
		return
	}
	database.GetTraining(id, &w)
}

func UpdateQuestions(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	question := models.UploadQuestion{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&question)
	if err != nil {
		panic(err)
	}
	database.UploadQuestions(question, &writer)
}

func UploadForm(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	form := models.Form{}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&form)
	fmt.Println(form)
	if err != nil {
		panic(err)
	}
	database.UploadForm(form, &writer)
}

func GetFeedbackFromCentre(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	idS := vars["tcname"]

	database.GetFeedbackFromCentre(idS, &writer)
}

func GetTestFromCentre(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	idS := vars["tcname"]

	database.GetTestFromCentre(idS, &writer)
}

func GetAllTCentres(w http.ResponseWriter, r *http.Request) {

	database.GetAllTrainingCentre(&w)
}

func GetTCentres(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idS := vars["id"]

	id, err := strconv.ParseInt(idS, 32, 8)
	if err != nil {
		fmt.Fprint(w, "Invalid Request")
		w.Header().Set("Status-Code", string(400))
		return
	}
	database.GetTrainingCentre(id, &w)
}

func SubmitQuery(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	tag := r.Form.Get("tag")
	query := r.Form.Get("query")
	timestamp := r.Form.Get("timestamp")
	idS := r.Form.Get("id")

	id, err := strconv.ParseInt(idS, 32, 8)
	if err != nil {
		fmt.Fprint(w, "Invalid Request")
		w.Header().Set("Status-Code", string(400))
		return
	}

	database.UploadQuery(id, query, tag, timestamp, &w)
}

func GetIndiaStats(writer http.ResponseWriter, request *http.Request) {

	database.GetIndiaStats(&writer)
}

func GetInProgress(writer http.ResponseWriter, request *http.Request) {

	database.GetInProgress(&writer)
}

func GetCanInProgress(writer http.ResponseWriter, request *http.Request) {
	database.GetCanInProgress(&writer)
}

func GetQueries(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idS := vars["id"]

	database.GetQuery(idS, &w)
}

func AddNotif(w http.ResponseWriter, r *http.Request) {
	notif := models.Notif{}
	decoder := json.NewDecoder(r.Body)

	decoder.Decode(&notif)
	database.AddNotif(notif, &w)
}

func GetNotifs(w http.ResponseWriter, r *http.Request) {
	database.GetNotifs(&w)
}

func GetNotifsSig(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	date, err := time.Parse("Fri Mar 31 20:03:16 GMT 2017", vars["lastPull"])
	if err != nil {
		fmt.Fprint(w, "Invalid Request")
		w.Header().Set("Status-Code", string(400))
		return
	}
	if database.CheckNotif(date) {
		fmt.Fprint(w, "1")
	} else {
		fmt.Fprint(w, "0")
	}
}

func GetTotalStats(writer http.ResponseWriter, request *http.Request) {
	database.GetTotalStats(&writer)
}
