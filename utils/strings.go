package utils

import (
	"regexp"
	"strings"
)

func ValidIDString(evalID string) string {
	newIDString := strings.ReplaceAll(evalID, " ", "_")
	return regexp.MustCompile(`[^a-zA-Z0-9_]+`).ReplaceAllString(newIDString, "")
}
