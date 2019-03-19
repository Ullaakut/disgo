package disgo

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func TestNewTerminal(t *testing.T) {
	term := NewTerminal()

	assert.Equal(t, term.defaultOutput, os.Stdout)
	assert.Equal(t, term.errorOutput, os.Stderr)
	assert.False(t, term.debug)
	assert.Nil(t, term.step)
}

func TestNewTerminalWithOptions(t *testing.T) {
	defaultOut := &bytes.Buffer{}
	errorOut := &bytes.Buffer{}

	term := NewTerminal(WithDefaultOutput(defaultOut), WithErrorOutput(errorOut), WithDebug(true), WithColors(false))

	assert.Equal(t, term.defaultOutput, defaultOut)
	assert.Equal(t, term.errorOutput, errorOut)
	assert.True(t, term.debug)
	assert.True(t, color.NoColor)
	assert.Nil(t, term.step)
}

func TestGlobalOptions(t *testing.T) {
	defaultOut := &bytes.Buffer{}
	errorOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(defaultOut), WithErrorOutput(errorOut), WithDebug(true), WithColors(false))

	assert.Equal(t, globalTerm.defaultOutput, defaultOut)
	assert.Equal(t, globalTerm.errorOutput, errorOut)
	assert.True(t, globalTerm.debug)
	assert.True(t, color.NoColor)
	assert.Nil(t, globalTerm.step)
}

func TestInfoWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
	}

	// Info should not append a newline after printing.
	term.Info("one sentence")
	term.Info("another sentence")
	assert.Contains(t, defaultOut.String(), "one sentenceanother sentence")
	assert.NotContains(t, defaultOut.String(), "one sentence\n")
	assert.NotContains(t, defaultOut.String(), "another sentence\n")

	// Info should not join arguments with a space.
	term.Info("element one", "element two")
	assert.Contains(t, defaultOut.String(), "element oneelement two")
	assert.NotContains(t, defaultOut.String(), "element one element two")

	assert.Equal(t, "one sentenceanother sentenceelement oneelement two", defaultOut.String())
}

func TestInfolnWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
	}

	// Info should append a newline after printing.
	term.Infoln("one sentence")
	term.Infoln("another sentence")
	assert.Contains(t, defaultOut.String(), "one sentence\nanother sentence\n")
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	// Info should join arguments with a space.
	term.Infoln("element one", "element two")
	assert.Contains(t, defaultOut.String(), "element one element two\n")
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	assert.Equal(t, "one sentence\nanother sentence\nelement one element two\n", defaultOut.String())
}

func TestInfofWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
	}

	// Infof should not append a newline after printing.
	term.Infof("one sentence")
	term.Infof("another sentence")
	assert.Contains(t, defaultOut.String(), "one sentenceanother sentence")
	assert.NotContains(t, defaultOut.String(), "one sentence\n")
	assert.NotContains(t, defaultOut.String(), "another sentence\n")

	// Info should not join arguments with a space.
	term.Infof("element one%s", "element two")
	assert.Contains(t, defaultOut.String(), "element oneelement two")
	assert.NotContains(t, defaultOut.String(), "element one element two")

	assert.Equal(t, "one sentenceanother sentenceelement oneelement two", defaultOut.String())
}

func TestDebugEnabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
		debug:         true,
	}

	// Debug should not append a newline after printing.
	term.Debug("one sentence")
	term.Debug("another sentence")
	assert.Contains(t, defaultOut.String(), "one sentenceanother sentence")
	assert.NotContains(t, defaultOut.String(), "one sentence\n")
	assert.NotContains(t, defaultOut.String(), "another sentence\n")

	// Debug should not join arguments with a space.
	term.Debug("element one", "element two")
	assert.Contains(t, defaultOut.String(), "element oneelement two")
	assert.NotContains(t, defaultOut.String(), "element one element two")

	assert.Equal(t, "one sentenceanother sentenceelement oneelement two", defaultOut.String())
}

func TestDebuglnEnabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
		debug:         true,
	}

	// Debug should append a newline after printing.
	term.Debugln("one sentence")
	term.Debugln("another sentence")
	assert.Contains(t, defaultOut.String(), "one sentence\nanother sentence\n")
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	// Debug should join arguments with a space.
	term.Debugln("element one", "element two")
	assert.Contains(t, defaultOut.String(), "element one element two\n")
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	assert.Equal(t, "one sentence\nanother sentence\nelement one element two\n", defaultOut.String())
}

func TestDebugfEnabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
		debug:         true,
	}

	// Debugf should not append a newline after printing.
	term.Debugf("one sentence")
	term.Debugf("another sentence")
	assert.Contains(t, defaultOut.String(), "one sentenceanother sentence")
	assert.NotContains(t, defaultOut.String(), "one sentence\n")
	assert.NotContains(t, defaultOut.String(), "another sentence\n")

	// Debug should not join arguments with a space.
	term.Debugf("element one%s", "element two")
	assert.Contains(t, defaultOut.String(), "element oneelement two")
	assert.NotContains(t, defaultOut.String(), "element one element two")

	assert.Equal(t, "one sentenceanother sentenceelement oneelement two", defaultOut.String())
}

func TestDebugDisabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
		debug:         false,
	}

	// Debug should not output anything when debug is disabled.
	term.Debug("one sentence")
	term.Debug("another sentence")
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	term.Debug("element one", "element two")
	assert.NotContains(t, defaultOut.String(), "element oneelement two")
}

func TestDebuglnDisabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
		debug:         false,
	}

	// Debugln should not output anything when debug is disabled.
	term.Debugln("one sentence")
	term.Debugln("another sentence")
	assert.NotContains(t, defaultOut.String(), "one sentence\nanother sentence\n")

	term.Debugln("element one", "element two")
	assert.NotContains(t, defaultOut.String(), "element one element two\n")
}

func TestDebugfDisabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
		debug:         false,
	}

	// Debugf should not output anything when debug is disabled.
	term.Debugf("one sentence")
	term.Debugf("another sentence")
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	term.Debugf("element one%s", "element two")
	assert.NotContains(t, defaultOut.String(), "element oneelement two")
}

func TestErrorWithoutStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	term := &Terminal{
		errorOutput:   errorOut,
		defaultOutput: ioutil.Discard,
	}

	// Error should not append a newline after printing.
	term.Error("one sentence")
	term.Error("another sentence")
	assert.Contains(t, errorOut.String(), "one sentenceanother sentence")
	assert.NotContains(t, errorOut.String(), "one sentence\n")
	assert.NotContains(t, errorOut.String(), "another sentence\n")

	// Error should not join arguments with a space.
	term.Error("element one", "element two")
	assert.Contains(t, errorOut.String(), "element oneelement two")
	assert.NotContains(t, errorOut.String(), "element one element two")

	assert.Equal(t, "one sentenceanother sentenceelement oneelement two", errorOut.String())
}

func TestErrorlnWithoutStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	term := &Terminal{
		errorOutput:   errorOut,
		defaultOutput: ioutil.Discard,
	}

	// Error should append a newline after printing.
	term.Errorln("one sentence")
	term.Errorln("another sentence")
	assert.Contains(t, errorOut.String(), "one sentence\nanother sentence\n")
	assert.NotContains(t, errorOut.String(), "one sentenceanother sentence")

	// Error should join arguments with a space.
	term.Errorln("element one", "element two")
	assert.Contains(t, errorOut.String(), "element one element two\n")
	assert.NotContains(t, errorOut.String(), "element oneelement two")

	assert.Equal(t, "one sentence\nanother sentence\nelement one element two\n", errorOut.String())
}

func TestErrorfWithoutStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	term := &Terminal{
		errorOutput:   errorOut,
		defaultOutput: ioutil.Discard,
	}

	// Errorf should not append a newline after printing.
	term.Errorf("one sentence")
	term.Errorf("another sentence")
	assert.Contains(t, errorOut.String(), "one sentenceanother sentence")
	assert.NotContains(t, errorOut.String(), "one sentence\n")
	assert.NotContains(t, errorOut.String(), "another sentence\n")

	// Error should not join arguments with a space.
	term.Errorf("element one%s", "element two")
	assert.Contains(t, errorOut.String(), "element oneelement two")
	assert.NotContains(t, errorOut.String(), "element one element two")

	assert.Equal(t, "one sentenceanother sentenceelement oneelement two", errorOut.String())
}

//*******************//
// Test global logger//
//*******************//

func TestGlobalTerminalInfoWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(defaultOut))

	// Info should not append a newline after printing.
	Info("one sentence")
	Info("another sentence")
	assert.Contains(t, defaultOut.String(), "one sentenceanother sentence")
	assert.NotContains(t, defaultOut.String(), "one sentence\n")
	assert.NotContains(t, defaultOut.String(), "another sentence\n")

	// Info should not join arguments with a space.
	Info("element one", "element two")
	assert.Contains(t, defaultOut.String(), "element oneelement two")
	assert.NotContains(t, defaultOut.String(), "element one element two")

	assert.Equal(t, "one sentenceanother sentenceelement oneelement two", defaultOut.String())
}

func TestGlobalTerminalInfolnWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(defaultOut))

	// Info should append a newline after printing.
	Infoln("one sentence")
	Infoln("another sentence")
	assert.Contains(t, defaultOut.String(), "one sentence\nanother sentence\n")
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	// Info should join arguments with a space.
	Infoln("element one", "element two")
	assert.Contains(t, defaultOut.String(), "element one element two\n")
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	assert.Equal(t, "one sentence\nanother sentence\nelement one element two\n", defaultOut.String())
}

func TestGlobalTerminalInfofWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(defaultOut))

	// Infof should not append a newline after printing.
	Infof("one sentence")
	Infof("another sentence")
	assert.Contains(t, defaultOut.String(), "one sentenceanother sentence")
	assert.NotContains(t, defaultOut.String(), "one sentence\n")
	assert.NotContains(t, defaultOut.String(), "another sentence\n")

	// Info should not join arguments with a space.
	Infof("element one%s", "element two")
	assert.Contains(t, defaultOut.String(), "element oneelement two")
	assert.NotContains(t, defaultOut.String(), "element one element two")

	assert.Equal(t, "one sentenceanother sentenceelement oneelement two", defaultOut.String())
}

func TestGlobalTerminalDebugEnabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(defaultOut), WithDebug(true))

	// Debug should not append a newline after printing.
	Debug("one sentence")
	Debug("another sentence")
	assert.Contains(t, defaultOut.String(), "one sentenceanother sentence")
	assert.NotContains(t, defaultOut.String(), "one sentence\n")
	assert.NotContains(t, defaultOut.String(), "another sentence\n")

	// Debug should not join arguments with a space.
	Debug("element one", "element two")
	assert.Contains(t, defaultOut.String(), "element oneelement two")
	assert.NotContains(t, defaultOut.String(), "element one element two")

	assert.Equal(t, "one sentenceanother sentenceelement oneelement two", defaultOut.String())
}

func TestGlobalTerminalDebuglnEnabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(defaultOut), WithDebug(true))

	// Debug should append a newline after printing.
	Debugln("one sentence")
	Debugln("another sentence")
	assert.Contains(t, defaultOut.String(), "one sentence\nanother sentence\n")
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	// Debug should join arguments with a space.
	Debugln("element one", "element two")
	assert.Contains(t, defaultOut.String(), "element one element two\n")
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	assert.Equal(t, "one sentence\nanother sentence\nelement one element two\n", defaultOut.String())
}

func TestGlobalTerminalDebugfEnabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(defaultOut), WithDebug(true))

	// Debugf should not append a newline after printing.
	Debugf("one sentence")
	Debugf("another sentence")
	assert.Contains(t, defaultOut.String(), "one sentenceanother sentence")
	assert.NotContains(t, defaultOut.String(), "one sentence\n")
	assert.NotContains(t, defaultOut.String(), "another sentence\n")

	// Debug should not join arguments with a space.
	Debugf("element one%s", "element two")
	assert.Contains(t, defaultOut.String(), "element oneelement two")
	assert.NotContains(t, defaultOut.String(), "element one element two")

	assert.Equal(t, "one sentenceanother sentenceelement oneelement two", defaultOut.String())
}

func TestGlobalTerminalDebugDisabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(defaultOut), WithDebug(false))

	// Debug should not output anything when debug is disabled.
	Debug("one sentence")
	Debug("another sentence")
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	Debug("element one", "element two")
	assert.NotContains(t, defaultOut.String(), "element oneelement two")
}

func TestGlobalTerminalDebuglnDisabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(defaultOut), WithDebug(false))

	// Debugln should not output anything when debug is disabled.
	Debugln("one sentence")
	Debugln("another sentence")
	assert.NotContains(t, defaultOut.String(), "one sentence\nanother sentence\n")

	Debugln("element one", "element two")
	assert.NotContains(t, defaultOut.String(), "element one element two\n")
}

func TestGlobalTerminalDebugfDisabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(defaultOut), WithDebug(false))

	// Debugf should not output anything when debug is disabled.
	Debugf("one sentence")
	Debugf("another sentence")
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	Debugf("element one%s", "element two")
	assert.NotContains(t, defaultOut.String(), "element oneelement two")
}

func TestGlobalTerminalErrorWithoutStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(ioutil.Discard), WithErrorOutput(errorOut))

	// Error should not append a newline after printing.
	Error("one sentence")
	Error("another sentence")
	assert.Contains(t, errorOut.String(), "one sentenceanother sentence")
	assert.NotContains(t, errorOut.String(), "one sentence\n")
	assert.NotContains(t, errorOut.String(), "another sentence\n")

	// Error should not join arguments with a space.
	Error("element one", "element two")
	assert.Contains(t, errorOut.String(), "element oneelement two")
	assert.NotContains(t, errorOut.String(), "element one element two")

	assert.Equal(t, "one sentenceanother sentenceelement oneelement two", errorOut.String())
}

func TestGlobalTerminalErrorlnWithoutStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(ioutil.Discard), WithErrorOutput(errorOut))

	// Error should append a newline after printing.
	Errorln("one sentence")
	Errorln("another sentence")
	assert.Contains(t, errorOut.String(), "one sentence\nanother sentence\n")
	assert.NotContains(t, errorOut.String(), "one sentenceanother sentence")

	// Error should join arguments with a space.
	Errorln("element one", "element two")
	assert.Contains(t, errorOut.String(), "element one element two\n")
	assert.NotContains(t, errorOut.String(), "element oneelement two")

	assert.Equal(t, "one sentence\nanother sentence\nelement one element two\n", errorOut.String())
}

func TestGlobalTerminalErrorfWithoutStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(ioutil.Discard), WithErrorOutput(errorOut))

	// Errorf should not append a newline after printing.
	Errorf("one sentence")
	Errorf("another sentence")
	assert.Contains(t, errorOut.String(), "one sentenceanother sentence")
	assert.NotContains(t, errorOut.String(), "one sentence\n")
	assert.NotContains(t, errorOut.String(), "another sentence\n")

	// Error should not join arguments with a space.
	Errorf("element one%s", "element two")
	assert.Contains(t, errorOut.String(), "element oneelement two")
	assert.NotContains(t, errorOut.String(), "element one element two")

	assert.Equal(t, "one sentenceanother sentenceelement oneelement two", errorOut.String())
}

//******************//
// Tests with Steps //
//******************//

func TestInfoWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
		step:          &step{},
	}

	term.Info("one sentence")
	term.Info("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, term.step.queue, stepOutput{content: "one sentence", level: levelInfo})
	assert.Contains(t, term.step.queue, stepOutput{content: "another sentence", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	term.Info("element one", "element two")
	assert.Contains(t, term.step.queue, stepOutput{content: "element oneelement two", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	term.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element oneelement two")
}

func TestInfolnWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
		step:          &step{},
	}

	term.Infoln("one sentence")
	term.Infoln("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, term.step.queue, stepOutput{content: "one sentence\n", level: levelInfo})
	assert.Contains(t, term.step.queue, stepOutput{content: "another sentence\n", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	term.Infoln("element one", "element two")
	assert.Contains(t, term.step.queue, stepOutput{content: "element one element two\n", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "element one element two")

	term.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element one element two")
}

func TestInfofWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
		step:          &step{},
	}

	term.Infof("one sentence")
	term.Infof("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, term.step.queue, stepOutput{content: "one sentence", level: levelInfo})
	assert.Contains(t, term.step.queue, stepOutput{content: "another sentence", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	term.Infof("element one%s", "element two")
	assert.Contains(t, term.step.queue, stepOutput{content: "element oneelement two", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	term.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element oneelement two")
}

func TestDebugEnabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
		step:          &step{},
		debug:         true,
	}

	term.Debug("one sentence")
	term.Debug("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, term.step.queue, stepOutput{content: "one sentence", level: levelDebug})
	assert.Contains(t, term.step.queue, stepOutput{content: "another sentence", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	term.Debug("element one", "element two")
	assert.Contains(t, term.step.queue, stepOutput{content: "element oneelement two", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	term.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element oneelement two")
}

func TestDebuglnEnabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
		step:          &step{},
		debug:         true,
	}

	term.Debugln("one sentence")
	term.Debugln("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, term.step.queue, stepOutput{content: "one sentence\n", level: levelDebug})
	assert.Contains(t, term.step.queue, stepOutput{content: "another sentence\n", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	term.Debugln("element one", "element two")
	assert.Contains(t, term.step.queue, stepOutput{content: "element one element two\n", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element one element two")

	term.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element one element two")
}

func TestDebugfEnabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
		step:          &step{},
		debug:         true,
	}

	term.Debugf("one sentence")
	term.Debugf("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, term.step.queue, stepOutput{content: "one sentence", level: levelDebug})
	assert.Contains(t, term.step.queue, stepOutput{content: "another sentence", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	term.Debugf("element one%s", "element two")
	assert.Contains(t, term.step.queue, stepOutput{content: "element oneelement two", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	term.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element oneelement two")
}

func TestDebugDisabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
		step:          &step{},
		debug:         false,
	}

	term.Debug("one sentence")
	term.Debug("another sentence")
	// Since debug is disabled, even if a step is in progress, debug logs should not be queued.
	assert.NotContains(t, term.step.queue, stepOutput{content: "one sentence", level: levelDebug})
	assert.NotContains(t, term.step.queue, stepOutput{content: "another sentence", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	term.Debug("element one", "element two")
	assert.NotContains(t, term.step.queue, stepOutput{content: "element oneelement two", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	term.EndStep()

	// Since debug is disabled, debug level logs are not shown.
	assert.NotContains(t, defaultOut.String(), "> one sentence")
	assert.NotContains(t, defaultOut.String(), "> another sentence")
	assert.NotContains(t, defaultOut.String(), "> element oneelement two")
}

func TestDebuglnDisabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
		step:          &step{},
		debug:         false,
	}

	term.Debugln("one sentence")
	term.Debugln("another sentence")
	// Since debug is disabled, even if a step is in progress, debug logs should not be queued.
	assert.NotContains(t, term.step.queue, stepOutput{content: "one sentence\n", level: levelDebug})
	assert.NotContains(t, term.step.queue, stepOutput{content: "another sentence\n", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	term.Debugln("element one", "element two")
	assert.NotContains(t, term.step.queue, stepOutput{content: "element one element two\n", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element one element two")

	term.EndStep()

	// Since debug is disabled, debug level logs are not shown.
	assert.NotContains(t, defaultOut.String(), "> one sentence")
	assert.NotContains(t, defaultOut.String(), "> another sentence")
	assert.NotContains(t, defaultOut.String(), "> element one element two")
}

func TestDebugfDisabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
		step:          &step{},
		debug:         false,
	}

	term.Debugf("one sentence")
	term.Debugf("another sentence")
	// Since debug is disabled, even if a step is in progress, debug logs should not be queued.
	assert.NotContains(t, term.step.queue, stepOutput{content: "one sentence", level: levelDebug})
	assert.NotContains(t, term.step.queue, stepOutput{content: "another sentence", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	term.Debugf("element one%s", "element two")
	assert.NotContains(t, term.step.queue, stepOutput{content: "element oneelement two", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	term.EndStep()

	// Since debug is disabled, debug level logs are not shown.
	assert.NotContains(t, defaultOut.String(), "> one sentence")
	assert.NotContains(t, defaultOut.String(), "> another sentence")
	assert.NotContains(t, defaultOut.String(), "> element oneelement two")
}

func TestErrorWithStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	term := &Terminal{
		errorOutput:   errorOut,
		defaultOutput: ioutil.Discard,
		step:          &step{},
	}

	term.Error("one sentence")
	term.Error("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, term.step.queue, stepOutput{content: "one sentence", level: levelError})
	assert.Contains(t, term.step.queue, stepOutput{content: "another sentence", level: levelError})
	assert.NotContains(t, errorOut.String(), "one sentenceanother sentence")

	term.Error("element one", "element two")
	assert.Contains(t, term.step.queue, stepOutput{content: "element oneelement two", level: levelError})
	assert.NotContains(t, errorOut.String(), "element oneelement two")

	term.EndStep()

	assert.Contains(t, errorOut.String(), "> one sentence")
	assert.Contains(t, errorOut.String(), "> another sentence")
	assert.Contains(t, errorOut.String(), "> element oneelement two")
}

func TestErrorlnWithStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	term := &Terminal{
		errorOutput:   errorOut,
		defaultOutput: ioutil.Discard,
		step:          &step{},
	}

	term.Errorln("one sentence")
	term.Errorln("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, term.step.queue, stepOutput{content: "one sentence\n", level: levelError})
	assert.Contains(t, term.step.queue, stepOutput{content: "another sentence\n", level: levelError})
	assert.NotContains(t, errorOut.String(), "one sentenceanother sentence")

	term.Errorln("element one", "element two")
	assert.Contains(t, term.step.queue, stepOutput{content: "element one element two\n", level: levelError})
	assert.NotContains(t, errorOut.String(), "element one element two")

	term.EndStep()

	assert.Contains(t, errorOut.String(), "> one sentence")
	assert.Contains(t, errorOut.String(), "> another sentence")
	assert.Contains(t, errorOut.String(), "> element one element two")
}

func TestErrorfWithStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	term := &Terminal{
		errorOutput:   errorOut,
		defaultOutput: ioutil.Discard,
		step:          &step{},
	}

	term.Errorf("one sentence")
	term.Errorf("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, term.step.queue, stepOutput{content: "one sentence", level: levelError})
	assert.Contains(t, term.step.queue, stepOutput{content: "another sentence", level: levelError})
	assert.NotContains(t, errorOut.String(), "one sentenceanother sentence")

	term.Errorf("element one%s", "element two")
	assert.Contains(t, term.step.queue, stepOutput{content: "element oneelement two", level: levelError})
	assert.NotContains(t, errorOut.String(), "element oneelement two")

	term.EndStep()

	assert.Contains(t, errorOut.String(), "> one sentence")
	assert.Contains(t, errorOut.String(), "> another sentence")
	assert.Contains(t, errorOut.String(), "> element oneelement two")
}

//*********************//
// Test Global Terminal //
//*********************//

func TestGlobalInfoWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(defaultOut))
	StartStep("test task")

	globalTerm.Info("one sentence")
	globalTerm.Info("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "one sentence", level: levelInfo})
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "another sentence", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	globalTerm.Info("element one", "element two")
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "element oneelement two", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	globalTerm.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element oneelement two")
}

func TestGlobalInfolnWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(defaultOut))
	StartStep("test task")

	globalTerm.Infoln("one sentence")
	globalTerm.Infoln("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "one sentence\n", level: levelInfo})
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "another sentence\n", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	globalTerm.Infoln("element one", "element two")
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "element one element two\n", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "element one element two")

	globalTerm.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element one element two")
}

func TestGlobalInfofWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(defaultOut))
	StartStep("test task")

	globalTerm.Infof("one sentence")
	globalTerm.Infof("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "one sentence", level: levelInfo})
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "another sentence", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	globalTerm.Infof("element one%s", "element two")
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "element oneelement two", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	globalTerm.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element oneelement two")
}

func TestGlobalDebugEnabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(defaultOut), WithDebug(true))
	StartStep("test task")

	globalTerm.Debug("one sentence")
	globalTerm.Debug("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "one sentence", level: levelDebug})
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "another sentence", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	globalTerm.Debug("element one", "element two")
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "element oneelement two", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	globalTerm.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element oneelement two")
}

func TestGlobalDebuglnEnabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(defaultOut), WithDebug(true))
	StartStep("test task")

	globalTerm.Debugln("one sentence")
	globalTerm.Debugln("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "one sentence\n", level: levelDebug})
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "another sentence\n", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	globalTerm.Debugln("element one", "element two")
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "element one element two\n", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element one element two")

	globalTerm.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element one element two")
}

func TestGlobalDebugfEnabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(defaultOut), WithDebug(true))
	StartStep("test task")

	globalTerm.Debugf("one sentence")
	globalTerm.Debugf("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "one sentence", level: levelDebug})
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "another sentence", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	globalTerm.Debugf("element one%s", "element two")
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "element oneelement two", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	globalTerm.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element oneelement two")
}

func TestGlobalDebugDisabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(defaultOut), WithDebug(false))
	StartStep("test task")

	globalTerm.Debug("one sentence")
	globalTerm.Debug("another sentence")
	// Since debug is disabled, even if a step is in progress, debug logs should not be queued.
	assert.NotContains(t, globalTerm.step.queue, stepOutput{content: "one sentence", level: levelDebug})
	assert.NotContains(t, globalTerm.step.queue, stepOutput{content: "another sentence", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	globalTerm.Debug("element one", "element two")
	assert.NotContains(t, globalTerm.step.queue, stepOutput{content: "element oneelement two", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	globalTerm.EndStep()

	// Since debug is disabled, debug level logs are not shown.
	assert.NotContains(t, defaultOut.String(), "> one sentence")
	assert.NotContains(t, defaultOut.String(), "> another sentence")
	assert.NotContains(t, defaultOut.String(), "> element oneelement two")
}

func TestGlobalDebuglnDisabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(defaultOut), WithDebug(false))
	StartStep("test task")

	globalTerm.Debugln("one sentence")
	globalTerm.Debugln("another sentence")
	// Since debug is disabled, even if a step is in progress, debug logs should not be queued.
	assert.NotContains(t, globalTerm.step.queue, stepOutput{content: "one sentence\n", level: levelDebug})
	assert.NotContains(t, globalTerm.step.queue, stepOutput{content: "another sentence\n", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	globalTerm.Debugln("element one", "element two")
	assert.NotContains(t, globalTerm.step.queue, stepOutput{content: "element one element two\n", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element one element two")

	globalTerm.EndStep()

	// Since debug is disabled, debug level logs are not shown.
	assert.NotContains(t, defaultOut.String(), "> one sentence")
	assert.NotContains(t, defaultOut.String(), "> another sentence")
	assert.NotContains(t, defaultOut.String(), "> element one element two")
}

func TestGlobalDebugfDisabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(defaultOut), WithDebug(false))
	StartStep("test task")

	globalTerm.Debugf("one sentence")
	globalTerm.Debugf("another sentence")
	// Since debug is disabled, even if a step is in progress, debug logs should not be queued.
	assert.NotContains(t, globalTerm.step.queue, stepOutput{content: "one sentence", level: levelDebug})
	assert.NotContains(t, globalTerm.step.queue, stepOutput{content: "another sentence", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	globalTerm.Debugf("element one%s", "element two")
	assert.NotContains(t, globalTerm.step.queue, stepOutput{content: "element oneelement two", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	globalTerm.EndStep()

	// Since debug is disabled, debug level logs are not shown.
	assert.NotContains(t, defaultOut.String(), "> one sentence")
	assert.NotContains(t, defaultOut.String(), "> another sentence")
	assert.NotContains(t, defaultOut.String(), "> element oneelement two")
}

func TestGlobalErrorWithStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(ioutil.Discard), WithErrorOutput(errorOut))
	StartStep("test task")

	globalTerm.Error("one sentence")
	globalTerm.Error("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "one sentence", level: levelError})
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "another sentence", level: levelError})
	assert.NotContains(t, errorOut.String(), "one sentenceanother sentence")

	globalTerm.Error("element one", "element two")
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "element oneelement two", level: levelError})
	assert.NotContains(t, errorOut.String(), "element oneelement two")

	globalTerm.EndStep()

	assert.Contains(t, errorOut.String(), "> one sentence")
	assert.Contains(t, errorOut.String(), "> another sentence")
	assert.Contains(t, errorOut.String(), "> element oneelement two")
}

func TestGlobalErrorlnWithStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(ioutil.Discard), WithErrorOutput(errorOut))
	StartStep("test task")

	globalTerm.Errorln("one sentence")
	globalTerm.Errorln("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "one sentence\n", level: levelError})
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "another sentence\n", level: levelError})
	assert.NotContains(t, errorOut.String(), "one sentenceanother sentence")

	globalTerm.Errorln("element one", "element two")
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "element one element two\n", level: levelError})
	assert.NotContains(t, errorOut.String(), "element one element two")

	globalTerm.EndStep()

	assert.Contains(t, errorOut.String(), "> one sentence")
	assert.Contains(t, errorOut.String(), "> another sentence")
	assert.Contains(t, errorOut.String(), "> element one element two")
}

func TestGlobalErrorfWithStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	SetTerminalOptions(WithDefaultOutput(ioutil.Discard), WithErrorOutput(errorOut))
	StartStep("test task")

	globalTerm.Errorf("one sentence")
	globalTerm.Errorf("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "one sentence", level: levelError})
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "another sentence", level: levelError})
	assert.NotContains(t, errorOut.String(), "one sentenceanother sentence")

	globalTerm.Errorf("element one%s", "element two")
	assert.Contains(t, globalTerm.step.queue, stepOutput{content: "element oneelement two", level: levelError})
	assert.NotContains(t, errorOut.String(), "element oneelement two")

	globalTerm.EndStep()

	assert.Contains(t, errorOut.String(), "> one sentence")
	assert.Contains(t, errorOut.String(), "> another sentence")
	assert.Contains(t, errorOut.String(), "> element oneelement two")
}
