package system

/*
This file contains utilities to be used
*/

import (
	"os"
)

// GetEnv fetches an environment variable if exist else returns the default value
func GetEnv(key, defaultStr string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultStr
}
