package database

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/vyasgiridhar/server/models"
)

func createConn() *sql.DB {
	db, err := sql.Open("mysql", "smh2017:smh2017bro@/SMH")
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

	db.QueryRow("select count(*) from candidate where username = ? and password = ?", username, password).Scan(&i)
	if i != 0 {
		return true
	}
	return false
}

func GetCandidateID(id int64, w *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*w).Header().Set("Status-Code", string(400))
		fmt.Fprint((*w), "POOP")
		return
	}
	rows, err := db.Query("select * from candidate where AadharNo = ?", id)
	if err != nil {
		(*w).Header().Set("Status-Code", string(400))
		fmt.Fprint((*w), err.Error())
		return
	}
	makeStructJSON(rows, w)
	return
}

func GetCandidateName(Name string, w *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*w).Header().Set("Status-Code", string(400))
		return
	}
	rows, err := db.Query("select * from candidate where Name = ?", Name)
	if err != nil {
		(*w).Header().Set("Status-Code", string(400))
		return
	}
	makeStructJSON(rows, w)
	return
}

func GetIndiaStats(w *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*w).Header().Set("Status-Code", string(400))
		return
	}
	rows, err := db.Query("select state,count(*) from candidate group by state")
	if err != nil {
		(*w).Header().Set("Status-Code", string(400))
		return
	}
	makeStructJSON(rows, w)
	return
}

func GetCanInProgress(writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	rows, err := db.Query("select count(*) from candidate where status = \"under_training\"")
	if err != nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	makeStructJSON(rows, writer)
	return
}

func GetInProgress(w *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*w).Header().Set("Status-Code", string(400))
		return
	}
	rows, err := db.Query("select count(*) from training where status = \"ongoing\"")
	if err != nil {
		(*w).Header().Set("Status-Code", string(400))
		return
	}
	makeStructJSON(rows, w)
	return
}

func GetTrainingCentre(i int64, writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	rows, err := db.Query("SELECT * from trainingcentre where tid = ?",i)
	if err != nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	makeStructJSON(rows, writer)
}

func UploadQuery(id int64, query, tag, timestamp string, writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	_, err := db.Exec("INSERT INTO queries VALUE (?,?,?,?,?,?)", id, query, tag, timestamp, nil,nil)
	if err != nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	return
}

func GetQuery(tag string, writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	rows, err := db.Query("SELECT * FROM queries WHERE tag like ?", tag)
	if err != nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	makeStructJSON(rows, writer)
}

func GetCandidateStatus(i int64, writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	rows, err := db.Query("SELECT status FROM T_T_C where id = (select tid from candidate where adhaarno = ?)",i)
	if err != nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	makeStructJSON(rows, writer)
}

func UploadForm(form models.Form, w *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*w).Header().Set("Status-Code", string(400))
		return
	}
	for i := 0; i < len(form.Data); i += 1 {
		_, err := db.Exec("INSERT INTO feedback VALUES (?,?,?)", form.Aid, form.Data[i].Question, form.Data[i].Answer)
		if err != nil {
			(*w).Header().Set("Status-Code", string(400))
			return
		}
	}
	(*w).Header().Set("Status-Code", string(200))
}


func GetFeedbackFromCentre(tname string, writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	rows, err := db.Query("select * from feedback NATURAL JOIN questions where tag = \"feedback\" and aadharno in (select aadharno from candidate where tid in (select tid from training where name like ?))",tname)
	if err != nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	makeStructJSON(rows, writer)
}

func GetTestFromCentre(tname string, writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	rows, err := db.Query("select * from feedback NATURAL JOIN questions where tag = \"test\" and aadharno in (select aadharno from candidate where tid in (select tid from training where name like ?))",tname)
	if err != nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	makeStructJSON(rows, writer)

}

func GetCandidateDSkill(i int64, writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	rows, err := db.Query("")
	if err != nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	makeStructJSON(rows, writer)
}

func GetCandidateRegistration(i int64, writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	rows, err := db.Query("")
	if err != nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	makeStructJSON(rows, writer)
}


func GetCandidateTraining(name string, writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	rows, err := db.Query("select aadharno from candidate where tid in (select tid from training where name like ?)",name)
	if err != nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	makeStructJSON(rows, writer)
}

func GetTraining(i int64, writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	rows, err := db.Query("SELECT * FROM training where id = ?",i)
	if err != nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	makeStructJSON(rows, writer)
}

func AddNotif(notif models.Notif, writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	_, err := db.Exec("insert into notif values(?,?)",notif.Content,notif.Timestamp)
	if err != nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	(*writer).Header().Set("Status-Code", string(300))
}

func GetNotifs(writer *http.ResponseWriter) {
	db := createConn()
	if db == nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	rows, err := db.Query("SELECT content from notif")
	if err != nil {
		(*writer).Header().Set("Status-Code", string(400))
		return
	}
	makeStructJSON(rows, writer)
}

func CheckNotif(date time.Time) bool {
	i := 0
	db := createConn()
	if db == nil {
		return false
	}
	rows, err := db.Query("")
	if err != nil {
		return false
	}
	if rows.Scan(&i) != nil {
		return false
	}
	if i > 0 {
		return true
	}

	return false
}
