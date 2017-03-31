package database

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/vyasgiridhar/server/models"
)

func createConn() *sql.DB {
	db, err := sql.Open("mysql", "root:babababa@/ddugky")
	if err != nil {
		return nil
	}
	return db
}

func CheckLogin(username, password string) bool {
	db := createConn()
	if db == nil {
		return false
	}
	i := 0

	db.QueryRow("select count(*) from Users where username = ? and password = ?", username, password).Scan(&i)
	if i != 0 {
		return true
	}
	return false
}

func Signup(body models.SignupBody) bool {
	db := createConn()
	if db == nil {
		return false
	}
	_, err := db.Exec("insert into Users values(?,?,?,?,?,?,?,?,?,?,?,?,?)", body.Username, body.Password, body.Name, body.Address, body.DOB, body.Mail,
		body.MobileNo, body.PAN, body.AdhaarNO, body.City, body.Gname, body.District, body.Etype)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func GetCitizenID(id int64, w *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*w).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	rows, err := db.Query("select * from Citizen where AdhaarID = ?", id)
	if err != nil {
		(*w).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	makeStructJSON(rows, w)
	return
}

func GetCitizenName(Name string, w *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*w).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	rows, err := db.Query("select * from Citizen where Name = ?", Name)
	if err != nil {
		(*w).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	makeStructJSON(rows, w)
	return
}

func GetTrainingForJob(i int64, writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	rows, err := db.Query("")
	if err != nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	makeStructJSON(rows, writer)
}

func GetTrainingCentre(i int64, writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	rows, err := db.Query("")
	if err != nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	makeStructJSON(rows, writer)
}

func UploadQuery(i int64, i2 string, i3 string, i4 string, writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	rows, err := db.Query("")
	if err != nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	makeStructJSON(rows, writer)
}

func GetQuery(writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	rows, err := db.Query("")
	if err != nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	makeStructJSON(rows, writer)
}

func GetCentreTraining(i int64, writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	rows, err := db.Query("")
	if err != nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	makeStructJSON(rows, writer)
}

func GetJobsTraining(i int64, writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	rows, err := db.Query("")
	if err != nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	makeStructJSON(rows, writer)
}

func GetIndustryJobs(i string, writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	rows, err := db.Query("")
	if err != nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	makeStructJSON(rows, writer)
}

func GetCitizenDSkill(i int64, writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	rows, err := db.Query("")
	if err != nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	makeStructJSON(rows, writer)
}

func GetCitizenTrainingSkill(i int64, writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	rows, err := db.Query("")
	if err != nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	makeStructJSON(rows, writer)
}

func GetTraining(i int64, writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	rows, err := db.Query("")
	if err != nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	makeStructJSON(rows, writer)
}

func GetIndustry(name string, writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	rows, err := db.Query("")
	if err != nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	makeStructJSON(rows, writer)
}

func AddNotif(notif models.Notif, writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	rows, err := db.Query("")
	if err != nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	makeStructJSON(rows, writer)
}

func GetNotifs(writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	rows, err := db.Query("")
	if err != nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return
	}
	makeStructJSON(rows, writer)
}

func CheckNotif() bool {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return false
	}
	rows, err := db.Query("")
	if err != nil {
		(*writer).Header().Set("Status-Code", string(http.StatusBadRequest))
		return false
	}
	if rows.Scan(&i) > 0 {
		return true
	}

	return false
}
