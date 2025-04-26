package main

import (
	"gin/src/controllers"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Posts API endpoints
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPostByID)
	r.POST("/posts", controllers.CreatePost)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
