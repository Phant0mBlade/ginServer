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
var logger = log.GetCustomZapLogger()

// var logger2 = log.GetCustomZapLogger()

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
	// var x []interface{}
	var m = make(map[string]interface{})
	m["test"] = 22
	router.GET("/ping", func(c *gin.Context) {
		logger.Error("This is a error log", "myKey", m)
		logger.Info("ABCD", "myKey", m)
		// Uncomment to check if the logger and logger2 deeply equal and have the same address
		// fmt.Println(reflect.DeepEqual(logger, logger2))
		// fmt.Println(&logger, &logger2)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
