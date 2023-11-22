package main

import (
	"tmp2/controllers"
	"tmp2/initialisers"

	"github.com/gin-gonic/gin"
)

func init() {
	initialisers.LoadEnvVariables()
	initialisers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}
