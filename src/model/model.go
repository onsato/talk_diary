package model

type Users struct {
	Uid  int    `json:"Uid"`
	Name string `json:"name"`
}

type Dairies struct {
	Did   int    `json:"Did"`
	Date  int    `json:"date"`
	Title string `json:"title"`
	Text  string `json:"text"`
}
