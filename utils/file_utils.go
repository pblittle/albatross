package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"albatross/internal/models"
)

// ReplaceFileExtension replaces the extension of a filename with a new suffix
func ReplaceFileExtension(filename, newSuffix string) string {
	base := strings.TrimSuffix(filename, filepath.Ext(filename))
	return base + newSuffix
}

// WriteCSV writes processed shot data to a CSV file
func WriteCSV(filename string, data []models.ProcessedShotData) error {
	if len(data) == 0 {
		return fmt.Errorf("no data to write")
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("creating file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	if err := writer.Write([]string{"Club", "Type", "Target", "Total", "Side"}); err != nil {
		return fmt.Errorf("writing header: %w", err)
	}

	// Write data rows
	for _, d := range data {
		record := []string{
			d.Club,
			d.Type,
			fmt.Sprintf("%.2f", d.Target),
			fmt.Sprintf("%.2f", d.Total),
			fmt.Sprintf("%.2f", d.Side),
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("writing record: %w", err)
		}
	}
	return nil
}
