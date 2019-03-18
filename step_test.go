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

	console := &Console{
		defaultOutput: defaultOut,
	}

	// Outputs should be queued during a task, and printed afterwards.
	console.Infoln("Test step-by-step process")
	console.StartStep("Simulated task #1")
	console.Infoln("25%")
	console.Infoln("50%")
	console.Infoln("75%")
	console.Infoln("100%")
	console.EndStep()

	assert.Equal(t, "Test step-by-step process\nSimulated task #1...ok\n  > 25%\n  > 50%\n  > 75%\n  > 100%\n", defaultOut.String())
}

func TestStartStepfQueuesOutputs(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	console := &Console{
		defaultOutput: defaultOut,
	}

	// Outputs should be queued during a task, and printed afterwards.
	console.Infoln("Test step-by-step process")
	console.StartStepf("Simulated task #%d", 1)
	console.Infoln("25%")
	console.Infoln("50%")
	console.Infoln("75%")
	console.Infoln("100%")
	console.EndStep()

	assert.Equal(t, "Test step-by-step process\nSimulated task #1...ok\n  > 25%\n  > 50%\n  > 75%\n  > 100%\n", defaultOut.String())
}

func TestStepSystemStopsPreviousTaskWhenStartingNewOne(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	console := &Console{
		defaultOutput: defaultOut,
	}

	// Starting a task while another is in progress should consider the previous
	// task successfully ended.
	console.StartStep("Simulated task #1")
	console.StartStep("Simulated task #2")

	assert.Equal(t, "Simulated task #1...ok\nSimulated task #2...", defaultOut.String())
}

func TestFailStepReturnsError(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	console := &Console{
		defaultOutput: defaultOut,
	}

	console.StartStep("Simulated task #1")

	dummyError := errors.New("dummy error")
	err := console.FailStep(dummyError)

	// FailStep should return the given error without modifications.
	assert.Equal(t, dummyError, err)
}

func TestFailStepfCreatesError(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	console := &Console{
		defaultOutput: defaultOut,
	}

	console.StartStep("Simulated task #1")

	expectedError := fmt.Errorf("task %d failed", 1)
	err := console.FailStepf("task %d failed", 1)

	// FailStep should return the given error without modifications.
	assert.Equal(t, expectedError, err)
}

func TestFailStepOutput(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	console := &Console{
		defaultOutput: defaultOut,
	}

	// Starting a task while another is in progress should consider the previous
	// task successfully ended.
	console.StartStep("Simulated task #1")
	_ = console.FailStep(nil)

	assert.Equal(t, "Simulated task #1...ko\n", defaultOut.String())
}

func TestFailStepfOutput(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	console := &Console{
		defaultOutput: defaultOut,
	}

	// Starting a task while another is in progress should consider the previous
	// task successfully ended.
	console.StartStep("Simulated task #1")
	_ = console.FailStepf("task %d failed", 1)

	assert.Equal(t, "Simulated task #1...ko\n", defaultOut.String())
}

//*********************//
// Test Global Console //
//*********************//

func TestGlobalStartStepQueuesOutputs(t *testing.T) {
	defaultOut := &bytes.Buffer{}

	cnsl = &Console{
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

	cnsl = &Console{
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

	cnsl = &Console{
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

	cnsl = &Console{
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

	cnsl = &Console{
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

	cnsl = &Console{
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

	cnsl = &Console{
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
	cnsl = &Console{}

	defer (func() {
		if r := recover(); r != nil {
			assert.FailNow(t, "Calling EndStep while no step is set should not panic")
		}
	})()

	EndStep()
}

func TestFailStepReturnsWhenNoStep(t *testing.T) {
	cnsl = &Console{}

	defer (func() {
		if r := recover(); r != nil {
			assert.FailNow(t, "Calling FailStep while no step is set should not panic")
		}
	})()

	FailStepf("")
	FailStep(nil)
}
