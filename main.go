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
	
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/post/:id", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.PostsDelete)

	r.GET("/posts", controllers.PostsIndex)
	r.GET("/post/:id", controllers.PostsShow)

	r.Run() 
}