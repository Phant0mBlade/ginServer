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
	env := systemUtils.GetEnv("ENVIRONMENT", "DEV")
	// var zeroLogger ZeroLoggerStruct

	// we call zeroLogger creation exactly once
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
	// msg    string
	logger *zerolog.Event
}

// set level for logging
func (l *ZeroLoggerStruct) Debug() *ZeroLoggerStruct {
	// l.msg = msg
	// l.log.Debug().Msg(msg)
	l.logger = l.log.Debug()
	return l
}

// set level for logging
func (l *ZeroLoggerStruct) Info() *ZeroLoggerStruct {
	// l.log.Info().Msg(msg)
	l.logger = l.log.Info()
	return l
}

// set level for logging
func (l *ZeroLoggerStruct) Error() *ZeroLoggerStruct {
	// l.log.Error().Msg(msg)
	l.logger = l.log.Error()
	return l
}

// set any extra field for logging
func (l *ZeroLoggerStruct) Any(key string, val interface{}) *ZeroLoggerStruct {
	l.logger = l.logger.Any(key, val)
	return l
}

// log actual msg with extra field if passed
func (l *ZeroLoggerStruct) Msg(msg string) {
	l.logger.Msg(msg)
}

/*
func (l *ZeroLoggerStruct) Debug(msg, key string, val interface{}) {
	l.log.Debug().Any(key, val).Msg(msg)
}

func (l *ZeroLoggerStruct) Info(msg, key string, val interface{}) {
	l.log.Info().Any(key, val).Msg(msg)
}

func (l *ZeroLoggerStruct) Error(msg, key string, val interface{}) {
	l.log.Error().Any(key, val).Msg(msg)
}
*/
