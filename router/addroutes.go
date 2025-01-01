package router

import (
	albumcontroller "ginServer/controller/album"
	log "ginServer/init/log"
	constants "ginServer/utils/constant"

	"github.com/gin-gonic/gin"
)

// Get either the the zerolog or zap logger
//
// log.GetCustomZapLogger() or log.GetCustomZeroLogger()
var logger = log.GetCustomZeroLogger()

// AddRoutes Routes request to its request controller
func Addv1Routes(router *gin.RouterGroup) {
	// One do not necessary have to name them v1
	// they can `AddAlbumRoutes` and similaryly other names
	router.GET(constants.GetAlbumsURI, albumcontroller.GetAlbums)
	router.GET(constants.GetAlbumsIDURI, albumcontroller.GetAlbumByID)
	router.POST(constants.PostAlbumsURI, albumcontroller.PostAlbums)
}

// Health Check router
func AddHealthCheck(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		logger.Error("This is a error log")
		logger.Info("ABCD")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
