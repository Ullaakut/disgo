package disgo

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartStepQueuesOutputs(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
	}

	// Outputs should be queued during a task, and printed afterwards.
	term.Infoln("Test step-by-step process")
	term.StartStep("Simulated task #1")
	term.Infoln("25%")
	term.Infoln("50%")
	term.Infoln("75%")
	term.Infoln("100%")
	term.EndStep()

	assert.Equal(t, "Test step-by-step process\nSimulated task #1...ok\n  > 25%\n  > 50%\n  > 75%\n  > 100%\n", defaultOut.String())
}

func TestStartStepfQueuesOutputs(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
	}

	// Outputs should be queued during a task, and printed afterwards.
	term.Infoln("Test step-by-step process")
	term.StartStepf("Simulated task #%d", 1)
	term.Infoln("25%")
	term.Infoln("50%")
	term.Infoln("75%")
	term.Infoln("100%")
	term.EndStep()

	assert.Equal(t, "Test step-by-step process\nSimulated task #1...ok\n  > 25%\n  > 50%\n  > 75%\n  > 100%\n", defaultOut.String())
}

func TestStepSystemStopsPreviousTaskWhenStartingNewOne(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
	}

	// Starting a task while another is in progress should consider the previous
	// task successfully ended.
	term.StartStep("Simulated task #1")
	term.StartStep("Simulated task #2")

	assert.Equal(t, "Simulated task #1...ok\nSimulated task #2...", defaultOut.String())
}

func TestFailStepReturnsError(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
	}

	term.StartStep("Simulated task #1")

	dummyError := errors.New("dummy error")
	err := term.FailStep(dummyError)

	// FailStep should return the given error without modifications.
	assert.Equal(t, dummyError, err)
}

func TestFailStepfCreatesError(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
	}

	term.StartStep("Simulated task #1")

	expectedError := fmt.Errorf("task %d failed", 1)
	err := term.FailStepf("task %d failed", 1)

	// FailStep should return the given error without modifications.
	assert.Equal(t, expectedError, err)
}

func TestFailStepOutput(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
	}

	// Starting a task while another is in progress should consider the previous
	// task successfully ended.
	term.StartStep("Simulated task #1")
	_ = term.FailStep(nil)

	assert.Equal(t, "Simulated task #1...ko\n", defaultOut.String())
}

func TestFailStepfOutput(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	term := &Terminal{
		defaultOutput: defaultOut,
	}

	// Starting a task while another is in progress should consider the previous
	// task successfully ended.
	term.StartStep("Simulated task #1")
	_ = term.FailStepf("task %d failed", 1)

	assert.Equal(t, "Simulated task #1...ko\n", defaultOut.String())
}

//*********************//
// Test Global Terminal //
//*********************//

func TestGlobalStartStepQueuesOutputs(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	globalTerm = &Terminal{
		defaultOutput: defaultOut,
	}

	// Outputs should be queued during a task, and printed afterwards.
	Infoln("Test step-by-step process")
	StartStep("Simulated task #1")
	Infoln("25%")
	Infoln("50%")
	Infoln("75%")
	Infoln("100%")
	EndStep()

	assert.Equal(t, "Test step-by-step process\nSimulated task #1...ok\n  > 25%\n  > 50%\n  > 75%\n  > 100%\n", defaultOut.String())
}

func TestGlobalStartStepfQueuesOutputs(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	globalTerm = &Terminal{
		defaultOutput: defaultOut,
	}

	// Outputs should be queued during a task, and printed afterwards.
	Infoln("Test step-by-step process")
	StartStepf("Simulated task #%d", 1)
	Infoln("25%")
	Infoln("50%")
	Infoln("75%")
	Infoln("100%")
	EndStep()

	assert.Equal(t, "Test step-by-step process\nSimulated task #1...ok\n  > 25%\n  > 50%\n  > 75%\n  > 100%\n", defaultOut.String())
}

func TestGlobalStepSystemStopsPreviousTaskWhenStartingNewOne(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	globalTerm = &Terminal{
		defaultOutput: defaultOut,
	}

	// Starting a task while another is in progress should consider the previous
	// task successfully ended.
	StartStep("Simulated task #1")
	StartStep("Simulated task #2")

	assert.Equal(t, "Simulated task #1...ok\nSimulated task #2...", defaultOut.String())
}

func TestGlobalFailStepReturnsError(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	globalTerm = &Terminal{
		defaultOutput: defaultOut,
	}

	StartStep("Simulated task #1")

	dummyError := errors.New("dummy error")
	err := FailStep(dummyError)

	// FailStep should return the given error without modifications.
	assert.Equal(t, dummyError, err)
}

func TestGlobalFailStepfCreatesError(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	globalTerm = &Terminal{
		defaultOutput: defaultOut,
	}

	StartStep("Simulated task #1")

	expectedError := fmt.Errorf("task %d failed", 1)
	err := FailStepf("task %d failed", 1)

	// FailStep should return the given error without modifications.
	assert.Equal(t, expectedError, err)
}

func TestGlobalFailStepOutput(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	globalTerm = &Terminal{
		defaultOutput: defaultOut,
	}

	// Starting a task while another is in progress should consider the previous
	// task successfully ended.
	StartStep("Simulated task #1")
	_ = FailStep(nil)

	assert.Equal(t, "Simulated task #1...ko\n", defaultOut.String())
}

func TestGlobalFailStepfOutput(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	globalTerm = &Terminal{
		defaultOutput: defaultOut,
	}

	// Starting a task while another is in progress should consider the previous
	// task successfully ended.
	StartStep("Simulated task #1")
	_ = FailStepf("task %d failed", 1)

	assert.Equal(t, "Simulated task #1...ko\n", defaultOut.String())
}

//*************//
// Other tests //
//*************//

func TestEndStepReturnsWhenNoStep(t *testing.T) {
	globalTerm = &Terminal{}

	defer (func() {
		if r := recover(); r != nil {
			assert.FailNow(t, "Calling EndStep while no step is set should not panic")
		}
	})()

	EndStep()
}

func TestFailStepReturnsWhenNoStep(t *testing.T) {
	globalTerm = &Terminal{}

	defer (func() {
		if r := recover(); r != nil {
			assert.FailNow(t, "Calling FailStep while no step is set should not panic")
		}
	})()

	FailStepf("")
	FailStep(nil)
}
