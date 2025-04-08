package features

import (
	"github.com/gin-gonic/gin"
	"github.com/teteumachado/goapi/controllers"
	"github.com/teteumachado/goapi/core"
)

func UsersFeature(r *gin.Engine) gin.RouterGroup {
	feature := core.Feature{
		Route:      "/users",
		Controller: controllers.UsersController,
	}

	return *feature.RegisterFeature(r)
}
