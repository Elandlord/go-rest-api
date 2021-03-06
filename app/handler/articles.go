package handler

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"mentechmedia.nl/rest-api/app/model"
)

type articleNotFoundError struct {
	message string
}

func (error *articleNotFoundError) Error() string {
	return error.message
}

func AllArticles(db *sql.DB, writer http.ResponseWriter, request *http.Request) {
	articles := getAllArticles(db)

	RespondJSON(writer, http.StatusOK, articles)
}

func StoreArticle(db *sql.DB, writer http.ResponseWriter, request *http.Request) {
	requestBody, _ := ioutil.ReadAll(request.Body)

	var article model.Article
	json.Unmarshal(requestBody, &article)

	sqlStatement := "INSERT INTO articles (Content, Description, Title) VALUES (?, ?, ?)"
	_, err := db.Exec(sqlStatement, article.Content, article.Description, article.Title)

	if err != nil {
		RespondError(writer, http.StatusNotFound, err.Error())
		return
	}

	RespondJSON(writer, http.StatusCreated, article)
}

func FindArticle(db *sql.DB, writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["id"]

	article, err := getArticleOr404(db, key, writer, request)

	if err != nil {
		RespondError(writer, http.StatusNotFound, err.Error())
		return
	}

	if article == nil {
		RespondError(writer, http.StatusNotFound, err.Error())
		return
	}

	RespondJSON(writer, http.StatusOK, article)
}

func UpdateArticle(db *sql.DB, writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	requestBody, _ := ioutil.ReadAll(request.Body)

	var updatedArticle model.Article
	json.Unmarshal(requestBody, &updatedArticle)

	articleToUpdate, err := getArticleOr404(db, id, writer, request)

	if err != nil {
		RespondError(writer, http.StatusNotFound, err.Error())
		return
	}

	articleToUpdate.Title = updatedArticle.Title
	articleToUpdate.Description = updatedArticle.Description
	articleToUpdate.Content = updatedArticle.Content

	sqlStatement := "UPDATE articles SET Content = ?, Description = ?, Title = ? WHERE Id = ?"
	_, err = db.Exec(sqlStatement, articleToUpdate.Content, articleToUpdate.Description, articleToUpdate.Title, articleToUpdate.Id)

	if err != nil {
		RespondError(writer, http.StatusNotFound, err.Error())
		return
	}

	json.NewEncoder(writer).Encode(updatedArticle)
}

func DeleteArticle(db *sql.DB, writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	articleToDelete, err := getArticleOr404(db, id, writer, request)

	if err != nil {
		RespondError(writer, http.StatusNotFound, err.Error())
		return
	}

	sqlStatement := "DELETE FROM articles WHERE Id = ?"
	_, err = db.Exec(sqlStatement, articleToDelete.Id)
}

func getAllArticles(db *sql.DB) []model.Article {
	rows, err := db.Query("SELECT * FROM articles")

	if err != nil {
		log.Fatal(err)
	}

	var articles []model.Article

	for rows.Next() {

		var article model.Article
		err := rows.Scan(&article.Id, &article.Title, &article.Description, &article.Content)

		if err != nil {
			log.Fatal(err)
		}

		articles = append(articles, article)
	}

	return articles
}

func getArticleOr404(db *sql.DB, id string, w http.ResponseWriter, r *http.Request) (*model.Article, error) {
	sqlStatement := "SELECT * FROM articles WHERE Id = ?"

	foundArticle, err := db.Query(sqlStatement, id)

	if err != nil {
		return nil, err
	}

	article := model.Article{}

	for foundArticle.Next() {
		err := foundArticle.Scan(&article.Id, &article.Content, &article.Description, &article.Title)

		if err != nil {
			log.Fatal(err)
		}
	}

	if article.Id == "" {
		return nil, &articleNotFoundError{"Article not found."}
	}

	return &article, nil
}
