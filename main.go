package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

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

func allArticles(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint Hit: All Articles Endpoint")

	// TODO: Replace with an app structure that holds the DB connection
	config := config.GetConfig()
	dbConnection := dbConnect(config)

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

func storeArticle(writer http.ResponseWriter, request *http.Request) {
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

func findArticle(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["id"]

	// Loop over all Articles and return if the key matches the Article
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(writer).Encode(article)
		}
	}
}

func updateArticle(writer http.ResponseWriter, request *http.Request) {
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

func deleteArticle(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	// Loop over all Articles and return if the key matches the Article
	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint Hit: homePage")

	writer.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprint(writer, "<h1> Dit is een test! <h1>")
}

func dbConnect(config *config.Config) *sql.DB {
	dbConnection, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", config.DB.Username, config.DB.Password, config.DB.Name))

	if err != nil {
		panic(err)
	}

	dbConnection.SetConnMaxLifetime(time.Minute * 3)
	dbConnection.SetMaxOpenConns(10)
	dbConnection.SetMaxIdleConns(10)
	return dbConnection
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	registerRoutes(myRouter)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func registerRoutes(router *mux.Router) {
	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", allArticles)
	router.HandleFunc("/article", storeArticle).Methods("POST")
	router.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	router.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	router.HandleFunc("/articles/{id}", findArticle)
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	config := config.GetConfig()

	dbConnect(config)

	handleRequests()
}
