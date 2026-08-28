package maxsumsubarrayofsizek

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumSubarraySum(t *testing.T) {
	type testCase struct {
		inputNums      []int
		inputK         int
		expectedOutput int
	}

	testCases := []testCase{
		{
			[]int{2, 1, 5, 1, 3, 2},
			3,
			9,
		},
		{
			[]int{100, 200, 300, 400},
			1,
			400,
		},
		{
			[]int{100, 200, 300, 400},
			2,
			700,
		},
		{
			[]int{-100, -200, -300, -400},
			2,
			-300,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test: %d", i), func(t *testing.T) {
			output := maximumSubarraySum(tc.inputNums, tc.inputK)
			assert.Equal(t, tc.expectedOutput, output)
		})
	}
}
