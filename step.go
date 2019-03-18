package disgo

import (
	"fmt"
	"strings"
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

// StartStep sets a step in the console, which prints
// the step's label and makes the console queue outputs
// until the step is or failed. If a step was already in
// progress, it is considered to have been ended successfully.
func (c *Console) StartStep(label string) {
	if c.step != nil {
		c.EndStep()
	}

	fmt.Fprint(c.defaultOutput, Trace(label, "..."))

	c.step = &step{}
}

// StartStep sets a step in the global console, which prints
// the step's label and makes the console queue outputs
// until the step is or failed. If a step was already in
// progress, it is considered to have been ended successfully.
// Warning: This is not thread-safe.
func StartStep(label string) {
	cnsl.StartStep(label)
}

// StartStepf sets a step in the console, which prints
// the step's label and makes the console queue outputs
// until the step is or failed. If a step was already in
// progress, it is considered to have been ended successfully.
func (c *Console) StartStepf(format string, a ...interface{}) {
	c.StartStep(fmt.Sprintf(format, a...))
}

// StartStepf sets a step in the console, which prints
// the step's label and makes the console queue outputs
// until the step is or failed. If a step was already in
// progress, it is considered to have been ended successfully.
// Warning: This is not thread-safe.
func StartStepf(format string, a ...interface{}) {
	cnsl.StartStepf(format, a...)
}

// FailStep ends a step with a failure state. It then
// prints all of the outputs that were queued while the step
// was in progress, and returns the given error for error
// handling.
func (c *Console) FailStep(err error) error {
	if c.step == nil {
		return err
	}

	fmt.Fprintln(c.defaultOutput, Failure("ko"))

	c.printQueue()

	c.step = nil
	return err
}

// FailStep ends a step with a failure state. It then
// prints all of the outputs that were queued while the step
// was in progress, and returns the given error for error
// handling.
// Warning: This is not thread-safe.
func FailStep(err error) error {
	return cnsl.FailStep(err)
}

// FailStepf ends a step with a failure state. It then
// prints all of the outputs that were queued while the step
// was in progress, and returns an error created from the
// given format and arguments.
func (c *Console) FailStepf(format string, a ...interface{}) error {
	return c.FailStep(fmt.Errorf(format, a...))
}

// FailStepf ends a step with a failure state on the global.
// console. It then prints all of the outputs that were queued
// while the step was in progress, and returns an error
// created from the given format and arguments.
// Warning: This is not thread-safe.
func FailStepf(format string, a ...interface{}) error {
	return cnsl.FailStepf(format, a...)
}

// EndStep ends a step with a success state. It then
// prints all of the outputs that were queued while the step
// was in progress.
func (c *Console) EndStep() {
	if c.step == nil {
		return
	}

	fmt.Fprintln(c.defaultOutput, Success("ok"))

	c.printQueue()

	c.step = nil
}

// EndStep ends a step with a success state on the global.
// console. It then prints all of the outputs that were queued
// while the step was in progress.
// Warning: This is not thread-safe.
func EndStep() {
	cnsl.EndStep()
}

// printQueue prints all of the outputs that were queued during a step,
// neatly indented after that step is ended.
func (c Console) printQueue() {
	for _, output := range c.step.queue {
		// Trim the last newline from the output's content.
		output.content = strings.TrimSuffix(output.content, "\n")
		// Indent the content to make it obvious that it is
		// part of a step's processing.
		output.content = strings.Replace(output.content, "\n", "\n    ", -1)

		// Print the output on the proper writer.
		switch output.level {
		case levelDebug:
			if c.debug {
				fmt.Fprintf(c.defaultOutput, "  > %s\n", Trace(output.content))
			}
		case levelInfo:
			fmt.Fprintf(c.defaultOutput, "  > %s\n", Trace(output.content))
		case levelError:
			fmt.Fprintf(c.errorOutput, "  > %s\n", Failure(output.content))
		}
	}
}
