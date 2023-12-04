package main

import (
	"github.com/go-crud/initializers"
	"github.com/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// Migrate the schema
	// to access the database
	initializers.DB.AutoMigrate(&models.Post{})
}
