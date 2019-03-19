package disgo

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

var globalTerm *Terminal

func init() {
	globalTerm = NewTerminal()
}

// Terminal represents a disgo Terminal.
// It can write to and prompt users in a command-line interface.
type Terminal struct {
	// Writer on which the Info and Debug outputs are written.
	defaultOutput io.Writer
	// Writer on which the Error outputs are written.
	errorOutput io.Writer
	// Reader from which the user's response to the prompt is
	// read.
	reader *bufio.Reader

	// Current step that is in progress. When a task is in
	// progress, outputs are queued and will be printed once
	// the task is over.
	step *step

	// Whether or not Debug outputs are enabled. If this is
	// disabled, Debug outputs are not shown to the user.
	debug bool

	// Whether or not this terminal should be interactive. If this is
	// set to false, the users will never be prompted and calls to prompting
	// methods will always return default values.
	// prompts will never prompt the user and always
	// return default values. This can be useful for running this code
	// outside of a TTY for example.
	interactive bool
}

// NewTerminal creates a new Terminal.
func NewTerminal(options ...func(*Terminal)) *Terminal {
	term := Terminal{
		defaultOutput: os.Stdout,
		errorOutput:   os.Stderr,
		reader:        bufio.NewReader(os.Stdin),
		interactive:   true,
	}

	for _, option := range options {
		option(&term)
	}

	return &term
}

// WithDebug enables or disables the terminal debug mode.
func WithDebug(enabled bool) func(*Terminal) {
	return func(term *Terminal) {
		term.debug = enabled
	}
}

// WithDefaultOutput sets the default writer on the terminal.
func WithDefaultOutput(writer io.Writer) func(*Terminal) {
	return func(term *Terminal) {
		term.defaultOutput = writer
	}
}

// WithErrorOutput sets the error writer on the terminal.
func WithErrorOutput(writer io.Writer) func(*Terminal) {
	return func(term *Terminal) {
		term.errorOutput = writer
	}
}

// WithReader sets the reader on the Terminal. By default, if this
// option is not used, the default reader will be os.Stdin.
func WithReader(reader io.Reader) func(*Terminal) {
	return func(term *Terminal) {
		term.reader = bufio.NewReader(reader)
	}
}

// WithColors sets the use of colors in the terminal. By default, whether or not
// colors are enabled depends on the user's TTY, but this option can be used
// to force colors to be enabled or disabled.
func WithColors(enabled bool) func(*Terminal) {
	return func(_ *Terminal) {
		color.NoColor = !enabled
	}
}

// WithInteractive enables or disables the terminal interactive mode.
func WithInteractive(enabled bool) func(*Terminal) {
	return func(term *Terminal) {
		term.interactive = enabled
	}
}

// SetTerminalOptions applies options to the global terminal.
func SetTerminalOptions(options ...func(*Terminal)) {
	for _, option := range options {
		option(globalTerm)
	}
}

// Info writes an info output on the terminal's default writer.
func (t Terminal) Info(a ...interface{}) {
	if t.step != nil {
		t.step.pushInfo(fmt.Sprint(a...))
		return
	}

	fmt.Fprint(t.defaultOutput, a...)
}

// Info writes an info output on the global terminal's default writer.
func Info(a ...interface{}) {
	globalTerm.Info(a...)
}

// Infoln writes an info output on the terminal's default writer
// and appends a newline to its input.
func (t Terminal) Infoln(a ...interface{}) {
	if t.step != nil {
		t.step.pushInfo(fmt.Sprintln(a...))
		return
	}

	fmt.Fprintln(t.defaultOutput, a...)
}

// Infoln writes an info output on the global terminal's default writer
// and appends a newline to its input.
func Infoln(a ...interface{}) {
	globalTerm.Infoln(a...)
}

// Infof formats according to a format specifier and writes
// to the terminal's default writer.
func (t Terminal) Infof(format string, a ...interface{}) {
	if t.step != nil {
		t.step.pushInfo(fmt.Sprintf(format, a...))
		return
	}

	fmt.Fprintf(t.defaultOutput, format, a...)
}

// Infof formats according to a format specifier and writes
// to the global terminal's default writer.
func Infof(format string, a ...interface{}) {
	globalTerm.Infof(format, a...)
}

// Debug writes a debug output on the terminal's default writer if
// the debug outputs are enabled.
func (t Terminal) Debug(a ...interface{}) {
	if !t.debug {
		return
	}

	if t.step != nil {
		t.step.pushDebug(fmt.Sprint(a...))
		return
	}

	fmt.Fprint(t.defaultOutput, a...)
}

// Debug writes a debug output on the global terminal's default writer if
// the debug outputs are enabled.
func Debug(a ...interface{}) {
	globalTerm.Debug(a...)
}

// Debugln writes a debug output on the terminal's default writer if
// the debug outputs are enabled and appends a newline to its input.
func (t Terminal) Debugln(a ...interface{}) {
	if !t.debug {
		return
	}

	if t.step != nil {
		t.step.pushDebug(fmt.Sprintln(a...))
		return
	}

	fmt.Fprintln(t.defaultOutput, a...)
}

// Debugln writes a debug output on the global terminal's default writer if
// the debug outputs are enabled and appends a newline to its input.
func Debugln(a ...interface{}) {
	globalTerm.Debugln(a...)
}

// Debugf formats according to a format specifier and writes
// to the terminal's default writer if the debug outputs are enabled.
func (t Terminal) Debugf(format string, a ...interface{}) {
	if !t.debug {
		return
	}

	if t.step != nil {
		t.step.pushDebug(fmt.Sprintf(format, a...))
		return
	}

	fmt.Fprintf(t.defaultOutput, format, a...)
}

// Debugf formats according to a format specifier and writes
// to the global terminal's default writer if the debug outputs are enabled.
func Debugf(format string, a ...interface{}) {
	globalTerm.Debugf(format, a...)
}

// Error writes an error output on the terminal's error writer.
func (t Terminal) Error(a ...interface{}) {
	if t.step != nil {
		t.step.pushError(fmt.Sprint(a...))
		return
	}

	fmt.Fprint(t.errorOutput, a...)
}

// Error writes an error output on the global terminal's error writer.
func Error(a ...interface{}) {
	globalTerm.Error(a...)
}

// Errorln writes an error output on the terminal's error writer.
// It appends a newline to its input.
func (t Terminal) Errorln(a ...interface{}) {
	if t.step != nil {
		t.step.pushError(fmt.Sprintln(a...))
		return
	}

	fmt.Fprintln(t.errorOutput, a...)
}

// Errorln writes an error output on the global terminal's error writer.
// It appends a newline to its input.
func Errorln(a ...interface{}) {
	globalTerm.Errorln(a...)
}

// Errorf formats according to a format specifier and writes
// to the terminal's error writer.
func (t Terminal) Errorf(format string, a ...interface{}) {
	if t.step != nil {
		t.step.pushError(fmt.Sprintf(format, a...))
		return
	}

	fmt.Fprintf(t.errorOutput, format, a...)
}

// Errorf formats according to a format specifier and writes
// to the global terminal's error writer.
func Errorf(format string, a ...interface{}) {
	globalTerm.Errorf(format, a...)
}
