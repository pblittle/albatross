package processors

import (
	"testing"
)

func TestNormalizeClubType(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Driver", "driver", "Dr"},
		{"D", "d", "Dr"},
		{"DR", "dr", "Dr"},
		{"3 Wood", "3w", "3W"},
		{"5 Wood", "5Wood", "5W"},
		{"7 Wood", "7 wood", "7W"},
		{"4 Hybrid", "4h", "4Hy"},
		{"5 Hybrid", "5 Hybrid", "5Hy"},
		{"6 Hybrid", "6HY", "6Hy"},
		{"7 Iron", "7i", "7i"},
		{"8 Iron", "8Iron", "8i"},
		{"9 Iron", "9 iron", "9i"},
		{"Pitching Wedge", "pw", "Pw"},
		{"Sand Wedge", "SW", "Sw"},
		{"Gap Wedge", "Gap Wedge", "Gw"},
		{"Lob Wedge", "LobWedge", "Lw"},
		{"Putter", "putter", "Putter"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeClubType(tt.input)
			if result != tt.expected {
				t.Errorf("NormalizeClubType(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestDetermineShotType(t *testing.T) {
	tests := []struct {
		name     string
		clubType string
		expected string
	}{
		{"Driver", "Dr", "Tee"},
		{"3 Wood", "3W", "Tee"},
		{"5 Wood", "5W", "Tee"},
		{"3 Iron", "3i", "Approach"},
		{"7 Iron", "7i", "Approach"},
		{"Pitching Wedge", "Pw", "Approach"},
		{"Sand Wedge", "Sw", "Approach"},
		{"4 Hybrid", "4Hy", "Approach"},
		{"Putter", "Putter", "Approach"}, // Assuming default is "Approach"
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DetermineShotType(tt.clubType)
			if result != tt.expected {
				t.Errorf("DetermineShotType(%q) = %q, want %q", tt.clubType, result, tt.expected)
			}
		})
	}
}

func TestExtractNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"No number", "wood", ""},
		{"Number at start", "3wood", "3"},
		{"Number in middle", "iron7", "7"},
		{"Number at end", "hybrid4", "4"},
		{"Multiple numbers", "5iron9", "5"},
		{"No digits", "iron", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractNumber(tt.input)
			if result != tt.expected {
				t.Errorf("extractNumber(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCapitalizeFirst(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", ""},
		{"Single letter", "a", "A"},
		{"Multiple letters", "hello", "Hello"},
		{"Already capitalized", "World", "World"},
		{"With numbers", "7iron", "7iron"},
		{"With spaces", "gap wedge", "Gap wedge"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := capitalizeFirst(tt.input)
			if result != tt.expected {
				t.Errorf("capitalizeFirst(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkNormalizeClubType(b *testing.B) {
	clubTypes := []string{"driver", "3w", "5 Hybrid", "7 iron", "pw", "SandWedge"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, clubType := range clubTypes {
			NormalizeClubType(clubType)
		}
	}
}

func BenchmarkDetermineShotType(b *testing.B) {
	clubTypes := []string{"Dr", "3W", "5Hy", "7i", "Pw", "Sw"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, clubType := range clubTypes {
			DetermineShotType(clubType)
		}
	}
}
