package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vyasgiridhar/server/database"
	"github.com/vyasgiridhar/server/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Authorized")
	w.Header().Add("Status-Code", string(http.StatusOK))
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	signup := models.SignupBody{}
	decoder := json.NewDecoder(r.Body)

	decoder.Decode(&signup)
	database.Signup(signup)
}

func GetCitizen(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idS := vars["ID"]

	id, err := strconv.ParseInt(idS, 32, 8)
	if err != nil {
		fmt.Fprint(w, "Invalid Request")
		w.Header().Set("Status-Code", string(http.StatusBadRequest))
	}
	database.GetCitizenDID(int(id), &w)

}
