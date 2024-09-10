package utils

import (
	"os"
	"testing"

	"albatross/internal/models"
)

func TestReplaceFileExtension(t *testing.T) {
	tests := []struct {
		name      string
		filename  string
		newSuffix string
		expected  string
	}{
		{"Replace .csv", "mlm2pro.csv", "_processed.csv", "mlm2pro_processed.csv"},
		{"Replace .txt", "report.txt", ".pdf", "report.pdf"},
		{"No extension", "datafile", ".csv", "datafile.csv"},
		{"Multiple dots", "data.backup.csv", "_new.csv", "data.backup_new.csv"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReplaceFileExtension(tt.filename, tt.newSuffix)
			if result != tt.expected {
				t.Errorf("ReplaceFileExtension(%q, %q) = %q, want %q", tt.filename, tt.newSuffix, result, tt.expected)
			}
		})
	}
}

func TestWriteCSV(t *testing.T) {
	testFile := "test_output.csv"
	testData := []models.ProcessedShotData{
		{Club: "Dr", Type: "Tee", Target: 250.0, Total: 260.0, Side: 5.0},
		{Club: "7i", Type: "Approach", Target: 150.0, Total: 155.0, Side: -2.0},
	}

	// Clean up the test file after the test
	defer os.Remove(testFile)

	err := WriteCSV(testFile, testData)
	if err != nil {
		t.Fatalf("WriteCSV failed: %v", err)
	}

	// Check if file was created
	_, err = os.Stat(testFile)
	if os.IsNotExist(err) {
		t.Fatalf("Output file was not created")
	}

	// Read the contents of the file and verify
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	expectedContent := `Club,Type,Target,Total,Side
Dr,Tee,250.00,260.00,5.00
7i,Approach,150.00,155.00,-2.00
`

	if string(content) != expectedContent {
		t.Errorf("File content mismatch.\nExpected:\n%s\nGot:\n%s", expectedContent, string(content))
	}

	// Test with empty data
	err = WriteCSV(testFile, []models.ProcessedShotData{})
	if err == nil {
		t.Errorf("Expected error for empty data, got nil")
	}
}

func TestWriteCSVErrors(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		data     []models.ProcessedShotData
		wantErr  bool
	}{
		{"Invalid filename", "/invalid/path/file.csv", []models.ProcessedShotData{{Club: "Dr"}}, true},
		{"Empty data", "empty.csv", []models.ProcessedShotData{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteCSV(tt.filename, tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
