package writer

import (
	"os"
	"testing"

	"albatross/internal/models"
)

func TestShotPatternWriter(t *testing.T) {
	writer := ShotPatternWriter{}
	testFile := "test_output.csv"

	// Clean up the test file after the test
	defer os.Remove(testFile)

	testData := []models.ProcessedShotData{
		{Club: "Dr", Type: "Tee", Target: 250, Total: 260, Side: 5},
		{Club: "7i", Type: "Approach", Target: 150, Total: 155, Side: -2},
	}

	err := writer.Write(testFile, testData)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
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
	err = writer.Write(testFile, []models.ProcessedShotData{})
	if err == nil {
		t.Fatalf("Expected error for empty data, got nil")
	}
}
