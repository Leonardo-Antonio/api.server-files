package handler

import (
	"log"
	"net/http"

	"github.com/Leonardo-Antonio/api.server-files/src/helper"
	"github.com/Leonardo-Antonio/api.server-files/src/helper/response"
	"github.com/gofiber/fiber/v2"
)

type Pdf struct{}

func (i *Pdf) Upload(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.Err(err.Error(), nil))
	}

	log.Println(fileHeader.Header.Get("Content-Type"))
	if fileHeader.Header.Get("Content-Type") != "application/pdf" {
		return c.Status(http.StatusBadRequest).JSON(response.Err("el formato ingresado no es correcto", nil))
	}

	fileName, err := helper.SaveFile(fileHeader, "static/pdfs/", "pdf")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.Err(err.Error(), nil))
	}

	return c.Status(http.StatusCreated).JSON(response.Successful("se guardo correctamente", map[string]string{
		"url": "http://" + c.Hostname() + string(c.Request().URI().Path()) + "/" + fileName,
	}))
}
