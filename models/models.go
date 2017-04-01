package models

type Notif struct {
	Content string `json:"content"`
}

type Form struct {
	Aid  int64      `json:"aadharid"`
	Data []FormData `json:"data"`
}

type FormData struct {
	Question int    `json:"question"`
	Answer   string `json:"answer"`
}

type UploadQuestion struct {
	Tcid int      `json:"tcid"`
	Data FormData `json:"qa"`
}

type Qlist struct {
	list []int `json:"list"`
}
