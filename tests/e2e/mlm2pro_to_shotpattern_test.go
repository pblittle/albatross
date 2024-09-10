package e2e

import (
	"os"
	"path/filepath"
	"testing"

	"albatross/internal/calculators"
	"albatross/internal/parsers"
	"albatross/internal/writer"
)

func TestMLM2ProToShotPattern(t *testing.T) {
	tests := []struct {
		name           string
		inputFile      string
		expectedOutput string
	}{
		{
			name:           "MLM2Pro to ShotPattern Sample 1",
			inputFile:      "../../examples/input/mlm2pro.csv",
			expectedOutput: "../../examples/expected_output/mlm2pro_processed.csv",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Process the input file
			shotData, err := parsers.ProcessShotData(tt.inputFile, "mlm2pro")
			if err != nil {
				t.Fatalf("Failed to process shot data: %v", err)
			}

			// Calculate targets
			calculators.CalculateTargets(&shotData)

			// Create a temporary directory for our test output
			tempDir, err := os.MkdirTemp("", "albatross_test")
			if err != nil {
				t.Fatalf("Failed to create temporary test directory: %v", err)
			}
			defer os.RemoveAll(tempDir)

			// Define the output file path
			outputFile := filepath.Join(tempDir, "output_shotpattern.csv")

			// Write processed data to the output file
			writer := writer.ShotPatternWriter{}
			if err := writer.Write(outputFile, shotData); err != nil {
				t.Fatalf("Failed to write ShotPattern data to output file '%s': %v", outputFile, err)
			}

			// Compare the output with the expected output
			actualContent, err := os.ReadFile(outputFile)
			if err != nil {
				t.Fatalf("Failed to read actual ShotPattern output file '%s': %v", outputFile, err)
			}

			expectedContent, err := os.ReadFile(tt.expectedOutput)
			if err != nil {
				t.Fatalf("Failed to read expected ShotPattern output file '%s': %v", tt.expectedOutput, err)
			}

			if string(actualContent) != string(expectedContent) {
				t.Errorf("ShotPattern output mismatch.\nExpected:\n%s\nGot:\n%s", expectedContent, actualContent)
			}
		})
	}
}
