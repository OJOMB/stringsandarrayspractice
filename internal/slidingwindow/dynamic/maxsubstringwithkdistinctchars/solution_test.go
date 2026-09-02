package maxsubstringwithkdistinctchars

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstringKDistinct(t *testing.T) {
	testcases := []struct {
		s      string
		k      int
		expect int
	}{
		{"", 2, 0},
		{"a", 0, 0},
		{"a", 1, 1},
		{"eceba", 2, 3},
		{"aa", 1, 2},
		{"abcabcabc", 2, 2},
		{"abaccc", 2, 4},
	}

	for _, tc := range testcases {
		result := lengthOfLongestSubstringKDistinct(tc.s, tc.k)
		assert.Equal(t, tc.expect, result)
	}
}
