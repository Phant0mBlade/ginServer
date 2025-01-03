package log

import (
	"fmt"
	"os"
	"sync"
	"time"

	systemUtils "ginServer/utils/system"

	"github.com/rs/zerolog"
	// "github.com/rs/zerolog/log"
)

var once sync.Once

// to ensure that whenever a new logger is required
// the same one is returned
var zeroLogger ZeroLoggerStruct

// return a once initialised logger
func GetCustomZeroLogger() ZeroLoggerStruct {
	env := systemUtils.GetEnv("ENVIRONMENT", "DEV1")
	// var zeroLogger ZeroLoggerStruct
	once.Do(func() {
		zeroLogger = createZeroLogger(env)
	})
	return zeroLogger
}

// Defines a new custom logger
func createZeroLogger(env string) ZeroLoggerStruct {
	fmt.Println("createCustomLogger running...")
	var zeroLogger zerolog.Logger
	if env == "DEV" {
		zeroLogger = zerolog.
			New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
			Level(zerolog.DebugLevel).
			With().
			Timestamp().
			Caller().
			Str("App Version", systemUtils.GetEnv("App_Version", "dev_alpha")).
			Logger()
	} else { // for production
		zeroLogger = zerolog.
			New(os.Stdout).
			Level(zerolog.InfoLevel).
			With().
			Timestamp().
			Str("App Version", systemUtils.GetEnv("App_Version", "production 0.1")).
			Logger()
	}
	fmt.Println("createCustomLogger done.")
	return ZeroLoggerStruct{log: zeroLogger}
}

type ZeroLoggerStruct struct {
	log zerolog.Logger
}

func (l *ZeroLoggerStruct) Debug(msg, key string, val interface{}) {
	l.log.Debug().Any(key, val).Msg(msg)
}

func (l *ZeroLoggerStruct) Info(msg, key string, val interface{}) {
	l.log.Info().Any(key, val).Msg(msg)
}

func (l *ZeroLoggerStruct) Error(msg, key string, val interface{}) {
	l.log.Error().Any(key, val).Msg(msg)
}
