package util

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
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

func LoadingAnimation(done chan bool, loadingText string) {
	for {
		select {
		case <-done:
			fmt.Printf("\r%50s\r", " ") // Clear the line with enough spaces
			return
		default:
			fmt.Printf("\r%s", loadingText) // Print the static text
			time.Sleep(750 * time.Millisecond)
			for i := 0; i < 3; i++ {
				fmt.Print(".")
				time.Sleep(750 * time.Millisecond)
			}
			fmt.Printf("\r%80s\r", " ") // Clear the line after printing dots
		}
	}
}
