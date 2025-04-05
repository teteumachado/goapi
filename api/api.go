package api

import (
	"github.com/gin-gonic/gin"
	"github.com/teteumachado/goapi/features"
)

func Start() {
	r := gin.Default()

	features.UsersFeature(r)

	r.Run()
}
