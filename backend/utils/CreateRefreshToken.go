package utils

import (
	"time"

	"github.com/alisalimli/bookstore/backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func CreateRefreshToken(context *fiber.Ctx, user models.User) (string, error) {
	refreshTokenClaims := jwt.MapClaims{
		"user_id":   user.ID,
		"user_name": user.Name,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // 30 min
	}

	// Create token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	// Generate encoded token and send it as response.
	encodedRefreshToken, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return encodedRefreshToken, nil
}
