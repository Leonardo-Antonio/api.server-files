package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/Leonardo-Antonio/api.server-files/src/helper/response"
	"github.com/gofiber/fiber/v2"
)

func ValidHostname(c *fiber.Ctx) error {
	log.Println(c.Context().RemoteIP().String())
	if os.Getenv("HOSTNAME") != c.Context().RemoteIP().String() {
		return c.Status(http.StatusUnauthorized).JSON(response.Err("no tiene autorización para ver este contenido", nil))
	}
	return c.Next()
}
