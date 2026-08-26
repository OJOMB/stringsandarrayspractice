package minimumsizesubarraysum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSubArrayLen(t *testing.T) {
	type testCase struct {
		inputNums      []int
		inputTarget    int
		expectedOutput int
	}

	testCases := []testCase{
		{
			[]int{2, 3, 1, 2, 4, 3},
			7,
			2,
		},
		{
			[]int{1, 4, 4},
			4,
			1,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test: %d", i), func(t *testing.T) {
			output := minSubArrayLen(tc.inputTarget, tc.inputNums)
			assert.Equal(t, tc.expectedOutput, output)
		})
	}
}
