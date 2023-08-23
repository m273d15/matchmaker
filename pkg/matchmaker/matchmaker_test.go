package matchmaker_test

import (
	"fmt"
	mm "github.com/m273d15/matchmaker/pkg/matchmaker"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	doesMatch bool
	value     string
	options   []mm.MatchOption
}

func TestCharacterClasses(t *testing.T) {
	tData := []testData{
		// Digit
		{
			doesMatch: true,
			value:     "123",
			options: []mm.MatchOption{
				mm.ZeroOrMorePreferMore(mm.Digit),
			},
		},
		{
			doesMatch: false,
			value:     "aa",
			options: []mm.MatchOption{
				mm.ZeroOrMorePreferMore(mm.Digit),
			},
		},
		// NoDigit
		{
			doesMatch: true,
			value:     "ab-)",
			options: []mm.MatchOption{
				mm.ZeroOrMorePreferMore(mm.NoDigit),
			},
		},
		{
			doesMatch: false,
			value:     "a1a",
			options: []mm.MatchOption{
				mm.ZeroOrMorePreferMore(mm.NoDigit),
			},
		},
		// Whitespace
		{
			doesMatch: true,
			value:     " ", // space
			options: []mm.MatchOption{
				mm.ZeroOrMorePreferMore(mm.Whitespace),
			},
		},
		{
			doesMatch: true,
			value:     "\t", // tab
			options: []mm.MatchOption{
				mm.ZeroOrMorePreferMore(mm.Whitespace),
			},
		},
		{
			doesMatch: false,
			value:     "a",
			options: []mm.MatchOption{
				mm.ZeroOrMorePreferMore(mm.Whitespace),
			},
		},
		// NoWhitespace
		{
			doesMatch: false,
			value:     " ", // space
			options: []mm.MatchOption{
				mm.ZeroOrMorePreferMore(mm.NoWhitespace),
			},
		},
		{
			doesMatch: false,
			value:     "\t", // tab
			options: []mm.MatchOption{
				mm.ZeroOrMorePreferMore(mm.NoWhitespace),
			},
		},
		{
			doesMatch: true,
			value:     "a",
			options: []mm.MatchOption{
				mm.ZeroOrMorePreferMore(mm.NoWhitespace),
			},
		},
	}

	matchExactlyTest(t, tData)
}

func TestMatchExactlyZeroOrMore(t *testing.T) {
	tData := []testData{
		{
			doesMatch: true,
			value:     "123",
			options: []mm.MatchOption{
				mm.ZeroOrMorePreferMore(mm.Digit),
			},
		},
	}

	matchExactlyTest(t, tData)
}

func TestMatchExactlyGroup(t *testing.T) {
  m := mm.MatchExactly(
        mm.OneOrMorePreferMore(mm.Digit),
        mm.Group(
          mm.OneOrMorePreferMore(mm.Alphabetic),
        ),
        mm.OneOrMorePreferMore(mm.Digit),
  )
  match := m.MatchGroups("123abc321")

  assert.Len(t, match, 2)
  assert.Equal(t, "abc", match[1])
}

func matchExactlyTest(t *testing.T, tData []testData) {
	t.Parallel()
	for _, data := range tData {
		t.Run(fmt.Sprintf("Match %s", data.value), func(tt *testing.T) {
			m := mm.MatchExactly(data.options...)
			matched := m.MatchString(data.value)
			assert.Equal(tt, data.doesMatch, matched)
		})
	}
}

// Matchmaker equivalent to regex representations "^[[:digit:]]+([[:alpha:]]+)[[:digit:]]*$"
func Example() {
  m := mm.MatchExactly(
        mm.OneOrMorePreferMore(mm.Digit),
        mm.Group(
          mm.OneOrMorePreferMore(mm.Alphabetic),
        ),
        mm.ZeroOrMorePreferMore(mm.Digit),
  )
  match := m.MatchGroups("123abc321")
  fmt.Println(match[1])
  // Output: abc

}
