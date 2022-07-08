package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	recommendation, _ := utf8.DecodeRuneInString("❗")
	search, _ := utf8.DecodeLastRuneInString("🔍")
	weather, _ := utf8.DecodeLastRuneInString("☀")

	for _, value := range log {

		switch value {
		case recommendation:
			return "recommendation"
		case search:
			return "search"
		case weather:
			return "weather"
		default:
			continue

		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runes := []rune{}

	for _, value := range log {
		if value == oldRune {
			runes = append(runes, newRune)
		} else {
			runes = append(runes, value)
		}
	}

	return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
