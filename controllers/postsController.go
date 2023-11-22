package controllers

import (
	"tmp2/initialisers"
	"tmp2/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initialisers.DB.Create(&post)

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
	initialisers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
	// Respond with them
}

func PostsShow(c *gin.Context) {
	// Get ID off url
	id := c.Param("id")

	var post models.Post
	initialisers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get the ID off the url
	id := c.Param("id")

	// Get the data off request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	// Find the post we're updating
	var post models.Post
	initialisers.DB.First(&post, id)

	// Update it
	initialisers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the ID off the url
	id := c.Param("id")

	// Delete the post
	initialisers.DB.Delete(&models.Post{}, id)

	// Repond with it
	c.Status(200)
}
