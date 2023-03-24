package main

import (
	"fmt"
	"log"
	"os"

	"github.com/elisalimli/bookstore/backend/controllers"
	"github.com/elisalimli/bookstore/backend/initializers"
	"github.com/elisalimli/bookstore/backend/middleware"
	"github.com/elisalimli/bookstore/backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v4"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}
func me(c *fiber.Ctx) error {
	token := c.Locals("jwt").(*jwt.Token)
	// Extract the user ID from the token's claims
	claims := token.Claims.(jwt.MapClaims)

	userID := claims["user_id"].(string)
	var user models.User

	result := initializers.DB.First(&user, "id = ?", userID)

	fmt.Println("User :", result)

	return c.JSON(fiber.Map{"ok": true, "user": user})
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Post("/post", controllers.CreatePost)
	app.Post("/signup", controllers.SignUpUser)
	app.Post("/signin", controllers.LoginUser)
	app.Post("/refresh_token", controllers.RefreshToken)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("works!")
	})

	// Restricted Routes
	app.Get("/me", middleware.JWTProtected(), me)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))

}
