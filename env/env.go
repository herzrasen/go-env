package env

import (
	"fmt"
	"os"
)

func GetOrError(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return "", fmt.Errorf("environment variable not defined: %s", key)
	}
	return v, nil
}

func GetOrElse(key string, orElse string) string {
	v := os.Getenv(key)
	if v == "" {
		return orElse
	}
	return v
}
