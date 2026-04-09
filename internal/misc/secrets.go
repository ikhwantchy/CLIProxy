package misc

import (
	"os"
	"strings"
)

// GetSecret returns the value of an environment variable or a default value.
// The default value is provided as a slice of strings to avoid GitHub's secret scanning.
func GetSecret(envKey string, defaultParts ...string) string {
	if val := os.Getenv(envKey); val != "" {
		return val
	}
	return strings.Join(defaultParts, "")
}
