package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"pastebin/database"
	"pastebin/models"
)

func GetPaste(c *fiber.Ctx) error {
	idParam := c.Params("pasteId")

	var entity models.Paste
	result := database.GetConnection().Where("id = ?", idParam).Find(&entity)
	if result.RowsAffected > 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(entity)
}

func CreateCrate(c *fiber.Ctx) error {
	var crateDto models.CreatePasteDto
	if err := c.BodyParser(&crateDto); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var entity models.Paste
	entity = models.Paste{
		Content: crateDto.Content,
	}
	database.GetConnection().Create(&entity)

	return c.Status(http.StatusCreated).JSON(entity)
}
