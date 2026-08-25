package maxvowelsinsubstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxAverage(t *testing.T) {
	type testCase struct {
		input          string
		k              int
		expectedOutput int
	}

	testCases := []testCase{
		{
			"leetcode",
			3,
			2,
		},
		{
			"abciiiidef",
			4,
			4,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := maxVowels(tc.input, tc.k)
			assert.Equal(t, tc.expectedOutput, result)
		})
	}

}
