package processors

import (
	"regexp"
	"strings"
	"unicode"
)

// NormalizeClubType standardizes the club type notation.
// It converts various input formats to a consistent output format.
// For example: "driver" -> "Dr", "3 wood" -> "3W", "7 iron" -> "7i"
func NormalizeClubType(clubType string) string {
	clubType = strings.TrimSpace(strings.ToLower(clubType))

	// Handle driver variations
	if clubType == "driver" || clubType == "d" || clubType == "dr" {
		return "Dr"
	}

	// Handle woods
	if strings.Contains(clubType, "w") || strings.Contains(clubType, "wood") {
		number := extractNumber(clubType)
		if number != "" {
			return number + "W"
		}
	}

	// Handle hybrids
	if strings.Contains(clubType, "h") || strings.Contains(clubType, "hybrid") || strings.Contains(clubType, "hy") {
		number := extractNumber(clubType)
		if number != "" {
			return number + "Hy"
		}
	}

	// Handle irons
	if strings.Contains(clubType, "i") || strings.Contains(clubType, "iron") {
		number := extractNumber(clubType)
		if number != "" {
			return number + "i"
		}
	}

	// Handle wedges
	wedgeMap := map[string]string{
		"pw": "Pw", "pitching": "Pw",
		"sw": "Sw", "sand": "Sw",
		"gw": "Gw", "gap": "Gw",
		"lw": "Lw", "lob": "Lw",
	}
	for key, value := range wedgeMap {
		if strings.Contains(clubType, key) {
			return value
		}
	}

	// If no specific rule applies, capitalize the first letter
	return capitalizeFirst(clubType)
}

// DetermineShotType categorizes the shot based on the club type.
// It returns either "Tee" or "Approach" depending on the club used.
func DetermineShotType(clubType string) string {
	clubTypeMap := map[*regexp.Regexp]string{
		regexp.MustCompile(`(?i)^d`):                     "Tee",
		regexp.MustCompile(`(?i)^[2-9]w(ood)?$`):         "Tee",
		regexp.MustCompile(`(?i)^\d+i(ron)?$`):           "Approach",
		regexp.MustCompile(`(?i)^[psgl]w(edge)?$`):       "Approach",
		regexp.MustCompile(`(?i)^\d+\s*(h(ybrid)?|hy)$`): "Approach",
	}

	for pattern, shotType := range clubTypeMap {
		if pattern.MatchString(clubType) {
			return shotType
		}
	}
	return "Approach" // Default to Approach if no match is found
}

// extractNumber retrieves the first numeric value from a string.
// For example, "3wood" would return "3".
func extractNumber(s string) string {
	re := regexp.MustCompile(`\d+`)
	return re.FindString(s)
}

// capitalizeFirst capitalizes the first letter of a string.
// It's used as a fallback for club types that don't match any specific rules.
func capitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
