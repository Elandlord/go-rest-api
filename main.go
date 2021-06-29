package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"mentechmedia.nl/config"
	"mentechmedia.nl/model"
)

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint Hit: homePage")

	writer.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprint(writer, "<h1> Dit is een test! <h1>")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	registerRoutes(myRouter)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func registerRoutes(router *mux.Router) {
	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", model.AllArticles)
	router.HandleFunc("/article", model.StoreArticle).Methods("POST")
	router.HandleFunc("/articles/{id}", model.UpdateArticle).Methods("PUT")
	router.HandleFunc("/articles/{id}", model.DeleteArticle).Methods("DELETE")
	router.HandleFunc("/articles/{id}", model.FindArticle)
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	configFile := config.GetConfig()

	config.DbConnect(configFile)

	handleRequests()
}
