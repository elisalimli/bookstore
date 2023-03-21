package main

import (
	"github.com/alisalimli/bookstore/backend/initializers"
	"github.com/alisalimli/bookstore/backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})

}
