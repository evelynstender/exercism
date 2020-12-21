// Package twofer calculates the sharing
package twofer

import (
	"fmt"
	"strings"
)

// ShareWith returns the sharing
func ShareWith(name string) string {
	if strings.Trim(name, " ") != "" {
		return fmt.Sprint("One for ", name, ", one for me.")
	}
	return "One for you, one for me."
}
