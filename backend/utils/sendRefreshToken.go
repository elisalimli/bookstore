package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func SendRefreshToken(c *fiber.Ctx, refreshToken string) {
	cookie := fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		// Secure:   true,
	}

	c.Cookie(&cookie)
}
