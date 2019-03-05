package colog

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrompterConfirm(t *testing.T) {
	testCases := []struct {
		desc           string
		input          string
		defaultValue   bool
		prompt         string
		expectedResult bool
		expectsError   bool
	}{
		{
			desc:           "returns true",
			input:          "y",
			defaultValue:   false,
			prompt:         "Where's waldo ?",
			expectedResult: true,
			expectsError:   false,
		},
		{
			desc:           "returns false",
			input:          "n",
			defaultValue:   false,
			prompt:         "Where's waldo ?",
			expectedResult: false,
			expectsError:   false,
		},
		{
			desc:           "returns defaultValue",
			input:          "",
			defaultValue:   true,
			prompt:         "Where's waldo ?",
			expectedResult: true,
			expectsError:   false,
		},
		{
			desc:           "returns an error",
			input:          "NOOOOOOOOOOOOOOOOOOOOOOO",
			defaultValue:   false,
			prompt:         "Where's Padme ?",
			expectedResult: false,
			expectsError:   true,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			in := bytes.Buffer{}
			out := bytes.Buffer{}

			_, err := in.Write(append([]byte(test.input), '\n'))
			require.NoError(t, err)

			subject := NewPrompter(&out, &in)
			result, err := subject.Confirm(test.prompt, test.defaultValue, DefaultConfirmation)

			if test.expectsError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.expectedResult, result)
			assert.Contains(t, out.String(), test.prompt)
		})
	}
}
