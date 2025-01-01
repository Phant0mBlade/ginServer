package log

import (
	"fmt"
	"sync"

	// "log/slog"

	systemUtils "ginServer/utils/system"

	"go.uber.org/zap"
)

var zapOnce sync.Once

// we import this logger in every package/file of our project
var zapLogger *zap.Logger

// return a once initialised logger
func GetCustomZapLogger() *zap.Logger {
	env := systemUtils.GetEnv("ENVIRONMENT", "DEV1")
	zapOnce.Do(func() {
		createCustomZapLogger(env)
	})
	return zapLogger
}

// Defines a new custom logger over the global logger variable
func createCustomZapLogger(env string) {
	fmt.Println("createCustomLogger running...")
	var config zap.Config
	if env == "DEV" {
		config = zap.NewDevelopmentConfig()
		// zapLogger, _ = zap.NewDevelopment()
	} else { // for production
		config = zap.NewProductionConfig()
		// zapLogger, _ = zap.NewProduction()
	}
	config.DisableStacktrace = true
	zapLogger, _ = config.Build()
	fmt.Println("createCustomLogger done.")
}
