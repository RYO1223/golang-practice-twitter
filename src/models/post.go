package models

import (
	"time"
)

// Post represents a post in the system
type Post struct {
	ID        int       `json:"id"`
	Content   string    `json:"content" binding:"required"`
	Author    string    `json:"author" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

// Posts slice to store post data
var Posts = []Post{}

// PostID counter for posts
var PostID int = 1
