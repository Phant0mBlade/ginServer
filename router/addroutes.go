package addroutes

import (
	albumcontroller "ginServer/controller/album"
	constants "ginServer/utils/constant"

	"github.com/gin-gonic/gin"
)

const (
	v1 = constants.APIVersionV1
)

// AddRoutes Routes request to its request controller
func AddRoutes(router *gin.RouterGroup) {
	router.GET(v1+constants.GetAlbumsURI, albumcontroller.GetAlbums)
	router.GET(v1+constants.GetAlbumsIDURI, albumcontroller.GetAlbumByID)
	router.POST(v1+constants.PostAlbumsURI, albumcontroller.PostAlbums)
}
