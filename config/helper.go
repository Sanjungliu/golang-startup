package config

import (
	"os"
	"strconv"
)

func getStringOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getIntOrDefault(key string, defaultValue int) int {
	value := os.Getenv(key)
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil || intValue == 0 {
		return defaultValue
	}
	return int(intValue)
}

func getBoolOrDefault(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return boolValue
}
