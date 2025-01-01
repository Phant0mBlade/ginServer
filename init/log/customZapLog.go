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
// var zapLogger *zap.Logger

// return a once initialised logger
func GetCustomZapLogger() ZapLogger {
	env := systemUtils.GetEnv("ENVIRONMENT", "DEV1")
	var zapLogger ZapLogger
	zapOnce.Do(func() {
		zapLogger = *createZapLogger(env)
	})
	return zapLogger
}

// Defines a new custom logger over the global logger variable
func createZapLogger(env string) *ZapLogger {
	fmt.Println("createCustomLogger running...")

	var config zap.Config
	var zapLogger *zap.Logger

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

	return &ZapLogger{log: zapLogger}
}

type ZapLogger struct {
	log *zap.Logger
}

func (l *ZapLogger) Debug(msg, key string, val interface{}) {
	l.log.Debug(msg, zap.Any(key, val))
}

func (l *ZapLogger) Info(msg, key string, val interface{}) {
	l.log.Info(msg, zap.Any(key, val))
}

func (l *ZapLogger) Error(msg, key string, val interface{}) {
	l.log.Error(msg, zap.Any(key, val))
}
