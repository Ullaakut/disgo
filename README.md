# Disgo

<p align="center">
    <img width="400px" src="images/logo-with-label-light.png">
</p>
<p align="center">
    <a href="#license">
        <img src="https://img.shields.io/badge/license-MIT-blue.svg?style=flat" />
    </a>
    <a href="https://godoc.org/github.com/Ullaakut/disgo">
        <img src="https://godoc.org/github.com/Ullaakut/disgo?status.svg" />
    </a>
    <a href="https://goreportcard.com/report/github.com/Ullaakut/disgo">
        <img src="https://goreportcard.com/badge/github.com/Ullaakut/disgo" />
    </a>
    <a href="https://github.com/Ullaakut/disgo/releases/latest">
        <img src="https://img.shields.io/github/release/Ullaakut/disgo.svg?style=flat" />
    </a>
</p>

Simple console output library for Go command-line interfaces.

Disgo provides four essential features for most user-friendly CLI applications:

1. Simple output levels (in `github.com/ullaakut/disgo/console`)
2. Output formatting (in `github.com/ullaakut/disgo/console`)
3. Step-by-step outputs (in `github.com/ullaakut/disgo/console`)
4. Simple user prompting (in `github.com/ullaakut/disgo/prompter`)

## Console package

The console package provides an idiomatic library to build user-friendly command-line interfaces.

You can use it globally within your application, or you can instantiate your own `Console`.

### Console options

When creating a `Console` instance or when using the global `Console` that this package provides, you might want to give it some options, such as:

- **`WithDebug`**, which lets you enable or disable the debug output _(it is disabled by default)_
- **`WithDefaultWriter`**, which lets you specify an `io.Writer` on which `Debug` and `Info`-level outputs should be written _(it is set to `os.Stdout` by default)_
- **`WithErrorWriter`**, which lets you specify an `io.Writer` on which `Error`-level outputs should be written _(it is set to `os.Stderr` by default)_
- **`WithColors`**, which lets you explicitely enable or disable colors in your output _(it is enabled by default)_

You can either pass those options to `console.New()` when creating a `Console` instance, like so:

```go
    myConsole := console.New(console.WithDebug(true))
```

Or, if you are using the global console, you will simply need to call the `SetGlobalOptions` function:

```go
    console.SetGlobalOptions(console.WithDebug(true))
```

### Writing to the console

Now that your console is set up, you can start writing on it. Printing functions behave idiomatically, like you would expect.

Here is how to use them on a local console:

```go
    // All of those give the same output:
    // "Number of days in a year: 365" followed by a newline.
    myConsole.Infoln("Number of days in a year:", 365)
    myConsole.Infof("Number of days in a year: %d\n", 365)
    myConsole.Info("Number of days in a year: 365\n")

    // Debug methods are similar to info, except that they are not printed
    // if debug outputs are not enabled on the console.
    myConsole.Debugln("Number of days in a year:", 365)
    myConsole.Debugf("Number of days in a year: %d\n", 365)
    myConsole.Debug("Number of days in a year: 365\n")


    // Error methods are similar to info, except that they are written on
    // the error writer (os.Stderr by default).
    myConsole.Debugln("Number of days in a year:", 365)
    myConsole.Debugf("Number of days in a year: %d\n", 365)
    myConsole.Debug("Number of days in a year: 365\n")
```

When using the global console, simply call the console printing functions directly:

```go
    // All of those give the same output:
    // "Number of days in a year: 365" followed by a newline.
    console.Infoln("Number of days in a year:", 365)
    console.Infof("Number of days in a year: %d\n", 365)
    console.Info("Number of days in a year: 365\n")

    // Debug methods are similar to info, except that they are not printed
    // if debug outputs are not enabled on the console.
    console.Debugln("Number of days in a year:", 365)
    console.Debugf("Number of days in a year: %d\n", 365)
    console.Debug("Number of days in a year: 365\n")


    // Error methods are similar to info, except that they are written on
    // the error writer (os.Stderr by default).
    console.Debugln("Number of days in a year:", 365)
    console.Debugf("Number of days in a year: %d\n", 365)
    console.Debug("Number of days in a year: 365\n")
```

### Output Formatting

Another feature provided by this package is **output formatting**. It exposes six different output formats, which will print an output with a specific color, font-weight and font-style depending on what the output's content should convey to the user. For example, if you want to attract a user's attention to an error, you might use the `console.Failure()` formatting function, like so:

```go
    if err := validateConfiguration; err != nil {
        console.Errorln("Invalid configuration detected:", console.Failure(err))
        return err
    }
```

<p align="center">
    <img src="images/output_failure.png" />
</p>

Other output formats include `Success`, `Trace`, `Important` and `Link`.

<p align="center">
    <img src="images/output_all.png" />
</p>

You can of course combine those formats in elegant ways, like shown in the [examples](#examples) section.

### Step-by-step processes

A lot of command-line interfaces describe step-by-step processes to the user, but it's difficult to combine clean code, clear output and elegant user interfaces. Disgo attempts to solve that problem by associating _steps_ to its console.

For example, when beginning a task, you can use `StartStep` and specify the description of that step. Then, until that task is over, all calls to Disgo's printing functions will be queued. Once the task is complete (by calling `EndStep`, `FailStep` or by starting another step with `StartStep`), the task status is printed and all of the outputs that were queued during the task are printed with an indent, under the task, like so:

<p align="center">
    <img src="images/example_step_by_step.png" />
</p>

It is also important to note that `FailStep` and `FailStepf` can return errors at the same time as they report a step as having failed. This allows you to write:

```go
    console.StartStep("Doing something")
    if err := doSomething(); err != nil {
        return console.FailStepf("unable to do something: %v", err)
    }
```

Instead of having to call `FailStep` in your error handling before returning. You are still free to do so if you prefer, though.

Using the global console for step management is not thread-safe though, as it was built with simplicity in mind and can only handle one step at a time.

## Examples

Here are a few examples of Disgo's output, using this repository's example program:

<p align="center">
    <img src="images/example_success.png" /><br/>
    <img src="images/example_failure.png" /><br/>
    <img src="images/example_failure_prompt.png" />
</p>

Disgo is also used in other projects, here are some examples:

<p align="center">
    <img src="https://raw.githubusercontent.com/Ullaakut/cameradar/master/images/Cameradar.gif" />
    <img src="https://raw.githubusercontent.com/Ullaakut/Gorsair/master/images/gorsair.gif" />
</p>
