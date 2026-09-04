package numberofislands

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumIslands(t *testing.T) {
	type testCase struct {
		inputGrid      [][]byte
		expectedOutput int
	}

	testCases := []testCase{
		{
			[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			1,
		},
		{
			[][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			3,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test dfs: %d", i), func(t *testing.T) {
			output := numIslandsDFS(tc.inputGrid)
			assert.Equal(t, tc.expectedOutput, output)
		})

		t.Run(fmt.Sprintf("test bfs: %d", i), func(t *testing.T) {
			output := numIslandsBFS(tc.inputGrid)
			assert.Equal(t, tc.expectedOutput, output)
		})
	}
}
