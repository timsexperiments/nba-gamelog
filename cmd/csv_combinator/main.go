package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type strslice []string

func (s *strslice) String() string {
	return fmt.Sprintf("%s", *s)
}

func (s *strslice) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func main() {
	var inputDir string
	var outfile string
	var files strslice
	flag.StringVar(&inputDir, "dir", "", "Directory containing CSV files")
	flag.Var(&files, "files", "List of CSV files (can be specified multiple times)")
	flag.StringVar(&outfile, "out", "combined.csv", "Name of the file to output.")
	flag.Parse()

	if inputDir == "" && len(files) == 0 {
		fmt.Println("Error: Please specify either an input directory using the -dir flag or individual files using the -files flag")
		os.Exit(1)
	}

	if inputDir != "" && len(files) > 0 {
		fmt.Println("Error: Please specify either an input directory using the -dir flag or individual files using the -files flag, but not both")
		os.Exit(1)
	}

	// If directory is provided, read all CSV files from the directory
	if inputDir != "" {
		files, _ = readCSVFilesFromDir(inputDir)
	}

	combinedData, err := combineCSVFiles(files)
	if err != nil {
		fmt.Println("Error combining CSV files:", err)
		os.Exit(1)
	}

	err = writeCSV(outfile, combinedData)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		os.Exit(1)
	}

	fmt.Println("Combined CSV file created successfully:", outfile)
}

func readCSVFilesFromDir(dir string) (strslice, error) {
	var files strslice
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range fileInfos {
		if strings.HasSuffix(fileInfo.Name(), ".csv") {
			files = append(files, filepath.Join(dir, fileInfo.Name()))
		}
	}
	return files, nil
}

func combineCSVFiles(filePaths []string) ([][]string, error) {
	headersMap := make(map[string]int)
	var combinedData [][]string

	for _, filePath := range filePaths {
		file, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			return nil, err
		}

		// Process headers
		for _, header := range records[0] {
			if _, exists := headersMap[header]; !exists {
				headersMap[header] = len(headersMap)
			}
		}

		// Process data
		for _, record := range records[1:] {
			row := make([]string, len(headersMap))
			for i, value := range record {
				header := records[0][i]
				row[headersMap[header]] = value
			}
			combinedData = append(combinedData, row)
		}
	}

	// Create headers row
	headersRow := make([]string, len(headersMap))
	for header, index := range headersMap {
		headersRow[index] = header
	}

	// Prepend headers row to combined data
	combinedData = append([][]string{headersRow}, combinedData...)

	return combinedData, nil
}

func writeCSV(filePath string, data [][]string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range data {
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}
