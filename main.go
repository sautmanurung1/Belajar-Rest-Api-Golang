package main

import (
	"REST-api/config"
	"REST-api/routes"
)

func main() {
	config.InitDBBook()
	config.InitDBUser()

	e := routes.New()
	e.Logger.Fatal(e.Start(":1234"))
}