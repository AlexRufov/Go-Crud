package main

import (
	"Go-Crud/controllers"
	"Go-Crud/inits"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.CreatePosts)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/post/:id", controllers.GetPost)
	r.PUT("/posts/:id", controllers.UpdatePosts)
	r.DELETE("/post/:id", controllers.DeletePost)

	err := r.Run()
	if err != nil {
		log.Fatal("Error starting application")
	}
}

func init() {
	inits.LoadEnvVariables()
	inits.ConnectToDB()
}
