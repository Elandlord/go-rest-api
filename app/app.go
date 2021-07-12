package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"mentechmedia.nl/rest-api/app/handler"
	"mentechmedia.nl/rest-api/auth"
	"mentechmedia.nl/rest-api/config"
)

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
	app.registerRoutes()
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint Hit: homePage")

	http.ServeFile(writer, request, request.URL.Path+"go/rest-api/static/index.html")
}

func (app *App) registerRoutes() {
	app.Get("/", homePage)
	app.Post("/authenticate", auth.Authenticate)
	app.GetWithMiddleware("/articles", app.handleRequest(handler.AllArticles))
	app.PostWithMiddleware("/articles", app.handleRequest(handler.StoreArticle))
	app.PutWithMiddleware("/articles/{id}", app.handleRequest(handler.UpdateArticle))
	app.DeleteWithMiddleware("/articles/{id}", app.handleRequest(handler.DeleteArticle))
	app.GetWithMiddleware("/articles/{id}", app.handleRequest(handler.FindArticle))
}

// Get wraps the router for GET method, add AuthMiddleware
func (app *App) GetWithMiddleware(path string, f http.HandlerFunc) {
	app.Router.Handle(path, auth.AuthMiddleware(f)).Methods("GET")
}

// Get wraps the router for POST method, add AuthMiddleware
func (app *App) PostWithMiddleware(path string, f http.HandlerFunc) {
	app.Router.Handle(path, auth.AuthMiddleware(f)).Methods("POST")
}

// Put wraps the router for PUT method, add AuthMiddleware
func (app *App) PutWithMiddleware(path string, f http.HandlerFunc) {
	app.Router.Handle(path, auth.AuthMiddleware(f)).Methods("PUT")
}

// Delete wraps the router for DELETE method, add AuthMiddleware
func (app *App) DeleteWithMiddleware(path string, f http.HandlerFunc) {
	app.Router.Handle(path, auth.AuthMiddleware(f)).Methods("DELETE")
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
