package stocks

import (
	"fmt"
	"os"
)

// GetEnvVar retrieves an environment variable
func GetEnvVar(key string) string {
	value := os.Getenv(key)
	if value == "" {
		fmt.Println("Warning: Env var is not set")
	}
	return value
}
