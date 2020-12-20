// Package acronym converts a phrase to its acronym
package acronym

import (
	"strings"
)

// Abbreviate is the main package function
func Abbreviate(s string) string {
	var abbreviation string
	replacements := strings.NewReplacer("-", " ", "_", " ")
	s = replacements.Replace(s)
	words := strings.Split(s, " ")

	for _, word := range words {
		if word != "" {
			abbreviation += strings.ToUpper(string(word[0]))
		}
	}

	return abbreviation
}
