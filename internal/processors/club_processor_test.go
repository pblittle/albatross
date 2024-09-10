package processors

import (
	"testing"
)

func TestNormalizeClubType(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"driver", "Dr"},
		{"d", "Dr"},
		{"dr", "Dr"},
		{"3w", "3W"},
		{"5Wood", "5W"},
		{"7 wood", "7W"},
		{"4h", "4Hy"},
		{"5 Hybrid", "5Hy"},
		{"6HY", "6Hy"},
		{"7i", "7i"},
		{"8Iron", "8i"},
		{"9 iron", "9i"},
		{"pw", "Pw"},
		{"SW", "Sw"},
		{"Gap Wedge", "Gw"},
		{"LobWedge", "Lw"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := NormalizeClubType(tt.input)
			if result != tt.expected {
				t.Errorf("NormalizeClubType(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestDetermineShotType(t *testing.T) {
	tests := []struct {
		clubType string
		expected string
	}{
		{"Dr", "Tee"},
		{"3W", "Tee"},
		{"5W", "Tee"},
		{"3i", "Approach"},
		{"7i", "Approach"},
		{"Pw", "Approach"},
		{"Sw", "Approach"},
		{"4Hy", "Approach"},
		{"Putter", "Approach"}, // Assuming default is "Approach"
	}

	for _, tt := range tests {
		t.Run(tt.clubType, func(t *testing.T) {
			result := DetermineShotType(tt.clubType)
			if result != tt.expected {
				t.Errorf("DetermineShotType(%q) = %q, want %q", tt.clubType, result, tt.expected)
			}
		})
	}
}
