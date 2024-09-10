package reader

import (
	"fmt"
	"strconv"
	"strings"

	"albatross/internal/models"
	"albatross/internal/processors"
)

// MLM2ProLaunchMonitor implements the LaunchMonitor interface for the MLM2Pro launch monitor
type MLM2ProLaunchMonitor struct{}

// NewMLM2ProLaunchMonitor creates and returns a new MLM2ProLaunchMonitor instance
func NewMLM2ProLaunchMonitor() models.LaunchMonitor {
	return &MLM2ProLaunchMonitor{}
}

// ParseRow converts a row of strings into a RawShotData struct for MLM2Pro data
func (launchMonitor MLM2ProLaunchMonitor) ParseRow(row []string, headers []string) (models.RawShotData, error) {
	if len(row) < len(headers) {
		return models.RawShotData{}, fmt.Errorf("insufficient columns")
	}

	data := make(map[string]string)
	for i, header := range headers {
		data[header] = strings.TrimSpace(row[i])
	}

	return models.RawShotData{
		LaunchMonitorType: "MLM2Pro",
		Data:              data,
	}, nil
}

// ProcessRawData converts RawShotData into ProcessedShotData for MLM2Pro data
func (launchMonitor MLM2ProLaunchMonitor) ProcessRawData(rawData models.RawShotData) models.ProcessedShotData {
	clubType := rawData.Data["club type"]
	totalDistance, _ := strconv.ParseFloat(rawData.Data["total distance"], 64)
	sideCarry, _ := strconv.ParseFloat(rawData.Data["side carry"], 64)

	normalizedClub := processors.NormalizeClubType(clubType)
	shotType := processors.DetermineShotType(normalizedClub)

	processed := models.ProcessedShotData{
		Club:  normalizedClub,
		Type:  shotType,
		Total: totalDistance,
		Side:  sideCarry,
	}

	return processed
}
