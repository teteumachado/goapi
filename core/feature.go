package core

import "github.com/gin-gonic/gin"

type Controller func(r *gin.RouterGroup)

type Feature struct {
	Route      string
	Controller Controller
}

func (f *Feature) RegisterFeature(router *gin.Engine) *gin.RouterGroup {
	group := router.Group(f.Route)
	f.Controller(group)
	return group
}
