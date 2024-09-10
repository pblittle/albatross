package reader

import (
	"reflect"
	"testing"

	"albatross/internal/models"
)

func TestMLM2ProLaunchMonitorParseRow(t *testing.T) {
	launchMonitor := MLM2ProLaunchMonitor{}
	headers := []string{"club type", "total distance", "side carry"}
	row := []string{"Dr", "250", "10"}

	expected := models.RawShotData{
		LaunchMonitorType: "MLM2Pro",
		Data: map[string]string{
			"club type":      "Dr",
			"total distance": "250",
			"side carry":     "10",
		},
	}

	result, err := launchMonitor.ParseRow(row, headers)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ParseRow() = %v, want %v", result, expected)
	}
}

func TestMLM2ProLaunchMonitorProcessRawData(t *testing.T) {
	launchMonitor := MLM2ProLaunchMonitor{}
	rawData := models.RawShotData{
		LaunchMonitorType: "MLM2Pro",
		Data: map[string]string{
			"club type":      "Dr",
			"total distance": "250",
			"side carry":     "10",
		},
	}

	expected := models.ProcessedShotData{
		Club:  "Dr",
		Type:  "Tee",
		Total: 250,
		Side:  10,
	}

	result := launchMonitor.ProcessRawData(rawData)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ProcessRawData() = %v, want %v", result, expected)
	}
}
