package features

import (
	"github.com/gin-gonic/gin"
	"github.com/teteumachado/goapi/core"
)

func UsersController(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get all users",
		})
	})
}

func UsersFeature(r *gin.Engine) gin.RouterGroup {
	feature := core.Feature{
		Route:      "/users",
		Controller: UsersController,
	}

	return *feature.RegisterFeature(r)
}
