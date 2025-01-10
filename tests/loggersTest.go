package test

import (
	"ginServer/init/log"

	"github.com/gin-gonic/gin"
)

// TestLoggers handler to test loggers
func TestLoggers(c *gin.Context) {
	var m = make(map[string]interface{})
	m["test"] = 22
	log.GlobalLogger.Error().Any("mykey", m).Msg("This is an error mesg")
	log.GlobalLogger.Error().Msg("This is an error msg")
	log.GlobalLogger.Error().Any("mykey", m).Msg("This is an error mesg")
	log.GlobalLogger.Info().Msg("info logger")
	log.GlobalLogger.Error().Any("mykey", m).Msg("This is an error mesg")

	c.JSON(200, gin.H{
		"message": "success",
	})
}
