# mylogger

`mylogger` is a simple logging utility for Go projects. Using go's `slog` but
with separated outputs for `stdout` and `stderr`.

## Installation

You can install `mylogger` using `go get`:

```bash
go get github.com/juparave/mylogger
```

## Usage

```go
package main

import (
	"github.com/juparave/mylogger"
)

func main() {
	// Create a new logger
	logger := mylogger.NewLogger()

	// Log some messages
	logger.Info("This is an informational message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
	logger.Debug("This is a debug message")
}
```

You can also create a logger with custom output buffers for testing purposes
using NewLoggerBuffers:

```go
package main

import (
	"bytes"
	"github.com/juparave/mylogger"
)

func main() {
	// Create buffers for capturing log output
	var stdOut, errOut bytes.Buffer

	// Create a logger with custom output buffers
	logger := mylogger.NewLoggerBuffers(&stdOut, &errOut)

	// Log some messages
	logger.Info("This is an informational message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
	logger.Debug("This is a debug message")

	// Check the output captured in buffers
	println("Standard output:", stdOut.String())
	println("Error output:", errOut.String())
}
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.


