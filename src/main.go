package main

import (
	"log"

	"github.com/Leonardo-Antonio/api.server-files/src/app"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	app := app.New()
	app.Middlewares()
	app.Routers()
	app.Listening()
}
