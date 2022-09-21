package main

import (
	"Go-Crud/inits"
	"Go-Crud/models"
)

func init() {
	inits.LoadEnvVariables()
	inits.ConnectToDB()
}

func main() {
	inits.DB.AutoMigrate(&models.Post{})
}
