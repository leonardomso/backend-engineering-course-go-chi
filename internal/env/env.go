package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return value
}

func GetInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	valueAsInt, error := strconv.Atoi(value)

	if error != nil {
		return fallback
	}

	return valueAsInt
}
