package porcelain_test

import (
	"github.com/m273d15/matchmaker/pkg/matchmaker"
	"github.com/m273d15/matchmaker/pkg/porcelain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractGreedy(t *testing.T) {
  extracted, err := porcelain.Extract("123abc321",
    matchmaker.OneOrMorePreferMore(matchmaker.Alphabetic),
  )
  assert.NoError(t, err)
  assert.Equal(t, "abc", extracted)
}

func TestExtractNonGreedy(t *testing.T) {
  extracted, err := porcelain.Extract("123abc321",
    matchmaker.OneOrMorePreferFewer(matchmaker.Alphabetic),
  )
  assert.NoError(t, err)
  assert.Equal(t, "a", extracted)
}
