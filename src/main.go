package main

import (
	"github.com/Leonardo-Antonio/api.server-files/src/app"
)

func main() {
	app := app.New(app.DEV)
	app.Config()
	app.Middlewares()
	app.Routers()
	app.Listening()
}
