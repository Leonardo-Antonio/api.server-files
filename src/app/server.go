package app

import (
	"log"

	"github.com/Leonardo-Antonio/api.server-files/src/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type app struct {
	app *fiber.App
}

func New() *app {
	return &app{fiber.New()}
}

func (a *app) Middlewares() {
	a.app.Use(cors.New())
	a.app.Use(logger.New())
}

func (a *app) Routers() {
	router.Image(a.app)
	router.Pdfs(a.app)
}

func (a *app) Listening() {
	if err := a.app.Listen(":8080"); err != nil {
		log.Fatalln(err)
	}
}
