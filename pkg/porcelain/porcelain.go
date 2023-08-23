// Following the git plumbing and porcelain design pattern (https://git-scm.com/book/en/v2/Git-Internals-Plumbing-and-Porcelain), this package provides
// high-level methods leveraging the low-level matcher functions defined by matchmaker
package porcelain

import (
	"errors"
	"github.com/m273d15/matchmaker/pkg/matchmaker"
)

// Extracts a substring from a string.
//
// Example:
//   extracted, _ := porcelain.Extract("123abc321",
//      matchmaker.OneOrMorePreferMore(matchmaker.Alphabetic),
//    )
//   fmt.Println(extracted)
//   // Output: "abc"

func Extract(whereToExtract string, whatToExtract ...matchmaker.MatchOption) (string, error) {
  if len(whatToExtract) == 0 {
    return "", errors.New("At least one match option is needed")
  }
  m := matchmaker.MatchSubstring(
    whatToExtract...
  )
  g := m.MatchGroups(whereToExtract)
  if len(g) == 0 {
    return "", errors.New("No match to extract")
  }
  return g[0], nil
}
