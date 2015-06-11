package gstrings

import (
	"strings"
	"unicode"
)

// Rstrip returns a copy of s
// with trailing whitespace removed.
// See also gstrings.Lstrip and gstrings.Strip
func Rstrip(s string) string {
	return strings.TrimRightFunc(s, unicode.IsSpace)
}