package calculators

import (
	"reflect"
	"testing"

	"albatross/internal/models"
)

func TestCalculateTargets(t *testing.T) {
	tests := []struct {
		name     string
		input    []models.ProcessedShotData
		expected []models.ProcessedShotData
	}{
		{
			name: "Single club type",
			input: []models.ProcessedShotData{
				{Club: "7i", Total: 150},
				{Club: "7i", Total: 160},
				{Club: "7i", Total: 155},
			},
			expected: []models.ProcessedShotData{
				{Club: "7i", Total: 150, Target: 155},
				{Club: "7i", Total: 160, Target: 155},
				{Club: "7i", Total: 155, Target: 155},
			},
		},
		{
			name: "Multiple club types",
			input: []models.ProcessedShotData{
				{Club: "Dr", Total: 250},
				{Club: "Dr", Total: 260},
				{Club: "3W", Total: 220},
				{Club: "3W", Total: 230},
				{Club: "7i", Total: 150},
				{Club: "7i", Total: 160},
				{Club: "7i", Total: 155},
			},
			expected: []models.ProcessedShotData{
				{Club: "Dr", Total: 250, Target: 255},
				{Club: "Dr", Total: 260, Target: 255},
				{Club: "3W", Total: 220, Target: 225},
				{Club: "3W", Total: 230, Target: 225},
				{Club: "7i", Total: 150, Target: 155},
				{Club: "7i", Total: 160, Target: 155},
				{Club: "7i", Total: 155, Target: 155},
			},
		},
		{
			name: "Even number of shots",
			input: []models.ProcessedShotData{
				{Club: "5i", Total: 180},
				{Club: "5i", Total: 185},
				{Club: "5i", Total: 190},
				{Club: "5i", Total: 195},
			},
			expected: []models.ProcessedShotData{
				{Club: "5i", Total: 180, Target: 187.5},
				{Club: "5i", Total: 185, Target: 187.5},
				{Club: "5i", Total: 190, Target: 187.5},
				{Club: "5i", Total: 195, Target: 187.5},
			},
		},
		{
			name: "Hybrid and wedge clubs",
			input: []models.ProcessedShotData{
				{Club: "4Hy", Total: 200},
				{Club: "4Hy", Total: 210},
				{Club: "4Hy", Total: 205},
				{Club: "Pw", Total: 120},
				{Club: "Pw", Total: 125},
				{Club: "Sw", Total: 100},
				{Club: "Sw", Total: 105},
			},
			expected: []models.ProcessedShotData{
				{Club: "4Hy", Total: 200, Target: 205},
				{Club: "4Hy", Total: 210, Target: 205},
				{Club: "4Hy", Total: 205, Target: 205},
				{Club: "Pw", Total: 120, Target: 122.5},
				{Club: "Pw", Total: 125, Target: 122.5},
				{Club: "Sw", Total: 100, Target: 102.5},
				{Club: "Sw", Total: 105, Target: 102.5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CalculateTargets(&tt.input)
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("CalculateTargets() = %v, want %v", tt.input, tt.expected)
			}
		})
	}
}

func TestCalculateMedian(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		expected float64
	}{
		{"Empty slice", []float64{}, 0},
		{"Single element", []float64{5}, 5},
		{"Odd number of elements", []float64{1, 3, 5}, 3},
		{"Even number of elements", []float64{1, 3, 5, 7}, 4},
		{"Unsorted slice", []float64{5, 2, 8, 1, 9}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateMedian(tt.input)
			if result != tt.expected {
				t.Errorf("calculateMedian(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
