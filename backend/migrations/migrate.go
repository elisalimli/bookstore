package main

import (
	"log"
	"os"

	"github.com/alisalimli/bookstore/backend/initializers"
	"github.com/alisalimli/bookstore/backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	log.Println("Running Migrations")
	err := initializers.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Migration Failed:  \n", err.Error())
		os.Exit(1)
	}

}
