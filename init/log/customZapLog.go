package log

import (
	"fmt"
	"sync"

	// "log/slog"

	systemUtils "ginServer/utils/system"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
	log   *zap.Logger
	zapF  zapcore.Field
	level string
}

// sets level of loggin
func (l *ZapLoggerStruct) Debug() *ZapLoggerStruct {
	// Need to find a better way for key-val log output
	l.level = "debug"
	return l
}

// sets level of loggin
func (l *ZapLoggerStruct) Info() *ZapLoggerStruct {
	// Need to find a better way for key-val log output
	l.level = "info"
	return l
}

// sets level of loggin
func (l *ZapLoggerStruct) Error() *ZapLoggerStruct {
	// Need to find a better way for key-val log output
	l.level = "error"
	return l
}

// set any extra field for logging
func (l *ZapLoggerStruct) Any(key string, val interface{}) *ZapLoggerStruct {
	// l.msg = msg
	// l.log.Debug().Msg(msg)
	l.zapF = zap.Any(key, val)
	return l
}

// actual logging with msg and the extra fields passed if any
func (l *ZapLoggerStruct) Msg(msg string) {
	// TODO: Need to find a better way
	// zerolog supports chaining itself so it was straight forward and easy
	// zaplog does not, so if l.zapF is undefined(nill and 0 values for its own field, check zapcore.Field)
	// we cannot use l.zapF without causing an error
	if l.zapF.Interface != nil {
		switch l.level {
		case "debug":
			fmt.Println("this is zapf", l.zapF)
			l.log.Debug(msg, l.zapF)
		case "info":
			l.log.Info(msg, l.zapF)
		case "error":
			l.log.Error(msg, l.zapF)
		}
	} else {
		switch l.level {
		case "debug":
			l.log.Debug(msg)
		case "info":
			l.log.Info(msg)
		case "error":
			l.log.Error(msg)
		}
	}
	l.zapF = zapcore.Field{}
}

/*
func (l *ZapLoggerStruct) Debug(msg, key string, val interface{}) {
	l.log.Debug(msg, zap.Any(key, val))
}

func (l *ZapLoggerStruct) Info(msg, key string, val interface{}) {
	l.log.Info(msg, zap.Any(key, val))
}

func (l *ZapLoggerStruct) Error(msg, key string, val interface{}) {
	l.log.Error(msg, zap.Any(key, val))
}
*/
