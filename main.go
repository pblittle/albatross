package main

import (
	"log"
	"os"
	"strings"

	"albatross/internal/calculators"
	"albatross/internal/parsers"
	"albatross/internal/writer"
	"albatross/utils"
)

func main() {
	// Parse command-line arguments
	if len(os.Args) != 3 {
		log.Fatal("Usage: go run main.go <launch_monitor_type> <input_csv_file>")
	}

	launchMonitorType := normalizeLaunchMonitorType(os.Args[1])
	inputFile := os.Args[2]

	// Validate launch monitor type
	if !isValidLaunchMonitorType(launchMonitorType) {
		log.Fatalf("Error: Invalid launch monitor type '%s'. Supported type is mlm2pro.", os.Args[1])
	}

	// Process shot data
	shotData, err := parsers.ProcessShotData(inputFile, launchMonitorType)
	if err != nil {
		log.Fatalf("Error processing shot data: %v", err)
	}

	log.Printf("Processed shot data: %+v", shotData)

	// Calculate targets
	calculators.CalculateTargets(&shotData)

	log.Printf("Calculated targets: %+v", shotData)

	// Write processed data to output file
	outputFile := utils.ReplaceFileExtension(inputFile, "_processed.csv")
	writer := writer.ShotPatternWriter{}
	if err := writer.Write(outputFile, shotData); err != nil {
		log.Fatalf("Error writing output file: %v", err)
	}

	log.Printf("Successfully processed %d shots and saved results to %s", len(shotData), outputFile)
}

// normalizeLaunchMonitorType converts the launch monitor type to lowercase for consistency
func normalizeLaunchMonitorType(launchMonitorType string) string {
	return strings.ToLower(launchMonitorType)
}

// isValidLaunchMonitorType checks if the provided launch monitor type is supported
func isValidLaunchMonitorType(launchMonitorType string) bool {
	return launchMonitorType == "mlm2pro"
}
