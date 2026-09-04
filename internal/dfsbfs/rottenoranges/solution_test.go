package rottenoranges

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrangesRottingBFS(t *testing.T) {
	type testCase struct {
		grid     [][]int
		expected int
	}

	tests := []testCase{
		{
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			expected: 4,
		},
		{
			grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			expected: -1,
		},
		{
			grid:     [][]int{{0, 2}},
			expected: 0,
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			result := orangesRotting(tc.grid)
			assert.Equal(t, tc.expected, result)
		})
	}
}
