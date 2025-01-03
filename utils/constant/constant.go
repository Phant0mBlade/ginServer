package constant

import (
	systemConfig "ginServer/utils/system"
)

var (
	//ServerPort The port on which the application will run
	ServerPort = systemConfig.GetEnv("SERVER_PORT", ":8080")
)

const (
	//Albums callback URI
	GetAlbumsURI = "/albums"
	//Albums by ID callback URI
	GetAlbumsIDURI = "/albums/:id"
	//Push Bulk callback URI
	PostAlbumsURI = "/albums"
	//APIVersionV1 returns API Version v1
	APIVersionV1 = "v1/"
)
