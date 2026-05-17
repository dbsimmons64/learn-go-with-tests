package iterations

import (
	"strings"
)

// Repeat returns a string containing the character repeated cycle times.
func Repeat(character string, cycles int) string {
	var repeated strings.Builder

	for range cycles {
		repeated.WriteString(character)
	}
	return repeated.String()
}
