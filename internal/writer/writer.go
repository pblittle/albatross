package writer

import "albatross/internal/models"

// Writer interface defines the method that any writer should implement
type Writer interface {
	Write(filename string, data []models.ProcessedShotData) error
}
