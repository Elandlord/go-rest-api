package model

type Article struct {
	Id          string `json:"Id"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Content     string `json:"Content"`
}

var Articles []Article
