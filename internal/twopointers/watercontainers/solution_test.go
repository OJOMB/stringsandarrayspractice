package watercontainers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	type testCase struct {
		inputHeight    []int
		expectedOutput int
	}

	testCases := []testCase{
		{
			[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			49,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test: %d", i), func(t *testing.T) {
			output := maxArea(tc.inputHeight)
			assert.Equal(t, tc.expectedOutput, output)
		})

	}
}
