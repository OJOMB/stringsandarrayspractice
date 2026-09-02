package longestrepeatingcharacterreplacement

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacterReplacement(t *testing.T) {
	testCases := []struct {
		inputS         string
		inputK         int
		expectedOutput int
	}{
		{
			"ABAB",
			2,
			4,
		},
		{
			"ABAB",
			0,
			1,
		},
		{
			"ABAB",
			1,
			3,
		},
		{
			"ABCBBB",
			1,
			5,
		},
		{
			"EQQEJDOBDPDPFPEIAQLQGDNIRDDGEHJIORMJPKGPLCPDFMIGHJNIIRSDSBRNJNROBALNSHCRFBASTLRMENCCIBJLGAITBFCSMPRO",
			2,
			5,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test: %d", i), func(t *testing.T) {
			output := characterReplacement(tc.inputS, tc.inputK)
			assert.Equal(t, tc.expectedOutput, output)
		})
	}
}
