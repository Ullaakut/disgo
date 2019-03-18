package disgo

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func TestNewConsole(t *testing.T) {
	console := New()

	assert.Equal(t, console.defaultOutput, os.Stdout)
	assert.Equal(t, console.errorOutput, os.Stderr)
	assert.False(t, console.debug)
	assert.Nil(t, console.step)
}

func TestNewConsoleWithOptions(t *testing.T) {
	defaultOut := &bytes.Buffer{}
	errorOut := &bytes.Buffer{}

	console := New(WithDefaultOutput(defaultOut), WithErrorOutput(errorOut), WithDebug(true), WithColors(false))

	assert.Equal(t, console.defaultOutput, defaultOut)
	assert.Equal(t, console.errorOutput, errorOut)
	assert.True(t, console.debug)
	assert.True(t, color.NoColor)
	assert.Nil(t, console.step)
}

func TestGlobalOptions(t *testing.T) {
	defaultOut := &bytes.Buffer{}
	errorOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(defaultOut), WithErrorOutput(errorOut), WithDebug(true), WithColors(false))

	assert.Equal(t, cnsl.defaultOutput, defaultOut)
	assert.Equal(t, cnsl.errorOutput, errorOut)
	assert.True(t, cnsl.debug)
	assert.True(t, color.NoColor)
	assert.Nil(t, cnsl.step)
}

func TestInfoWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	console := &Console{
		defaultOutput: defaultOut,
	}

	// Info should not append a newline after printing.
	console.Info("one sentence")
	console.Info("another sentence")
	assert.Contains(t, defaultOut.String(), "one sentenceanother sentence")
	assert.NotContains(t, defaultOut.String(), "one sentence\n")
	assert.NotContains(t, defaultOut.String(), "another sentence\n")

	// Info should not join arguments with a space.
	console.Info("element one", "element two")
	assert.Contains(t, defaultOut.String(), "element oneelement two")
	assert.NotContains(t, defaultOut.String(), "element one element two")

	assert.Equal(t, "one sentenceanother sentenceelement oneelement two", defaultOut.String())
}

func TestInfolnWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	console := &Console{
		defaultOutput: defaultOut,
	}

	// Info should append a newline after printing.
	console.Infoln("one sentence")
	console.Infoln("another sentence")
	assert.Contains(t, defaultOut.String(), "one sentence\nanother sentence\n")
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	// Info should join arguments with a space.
	console.Infoln("element one", "element two")
	assert.Contains(t, defaultOut.String(), "element one element two\n")
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	assert.Equal(t, "one sentence\nanother sentence\nelement one element two\n", defaultOut.String())
}

func TestInfofWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	console := &Console{
		defaultOutput: defaultOut,
	}

	// Infof should not append a newline after printing.
	console.Infof("one sentence")
	console.Infof("another sentence")
	assert.Contains(t, defaultOut.String(), "one sentenceanother sentence")
	assert.NotContains(t, defaultOut.String(), "one sentence\n")
	assert.NotContains(t, defaultOut.String(), "another sentence\n")

	// Info should not join arguments with a space.
	console.Infof("element one%s", "element two")
	assert.Contains(t, defaultOut.String(), "element oneelement two")
	assert.NotContains(t, defaultOut.String(), "element one element two")

	assert.Equal(t, "one sentenceanother sentenceelement oneelement two", defaultOut.String())
}

func TestDebugEnabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	console := &Console{
		defaultOutput: defaultOut,
		debug:         true,
	}

	// Debug should not append a newline after printing.
	console.Debug("one sentence")
	console.Debug("another sentence")
	assert.Contains(t, defaultOut.String(), "one sentenceanother sentence")
	assert.NotContains(t, defaultOut.String(), "one sentence\n")
	assert.NotContains(t, defaultOut.String(), "another sentence\n")

	// Debug should not join arguments with a space.
	console.Debug("element one", "element two")
	assert.Contains(t, defaultOut.String(), "element oneelement two")
	assert.NotContains(t, defaultOut.String(), "element one element two")

	assert.Equal(t, "one sentenceanother sentenceelement oneelement two", defaultOut.String())
}

func TestDebuglnEnabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	console := &Console{
		defaultOutput: defaultOut,
		debug:         true,
	}

	// Debug should append a newline after printing.
	console.Debugln("one sentence")
	console.Debugln("another sentence")
	assert.Contains(t, defaultOut.String(), "one sentence\nanother sentence\n")
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	// Debug should join arguments with a space.
	console.Debugln("element one", "element two")
	assert.Contains(t, defaultOut.String(), "element one element two\n")
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	assert.Equal(t, "one sentence\nanother sentence\nelement one element two\n", defaultOut.String())
}

func TestDebugfEnabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	console := &Console{
		defaultOutput: defaultOut,
		debug:         true,
	}

	// Debugf should not append a newline after printing.
	console.Debugf("one sentence")
	console.Debugf("another sentence")
	assert.Contains(t, defaultOut.String(), "one sentenceanother sentence")
	assert.NotContains(t, defaultOut.String(), "one sentence\n")
	assert.NotContains(t, defaultOut.String(), "another sentence\n")

	// Debug should not join arguments with a space.
	console.Debugf("element one%s", "element two")
	assert.Contains(t, defaultOut.String(), "element oneelement two")
	assert.NotContains(t, defaultOut.String(), "element one element two")

	assert.Equal(t, "one sentenceanother sentenceelement oneelement two", defaultOut.String())
}

func TestDebugDisabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	console := &Console{
		defaultOutput: defaultOut,
		debug:         false,
	}

	// Debug should not output anything when debug is disabled.
	console.Debug("one sentence")
	console.Debug("another sentence")
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	console.Debug("element one", "element two")
	assert.NotContains(t, defaultOut.String(), "element oneelement two")
}

func TestDebuglnDisabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	console := &Console{
		defaultOutput: defaultOut,
		debug:         false,
	}

	// Debugln should not output anything when debug is disabled.
	console.Debugln("one sentence")
	console.Debugln("another sentence")
	assert.NotContains(t, defaultOut.String(), "one sentence\nanother sentence\n")

	console.Debugln("element one", "element two")
	assert.NotContains(t, defaultOut.String(), "element one element two\n")
}

func TestDebugfDisabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	console := &Console{
		defaultOutput: defaultOut,
		debug:         false,
	}

	// Debugf should not output anything when debug is disabled.
	console.Debugf("one sentence")
	console.Debugf("another sentence")
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	console.Debugf("element one%s", "element two")
	assert.NotContains(t, defaultOut.String(), "element oneelement two")
}

func TestErrorWithoutStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	console := &Console{
		errorOutput:   errorOut,
		defaultOutput: ioutil.Discard,
	}

	// Error should not append a newline after printing.
	console.Error("one sentence")
	console.Error("another sentence")
	assert.Contains(t, errorOut.String(), "one sentenceanother sentence")
	assert.NotContains(t, errorOut.String(), "one sentence\n")
	assert.NotContains(t, errorOut.String(), "another sentence\n")

	// Error should not join arguments with a space.
	console.Error("element one", "element two")
	assert.Contains(t, errorOut.String(), "element oneelement two")
	assert.NotContains(t, errorOut.String(), "element one element two")

	assert.Equal(t, "one sentenceanother sentenceelement oneelement two", errorOut.String())
}

func TestErrorlnWithoutStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	console := &Console{
		errorOutput:   errorOut,
		defaultOutput: ioutil.Discard,
	}

	// Error should append a newline after printing.
	console.Errorln("one sentence")
	console.Errorln("another sentence")
	assert.Contains(t, errorOut.String(), "one sentence\nanother sentence\n")
	assert.NotContains(t, errorOut.String(), "one sentenceanother sentence")

	// Error should join arguments with a space.
	console.Errorln("element one", "element two")
	assert.Contains(t, errorOut.String(), "element one element two\n")
	assert.NotContains(t, errorOut.String(), "element oneelement two")

	assert.Equal(t, "one sentence\nanother sentence\nelement one element two\n", errorOut.String())
}

func TestErrorfWithoutStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	console := &Console{
		errorOutput:   errorOut,
		defaultOutput: ioutil.Discard,
	}

	// Errorf should not append a newline after printing.
	console.Errorf("one sentence")
	console.Errorf("another sentence")
	assert.Contains(t, errorOut.String(), "one sentenceanother sentence")
	assert.NotContains(t, errorOut.String(), "one sentence\n")
	assert.NotContains(t, errorOut.String(), "another sentence\n")

	// Error should not join arguments with a space.
	console.Errorf("element one%s", "element two")
	assert.Contains(t, errorOut.String(), "element oneelement two")
	assert.NotContains(t, errorOut.String(), "element one element two")

	assert.Equal(t, "one sentenceanother sentenceelement oneelement two", errorOut.String())
}

//*******************//
// Test global logger//
//*******************//

func TestGlobalConsoleInfoWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(defaultOut))

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

func TestGlobalConsoleInfolnWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(defaultOut))

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

func TestGlobalConsoleInfofWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(defaultOut))

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

func TestGlobalConsoleDebugEnabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(defaultOut), WithDebug(true))

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

func TestGlobalConsoleDebuglnEnabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(defaultOut), WithDebug(true))

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

func TestGlobalConsoleDebugfEnabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(defaultOut), WithDebug(true))

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

func TestGlobalConsoleDebugDisabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(defaultOut), WithDebug(false))

	// Debug should not output anything when debug is disabled.
	Debug("one sentence")
	Debug("another sentence")
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	Debug("element one", "element two")
	assert.NotContains(t, defaultOut.String(), "element oneelement two")
}

func TestGlobalConsoleDebuglnDisabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(defaultOut), WithDebug(false))

	// Debugln should not output anything when debug is disabled.
	Debugln("one sentence")
	Debugln("another sentence")
	assert.NotContains(t, defaultOut.String(), "one sentence\nanother sentence\n")

	Debugln("element one", "element two")
	assert.NotContains(t, defaultOut.String(), "element one element two\n")
}

func TestGlobalConsoleDebugfDisabledWithoutStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(defaultOut), WithDebug(false))

	// Debugf should not output anything when debug is disabled.
	Debugf("one sentence")
	Debugf("another sentence")
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	Debugf("element one%s", "element two")
	assert.NotContains(t, defaultOut.String(), "element oneelement two")
}

func TestGlobalConsoleErrorWithoutStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(ioutil.Discard), WithErrorOutput(errorOut))

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

func TestGlobalConsoleErrorlnWithoutStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(ioutil.Discard), WithErrorOutput(errorOut))

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

func TestGlobalConsoleErrorfWithoutStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(ioutil.Discard), WithErrorOutput(errorOut))

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

	console := &Console{
		defaultOutput: defaultOut,
		step:          &step{},
	}

	console.Info("one sentence")
	console.Info("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, console.step.queue, stepOutput{content: "one sentence", level: levelInfo})
	assert.Contains(t, console.step.queue, stepOutput{content: "another sentence", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	console.Info("element one", "element two")
	assert.Contains(t, console.step.queue, stepOutput{content: "element oneelement two", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	console.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element oneelement two")
}

func TestInfolnWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	console := &Console{
		defaultOutput: defaultOut,
		step:          &step{},
	}

	console.Infoln("one sentence")
	console.Infoln("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, console.step.queue, stepOutput{content: "one sentence\n", level: levelInfo})
	assert.Contains(t, console.step.queue, stepOutput{content: "another sentence\n", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	console.Infoln("element one", "element two")
	assert.Contains(t, console.step.queue, stepOutput{content: "element one element two\n", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "element one element two")

	console.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element one element two")
}

func TestInfofWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	console := &Console{
		defaultOutput: defaultOut,
		step:          &step{},
	}

	console.Infof("one sentence")
	console.Infof("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, console.step.queue, stepOutput{content: "one sentence", level: levelInfo})
	assert.Contains(t, console.step.queue, stepOutput{content: "another sentence", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	console.Infof("element one%s", "element two")
	assert.Contains(t, console.step.queue, stepOutput{content: "element oneelement two", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	console.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element oneelement two")
}

func TestDebugEnabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	console := &Console{
		defaultOutput: defaultOut,
		step:          &step{},
		debug:         true,
	}

	console.Debug("one sentence")
	console.Debug("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, console.step.queue, stepOutput{content: "one sentence", level: levelDebug})
	assert.Contains(t, console.step.queue, stepOutput{content: "another sentence", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	console.Debug("element one", "element two")
	assert.Contains(t, console.step.queue, stepOutput{content: "element oneelement two", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	console.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element oneelement two")
}

func TestDebuglnEnabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	console := &Console{
		defaultOutput: defaultOut,
		step:          &step{},
		debug:         true,
	}

	console.Debugln("one sentence")
	console.Debugln("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, console.step.queue, stepOutput{content: "one sentence\n", level: levelDebug})
	assert.Contains(t, console.step.queue, stepOutput{content: "another sentence\n", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	console.Debugln("element one", "element two")
	assert.Contains(t, console.step.queue, stepOutput{content: "element one element two\n", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element one element two")

	console.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element one element two")
}

func TestDebugfEnabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	console := &Console{
		defaultOutput: defaultOut,
		step:          &step{},
		debug:         true,
	}

	console.Debugf("one sentence")
	console.Debugf("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, console.step.queue, stepOutput{content: "one sentence", level: levelDebug})
	assert.Contains(t, console.step.queue, stepOutput{content: "another sentence", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	console.Debugf("element one%s", "element two")
	assert.Contains(t, console.step.queue, stepOutput{content: "element oneelement two", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	console.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element oneelement two")
}

func TestDebugDisabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	console := &Console{
		defaultOutput: defaultOut,
		step:          &step{},
		debug:         false,
	}

	console.Debug("one sentence")
	console.Debug("another sentence")
	// Since debug is disabled, even if a step is in progress, debug logs should not be queued.
	assert.NotContains(t, console.step.queue, stepOutput{content: "one sentence", level: levelDebug})
	assert.NotContains(t, console.step.queue, stepOutput{content: "another sentence", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	console.Debug("element one", "element two")
	assert.NotContains(t, console.step.queue, stepOutput{content: "element oneelement two", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	console.EndStep()

	// Since debug is disabled, debug level logs are not shown.
	assert.NotContains(t, defaultOut.String(), "> one sentence")
	assert.NotContains(t, defaultOut.String(), "> another sentence")
	assert.NotContains(t, defaultOut.String(), "> element oneelement two")
}

func TestDebuglnDisabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	console := &Console{
		defaultOutput: defaultOut,
		step:          &step{},
		debug:         false,
	}

	console.Debugln("one sentence")
	console.Debugln("another sentence")
	// Since debug is disabled, even if a step is in progress, debug logs should not be queued.
	assert.NotContains(t, console.step.queue, stepOutput{content: "one sentence\n", level: levelDebug})
	assert.NotContains(t, console.step.queue, stepOutput{content: "another sentence\n", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	console.Debugln("element one", "element two")
	assert.NotContains(t, console.step.queue, stepOutput{content: "element one element two\n", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element one element two")

	console.EndStep()

	// Since debug is disabled, debug level logs are not shown.
	assert.NotContains(t, defaultOut.String(), "> one sentence")
	assert.NotContains(t, defaultOut.String(), "> another sentence")
	assert.NotContains(t, defaultOut.String(), "> element one element two")
}

func TestDebugfDisabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	console := &Console{
		defaultOutput: defaultOut,
		step:          &step{},
		debug:         false,
	}

	console.Debugf("one sentence")
	console.Debugf("another sentence")
	// Since debug is disabled, even if a step is in progress, debug logs should not be queued.
	assert.NotContains(t, console.step.queue, stepOutput{content: "one sentence", level: levelDebug})
	assert.NotContains(t, console.step.queue, stepOutput{content: "another sentence", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	console.Debugf("element one%s", "element two")
	assert.NotContains(t, console.step.queue, stepOutput{content: "element oneelement two", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	console.EndStep()

	// Since debug is disabled, debug level logs are not shown.
	assert.NotContains(t, defaultOut.String(), "> one sentence")
	assert.NotContains(t, defaultOut.String(), "> another sentence")
	assert.NotContains(t, defaultOut.String(), "> element oneelement two")
}

func TestErrorWithStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	console := &Console{
		errorOutput:   errorOut,
		defaultOutput: ioutil.Discard,
		step:          &step{},
	}

	console.Error("one sentence")
	console.Error("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, console.step.queue, stepOutput{content: "one sentence", level: levelError})
	assert.Contains(t, console.step.queue, stepOutput{content: "another sentence", level: levelError})
	assert.NotContains(t, errorOut.String(), "one sentenceanother sentence")

	console.Error("element one", "element two")
	assert.Contains(t, console.step.queue, stepOutput{content: "element oneelement two", level: levelError})
	assert.NotContains(t, errorOut.String(), "element oneelement two")

	console.EndStep()

	assert.Contains(t, errorOut.String(), "> one sentence")
	assert.Contains(t, errorOut.String(), "> another sentence")
	assert.Contains(t, errorOut.String(), "> element oneelement two")
}

func TestErrorlnWithStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	console := &Console{
		errorOutput:   errorOut,
		defaultOutput: ioutil.Discard,
		step:          &step{},
	}

	console.Errorln("one sentence")
	console.Errorln("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, console.step.queue, stepOutput{content: "one sentence\n", level: levelError})
	assert.Contains(t, console.step.queue, stepOutput{content: "another sentence\n", level: levelError})
	assert.NotContains(t, errorOut.String(), "one sentenceanother sentence")

	console.Errorln("element one", "element two")
	assert.Contains(t, console.step.queue, stepOutput{content: "element one element two\n", level: levelError})
	assert.NotContains(t, errorOut.String(), "element one element two")

	console.EndStep()

	assert.Contains(t, errorOut.String(), "> one sentence")
	assert.Contains(t, errorOut.String(), "> another sentence")
	assert.Contains(t, errorOut.String(), "> element one element two")
}

func TestErrorfWithStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	console := &Console{
		errorOutput:   errorOut,
		defaultOutput: ioutil.Discard,
		step:          &step{},
	}

	console.Errorf("one sentence")
	console.Errorf("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, console.step.queue, stepOutput{content: "one sentence", level: levelError})
	assert.Contains(t, console.step.queue, stepOutput{content: "another sentence", level: levelError})
	assert.NotContains(t, errorOut.String(), "one sentenceanother sentence")

	console.Errorf("element one%s", "element two")
	assert.Contains(t, console.step.queue, stepOutput{content: "element oneelement two", level: levelError})
	assert.NotContains(t, errorOut.String(), "element oneelement two")

	console.EndStep()

	assert.Contains(t, errorOut.String(), "> one sentence")
	assert.Contains(t, errorOut.String(), "> another sentence")
	assert.Contains(t, errorOut.String(), "> element oneelement two")
}

//*********************//
// Test Global Console //
//*********************//

func TestGlobalInfoWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(defaultOut))
	StartStep("test task")

	cnsl.Info("one sentence")
	cnsl.Info("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "one sentence", level: levelInfo})
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "another sentence", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	cnsl.Info("element one", "element two")
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "element oneelement two", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	cnsl.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element oneelement two")
}

func TestGlobalInfolnWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(defaultOut))
	StartStep("test task")

	cnsl.Infoln("one sentence")
	cnsl.Infoln("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "one sentence\n", level: levelInfo})
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "another sentence\n", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	cnsl.Infoln("element one", "element two")
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "element one element two\n", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "element one element two")

	cnsl.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element one element two")
}

func TestGlobalInfofWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(defaultOut))
	StartStep("test task")

	cnsl.Infof("one sentence")
	cnsl.Infof("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "one sentence", level: levelInfo})
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "another sentence", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	cnsl.Infof("element one%s", "element two")
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "element oneelement two", level: levelInfo})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	cnsl.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element oneelement two")
}

func TestGlobalDebugEnabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(defaultOut), WithDebug(true))
	StartStep("test task")

	cnsl.Debug("one sentence")
	cnsl.Debug("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "one sentence", level: levelDebug})
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "another sentence", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	cnsl.Debug("element one", "element two")
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "element oneelement two", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	cnsl.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element oneelement two")
}

func TestGlobalDebuglnEnabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(defaultOut), WithDebug(true))
	StartStep("test task")

	cnsl.Debugln("one sentence")
	cnsl.Debugln("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "one sentence\n", level: levelDebug})
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "another sentence\n", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	cnsl.Debugln("element one", "element two")
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "element one element two\n", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element one element two")

	cnsl.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element one element two")
}

func TestGlobalDebugfEnabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(defaultOut), WithDebug(true))
	StartStep("test task")

	cnsl.Debugf("one sentence")
	cnsl.Debugf("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "one sentence", level: levelDebug})
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "another sentence", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	cnsl.Debugf("element one%s", "element two")
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "element oneelement two", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	cnsl.EndStep()

	assert.Contains(t, defaultOut.String(), "> one sentence")
	assert.Contains(t, defaultOut.String(), "> another sentence")
	assert.Contains(t, defaultOut.String(), "> element oneelement two")
}

func TestGlobalDebugDisabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(defaultOut), WithDebug(false))
	StartStep("test task")

	cnsl.Debug("one sentence")
	cnsl.Debug("another sentence")
	// Since debug is disabled, even if a step is in progress, debug logs should not be queued.
	assert.NotContains(t, cnsl.step.queue, stepOutput{content: "one sentence", level: levelDebug})
	assert.NotContains(t, cnsl.step.queue, stepOutput{content: "another sentence", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	cnsl.Debug("element one", "element two")
	assert.NotContains(t, cnsl.step.queue, stepOutput{content: "element oneelement two", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	cnsl.EndStep()

	// Since debug is disabled, debug level logs are not shown.
	assert.NotContains(t, defaultOut.String(), "> one sentence")
	assert.NotContains(t, defaultOut.String(), "> another sentence")
	assert.NotContains(t, defaultOut.String(), "> element oneelement two")
}

func TestGlobalDebuglnDisabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(defaultOut), WithDebug(false))
	StartStep("test task")

	cnsl.Debugln("one sentence")
	cnsl.Debugln("another sentence")
	// Since debug is disabled, even if a step is in progress, debug logs should not be queued.
	assert.NotContains(t, cnsl.step.queue, stepOutput{content: "one sentence\n", level: levelDebug})
	assert.NotContains(t, cnsl.step.queue, stepOutput{content: "another sentence\n", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	cnsl.Debugln("element one", "element two")
	assert.NotContains(t, cnsl.step.queue, stepOutput{content: "element one element two\n", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element one element two")

	cnsl.EndStep()

	// Since debug is disabled, debug level logs are not shown.
	assert.NotContains(t, defaultOut.String(), "> one sentence")
	assert.NotContains(t, defaultOut.String(), "> another sentence")
	assert.NotContains(t, defaultOut.String(), "> element one element two")
}

func TestGlobalDebugfDisabledWithStep(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(defaultOut), WithDebug(false))
	StartStep("test task")

	cnsl.Debugf("one sentence")
	cnsl.Debugf("another sentence")
	// Since debug is disabled, even if a step is in progress, debug logs should not be queued.
	assert.NotContains(t, cnsl.step.queue, stepOutput{content: "one sentence", level: levelDebug})
	assert.NotContains(t, cnsl.step.queue, stepOutput{content: "another sentence", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "one sentenceanother sentence")

	cnsl.Debugf("element one%s", "element two")
	assert.NotContains(t, cnsl.step.queue, stepOutput{content: "element oneelement two", level: levelDebug})
	assert.NotContains(t, defaultOut.String(), "element oneelement two")

	cnsl.EndStep()

	// Since debug is disabled, debug level logs are not shown.
	assert.NotContains(t, defaultOut.String(), "> one sentence")
	assert.NotContains(t, defaultOut.String(), "> another sentence")
	assert.NotContains(t, defaultOut.String(), "> element oneelement two")
}

func TestGlobalErrorWithStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(ioutil.Discard), WithErrorOutput(errorOut))
	StartStep("test task")

	cnsl.Error("one sentence")
	cnsl.Error("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "one sentence", level: levelError})
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "another sentence", level: levelError})
	assert.NotContains(t, errorOut.String(), "one sentenceanother sentence")

	cnsl.Error("element one", "element two")
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "element oneelement two", level: levelError})
	assert.NotContains(t, errorOut.String(), "element oneelement two")

	cnsl.EndStep()

	assert.Contains(t, errorOut.String(), "> one sentence")
	assert.Contains(t, errorOut.String(), "> another sentence")
	assert.Contains(t, errorOut.String(), "> element oneelement two")
}

func TestGlobalErrorlnWithStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(ioutil.Discard), WithErrorOutput(errorOut))
	StartStep("test task")

	cnsl.Errorln("one sentence")
	cnsl.Errorln("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "one sentence\n", level: levelError})
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "another sentence\n", level: levelError})
	assert.NotContains(t, errorOut.String(), "one sentenceanother sentence")

	cnsl.Errorln("element one", "element two")
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "element one element two\n", level: levelError})
	assert.NotContains(t, errorOut.String(), "element one element two")

	cnsl.EndStep()

	assert.Contains(t, errorOut.String(), "> one sentence")
	assert.Contains(t, errorOut.String(), "> another sentence")
	assert.Contains(t, errorOut.String(), "> element one element two")
}

func TestGlobalErrorfWithStep(t *testing.T) {
	errorOut := &bytes.Buffer{}

	SetGlobalOptions(WithDefaultOutput(ioutil.Discard), WithErrorOutput(errorOut))
	StartStep("test task")

	cnsl.Errorf("one sentence")
	cnsl.Errorf("another sentence")
	// Since a step is in progress, outputs should be queued and not printed.
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "one sentence", level: levelError})
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "another sentence", level: levelError})
	assert.NotContains(t, errorOut.String(), "one sentenceanother sentence")

	cnsl.Errorf("element one%s", "element two")
	assert.Contains(t, cnsl.step.queue, stepOutput{content: "element oneelement two", level: levelError})
	assert.NotContains(t, errorOut.String(), "element oneelement two")

	cnsl.EndStep()

	assert.Contains(t, errorOut.String(), "> one sentence")
	assert.Contains(t, errorOut.String(), "> another sentence")
	assert.Contains(t, errorOut.String(), "> element oneelement two")
}
