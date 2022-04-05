package main

import (
	"REST-api/config"
	"REST-api/routes"
)

func main() {
	config.Connect()

	e := routes.New()
	e.Logger.Fatal(e.Start(":1234"))
}