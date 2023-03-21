package controllers

import (
	"fmt"
	"strings"

	"github.com/alisalimli/bookstore/backend/initializers"
	"github.com/alisalimli/bookstore/backend/models"
	"github.com/alisalimli/bookstore/backend/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUser(context *fiber.Ctx) error {
	var payload *models.SignUpInput

	if err := context.BodyParser(&payload); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "message": err.Error()})
	}

	if err := utils.ValidateStruct(payload); err != nil {
		// If validation fails, return custom error messages
		fmt.Println(err)
		return context.JSON(fiber.Map{
			"ok":     false,
			"errors": err,
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "message": err.Error()})
	}

	newUser := models.User{
		Name:     payload.Name,
		Email:    strings.ToLower(payload.Email),
		Password: string(hashedPassword),
		// Photo:    &payload.Photo,
	}

	result := initializers.DB.Create(&newUser)
	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return context.Status(fiber.StatusConflict).JSON(fiber.Map{"ok": false, "message": "User with that email already exists"})
	} else if result.Error != nil {
		return context.Status(fiber.StatusBadGateway).JSON(fiber.Map{"ok": false, "message": "Something went wrong!"})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{"ok": true, "data": fiber.Map{"user": models.FilterUserRecord(&newUser)}})

}
