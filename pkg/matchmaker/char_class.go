package matchmaker

import (
	"fmt"
	"strings"
)

const (
	Any                 = `.`
	AlphaNumeric        = `[[:alnum:]]`
	Alphabetic          = `[[:alpha:]]`
	AlphabeticLowerCase = `[[:lower:]]`
	AlphabeticUpperCase = `[[:upper:]]`
	ASCII               = `[[:ascii:]]`
	Digit               = `[[:digit:]]`
	HexDigit            = `[[:xdigit:]]`
	NoDigit             = `\D`
	Whitespace          = `[[:space:]]`
	NoWhitespace        = `\S`
	Word                = `[[:word:]]`
	NoWord              = `\W`
)

// Converts a list of characters to a negated char class
//
// Example:
//   // "[^abc]"
//   ToCharClass([]string{"a", "b", "c"})
func ToNegatedCharClass(chars []string) string {
	if len(chars) == 1 {
		return chars[0]
	}
	return fmt.Sprintf("[^%s]", strings.Join(chars, ""))
}

// Converts a list of characters to a char class
//
// Example:
//   // "[abc]"
//   ToCharClass([]string{"a", "b", "c"})
func ToCharClass(chars []string) string {
	if len(chars) == 1 {
		return chars[0]
	}
	return fmt.Sprintf("[%s]", strings.Join(chars, ""))
}
