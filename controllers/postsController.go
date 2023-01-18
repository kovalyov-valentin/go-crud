package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-crud/initializers"
	"github.com/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	// Get data off request body
	var body models.Post
	c.Bind(&body)
	// Create post 
	post := models.Post{Title: body.Title, Body: body.Body,}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}
	// Return it 
	c.JSON(200, gin.H{
		"post": post,
	})
}


func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	// Respond with them 
	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get id off URL
	id := c.Param("id")

	// Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with them 
	c.JSON(200, gin.H{
		"post": post,
	})


}

func PostUpdate(c *gin.Context) {
	// Get the id off the URL 
	id := c.Param("id")
	// Get the data off request body 
	var body models.Post
	c.Bind(&body)
	// Find the post were updating 
	var post models.Post 
	initializers.DB.First(&post, id)
	// Update it 
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})
	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the id off the URL
	id := c.Param("id")

	// Delete the post
	var post models.Post
	initializers.DB.Delete(&post, id)

	// Respond 
	c.Status(200)
}