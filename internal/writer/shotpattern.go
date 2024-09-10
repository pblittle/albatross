package writer

import (
	"fmt"

	"albatross/internal/models"
	"albatross/utils"
)

type ShotPatternWriter struct{}

func (w ShotPatternWriter) Write(filename string, data []models.ProcessedShotData) error {
	if len(data) == 0 {
		return fmt.Errorf("no data to write")
	}
	return utils.WriteCSV(filename, data)
}
