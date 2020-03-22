package strings

import (
	"regexp"
	"strings"
)

var space = regexp.MustCompile(`\s+`)

// minifyString removes leading, trailing and extra spaces in the str.
func minifyString(str string) string {
	str = space.ReplaceAllString(str, " ")
	str = strings.TrimSpace(str)
	return str
}

func Minify(str string) string {
	return minifyString(str)
}
