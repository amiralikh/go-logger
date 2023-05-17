package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

// LogLevel represents the logging levels
type LogLevel int

const (
	// Info level
	Info LogLevel = iota
	// Error level
	Error
	// Warning level
	Warning
	// Debug level
	Debug
)

// Logger represents the logger object
type Logger struct {
	level     LogLevel
	logToFile bool
	logFile   *os.File
}

// NewLogger creates a new logger with the specified level and log output options
func NewLogger(level LogLevel, logToFile bool) (*Logger, error) {
	logger := &Logger{
		level:     level,
		logToFile: logToFile,
	}

	if logToFile {
		file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, fmt.Errorf("failed to open log file: %v", err)
		}
		logger.logFile = file
	}

	return logger, nil
}

// Log logs the message with the specified level
func (l *Logger) Log(level LogLevel, message string) {
	if level < l.level {
		return
	}

	logMessage := fmt.Sprintf("[%s] %s - %s\n", levelToString(level), time.Now().Format(time.RFC3339), message)

	if l.logToFile {
		if _, err := l.logFile.WriteString(logMessage); err != nil {
			log.Printf("Failed to write to log file: %v", err)
		}
	}

	log.Print(logMessage)
}

// Close closes the logger and the log file
func (l *Logger) Close() {
	if l.logToFile {
		l.logFile.Close()
	}
}

// Helper function to convert log level to string
func levelToString(level LogLevel) string {
	switch level {
	case Info:
		return "INFO"
	case Error:
		return "ERROR"
	case Warning:
		return "WARNING"
	case Debug:
		return "DEBUG"
	default:
		return "UNKNOWN"
	}
}
