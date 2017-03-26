package models

import "time"

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupBody struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	Address  string    `json:"address"`
	DOB      time.Time `json:"dob"`
	Mail     string    `json:"mail"`
	MobileNo uint32    `json:"mobileno"`
	PAN      int       `json:"pan"`
	AdhaarNO uint32    `json:"adhaarno"`
	City     string    `json:"city"`
	Gname    string    `json:"gname"`
	District string    `json:"district"`
	Etype    int       `json:"etype"`
}
