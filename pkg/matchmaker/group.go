package matchmaker

import "fmt"

func regexGroup(modificator string, opts ...MatchOption) MatchOption {
	return func(m *Matcher) {
    innerMatcher := &Matcher{regex: fmt.Sprintf("(%s)", modificator + "%s")}
		for _, opt := range opts {
			opt(innerMatcher)
		}

	  innerMatcher.regex = fmt.Sprintf(innerMatcher.regex, "")
    m.regex = fmt.Sprintf(m.regex, innerMatcher.regex + "%s")
	}
}

// Group statements that should be returned together in a match.
// Equivalent to the regex string representation "(statements...)".
//
// Example:
//    // "([ab])"
//    Group(
//      One(ToCharClass([]string{"a", "b"}))
//    )
func Group(opts ...MatchOption) MatchOption {
	return regexGroup("", opts...)
}

// Group statements that should not be returned together in a match.
// Equivalent to the regex string representation "(?:statements...)".
//
// Example:
//    // "(?:[ab])"
//    NamedGroup(
//      One(ToCharClass([]string{"a", "b"}))
//    )
func NonCapturingGroup(opts ...MatchOption) MatchOption {
  return regexGroup("?:", opts...)
}

// Group statements that should be returned together in a match identified by a name..
// Equivalent to the regex string representation "(?<name>statements...)".
//
// Example:
//    // "(?<name>[ab])"
//    NamedGroup(
//      One(ToCharClass([]string{"a", "b"}))
//    )
func NamedGroup(name string, opts ...MatchOption) MatchOption {
  modificator := fmt.Sprintf("?<%s>", name)
	return regexGroup(modificator, opts...)
}
