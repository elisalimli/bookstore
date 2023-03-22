package controllers

import (
	"fmt"
	"strings"

	"github.com/elisalimli/bookstore/backend/initializers"
	"github.com/elisalimli/bookstore/backend/models"
	"github.com/elisalimli/bookstore/backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUser(context *fiber.Ctx) error {
	var payload *models.SignUpInput

	// parsing body
	if err := context.BodyParser(&payload); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "message": err.Error()})
	}
	// validating input fields
	if errors := utils.ValidateStruct(payload); errors != nil {
		// If validation fails, return custom error messages
		fmt.Println(errors)
		return context.JSON(fiber.Map{
			"ok":     false,
			"errors": errors,
		})
	}
	// hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "message": err.Error()})
	}
	// creating a new user
	newUser := models.User{
		Name:     payload.Name,
		Email:    strings.ToLower(payload.Email),
		Password: string(hashedPassword),
		// Photo:    &payload.Photo,
	}
	// saving user to database
	result := initializers.DB.Create(&newUser)
	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicated key") {
		return context.Status(fiber.StatusConflict).JSON(fiber.Map{"ok": false, "message": "User with that email already exists"})
	} else if result.Error != nil {
		return context.Status(fiber.StatusBadGateway).JSON(fiber.Map{"ok": false, "message": "Something went wrong!"})
	}
	// success
	return context.Status(fiber.StatusCreated).JSON(fiber.Map{"ok": true, "data": fiber.Map{"user": models.FilterUserRecord(&newUser)}})

}
func LoginUser(c *fiber.Ctx) error {
	var payload *models.SignInInput
	// parsing body
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "message": err.Error()})
	}
	// validating input fields
	if errors := utils.ValidateStruct(payload); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "errors": errors})
	}

	var user models.User
	// querying the user
	result := initializers.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "message": "Invalid email or Password"})
	}
	// comparing password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "message": "Invalid email or Password"})
	}

	encodedRefreshToken, err := utils.CreateRefreshToken(c, user)
	if err != nil {
		return utils.JwtError(c, err)
	}
	// saving refresh token in cookie
	utils.SendRefreshToken(c, encodedRefreshToken)

	encodedAccessToken, err := utils.CreateAccessToken(c, user)

	if err != nil {
		return utils.JwtError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"ok": true, "access_token": encodedAccessToken})
}

func RefreshToken(c *fiber.Ctx) error {
	refreshTokenCookie := c.Cookies("refresh_token")

	// Verify that the refresh token is valid
	claims := jwt.MapClaims{}
	refreshToken, err := jwt.ParseWithClaims(refreshTokenCookie, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil // Replace with your actual secret key
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "message": "Invalid refresh token"})
	}
	if !refreshToken.Valid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "message": "Invalid refresh token"})
	}
	// Extract the user ID from the refresh token
	userID, ok := claims["user_id"]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "message": "Invalid refresh token"})
	}

	// Retrieve the user from the database
	var user models.User

	result := initializers.DB.First(&user).Where("id = ?", userID)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ok": false, "message": "User not found"})
	}

	encodedAccessToken, err := utils.CreateAccessToken(c, user)
	if err != nil {
		return utils.JwtError(c, err)
	}
	encodedRefreshToken, err := utils.CreateRefreshToken(c, user)
	if err != nil {
		return utils.JwtError(c, err)
	}
	utils.SendRefreshToken(c, encodedRefreshToken)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return the new access and refresh tokens in the response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"ok": true, "access_token": encodedAccessToken})
}
