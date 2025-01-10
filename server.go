/*
This is where my starting code will be,
Code structure is still in development so things may change
*/

package main

import (
	"github.com/gin-gonic/gin"

	addRoutes "ginServer/router"
	systemConst "ginServer/utils/constant"
)

func ginRouter() *gin.Engine {
	return gin.Default()
}

// add routes to their specific groups
func addRouterGroups(router *gin.Engine) {
	addRoutes.Addv1Routes(router.Group(systemConst.APIVersionV1))
	addRoutes.AddTestRoutes(router.Group(systemConst.TestURI))
	addRoutes.AddHealthCheck(router.Group(systemConst.HealthURI))
}

// func addLogging() {
// }

// add middlewares
// func addMiddlewares(router *gin.Engine) {
// 	// TODO:
// }

// create router and run server in gin
func server() {
	router := ginRouter()
	addRouterGroups(router)
	router.Run("0.0.0.0" + systemConst.ServerPort) // listen and serve on 0.0.0.0:8080
}

func main() {
	server()
}
