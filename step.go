package disgo

import (
	"fmt"
	"strings"

	"github.com/Ullaakut/disgo/style"
)

type outputLevel int

const (
	levelDebug outputLevel = iota
	levelInfo
	levelError
)

type stepOutput struct {
	content string
	level   outputLevel
}

type step struct {
	queue []stepOutput
}

func (s *step) pushDebug(content string) {
	s.queue = append(s.queue, stepOutput{
		level:   levelDebug,
		content: content,
	})
}

func (s *step) pushInfo(content string) {
	s.queue = append(s.queue, stepOutput{
		level:   levelInfo,
		content: content,
	})
}

func (s *step) pushError(content string) {
	s.queue = append(s.queue, stepOutput{
		level:   levelError,
		content: content,
	})
}

// StartStep sets a step in the terminal, which prints
// the step's label and makes the terminal queue outputs
// until the step is ended or failed. If a step was already in
// progress, it is considered to have been ended successfully.
func (t *Terminal) StartStep(label string) {
	if t.step != nil {
		t.EndStep()
	}

	fmt.Fprint(t.defaultOutput, label, "...")

	t.step = &step{}
}

// StartStep sets a step in the global terminal, which prints
// the step's label and makes the terminal queue outputs
// until the step is ended or failed. If a step was already in
// progress, it is considered to have been ended successfully.
// Warning: This is not thread-safe.
func StartStep(label string) {
	globalTerm.StartStep(label)
}

// StartStepf sets a step in the terminal, which prints
// the step's label and makes the terminal queue outputs
// until the step is ended or failed. If a step was already in
// progress, it is considered to have been ended successfully.
func (t *Terminal) StartStepf(format string, a ...interface{}) {
	t.StartStep(fmt.Sprintf(format, a...))
}

// StartStepf sets a step in the terminal, which prints
// the step's label and makes the terminal queue outputs
// until the step is ended or failed. If a step was already in
// progress, it is considered to have been ended successfully.
// Warning: This is not thread-safe.
func StartStepf(format string, a ...interface{}) {
	globalTerm.StartStepf(format, a...)
}

// FailStep ends a step with a failure state. It then
// prints all of the outputs that were queued while the step
// was in progress, and returns the given error for error
// handling.
func (t *Terminal) FailStep(err error) error {
	if t.step == nil {
		return err
	}

	fmt.Fprintln(t.defaultOutput, style.Failure("ko"))

	t.printQueue()

	t.step = nil
	return err
}

// FailStep ends a step with a failure state. It then
// prints all of the outputs that were queued while the step
// was in progress, and returns the given error for error
// handling.
// Warning: This is not thread-safe.
func FailStep(err error) error {
	return globalTerm.FailStep(err)
}

// FailStepf ends a step with a failure state. It then
// prints all of the outputs that were queued while the step
// was in progress, and returns an error created from the
// given format and arguments.
func (t *Terminal) FailStepf(format string, a ...interface{}) error {
	return t.FailStep(fmt.Errorf(format, a...))
}

// FailStepf ends a step with a failure state on the global.
// terminal. It then prints all of the outputs that were queued
// while the step was in progress, and returns an error
// created from the given format and arguments.
// Warning: This is not thread-safe.
func FailStepf(format string, a ...interface{}) error {
	return globalTerm.FailStepf(format, a...)
}

// EndStep ends a step with a success state. It then
// prints all of the outputs that were queued while the step
// was in progress.
func (t *Terminal) EndStep() {
	if t.step == nil {
		return
	}

	fmt.Fprintln(t.defaultOutput, style.Success("ok"))

	t.printQueue()

	t.step = nil
}

// EndStep ends a step with a success state on the global.
// terminal. It then prints all of the outputs that were queued
// while the step was in progress.
// Warning: This is not thread-safe.
func EndStep() {
	globalTerm.EndStep()
}

// printQueue prints all of the outputs that were queued during a step,
// neatly indented after that step is ended.
func (t Terminal) printQueue() {
	for _, output := range t.step.queue {
		// Trim the last newline from the output's content.
		output.content = strings.TrimSuffix(output.content, "\n")
		// Indent the content to make it obvious that it is
		// part of a step's processing.
		output.content = strings.Replace(output.content, "\n", "\n    ", -1)

		// Print the output on the proper writer.
		switch output.level {
		case levelDebug:
			if t.debug {
				fmt.Fprintf(t.defaultOutput, "  > %s\n", style.Trace(output.content))
			}
		case levelInfo:
			fmt.Fprintf(t.defaultOutput, "  > %s\n", style.Trace(output.content))
		case levelError:
			fmt.Fprintf(t.errorOutput, "  > %s\n", style.Failure(output.content))
		}
	}
}
