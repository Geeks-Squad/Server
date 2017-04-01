package models

import "time"

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Notif struct {
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

type Form struct {
	Aid  int64      `json:"aadharid"`
	Data []FormData `json:"data"`
}

type FormData struct {
	Question int `json:"question"`
	Answer   string `json:"answer"`
}
