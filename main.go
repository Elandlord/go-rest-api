package main

import (
	"fmt"

	"mentechmedia.nl/app"
)

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers with Docker")

	app := &app.App{}
	app.Initialize()
	app.Run(":8081")
}
