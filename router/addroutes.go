package addroutes

import (
	albumcontroller "ginServer/controller/album"
	log "ginServer/init/log"
	constants "ginServer/utils/constant"

	"github.com/gin-gonic/gin"
)

var logger = log.GetCustomLogger()

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
		logger.Info().Msg("This is an info log")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
