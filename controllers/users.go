package controllers

import (
	"github.com/gin-gonic/gin"
)

func UsersController(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get all users",
		})
	})
}
