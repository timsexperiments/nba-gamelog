package util

import (
	"fmt"
)

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func PrintWarning(message string) {
	const yellow = "\033[33m"
	const reset = "\033[0m"

	fmt.Printf("%s%s%s\n", yellow, message, reset)
}
