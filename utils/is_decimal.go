package utils

import "strings"

func IsDecimal(str string) bool {
	trimmedStr := strings.TrimSpace(str)
	splittedStr := strings.Split(trimmedStr, ".")
	if len(splittedStr) == 2 {
		return true
	}

	return false
}
