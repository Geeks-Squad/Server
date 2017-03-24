package server

type LoginBody struct {
	username string `json:"username"`
	password string `json:"password"`
}
