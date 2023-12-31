package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-crud/controllers"
	"github.com/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// Send requests
	r.POST("/posts", controllers.PostsCreate)

	// Read requests
	r.GET("/posts", controllers.PostsIndex)

	// Get request by id
	r.GET("/posts/:id", controllers.PostsShow)

	// Update request by id
	r.PUT("/posts/:id", controllers.PostsUpdate)

	// Delete request by id
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
}
