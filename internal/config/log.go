package config

// LogLevel represents the severity of log records.
type LogLevel string

const (
	// LogLevelDebug captures detailed information, typically useful for development
	// and debugging.
	LogLevelDebug LogLevel = "debug"

	// LogLevelInfo captures general information about the application's operation.
	LogLevelInfo LogLevel = "info"

	// LogLevelWarn captures non-critical events or potentially harmful situations.
	LogLevelWarn LogLevel = "warn"

	// LogLevelError captures critical events or errors that require immediate
	// attention.
	LogLevelError LogLevel = "error"
)

// LogFormat represents the encoding style of log records.
type LogFormat string

const (
	// LogFormatText renders log records as human-readable text.
	LogFormatText LogFormat = "text"

	// LogFormatJSON renders log records as structured JSON objects.
	LogFormatJSON LogFormat = "json"
)

// LogOutput represents the destination stream of log records.
type LogOutput string

const (
	// LogOutputStdout writes log records to the standard output stream (stdout).
	LogOutputStdout LogOutput = "stdout"

	// LogOutputStderr writes log records to the standard error stream (stderr).
	LogOutputStderr LogOutput = "stderr"
)

// Log defines the immutable application configuration for logging.
type Log interface {
	// Level returns the configured severity of log records.
	//
	// Environment variable: "LOG_LEVEL"
	//
	// Accepted values:
	//
	//  - [LogLevelDebug]
	//  - [LogLevelInfo]
	//  - [LogLevelWarn]
	//  - [LogLevelError]
	//
	// Default value: [LogLevelInfo]
	Level() LogLevel

	// Format returns the configured encoding style of log records.
	//
	// Environment variable: "LOG_FORMAT"
	//
	// Accepted values:
	//
	//  - [LogFormatText]
	//  - [LogFormatJSON]
	//
	// Default value: [LogFormatText]
	Format() LogFormat

	// Output returns the configured destination stream of log records.
	//
	// Environment variable: "LOG_OUTPUT"
	//
	// Accepted values:
	//
	//  - [LogOutputStdout]
	//  - [LogOutputStderr]
	//
	// Default value: [LogOutputStdout]
	Output() LogOutput
}

const (
	envLogLevel     = "LOG_LEVEL"
	defaultLogLevel = LogLevelInfo

	envLogFormat     = "LOG_FORMAT"
	defaultLogFormat = LogFormatText

	envLogOutput     = "LOG_OUTPUT"
	defaultLogOutput = LogOutputStdout
)

type log struct {
	level  LogLevel
	format LogFormat
	output LogOutput
}

func (l *log) Level() LogLevel {
	return l.level
}

func (l *log) Format() LogFormat {
	return l.format
}

func (l *log) Output() LogOutput {
	return l.output
}
