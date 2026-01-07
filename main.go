package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	type User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	var users = []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")

		for _, user := range users {
			if fmt.Sprint(user.ID) == id {
				c.JSON(200, user)
				return
			}
		}

		c.JSON(404, gin.H{
			"error": "User not found",
		})
	})

	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, users)
	})

	r.GET("/search", func(c *gin.Context) {
		keyword := c.Query("keyword")
		page := c.DefaultQuery("page", "1")

		c.JSON(200, gin.H{
			"keyword": keyword,
			"page":    page,
		})
	})

	r.POST("/users", func(c *gin.Context) {
		var newUser User
		if err := c.BindJSON(&newUser); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		newUser.ID = len(users) + 1
		users = append(users, newUser)
		c.JSON(201, newUser)
	})

	r.Run(":8080")
}
