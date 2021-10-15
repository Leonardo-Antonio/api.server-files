package middleware

import (
	"net/http"
	"os"

	"github.com/Leonardo-Antonio/api.server-files/src/helper/response"
	"github.com/gofiber/fiber/v2"
)

func Files(c *fiber.Ctx) error {
	apiKey := c.Query("key", "none")
	if apiKey == "none" {
		return c.Status(http.StatusForbidden).
			JSON(response.Err("ingrese un apikey", nil))
	}

	if apiKey != os.Getenv("KEY") {
		return c.Status(http.StatusUnauthorized).
			JSON(response.Err("no tiene acceso a subir una imagen", nil))
	}

	return c.Next()
}
