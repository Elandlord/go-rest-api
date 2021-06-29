package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"mentechmedia.nl/config"
)

type Article struct {
	Id          string `json:"Id"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Content     string `json:"Content"`
}

var Articles []Article

func AllArticles(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint Hit: All Articles Endpoint")

	// TODO: Replace with an app structure that holds the DB connection
	configFile := config.GetConfig()
	dbConnection := config.DbConnect(configFile)

	articles, err := dbConnection.Query("SELECT * FROM articles")

	defer dbConnection.Close()

	if err != nil {
		log.Fatal(err)
	}

	for articles.Next() {

		var article Article
		err := articles.Scan(&article.Id, &article.Title, &article.Description, &article.Content)

		if err != nil {
			log.Fatal(err)
		}

		Articles = append(Articles, article)
	}

	json.NewEncoder(writer).Encode(Articles)
}

func StoreArticle(writer http.ResponseWriter, request *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	requestBody, _ := ioutil.ReadAll(request.Body)

	var article Article
	json.Unmarshal(requestBody, &article)
	// update our global Articles array to include
	// our new Article
	Articles = append(Articles, article)

	json.NewEncoder(writer).Encode(article)
}

func FindArticle(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["id"]

	// Loop over all Articles and return if the key matches the Article
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(writer).Encode(article)
		}
	}
}

func UpdateArticle(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	requestBody, _ := ioutil.ReadAll(request.Body)

	var updatedArticle Article
	json.Unmarshal(requestBody, &updatedArticle)

	for index, article := range Articles {
		if article.Id == id {
			article.Title = updatedArticle.Title
			article.Description = updatedArticle.Description
			article.Content = updatedArticle.Content
			Articles[index] = article
			json.NewEncoder(writer).Encode(article)
		}
	}
}

func DeleteArticle(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	// Loop over all Articles and return if the key matches the Article
	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}
