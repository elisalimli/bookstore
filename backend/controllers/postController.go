package controllers

import (
	"log"

	"github.com/elisalimli/bookstore/backend/initializers"
	"github.com/elisalimli/bookstore/backend/models"
	"github.com/gofiber/fiber/v2"
)

type CreatePostInput struct {
	Title string
	Body  string
}

func CreatePost(context *fiber.Ctx) error {
	// Create new Book struct
	input := &CreatePostInput{}
	err := context.BodyParser(input)

	if err != nil {
		log.Fatal("Couldn't parse the body")
	}
	post := models.Post{Title: input.Title, Body: input.Body}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		return fiber.NewError(400, "Something went wrong!")
	}

	return context.JSON(fiber.Map{"post": post})
}
