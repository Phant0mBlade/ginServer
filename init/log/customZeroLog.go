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

// we import this logger in every package/file of our project
// var zeroLogger zerolog.Logger

// return a once initialised logger
func GetCustomZeroLogger() ZeroLogger {
	env := systemUtils.GetEnv("ENVIRONMENT", "DEV1")
	var zeroLogger ZeroLogger
	once.Do(func() {
		zeroLogger = createZeroLogger(env)
	})
	return zeroLogger
}

// Defines a new custom logger over the global logger variable
func createZeroLogger(env string) ZeroLogger {
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
			Level(zerolog.ErrorLevel).
			With().
			Timestamp().
			Str("App Version", systemUtils.GetEnv("App_Version", "production 0.1")).
			Logger()
	}
	fmt.Println("createCustomLogger done.")
	return ZeroLogger{log: zeroLogger}
}

type ZeroLogger struct {
	log zerolog.Logger
}

func (l *ZeroLogger) Debug(msg string, keyVal ...interface{}) {
	l.log.Debug().Fields(keyVal).Msg(msg)
}

func (l *ZeroLogger) Info(msg string, keyVal ...interface{}) {
	l.log.Info().Fields(keyVal).Msg(msg)
}

func (l *ZeroLogger) Error(msg string, keyVal ...interface{}) {
	l.log.Error().Fields(keyVal).Msg(msg)
}
