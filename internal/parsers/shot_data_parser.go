// Package parsers provides functionality for parsing shot data from various launch monitor types.
package parsers

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"albatross/internal/models"
	"albatross/internal/reader"
)

// headerPattern is a regular expression used to identify header rows in the CSV file.
var headerPattern = regexp.MustCompile(`(?i)(club type|total distance|side carry)`)

// ProcessShotData reads and processes shot data from a CSV file.
// It supports different launch monitor types and returns a slice of ProcessedShotData.
func ProcessShotData(inputFile string, launchMonitorType string) ([]models.ProcessedShotData, error) {
	// Open the input file
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	// Set up CSV reader
	csvReader := csv.NewReader(file)
	csvReader.Comma = ',' // Using comma as separator
	csvReader.LazyQuotes = true
	csvReader.FieldsPerRecord = -1 // Allow variable number of fields

	// Create appropriate launch monitor based on the type
	var launchMonitor models.LaunchMonitor
	switch launchMonitorType {
	case "mlm2pro":
		launchMonitor = reader.NewMLM2ProLaunchMonitor()
	default:
		return nil, fmt.Errorf("unsupported launch monitor type: %s", launchMonitorType)
	}

	var shotData []models.ProcessedShotData
	var headers []string
	inDataBlock := false

	// Read and process each row of the CSV file
	for {
		row, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("reading row: %w", err)
		}

		if len(row) == 0 {
			continue
		}

		// Check if the current row is a header row
		if isHeader(row) {
			headers = normalizeHeaders(row)
			inDataBlock = true
			log.Printf("Found headers: %v", headers)
			continue
		}

		if !inDataBlock {
			continue
		}

		// Check if we've reached the end of the data block
		if isEmptyRow(row) || strings.HasPrefix(strings.ToLower(row[0]), "average") {
			inDataBlock = false
			continue
		}

		// Parse and process the row data
		rawData, err := launchMonitor.ParseRow(row, headers)
		if err != nil {
			log.Printf("Skipping row due to error: %v", err)
			continue
		}

		processedData := launchMonitor.ProcessRawData(rawData)
		shotData = append(shotData, processedData)
	}

	if len(shotData) == 0 {
		return nil, fmt.Errorf("no valid data found in the file")
	}

	return shotData, nil
}

// isHeader checks if a row is a header row by matching against the headerPattern.
func isHeader(row []string) bool {
	if len(row) == 0 {
		return false
	}
	for _, cell := range row {
		if headerPattern.MatchString(cell) {
			return true
		}
	}
	return false
}

// normalizeHeaders standardizes header names by converting them to lowercase and trimming whitespace.
func normalizeHeaders(row []string) []string {
	normalized := make([]string, len(row))
	for i, header := range row {
		normalized[i] = strings.ToLower(strings.TrimSpace(header))
	}
	return normalized
}

// isEmptyRow checks if a row is empty by verifying that all cells are empty strings when trimmed.
func isEmptyRow(row []string) bool {
	for _, cell := range row {
		if strings.TrimSpace(cell) != "" {
			return false
		}
	}
	return true
}
