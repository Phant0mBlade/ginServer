package router

import (
	albumcontroller "ginServer/controller/album"
	test "ginServer/tests"
	constants "ginServer/utils/constant"

	"github.com/gin-gonic/gin"
)

// AddRoutes Routes request to its request controller
func Addv1Routes(router *gin.RouterGroup) {
	// One do not necessary have to name them v1
	// they can `AddAlbumRoutes` and similaryly other names
	router.GET(constants.GetAlbumsURI, albumcontroller.GetAlbums)
	router.GET(constants.GetAlbumsIDURI, albumcontroller.GetAlbumByID)
	router.POST(constants.PostAlbumsURI, albumcontroller.PostAlbums)
}

func AddTestRoutes(router *gin.RouterGroup) {
	router.GET(constants.GetLoggersURI, test.TestLoggers)
}

// Health Check router
func AddHealthCheck(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
