package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"mentechmedia.nl/config"
	"mentechmedia.nl/model"
)

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

		var article model.Article
		err := articles.Scan(&article.Id, &article.Title, &article.Description, &article.Content)

		if err != nil {
			log.Fatal(err)
		}

		model.Articles = append(model.Articles, article)
	}

	json.NewEncoder(writer).Encode(model.Articles)
}

func StoreArticle(writer http.ResponseWriter, request *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	requestBody, _ := ioutil.ReadAll(request.Body)

	var article model.Article
	json.Unmarshal(requestBody, &article)
	// update our global Articles array to include
	// our new Article
	model.Articles = append(model.Articles, article)

	json.NewEncoder(writer).Encode(article)
}

func FindArticle(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["id"]

	// Loop over all Articles and return if the key matches the Article
	for _, article := range model.Articles {
		if article.Id == key {
			json.NewEncoder(writer).Encode(article)
		}
	}
}

func UpdateArticle(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	requestBody, _ := ioutil.ReadAll(request.Body)

	var updatedArticle model.Article
	json.Unmarshal(requestBody, &updatedArticle)

	for index, article := range model.Articles {
		if article.Id == id {
			article.Title = updatedArticle.Title
			article.Description = updatedArticle.Description
			article.Content = updatedArticle.Content
			model.Articles[index] = article
			json.NewEncoder(writer).Encode(article)
		}
	}
}

func DeleteArticle(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	// Loop over all Articles and return if the key matches the Article
	for index, article := range model.Articles {
		if article.Id == id {
			model.Articles = append(model.Articles[:index], model.Articles[index+1:]...)
		}
	}
}
