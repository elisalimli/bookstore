package controllers

import (
	"fmt"

	"github.com/alisalimli/bookstore/backend/models"
	"github.com/gofiber/fiber/v2"
)

func SignUpUser(context *fiber.Ctx) error {
	var payload *models.SignUpInput

	if err := context.BodyParser(&payload); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	if err := models.ValidateStruct(payload); err != nil {
		// If validation fails, return custom error messages
		fmt.Println(err)
		return context.JSON(fiber.Map{
			"ok":     false,
			"errors": err,
		})
	}
	return context.SendString("hello")
}
