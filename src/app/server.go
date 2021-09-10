package app

import (
	"log"
	"os"

	"github.com/Leonardo-Antonio/api.server-files/src/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

type app struct {
	app   *fiber.App
	state string
}

func New(state string) *app {
	return &app{fiber.New(), state}
}

func (a *app) Config() {
	CreateDirStaticIsNotExist()
	if a.state == DEV {
		if err := godotenv.Load(); err != nil {
			log.Fatalln(err)
		}
	}
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
	if err := a.app.Listen(":" + os.Getenv("PORT")); err != nil {
		log.Fatalln(err)
	}
}
