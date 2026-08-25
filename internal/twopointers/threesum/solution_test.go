package threesum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	type testCase struct {
		input          []int
		expectedOutput [][]int
	}

	testCases := []testCase{
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			output := threeSum(tc.input)
			assert.Equal(t, tc.expectedOutput, output)
		})
	}
}
