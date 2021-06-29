package main

import (
	"fmt"

	"mentechmedia.nl/app"
)

// func handleRequests() {
// 	myRouter := mux.NewRouter().StrictSlash(true)

// 	registerRoutes(myRouter)

// 	log.Fatal(http.ListenAndServe(":10000", myRouter))
// }

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	app := &app.App{}
	app.Initialize()
	app.Run(":10000")
}
