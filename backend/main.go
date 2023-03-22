package main

import (
	"fmt"
	"log"
	"os"

	"github.com/elisalimli/bookstore/backend/controllers"
	"github.com/elisalimli/bookstore/backend/initializers"
	"github.com/elisalimli/bookstore/backend/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v4"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}
func me(c *fiber.Ctx) error {
	user := c.Locals("jwt").(*jwt.Token)
	fmt.Println("user", user)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["user_name"].(string)
	return c.SendString("Welcome " + name)
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
