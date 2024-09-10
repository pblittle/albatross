package models

import (
	"testing"
)

// MockLaunchMonitor implements the LaunchMonitor interface for testing
type MockLaunchMonitor struct{}

func (m MockLaunchMonitor) ParseRow(row []string, headers []string) (RawShotData, error) {
	return RawShotData{
		LaunchMonitorType: "Mock",
		Data: map[string]string{
			"club type":      row[0],
			"total distance": row[1],
			"side carry":     row[2],
		},
	}, nil
}

func (m MockLaunchMonitor) ProcessRawData(rawData RawShotData) ProcessedShotData {
	return ProcessedShotData{
		Club:  rawData.Data["club type"],
		Type:  "Test",
		Total: 0, // For simplicity, we're not converting strings to floats in this mock
		Side:  0,
	}
}

func TestLaunchMonitorInterface(t *testing.T) {
	var _ LaunchMonitor = (*MockLaunchMonitor)(nil)

	mock := MockLaunchMonitor{}
	headers := []string{"club type", "total distance", "side carry"}
	row := []string{"Driver", "250", "10"}

	rawData, err := mock.ParseRow(row, headers)
	if err != nil {
		t.Errorf("ParseRow returned unexpected error: %v", err)
	}

	if rawData.LaunchMonitorType != "Mock" {
		t.Errorf("Expected LaunchMonitorType to be 'Mock', got %s", rawData.LaunchMonitorType)
	}

	processedData := mock.ProcessRawData(rawData)
	if processedData.Club != "Driver" {
		t.Errorf("Expected Club to be 'Driver', got %s", processedData.Club)
	}
	if processedData.Type != "Test" {
		t.Errorf("Expected Type to be 'Test', got %s", processedData.Type)
	}
}
