package main

import (
	"log"
	"os"

	"github.com/alisalimli/bookstore/backend/controllers"
	"github.com/alisalimli/bookstore/backend/initializers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v4"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}
func restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["user_name"].(string)
	return c.SendString("Welcome " + name)
}

func main() {
	// app := gin.Default()
	// app.POST("/post", controllers.CreatePost)
	// app.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	app := fiber.New()
	app.Use(logger.New())
	// Get / reading out the encrypted cookie
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("value=" + c.Cookies("refresh_token"))
	})

	// Post / create the encrypted cookie
	app.Post("/", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:  "test",
			Value: "SomeThing",
		})
		return nil
	})
	app.Post("/post", controllers.CreatePost)
	app.Post("/signup", controllers.SignUpUser)
	app.Post("/signin", controllers.LoginUser)
	app.Post("/refresh_token", controllers.RefreshToken)

	// // JWT Middleware
	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: []byte("secret"),
	// }))

	// Restricted Routes
	app.Get("/restricted", restricted)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))

}
