package handlers

import (
	"encoding/base64"
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

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	signup := models.SignupBody{}
	decoder := json.NewDecoder(r.Body)

	decoder.Decode(&signup)
	database.Signup(signup)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	authorizationArray := r.Header["Authorization"]

	if len(authorizationArray) > 0 {
		authorization := strings.TrimSpace(authorizationArray[0])
		credentials := strings.Split(authorization, " ")

		if len(credentials) != 2 || credentials[0] != "Basic" {
			unauthorized(w)
			return
		}

		authstr, err := base64.StdEncoding.DecodeString(credentials[1])
		if err != nil {
			unauthorized(w)
			return
		}

		userpass := strings.Split(string(authstr), ":")
		if len(userpass) != 2 {
			unauthorized(w)
			return
		}

		if database.CheckLogin(userpass[0], userpass[1]) {
			fmt.Fprint(w, "1")
		} else {
			unauthorized(w)
		}
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

	id, err := strconv.ParseInt(idS, 32, 8)
	fmt.Println(id)
	if err != nil {
		fmt.Fprint(w, "Invalid Request")
		w.Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	database.GetCandidateID(id, &w)
}
func GetCandidateName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	database.GetCandidateName(name, &w)
}

func GetSkill(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	id, err := strconv.ParseInt(vars["cid"], 32, 8)
	if err != nil {
		fmt.Fprint(w, "Invalid Request")
		w.Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	database.GetCandidateDSkill(id, &w)
}

func GetCandidateRegistration(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 32, 8)
	if err != nil {
		fmt.Fprint(w, "Invalid Request")
		w.Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	database.GetCandidateDSkill(id, &w)
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
	database.GetCandidateTrainingSkill(id, &w)
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

	database.UploadQuery(id, query, tag, timestamp, &w)
}

func GetQueries(w http.ResponseWriter, r *http.Request) {
	database.GetQuery(&w)
}

func AddNotif(w http.ResponseWriter, r *http.Request) {
	notif := models.Notif{}
	//decoder := json.Decoder(r.Body)

	//decoder.Decode(&notif)
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
		w.Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	if database.CheckNotif(date) {
		fmt.Fprint(w, "1")
	} else {
		fmt.Fprint(w, "0")
	}
}
