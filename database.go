package main

import (
	"database/sql"
	"fmt"
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
	row := db.QueryRow("select count(*) from Users where username = ? and password = ?", username, password)
	if row != nil {
		return true
	}
	return false
}

func Signup(body SignupBody) bool {
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
