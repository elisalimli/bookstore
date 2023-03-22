package middleware

import (
	"github.com/elisalimli/bookstore/backend/utils"
	"github.com/gofiber/fiber/v2"

	jwtMiddleware "github.com/gofiber/jwt/v2"
)

// JWTProtected func for specify routes group with JWT authentication.
// See: https://github.com/gofiber/jwt
func JWTProtected() func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	config := jwtMiddleware.Config{
		// SigningKey:   []byte(os.Getenv("JWT_SECRET_KEY")),
		SigningKey:   []byte("secret"),
		ContextKey:   "jwt", // used in private routes
		ErrorHandler: utils.JwtError,
	}

	return jwtMiddleware.New(config)
}
