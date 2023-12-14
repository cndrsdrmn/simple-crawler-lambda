package utils

import (
	"os"
	"strconv"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

func Env(key string, defaultValue ...string) string {
	if v, exists := os.LookupEnv(key); exists {
		return v
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return ""
}

func EnvBool(key string, defaultValue ...bool) bool {
	str := strings.ToLower(Env(key))

	switch str {
	case "1", "t", "true":
		return true
	case "0", "f", "false":
		return false
	default:
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}

		return false
	}
}

func EnvFloat(key string, defaultValue ...float64) float64 {
	if str := Env(key); str != "" {
		if val, err := strconv.ParseFloat(str, 64); err == nil {
			return val
		}
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return 0.0
}

func EnvInt(key string, defaultValue ...int) int {
	str := Env(key)

	if value, err := strconv.Atoi(str); err == nil {
		return value
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return 0
}
