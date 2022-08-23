package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teris-io/shortid"
	"net/http"
	"pastebin/database"
	"pastebin/models"
)

func CreateCrate(c *fiber.Ctx) error {
	var dto models.CreatePasteDto
	if err := c.BodyParser(&dto); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var entity models.Paste
	entity = models.Paste{
		Content:  dto.Content,
		Password: dto.Password,
		SID:      shortid.MustGenerate(),
	}
	database.GetConnection().Create(&entity)

	return c.Status(http.StatusCreated).JSON(entity)
}

func GetPaste(c *fiber.Ctx) error {
	idParam := c.Params("pasteId")

	var entity models.Paste
	result := database.GetConnection().Where("s_id = ?", idParam).Find(&entity)
	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(entity)
}

func UpdatePaste(c *fiber.Ctx) error {
	idParam := c.Params("pasteId")

	var entity models.Paste
	result := database.GetConnection().Where("s_id = ?", idParam).Find(&entity)
	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	var dto models.UpdatePasteDto
	if err := c.BodyParser(&dto); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	entity.Content = dto.Content
	entity.Password = dto.Password
	database.GetConnection().Save(&entity)

	return c.Status(fiber.StatusOK).JSON(entity)
}

func DeletePaste(c *fiber.Ctx) error {
	idParam := c.Params("pasteId")

	var entity models.Paste
	result := database.GetConnection().Where("s_id = ?", idParam).Delete(&entity)
	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.SendStatus(fiber.StatusOK)
}
