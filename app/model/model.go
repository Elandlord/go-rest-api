package model

import (
	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Id          string `json:"Id"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Content     string `json:"Content"`
}

var Articles []Article
