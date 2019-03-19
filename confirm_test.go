package disgo

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrompterConfirm(t *testing.T) {
	testCases := []struct {
		desc               string
		input              string
		prompt             string
		choices            []string
		enableDefaultValue bool
		defaultValue       bool
		parser             ConfirmationParser
		expectedResult     bool
		expectsError       bool
	}{
		{
			desc:           "returns true",
			prompt:         "Where's waldo?",
			input:          "y",
			defaultValue:   false,
			expectedResult: true,
			expectsError:   false,
		},
		{
			desc:           "returns false",
			prompt:         "Where's waldo?",
			input:          "n",
			defaultValue:   false,
			expectedResult: false,
			expectsError:   false,
		},
		{
			desc:               "returns defaultValue",
			prompt:             "Where's waldo?",
			input:              "",
			enableDefaultValue: true,
			defaultValue:       true,
			expectedResult:     true,
			expectsError:       false,
		},
		{
			desc:               "doesn't use default value, errors",
			prompt:             "Where's waldo?",
			input:              "",
			enableDefaultValue: false,
			defaultValue:       false,
			expectsError:       true,
		},
		{
			desc:         "doesn't use default choices",
			prompt:       "Where's waldo?",
			input:        "p",
			choices:      []string{"p", "t", "x"},
			expectsError: true,
		},
		{
			desc:   "doesn't use default parser",
			prompt: "Where's waldo?",
			input:  "y",
			parser: func(input string) (bool, error) {
				return true, nil
			},
			expectedResult: true,
			expectsError:   false,
		},
		{
			desc:           "returns an error",
			prompt:         "Where's Padme?",
			input:          "NOOOOOOOOOOOOOOOOOOOOOOO",
			defaultValue:   false,
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

			subject := NewTerminal(WithReader(&in), WithDefaultOutput(&out))

			result, err := subject.Confirm(Confirmation{
				Label:              test.prompt,
				EnableDefaultValue: test.enableDefaultValue,
				DefaultValue:       test.defaultValue,
				Parser:             test.parser,
				Choices:            test.choices,
			})

			if test.expectsError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.expectedResult, result)
			assert.Contains(t, out.String(), test.prompt)

			if test.choices != nil {
				assert.Contains(t, out.String(), strings.Join(test.choices, "/"))
			} else {
				assert.Contains(t, out.String(), strings.Join(DefaultConfirmationChoices, "/"))
			}
		})
	}
}

type readerMock struct{}

func (w readerMock) Read([]byte) (int, error) {
	return 0, errors.New("mock error")
}

func (w readerMock) Write([]byte) (int, error) {
	return 0, nil
}

func TestReadError(t *testing.T) {
	in := readerMock{}
	out := bytes.Buffer{}

	_, err := in.Write(append([]byte("user input"), '\n'))
	require.NoError(t, err)

	subject := NewTerminal(WithReader(&in), WithDefaultOutput(&out))

	_, err = subject.Confirm(Confirmation{
		Label: "label",
	})

	// Ensure that if the reader fails, an error is returned.
	assert.Error(t, err)
}

func TestNonInteractivePrompter(t *testing.T) {
	in := readerMock{}
	out := bytes.Buffer{}

	subject := NewTerminal(WithReader(&in), WithDefaultOutput(&out), WithInteractive(false))

	value, err := subject.Confirm(Confirmation{
		Label:        "label",
		DefaultValue: true,
	})

	// Ensure that when the prompter is set to non-interactive,
	// it returns no error.
	assert.NoError(t, err)

	// Ensure that when the prompter is set to non-interactive,
	// it returns the default value.
	assert.Equal(t, true, value)
}

func TestGlobalPrompter(t *testing.T) {
	in := bytes.Buffer{}
	out := bytes.Buffer{}

	_, err := in.Write(append([]byte("y"), '\n'))
	require.NoError(t, err)

	SetTerminalOptions(WithReader(&in), WithDefaultOutput(&out))

	_, err = Confirm(Confirmation{
		Label:        "label",
		DefaultValue: true,
	})

	// Ensure that when the prompter is set to non-interactive,
	// it returns no error.
	assert.NoError(t, err)
}
