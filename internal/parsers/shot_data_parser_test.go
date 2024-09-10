package parsers

import (
	"os"
	"reflect"
	"testing"

	"albatross/internal/models"
)

func TestProcessShotData(t *testing.T) {
	// Create a temporary CSV file for testing
	tempFile, err := os.CreateTemp("", "test_shot_data_*.csv")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write test data to the temp file
	testData := `Club Type,Total Distance,Side Carry
Dr,250,10
7i,150,-5
`
	if _, err := tempFile.Write([]byte(testData)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tempFile.Close()

	// Test ProcessShotData function
	shotData, err := ProcessShotData(tempFile.Name(), "mlm2pro")
	if err != nil {
		t.Fatalf("ProcessShotData failed: %v", err)
	}

	expectedData := []models.ProcessedShotData{
		{Club: "Dr", Type: "Tee", Total: 250, Side: 10},
		{Club: "7i", Type: "Approach", Total: 150, Side: -5},
	}

	if !reflect.DeepEqual(shotData, expectedData) {
		t.Errorf("ProcessShotData result mismatch.\nGot: %+v\nWant: %+v", shotData, expectedData)
	}
}

func TestIsHeader(t *testing.T) {
	tests := []struct {
		name     string
		row      []string
		expected bool
	}{
		{"Valid header", []string{"Club Type", "Total Distance", "Side Carry"}, true},
		{"Invalid header", []string{"Name", "Age", "Score"}, false},
		{"Empty row", []string{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isHeader(tt.row)
			if result != tt.expected {
				t.Errorf("isHeader(%v) = %v, want %v", tt.row, result, tt.expected)
			}
		})
	}
}

func TestNormalizeHeaders(t *testing.T) {
	input := []string{"Club Type", "Total Distance", "Side Carry"}
	expected := []string{"club type", "total distance", "side carry"}

	result := normalizeHeaders(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("normalizeHeaders(%v) = %v, want %v", input, result, expected)
	}
}

func TestIsEmptyRow(t *testing.T) {
	tests := []struct {
		name     string
		row      []string
		expected bool
	}{
		{"Empty row", []string{"", "", ""}, true},
		{"Non-empty row", []string{"data", "", "more data"}, false},
		{"Whitespace row", []string{" ", "\t", "  "}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isEmptyRow(tt.row)
			if result != tt.expected {
				t.Errorf("isEmptyRow(%v) = %v, want %v", tt.row, result, tt.expected)
			}
		})
	}
}
