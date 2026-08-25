package twosum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	type testCase struct {
		inputNums      []int
		inputTarget    int
		expectedOutput []int
	}

	testCases := []testCase{
		{
			[]int{3, 2, 4},
			6,
			[]int{1, 2},
		},
		{
			[]int{0, 4, 3, 0},
			0,
			[]int{0, 3},
		},
		{
			[]int{0, 4, 3, 0},
			12,
			nil,
		},
		{
			[]int{-1, -2, -3, -4, -5},
			-8,
			[]int{2, 4},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := TwoSum(tc.inputNums, tc.inputTarget)
			assert.Equal(t, tc.expectedOutput, result)
		})
	}
}
