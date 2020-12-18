// This is a "stub" file.  It's a little start on your solution.
// Convert a phrase to its acronym.
package acronym

import (
	"strings"
)

// Abbreviate converts a phrase to its acronym.
func Abbreviate(s string) string {
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")
	var splitValue = strings.Fields(s)
	var acronym string

	for i := 0; i < len(splitValue); i++ {
		acronym += strings.SplitN(splitValue[i], "", 2)[0]
	}
	return strings.ToUpper(acronym)
}
