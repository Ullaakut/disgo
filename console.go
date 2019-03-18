package disgo

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

var cnsl *Console

func init() {
	cnsl = NewConsole()
}

// Console represents a disgo Console.
// It writes the output on a given io.Writer
// and can toggle debug outputs and have an error writer.
type Console struct {
	// Writer on which the Info and Debug outputs are written.
	defaultOutput io.Writer
	// Writer on which the Error outputs are written.
	errorOutput io.Writer

	// Current step that is in progress. When a task is in
	// progress, outputs are queued and will be printed once
	// the task is over.
	step *step

	// Whether or not Debug outputs are enabled. If this is
	// disabled, Debug outputs are not shown to the user.
	debug bool
}

// NewConsole creates a new Console and binds the given writer to its outputs.
func NewConsole(options ...func(*Console)) *Console {
	console := Console{
		defaultOutput: os.Stdout,
		errorOutput:   os.Stderr,
	}

	for _, option := range options {
		option(&console)
	}

	return &console
}

// WithDebug enables or disables the console debug mode.
func WithDebug(enabled bool) func(*Console) {
	return func(console *Console) {
		console.debug = enabled
	}
}

// WithDefaultOutput sets the default writer on the console.
func WithDefaultOutput(writer io.Writer) func(*Console) {
	return func(console *Console) {
		console.defaultOutput = writer
	}
}

// WithErrorOutput sets the error writer on the console.
func WithErrorOutput(writer io.Writer) func(*Console) {
	return func(console *Console) {
		console.errorOutput = writer
	}
}

// WithColors sets the use of colors in the console. By default, whether or not
// colors are enabled depends on the user's TTY, but this option can be used
// to force colors to be enabled or disabled.
func WithColors(enabled bool) func(*Console) {
	return func(_ *Console) {
		color.NoColor = !enabled
	}
}

// SetupGlobalConsole applies options to the global console.
func SetupGlobalConsole(options ...func(*Console)) {
	for _, option := range options {
		option(cnsl)
	}
}

// Info writes an info output on the console's default writer.
func (c Console) Info(a ...interface{}) {
	if c.step != nil {
		c.step.pushInfo(fmt.Sprint(a...))
		return
	}

	fmt.Fprint(c.defaultOutput, a...)
}

// Info writes an info output on the global console's default writer.
func Info(a ...interface{}) {
	cnsl.Info(a...)
}

// Infoln writes an info output on the console's default writer
// and appends a newline to its input.
func (c Console) Infoln(a ...interface{}) {
	if c.step != nil {
		c.step.pushInfo(fmt.Sprintln(a...))
		return
	}

	fmt.Fprintln(c.defaultOutput, a...)
}

// Infoln writes an info output on the global console's default writer
// and appends a newline to its input.
func Infoln(a ...interface{}) {
	cnsl.Infoln(a...)
}

// Infof formats according to a format specifier and writes
// to the console's default writer.
func (c Console) Infof(format string, a ...interface{}) {
	if c.step != nil {
		c.step.pushInfo(fmt.Sprintf(format, a...))
		return
	}

	fmt.Fprintf(c.defaultOutput, format, a...)
}

// Infof formats according to a format specifier and writes
// to the global console's default writer.
func Infof(format string, a ...interface{}) {
	cnsl.Infof(format, a...)
}

// Debug writes a debug output on the console's default writer if
// the debug outputs are enabled.
func (c Console) Debug(a ...interface{}) {
	if !c.debug {
		return
	}

	if c.step != nil {
		c.step.pushDebug(fmt.Sprint(a...))
		return
	}

	fmt.Fprint(c.defaultOutput, a...)
}

// Debug writes a debug output on the global console's default writer if
// the debug outputs are enabled.
func Debug(a ...interface{}) {
	cnsl.Debug(a...)
}

// Debugln writes a debug output on the console's default writer if
// the debug outputs are enabled and appends a newline to its input.
func (c Console) Debugln(a ...interface{}) {
	if !c.debug {
		return
	}

	if c.step != nil {
		c.step.pushDebug(fmt.Sprintln(a...))
		return
	}

	fmt.Fprintln(c.defaultOutput, a...)
}

// Debugln writes a debug output on the global console's default writer if
// the debug outputs are enabled and appends a newline to its input.
func Debugln(a ...interface{}) {
	cnsl.Debugln(a...)
}

// Debugf formats according to a format specifier and writes
// to the console's default writer if the debug outputs are enabled.
func (c Console) Debugf(format string, a ...interface{}) {
	if !c.debug {
		return
	}

	if c.step != nil {
		c.step.pushDebug(fmt.Sprintf(format, a...))
		return
	}

	fmt.Fprintf(c.defaultOutput, format, a...)
}

// Debugf formats according to a format specifier and writes
// to the global console's default writer if the debug outputs are enabled.
func Debugf(format string, a ...interface{}) {
	cnsl.Debugf(format, a...)
}

// Error writes an error output on the console's error writer.
func (c Console) Error(a ...interface{}) {
	if c.step != nil {
		c.step.pushError(fmt.Sprint(a...))
		return
	}

	fmt.Fprint(c.errorOutput, a...)
}

// Error writes an error output on the global console's error writer.
func Error(a ...interface{}) {
	cnsl.Error(a...)
}

// Errorln writes an error output on the console's error writer.
// It appends a newline to its input.
func (c Console) Errorln(a ...interface{}) {
	if c.step != nil {
		c.step.pushError(fmt.Sprintln(a...))
		return
	}

	fmt.Fprintln(c.errorOutput, a...)
}

// Errorln writes an error output on the global console's error writer.
// It appends a newline to its input.
func Errorln(a ...interface{}) {
	cnsl.Errorln(a...)
}

// Errorf formats according to a format specifier and writes
// to the console's error writer.
func (c Console) Errorf(format string, a ...interface{}) {
	if c.step != nil {
		c.step.pushError(fmt.Sprintf(format, a...))
		return
	}

	fmt.Fprintf(c.errorOutput, format, a...)
}

// Errorf formats according to a format specifier and writes
// to the global console's error writer.
func Errorf(format string, a ...interface{}) {
	cnsl.Errorf(format, a...)
}
