package constant

import (
	systemConfig "ginServer/utils/system"
)

var (
	//ServerPort The port on which the application will run
	ServerPort = systemConfig.GetEnv("SERVER_PORT", ":8080")
)

const (
	//Albums fetch URI
	GetAlbumsURI = "/albums"
	//Albums by ID fetch URI
	GetAlbumsIDURI = "/albums/:id"
	//Albums Bulk create URI
	PostAlbumsURI = "/albums"
	//Test test apis endpoint group
	TestURI = "/test"
	//Loggers test URI
	GetLoggersURI = "/loggers"
	//APIVersionV1 returns API Version v1 group
	APIVersionV1 = "v1/"
	//HealthCheck returns health check group
	HealthURI = "/health"
)
