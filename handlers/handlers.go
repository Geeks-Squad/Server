package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vyasgiridhar/server/database"
	"github.com/vyasgiridhar/server/models"
	"strings"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	signup := models.SignupBody{}
	decoder := json.NewDecoder(r.Body)

	decoder.Decode(&signup)
	database.Signup(signup)
}

func GetCitizen(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idS := vars["id"]

	id, err := strconv.ParseInt(idS, 32, 8)
	if err != nil {
		fmt.Fprint(w, "Invalid Request")
		w.Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	database.GetCitizenID(id, &w)

}
func GetCitizenName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	database.GetCitizenName(name, &w)
}

func GetSkill(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["cid"], 32, 8)
	if err != nil {
		fmt.Fprint(w, "Invalid Request")
		w.Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	database.GetCitizenDSkill(id, &w)
}

func GetCitizenRegistration(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 32, 8)
	if err != nil {
		fmt.Fprint(w, "Invalid Request")
		w.Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	database.GetCitizenDSkill(id, &w)
}

func GetTrainingCandidates(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idS := vars["training"]

	id, err := strconv.ParseInt(idS, 32, 8)
	if err != nil {
		fmt.Fprint(w, "Invalid Request")
		w.Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	database.GetCitizenTrainingSkill(id, &w)
}

func GetTraining(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idS := vars["id"]

	id, err := strconv.ParseInt(idS, 32, 8)
	if err != nil {
		fmt.Fprint(w, "Invalid Request")
		w.Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	database.GetTraining(id, &w)
}

func GetCentreTraining(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idS := vars["id"]

	id, err := strconv.ParseInt(idS, 32, 8)
	if err != nil {
		fmt.Fprint(w, "Invalid Request")
		w.Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	database.GetCentreTraining(id, &w)
}

func GetJobsTraining(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idS := vars["trainingid"]

	id, err := strconv.ParseInt(idS, 32, 8)
	if err != nil {
		fmt.Fprint(w, "Invalid Request")
		w.Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	database.GetJobsTraining(id, &w)
}

func GetTrainingForJob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idS := vars["jobid"]

	id, err := strconv.ParseInt(idS, 32, 8)
	if err != nil {
		fmt.Fprint(w, "Invalid Request")
		w.Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	database.GetTrainingForJob(id, &w)
}

func GetIndustry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	database.GetIndustry(name, &w)
}

func GetIndustryJobs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	industry := vars["industry"]

	database.GetIndustryJobs(industry, &w)
}

func GetTCentres(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idS := vars["id"]

	id, err := strconv.ParseInt(idS, 32, 8)
	if err != nil {
		fmt.Fprint(w, "Invalid Request")
		w.Header().Set("Status-Code", string(http.StatusBadRequest))
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
		w.Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}

	database.UploadQuery(id,query,tag,timestamp,&w)
}

func GetQueries(w http.ResponseWriter, r *http.Request) {
	database.GetQuery(&w)
}

func AddNotif(w http.ResponseWriter, r *http.Request) {
	notif := models.Notif{}
	decoder := json.Decoder(r.Body)

	decoder.Decode(&notif)
	database.AddNotif(notif,&w)
}

func GetNotifs(w http.ResponseWriter, r *http.Request) {
	database.GetNotifs(&w)
}