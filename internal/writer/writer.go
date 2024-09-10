// Package writer provides interfaces and implementations for writing processed shot data to various output formats.
package writer

import "albatross/internal/models"

// Writer interface defines the method that any writer should implement.
// This interface allows for different output formats to be used interchangeably,
// following the Strategy pattern.
type Writer interface {
	// Write takes a filename and a slice of ProcessedShotData, and writes the data to the specified file.
	// The exact format of the output is determined by the specific implementation of the Writer interface.
	// It returns an error if the writing process encounters any issues.
	Write(filename string, data []models.ProcessedShotData) error
}
