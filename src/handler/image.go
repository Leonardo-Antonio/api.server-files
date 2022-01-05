package handler

import (
	"net/http"

	"github.com/Leonardo-Antonio/api.server-files/src/helper"
	"github.com/Leonardo-Antonio/api.server-files/src/helper/response"
	"github.com/gofiber/fiber/v2"
)

type Image struct{}

func (i *Image) Upload(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.Err(err.Error(), nil))
	}

/* 	if fileHeader.Header.Get("Content-Type") != "image/png" {
		return c.Status(http.StatusBadRequest).JSON(response.Err("el formato de la imagen debe ser image/png", nil))
	} */

	fileName, err := helper.SaveFile(fileHeader, "static/images/", "png")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.Err(err.Error(), nil))
	}

	return c.Status(http.StatusCreated).JSON(response.Successful("se guardo correctamente", map[string]string{
		"url": "http://" + c.Hostname() + string(c.Request().URI().Path()) + "/" + fileName,
	}))
}
