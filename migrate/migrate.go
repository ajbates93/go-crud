package main

import (
	"tmp2/initialisers"
	"tmp2/models"
)

func init() {
	initialisers.ConnectToDB()
	initialisers.LoadEnvVariables()
}

func main() {
	initialisers.DB.AutoMigrate(&models.Post{})
}
