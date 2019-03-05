package colog

import (
	"fmt"
	"io"

	"github.com/fatih/color"
)

// Logger represents a colog Logger.
// It writes the output on a given io.Writer
// and can toggle debug logs.
type Logger struct {
	output io.Writer

	debug       bool
	logFilePath string
}

// NewLogger creates a new Logger.
func NewLogger(w io.Writer, options ...func(*Logger)) (*Logger, error) {
	logger := Logger{
		output: w,
	}

	for _, option := range options {
		option(&logger)
	}

	return &logger, nil
}

// WithDebug sets the logger debug mode to true.
func WithDebug() func(*Logger) {
	return func(logger *Logger) {
		logger.debug = true
	}
}

// WithNoColors disables colors in the logger.
func WithNoColors() func(*Logger) {
	return func(_ *Logger) {
		color.NoColor = true
	}
}

// Info writes an info log on the logger's writers.
func (l Logger) Info(a ...interface{}) {
	fmt.Fprint(l.output, a...)
}

// Infoln writes an info log on the logger's writers
// and appends a newline to its input.
func (l Logger) Infoln(a ...interface{}) {
	fmt.Fprintln(l.output, a...)
}

// Infof formats according to a format specifier and writes
// to the logger's writers.
func (l Logger) Infof(format string, a ...interface{}) {
	fmt.Fprintf(l.output, format, a...)
}

// Debug writes a debug log on the logger's writers  if the
// debug logs are enabled.
func (l Logger) Debug(a ...interface{}) {
	if !l.debug {
		return
	}

	fmt.Fprint(l.output, a...)
}

// Debugln writes a debug log on the logger's writers  if the
// debug logs are enabled and appends a newline to its input.
func (l Logger) Debugln(a ...interface{}) {
	if !l.debug {
		return
	}

	fmt.Fprintln(l.output, a...)
}

// Debugf formats according to a format specifier and writes
// to the logger's writers if the debug logs are enabled.
func (l Logger) Debugf(format string, a ...interface{}) {
	if !l.debug {
		return
	}

	fmt.Fprintf(l.output, format, a...)
}
