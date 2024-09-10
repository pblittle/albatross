package models

// LaunchMonitor interface defines the methods that any launch monitor type should implement
type LaunchMonitor interface {
	// ParseRow converts a row of strings into a RawShotData struct
	ParseRow(row []string, headers []string) (RawShotData, error)
	// ProcessRawData converts RawShotData into ProcessedShotData
	ProcessRawData(rawData RawShotData) ProcessedShotData
}

// RawShotData represents the raw data from any launch monitor
type RawShotData struct {
	LaunchMonitorType string
	Data              map[string]string
}

// ProcessedShotData represents the standardized processed shot data
type ProcessedShotData struct {
	Apex   float64 // The highest point of the shot's trajectory
	Carry  float64 // The carry distance of the shot
	Club   string  // The type of club used for the shot
	Roll   string  // The difference between total and carry distance
	Type   string  // The type of shot (e.g., "Tee" or "Approach")
	Target float64 // The target distance for this club type
	Total  float64 // The total distance of the shot
	Side   float64 // The side carry (lateral deviation) of the shot
}
