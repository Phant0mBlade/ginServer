package log

import (
	"fmt"
	"sync"

	// "log/slog"

	systemUtils "ginServer/utils/system"

	"go.uber.org/zap"
)

var zapOnce sync.Once

// to ensure that whenever a new logger is required
// the same one is returned
var zapLogger ZapLoggerStruct

// return a once initialised logger
func GetCustomZapLogger() ZapLoggerStruct {
	env := systemUtils.GetEnv("ENVIRONMENT", "DEV1")
	// var zapLogger ZapLoggerStruct
	zapOnce.Do(func() {
		zapLogger = *createZapLogger(env)
	})
	return zapLogger
}

// Defines a new custom logger over the global logger variable
func createZapLogger(env string) *ZapLoggerStruct {
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

	return &ZapLoggerStruct{log: zapLogger}
}

type ZapLoggerStruct struct {
	log *zap.Logger
}

func (l *ZapLoggerStruct) Debug(msg, key string, val interface{}) {
	// Need to find a better way for key-val log output
	l.log.Debug(msg, zap.Any(key, val))
}

func (l *ZapLoggerStruct) Info(msg, key string, val interface{}) {
	// Need to find a better way for key-val log output
	l.log.Info(msg, zap.Any(key, val))
}

func (l *ZapLoggerStruct) Error(msg, key string, val interface{}) {
	// Need to find a better way for key-val log output
	l.log.Error(msg, zap.Any(key, val))
}

/*
func (l *ZapLoggerStruct) Debug(msg string, keyVal ...interface{}) {
	l.log.Debug(msg)
}

func (l *ZapLoggerStruct) Info(msg string, keyVal ...interface{}) {
	l.log.Info(msg)
}

func (l *ZapLoggerStruct) Error(msg string, keyVal ...interface{}) {
	l.log.Error(msg)
}
*/
