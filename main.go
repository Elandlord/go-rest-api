package main

import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
)

type Article struct {
	Id string `json:"Id"`
	Title string `json:"Title"`
	Description string `json:"Description"`
	Content string `json:"Content"`
}

var Articles []Article

func allArticles(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint Hit: All Articles Endpoint")
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

func homePage(writer http.ResponseWriter, request *http.Request){
    fmt.Fprintf(writer, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/articles", allArticles)
	myRouter.HandleFunc("/article", storeArticle).Methods("POST")
	myRouter.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/articles/{id}", findArticle)

    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

    Articles = []Article{
        Article{Id: "1", Title: "Hello", Description: "Article Description", Content: "Article Content"},
        Article{Id: "2", Title: "Hello 2", Description: "Article Description", Content: "Article Content"},
    }

    handleRequests()
}