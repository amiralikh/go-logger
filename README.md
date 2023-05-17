Certainly! Here's an example of a README file for the logger module:

# Logger Module

The Logger module is a logging utility for Go that provides logging levels (info, error, warning, debug) and the ability to customize the output to the console or a file. Each log entry contains the level, message, and timestamp.

## Installation

To use the Logger module in your Go project, you can use the `go get` command:

```shell
go get github.com/amiralikh/logger
```

## Usage

Import the Logger module into your Go code:

```go
import (
    "log"
    "os"
    "github.com/amiralikh/logger"
)
```

### Creating a Logger

To create a new logger, use the `NewLogger` function:

```go
logger, err := logger.NewLogger(logger.Info, true)
if err != nil {
    log.Fatalf("Failed to create logger: %v", err)
}
defer logger.Close()
```

The `NewLogger` function takes two parameters: the logging level and a boolean value indicating whether to log to a file (`true`) or to the console (`false`). It returns a logger object and an error (if any).

### Logging Messages

You can log messages using different levels: info, error, warning, and debug. To log a message, use the `Log` method on the logger object:

```go
logger.Log(logger.Info, "This is an info message")
logger.Log(logger.Warning, "This is a warning message")
logger.Log(logger.Error, "This is an error message")
logger.Log(logger.Debug, "This is a debug message")
```

The `Log` method takes two parameters: the log level and the message string. The logger will only log messages that have a level greater than or equal to the logger's configured level.

### Customizing Output

By default, the logger outputs messages to the console. To log messages to a file, set the second parameter of `NewLogger` to `true` and provide the file path:

```go
logger, err := logger.NewLogger(logger.Info, true)
if err != nil {
    log.Fatalf("Failed to create logger: %v", err)
}
defer logger.Close()
```

This will create a log file named `log.txt` in the current directory. The log file is opened in append mode, so existing log entries are preserved.

### Closing the Logger

It's important to close the logger when you're done using it to ensure that any buffered log entries are written to the file and resources are properly released. Use the `Close` method on the logger object:

```go
logger.Close()
```

### Example

Here's an example usage of the Logger module:

```go
package main

import (
    "log"
    "os"
    "github.com/yourusername/logger"
)

func main() {
    // Create a new logger with level 'Info' and log to a file
    logger, err := logger.NewLogger(logger.Info, true)
    if err != nil {
        log.Fatalf("Failed to create logger: %v", err)
    }
    defer logger.Close()

    // Log messages with different levels
    logger.Log(logger.Info, "This is an info message")
    logger.Log(logger.Warning, "This is a warning message")
    logger.Log(logger.Error, "This is an error message")
    logger.Log(logger.Debug, "This is a debug message")

    // You can also log to the console by setting logToFile to false
    consoleLogger, err := logger.NewLogger(logger.Info, false)
    if err != nil {
        log.Fatalf("Failed to create console logger: %v", err)
    }
    defer consoleLogger.Close()

    consoleLogger

.Log(logger.Info, "This is an info message logged to the console")
    consoleLogger.Log(logger.Error, "This is an error message logged to the console")
}
```

## License

This project is licensed under the [MIT License](LICENSE).
