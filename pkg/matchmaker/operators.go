package matchmaker

import "fmt"

func regexOperator(chars string, op string) MatchOption {
	list := chars + op + "%s"
	return func(m *Matcher) {
		m.regex = fmt.Sprintf(m.regex, list)
	}
}

// Exactly onw occurance of a char in charClass.
//
// Example:
//   // "[abc]"
//   One(ToCharClass([]string{"a", "b", "c"}))
func One(charClass string) MatchOption {
  return regexOperator(charClass, "")
}

// n or m occurances of char in charClass, prefer more (Greedy).
// Equivalent to the regex string representation "statement{n,m}".
//
// Example:
//   // "[[:digit:]]{2,4}"
//   NOrMPreferMore(Digit, 2, 4)
func NOrMPreferMore(charClass string, n int, m int) MatchOption {
  // TODO: assert m>n
	return regexOperator(charClass, fmt.Sprintf("{%d,%d}", n, m))
}

// n or m occurances of char in charClass, prefer fewer (Non-Greedy).
// Equivalent to the regex string representation "statement{n,m}?".
//
// Example:
//   // "[[:digit:]]{2,4}?"
//   NOrMPreferFewer(Digit, 2, 4)
func NOrMPreferFewer(charClass string, n int, m int) MatchOption {
  // TODO: assert m>n
	return regexOperator(charClass, fmt.Sprintf("{%d,%d}?", n, m))
}

// Exactly n occurances of char in charClass.
// Equivalent to the regex string representation "statement{n}".
//
// Example:
//   // "[[:alnum:]]{4}"
//   ExactlyN(AlphaNumeric, 4)
func ExactlyN(charClass string, n int) MatchOption {
	return regexOperator(charClass, fmt.Sprintf("{%d}", n))
}

// n occurances of char in charClass or more, prefer more (Greedy).
// Equivalent to the regex string representation "statement{n,}".
//
// Example:
//   // "[[:alnum:]]{4,}"
//   NOrMorePreferMore(AlphaNumeric, 4)
func NOrMorePreferMore(charClass string, n int) MatchOption {
	return regexOperator(charClass, fmt.Sprintf("{%d,}", n))
}

// n occurances of char in charClass or more, prefer fewer (Non-Greedy).
// Equivalent to the regex string representation "statement{n,}?".
//
// Example:
//   // "[[:alnum:]]{4,}?"
//   NOrMorePreferFewer(AlphaNumeric, 4)
func NOrMorePreferFewer(charClass string, n int) MatchOption {
	return regexOperator(charClass, fmt.Sprintf("{%d,}?", n))
}

// Zero or one char in charClass, prefer one (Greedy).
// Equivalent to the regex string representation "statement?".
//
// Example:
//   // "[[:alnum:]]?"
//   ZeroOrOnePreferOne(AlphaNumeric)
func ZeroOrOnePreferOne(charClass string) MatchOption {
	return regexOperator(charClass, "?")
}

// Zero or one char in charClass, prefer zero (Non-Greedy).
// Equivalent to the regex string representation "statement??".
//
// Example:
//   // "[[:alnum:]]??"
//   ZeroOrOnePreferZero(AlphaNumeric)
func ZeroOrOnePreferZero(charClass string) MatchOption {
	return regexOperator(charClass, "??")
}

// Zero or more char in charClass, prefer more (Greedy).
// Equivalent to the regex string representation "statement*".
//
// Example:
//   // "[[:alnum:]]*"
//   ZeroOrMorePreferOne(AlphaNumeric)
func ZeroOrMorePreferMore(charClass string) MatchOption {
	return regexOperator(charClass, "*")
}

// Zero or more char in charClass, prefer fewer (Non-Greedy).
// Equivalent to the regex string representation "statement*?".
//
// Example:
//   // "[[:alnum:]]*?"
//   ZeroOrMorePreferFewer(AlphaNumeric)
func ZeroOrMorePreferFewer(charClass string) MatchOption {
	return regexOperator(charClass, "*?")
}

// One or more char in charClass, prefer more (Greedy).
// Equivalent to the regex string representation "statement+".
//
// Example:
//   // "[[:alnum:]]+"
//   OneOrMorePreferMore(AlphaNumeric)
func OneOrMorePreferMore(chars string) MatchOption {
	return regexOperator(chars, "+")
}

// One or more char in charClass, prefer fewer (Non-Greedy).
// Equivalent to the regex string representation "statement+?".
//
// Example:
//   // "[[:alnum:]]+?"
//   OneOrMorePreferFewer(AlphaNumeric)
func OneOrMorePreferFewer(chars string) MatchOption {
	return regexOperator(chars, "+?")
}
