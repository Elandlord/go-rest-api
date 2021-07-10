package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"mentechmedia.nl/rest-api/app/handler"
	"mentechmedia.nl/rest-api/auth"
	"mentechmedia.nl/rest-api/config"
)

var mySigningKey = []byte("ericgorestapi")

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("x-access-token")

		json.NewEncoder(w).Encode(r)
		header = strings.TrimSpace(header)

		if header == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Missing auth token")
			return
		} else {
			json.NewEncoder(w).Encode(fmt.Sprintf("Token found. Value %s", header))
		}
		next.ServeHTTP(w, r)
	})
}

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) Initialize() {

	configFile := config.GetConfig()
	app.DB = config.DbConnect(configFile)
	router := mux.NewRouter().StrictSlash(true)

	app.Router = router
	// Register middleware on every request
	secure := router.PathPrefix("/auth").Subrouter()
	secure.Use(auth.JwtVerify)

	app.registerRoutes()
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint Hit: homePage")

	writer.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprint(writer, "<h1> Dit is een test! <h1>")
}

func (app *App) registerRoutes() {
	app.Get("/", homePage)
	app.Get("/articles", app.handleRequest(handler.AllArticles))
	app.Post("/articles", app.handleRequest(handler.StoreArticle))
	app.Put("/articles/{id}", app.handleRequest(handler.UpdateArticle))
	app.Delete("/articles/{id}", app.handleRequest(handler.DeleteArticle))
	app.Get("/articles/{id}", app.handleRequest(handler.FindArticle))
}

// Get wraps the router for GET method
func (app *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (app *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (app *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (app *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (app *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, app.Router))
}

type RequestHandlerFunction func(db *sql.DB, w http.ResponseWriter, r *http.Request)

func (app *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		handler(app.DB, writer, request)
	}
}
