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
			myOutput := minWindow(tc.inputS, tc.inputT)
			assert.Equal(t, tc.expectedOutput, myOutput)
		})

		t.Run(fmt.Sprintf("claude test: %d", i), func(t *testing.T) {
			claudeOutput := claudeMinWindow(tc.inputS, tc.inputT)
			assert.Equal(t, tc.expectedOutput, claudeOutput)
		})
	}
}
