package router

import (
	albumcontroller "ginServer/controller/album"
	log "ginServer/init/log"
	constants "ginServer/utils/constant"

	"log/slog"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger = log.GetCustomLogger()
var zapLogger = log.GetCustomZapLogger()

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
		logger.Error().Int("x", 2).Msg("This is an Error log")
		zapLogger.Info("ABCD")
		zapLogger.Error("This is an error log", zap.Any("x", 2))
		slog.Info("This is slog info log")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
