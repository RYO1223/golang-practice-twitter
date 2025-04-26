package controllers

import (
	"net/http"
	"strconv"
	"time"

	"gin/src/models"

	"github.com/gin-gonic/gin"
)

// GetPosts returns all posts
func GetPosts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"posts": models.Posts,
	})
}

// GetPostByID returns a single post by ID
func GetPostByID(c *gin.Context) {
	idParam := c.Params.ByName("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	for _, post := range models.Posts {
		if post.ID == id {
			c.JSON(http.StatusOK, post)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
}

// CreatePost creates a new post
func CreatePost(c *gin.Context) {
	var newPost models.Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set post metadata
	newPost.ID = models.PostID
	newPost.CreatedAt = time.Now()

	// Increment ID counter
	models.PostID++

	// Add to posts slice
	models.Posts = append(models.Posts, newPost)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Post created successfully",
		"post":    newPost,
	})
}
