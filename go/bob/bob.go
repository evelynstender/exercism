/*
Package bob implements what bob would say to you
*/
package bob

import (
	"regexp"
	"strings"
)

var atLeastOneAlphaCharacter = regexp.MustCompile(`.*[a-zA-Z]+.*`)

// Hey shows bob's response to your stupid questions
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	switch true {
	case strings.TrimSpace(remark) == "":
		return "Fine. Be that way!"
	case !atLeastOneAlphaCharacter.MatchString(remark) && !strings.Contains(strings.ToUpper(remark), remark):
		return "Whatever."
	case strings.HasSuffix(remark, "?") && strings.Contains(strings.ToUpper(remark), remark) && atLeastOneAlphaCharacter.MatchString(remark):
		return "Calm down, I know what I'm doing!"
	case strings.HasSuffix(remark, "?"):
		return "Sure."
	case strings.ToUpper(remark) == remark && atLeastOneAlphaCharacter.MatchString(remark):
		return "Whoa, chill out!"

	default:
		return "Whatever."
	}
}
