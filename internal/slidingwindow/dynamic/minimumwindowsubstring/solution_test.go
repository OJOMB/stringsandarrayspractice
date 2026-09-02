package minimumwindowsubstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinWindow(t *testing.T) {
	type testCase struct {
		inputS         string
		inputT         string
		expectedOutput string
	}

	testCases := []testCase{
		{
			"ADOBECODEBANC",
			"ABC",
			"BANC",
		},
		{
			"a",
			"aa",
			"",
		},
		{
			"a",
			"a",
			"a",
		},
		{
			"ab",
			"b",
			"b",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test: %d", i), func(t *testing.T) {
			output := minWindow(tc.inputS, tc.inputT)
			assert.Equal(t, tc.expectedOutput, output)
		})
	}
}
