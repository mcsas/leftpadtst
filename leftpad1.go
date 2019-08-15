//This code starts wtih leftpad instead of main because it is not intended as
// the starting point of a program but a library used by other code
// matches the name of the directory that contains the package

package leftpad

import (
	"strings"
	"unicode/utf8"
)

// var at the pkg level - scoped to a pkg
// same as inside a func
// visible to everything within pkg
// cannot use := to declare them

var default_char = ' '

// Single line comment before ea. func is good practice

// Format takes in a string and an int and returns the string
// left-padded with spaces to the length of the int. If the
// string is already longer than the specified length, the
// original string is returned.

// Two levels of visibiliyt in Go
// Package Visible and Public (applies to :vars, funcs, and data types)
//   Public - can be accessed by other pkgs. (capitalize first letter)
//   Pkg visible - only available within the pkg (lowercase first letter)
//

func Format(s string, size int) string {
	return FormatRune(s, size, default_char)
}

// FormatRune takes in a string, an int, and a rune and returns the string
// left-padded with the specifed rune to the length of the int. If the
// string is already longer than the specified length, the
// original string is returned.
func FormatRune(s string, size int, r rune) string {
	l := utf8.RuneCountInString(s)
	if l >= size {
		return s
	}
	out := strings.Repeat(string(r), size-l) + s
	return out
}
