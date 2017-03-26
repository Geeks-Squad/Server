package database

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/vyasgiridhar/server/models"
)

func createConn() *sql.DB {
	db, err := sql.Open("mysql", "root:spd@/ddugky")
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

func GetCitizenDID(id int, w *http.ResponseWriter) bool {
	db := createConn()
	if db == nil {
		return false
	}
	rows, err := db.Query("select * from Citizen where AdhaarID = ?", id)
	if err != nil {
		return false
	}
	makeStructJSON(rows, w)
	return true
}

func GetCitizenDName(Name string, w *http.ResponseWriter) bool {
	db := createConn()
	if db == nil {
		return false
	}
	rows, err := db.Query("select * from Citizen where Name = ?", Name)
	if err != nil {
		return false
	}
	makeStructJSON(rows, w)
	return true
}

/*
func GetCitizenD(Name string, w *http.ResponseWriter) bool {
	db := createConn()
	if db == nil {
		return false
	}
	rows, err := db.Query("select * from Citizen where Name = ?", Name)
	makeStructJSON(rows, w)
	return true
}
*/
