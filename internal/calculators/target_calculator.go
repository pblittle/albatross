// Package calculators provides functions for performing calculations on shot data.
package calculators

import (
	"sort"

	"albatross/internal/models"
)

// CalculateTargets computes the median target distance for each club type
// and updates the Target field in each ProcessedShotData struct.
// This function modifies the input slice in-place.
func CalculateTargets(shotData *[]models.ProcessedShotData) {
	// Group shots by club type
	clubShots := make(map[string][]float64)
	for _, shot := range *shotData {
		clubShots[shot.Club] = append(clubShots[shot.Club], shot.Total)
	}

	// Calculate median for each club type
	for club, distances := range clubShots {
		median := calculateMedian(distances)

		// Update the Target field for all shots of this club type
		for i := range *shotData {
			if (*shotData)[i].Club == club {
				(*shotData)[i].Target = median
			}
		}
	}
}

// calculateMedian computes the median value from a slice of float64 numbers.
// It handles both odd and even-length slices.
// The function sorts the input slice and returns the middle value (or average of two middle values).
// If the input slice is empty, it returns 0.
func calculateMedian(numbers []float64) float64 {
	sort.Float64s(numbers)
	length := len(numbers)
	if length == 0 {
		return 0
	}
	if length%2 == 0 {
		// For even-length slices, average the two middle numbers
		return (numbers[length/2-1] + numbers[length/2]) / 2
	}
	// For odd-length slices, return the middle number
	return numbers[length/2]
}
