package matchmaker

import (
	"fmt"
	"regexp"
)

const (
	matchSomewhere = "%s"
	matchStart     = "^%s"
	matchEnd       = "%s$"
	matchExactly   = "^%s$"
)

type Matcher struct {
	regex string
}
type MatchOption func(*Matcher)

func (m *Matcher) debugRegex() string {
	return m.regex
}

// Return true if s matches the defined [matchmaker.Matcher].
func (m *Matcher) MatchString(s string) bool {
	matched, _ := regexp.MatchString(m.regex, s)
	return matched
}

// Return matched groups of s based on defined [matchmaker.Matcher].
func (m *Matcher) MatchGroups(s string) []string {
  re, _ := regexp.Compile(m.regex)
  return re.FindStringSubmatch(s)
}

func match(initialRegex string, opts ...MatchOption) *Matcher {
	m := &Matcher{
		regex: initialRegex,
	}

	for _, opt := range opts {
		opt(m)
	}
	m.regex = fmt.Sprintf(m.regex, "")
	return m
}

// Define a matcher that defines the whole string to match.
// Equivalent to the regex string representation "^statements$".
//
// Example:
//   // "^[[:digit:]]$"
//   MatchExactly(
//     One(Digit)
//   )
func MatchExactly(opts ...MatchOption) *Matcher {
	return match(matchExactly, opts...)
}

// Define a matcher that defines the prefix/start of a string to match.
// Equivalent to the regex string representation "^statements".
//
// Example:
//   // "^[[:digit:]]"
//   MatchPrefix(
//     One(Digit)
//   )
func MatchPrefix(opts ...MatchOption) *Matcher {
	return match(matchStart, opts...)
}

// Define a matcher that defines the postfix/end of a string to match.
// Equivalent to the regex string representation "statements$".
//
// Example:
//   // "[[:digit:]]$"
//   MatchPostfix(
//     One(Digit)
//   )
func MatchPostfix(opts ...MatchOption) *Matcher {
	return match(matchEnd, opts...)
}

// Define a matcher that defines a substring of a string to match.
// Equivalent to the regex string representation "statements".
//
// Example:
//   // "[[:digit:]]"
//   MatchSubstring(
//     One(Digit)
//   )
func MatchSubstring(opts ...MatchOption) *Matcher {
	return match(matchSomewhere, opts...)
}
