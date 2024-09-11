// Package main is the entry point for the Albatross application.
// It processes shot data from various launch monitors and calculates targets.
package main

import (
	"os"
	"strings"

	"albatross/internal/calculators"
	"albatross/internal/logging"
	"albatross/internal/parsers"
	"albatross/internal/writer"
	"albatross/utils"
)

// main is the entry point of the application. It handles command-line arguments,
// processes shot data, calculates targets, and writes the results to a file.
func main() {
	// Initialize the logger
	logging.InitLogger()

	// Parse command-line arguments
	if len(os.Args) != 3 {
		logging.Fatal("Usage: go run main.go <launch_monitor_type> <input_csv_file>", nil)
	}

	launchMonitorType := normalizeLaunchMonitorType(os.Args[1])
	inputFile := os.Args[2]

	// Validate launch monitor type
	if !isValidLaunchMonitorType(launchMonitorType) {
		logging.Fatal("Error: Invalid launch monitor type. Supported type is mlm2pro.", logging.Fields{
			"providedType": launchMonitorType,
		})
	}

	// Process shot data from the input file
	shotData, err := parsers.ProcessShotData(inputFile, launchMonitorType)
	if err != nil {
		logging.Error("Error processing shot data", err, logging.Fields{
			"inputFile":         inputFile,
			"launchMonitorType": launchMonitorType,
		})
		os.Exit(1)
	}

	logging.Info("Processed shot data", logging.Fields{
		"count": len(shotData),
	})

	// Calculate targets based on the processed shot data
	calculators.CalculateTargets(&shotData)

	logging.Debug("Calculated targets", logging.Fields{
		"shotData": shotData,
	})

	// Write processed data to an output file
	outputFile := utils.ReplaceFileExtension(inputFile, "_processed.csv")
	writer := writer.ShotPatternWriter{}
	if err := writer.Write(outputFile, shotData); err != nil {
		logging.Error("Error writing output file", err, logging.Fields{
			"outputFile": outputFile,
		})
		os.Exit(1)
	}

	logging.Info("Successfully processed shots and saved results", logging.Fields{
		"shotsProcessed": len(shotData),
		"outputFile":     outputFile,
	})
}

// normalizeLaunchMonitorType converts the launch monitor type to lowercase for consistency.
// This ensures that the type check is case-insensitive.
func normalizeLaunchMonitorType(launchMonitorType string) string {
	return strings.ToLower(launchMonitorType)
}

// isValidLaunchMonitorType checks if the provided launch monitor type is supported.
// Currently, only "mlm2pro" is supported.
func isValidLaunchMonitorType(launchMonitorType string) bool {
	return launchMonitorType == "mlm2pro"
}
