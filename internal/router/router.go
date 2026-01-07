package router

import (
	"github.com/hans312456/go-gin-project/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")

	user := controller.NewUserController()

	api.GET("/users/:id", user.GetByID)
	api.POST("/users", user.Create)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

}
