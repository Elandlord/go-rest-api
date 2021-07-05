package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"mentechmedia.nl/model"
)

func AllArticles(db *sql.DB, writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint Hit: All Articles Endpoint")

	// Create SQL statements that sanitizes the input
	sqlStatement := `INSERT INTO articles (Content, Description, Title) VALUES (?, ?, ?)`
	_, storeError := db.Exec(sqlStatement, "Eric Landheer", "Has created a functioning API in Go", "But improvements have to be made")

	if storeError != nil {
		log.Fatal(storeError)
	}

	articles, err := db.Query("SELECT * FROM articles")

	if err != nil {
		log.Fatal(err)
	}

	model.Articles = nil

	for articles.Next() {

		var article model.Article
		err := articles.Scan(&article.Id, &article.Title, &article.Description, &article.Content)

		if err != nil {
			log.Fatal(err)
		}

		model.Articles = append(model.Articles, article)
	}

	respondJSON(writer, http.StatusOK, model.Articles)
}

func StoreArticle(db *sql.DB, writer http.ResponseWriter, request *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	requestBody, _ := ioutil.ReadAll(request.Body)

	var article model.Article
	json.Unmarshal(requestBody, &article)

	// Create SQL statements that sanitizes the input
	sqlStatement := `INSERT INTO articles (Content, Description, Title) VALUES (?, ?, ?)`
	_, err := db.Exec(sqlStatement, "Eric Landheer", "Has created a functioning API in Go", "But improvements have to be made")

	if err != nil {
		log.Fatal(err)
	}

	respondJSON(writer, http.StatusCreated, article)
}

func FindArticle(db *sql.DB, writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint Hit: Find Article Endpoint")

	vars := mux.Vars(request)
	key := vars["id"]

	article := getArticleOr404(db, key, writer, request)

	if article == nil {
		return
	}

	respondJSON(writer, http.StatusOK, article)
}

func UpdateArticle(db *sql.DB, writer http.ResponseWriter, request *http.Request) {
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

func DeleteArticle(db *sql.DB, writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	// Loop over all Articles and return if the key matches the Article
	for index, article := range model.Articles {
		if article.Id == id {
			model.Articles = append(model.Articles[:index], model.Articles[index+1:]...)
		}
	}
}

func getArticleOr404(db *sql.DB, id string, w http.ResponseWriter, r *http.Request) *model.Article {
	// TODO: Input needs to be sanitized
	sqlStatement := `"SELECT * FROM articles WHERE Id = ?`

	foundArticle, err := db.Query(sqlStatement, id)

	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}

	article := model.Article{}

	for foundArticle.Next() {
		err := foundArticle.Scan(&article.Id, &article.Content, &article.Description, &article.Title)

		if err != nil {
			log.Fatal(err)
		}
	}

	return &article
}
