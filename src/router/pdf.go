package router

import (
	"github.com/Leonardo-Antonio/api.server-files/src/handler"
	"github.com/Leonardo-Antonio/api.server-files/src/middleware"
	"github.com/gofiber/fiber/v2"
)

func Pdfs(app *fiber.App) {
	pdf := new(handler.Pdf)
	group := app.Group("/api/v1/pdfs", middleware.ValidHostname)
	group.Static("", "./static/pdfs")
	group.Post("", middleware.Files, pdf.Upload)
}
