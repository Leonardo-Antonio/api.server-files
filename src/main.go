package main

import (
	"github.com/Leonardo-Antonio/api.server-files/src/app"
)

func main() {
	app := app.New(app.PROD)
	app.Config()
	app.Middlewares()
	app.Routers()
	app.Listening()
}
