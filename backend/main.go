package main

import (
	"log"
	"os"

	"github.com/alisalimli/bookstore/backend/controllers"
	"github.com/alisalimli/bookstore/backend/initializers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	// app := gin.Default()
	// app.POST("/post", controllers.CreatePost)
	// app.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	app := fiber.New()
	app.Use(logger.New())

	app.Post("/post", controllers.CreatePost)
	app.Post("/signup", controllers.SignUpUser)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))

}
