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
var logger zerolog.Logger

// return a once initialised logger
func GetCustomLogger() zerolog.Logger {
	env := systemUtils.GetEnv("ENVIRONMENT", "DEV1")
	once.Do(func() {
		createCustomLogger(env)
	})
	return logger
}

// Defines a new custom logger over the global logger variable
func createCustomLogger(env string) {
	fmt.Println("createCustomLogger running...")

	if env == "DEV" {
		logger = zerolog.
			New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
			Level(zerolog.DebugLevel).
			With().
			Timestamp().
			Caller().
			Str("App Version", systemUtils.GetEnv("App_Version", "dev_alpha")).
			Logger()
	} else { // for production
		logger = zerolog.
			New(os.Stdout).
			Level(zerolog.ErrorLevel).
			With().
			Timestamp().
			Str("App Version", systemUtils.GetEnv("App_Version", "production 0.1")).
			Logger()
	}
	fmt.Println("createCustomLogger done.")
}
