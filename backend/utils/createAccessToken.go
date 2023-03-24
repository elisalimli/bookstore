package utils

import (
	"time"

	"github.com/elisalimli/bookstore/backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func CreateAccessToken(context *fiber.Ctx, user models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   user.ID,
		"user_name": user.Name,
		"exp":       time.Now().Add(time.Minute * 30).Unix(), // 30 min
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	encodedAccessToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return encodedAccessToken, nil
}
