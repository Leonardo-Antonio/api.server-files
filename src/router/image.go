package router

import (
	"github.com/Leonardo-Antonio/api.server-files/src/handler"
	"github.com/Leonardo-Antonio/api.server-files/src/middleware"
	"github.com/gofiber/fiber/v2"
)

func Image(app *fiber.App) {
	image := new(handler.Image)
	group := app.Group("/api/v1/images")
	group.Static("", "./static/images")
	group.Post("", middleware.Files, image.Upload)
}
