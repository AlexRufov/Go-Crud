package controllers

import (
	"Go-Crud/inits"
	"Go-Crud/models"
	"github.com/gin-gonic/gin"
)

func CreatePosts(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}

	err := c.Bind(&body)
	if err != nil {
		return
	}

	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	result := inits.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	inits.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	inits.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePosts(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Body  string
		Title string
	}

	err := c.Bind(&body)
	if err != nil {
		return
	}
	var post models.Post
	inits.DB.First(&post, id)

	inits.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	inits.DB.Delete(&models.Post{}, id)
	c.Status(200)
}
