package config

import (
	"log"
	"os"
	"strconv"
)

// getEnv returns value for give key from environment
// if key is not present in environment it returns defaultValue
func getEnv(key, defaultValue string) string {
	v := os.Getenv(key)
	if len(v) > 0 {
		return v
	}
	return defaultValue
}

// getEnvInt returns integer value for give key from environment
// if key is not present in environment it returns defaultValue
// if key cannot be parsed to integer function will panic
func getEnvInt(key string, defaultValue int) int {
	v := os.Getenv(key)
	if len(v) == 0 {
		return defaultValue
	}

	valInteger, err := strconv.Atoi(v)
	if err != nil {
		log.Fatalf("variable `%s` cannot be parsed to INTEGER", key)
	}

	return valInteger
}
