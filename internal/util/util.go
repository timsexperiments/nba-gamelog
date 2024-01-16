package util

import (
	"encoding/csv"
	"fmt"
	"os"
)

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func WriteCSV(data [][]string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range data {
		if err := writer.Write(row); err != nil {
			return err // returns an error if there's an issue writing a row
		}
	}

	return nil
}

func PrintWarning(message string) {
	const yellow = "\033[33m"
	const reset = "\033[0m"

	fmt.Printf("%s%s%s\n", yellow, message, reset)
}
