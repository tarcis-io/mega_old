package config

type (
	// LogLevel represents the severity of log records.
	LogLevel string
)

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

const (
	// EnvLogLevel specifies the environment variable name for configuring the
	// [LogLevel].
	//
	// Expected values:
	//
	//  - [LogLevelDebug]
	//  - [LogLevelInfo]
	//  - [LogLevelWarn]
	//  - [LogLevelError]
	//
	// Default: [DefaultLogLevel]
	EnvLogLevel = "LOG_LEVEL"

	// DefaultLogLevel specifies the default [LogLevel], used as the fallback when
	// [EnvLogLevel] is unset or contains an invalid value.
	DefaultLogLevel = LogLevelInfo
)

type (
	// LogFormat represents the encoding style of log records.
	LogFormat string
)

const (
	// LogFormatText renders log records as human-readable text.
	LogFormatText LogFormat = "text"

	// LogFormatJSON renders log records as structured JSON objects.
	LogFormatJSON LogFormat = "json"
)

const (
	// EnvLogFormat specifies the environment variable name for configuring the
	// [LogFormat].
	//
	// Expected values:
	//
	//  - [LogFormatText]
	//  - [LogFormatJSON]
	//
	// Default: [DefaultLogFormat]
	EnvLogFormat = "LOG_FORMAT"

	// DefaultLogFormat specifies the default [LogFormat], used as the fallback when
	// [EnvLogFormat] is unset or contains an invalid value.
	DefaultLogFormat = LogFormatText
)

type (
	// LogOutput represents the destination stream of log records.
	LogOutput string
)

const (
	// LogOutputStdout writes log records to the standard output stream (stdout).
	LogOutputStdout LogOutput = "stdout"

	// LogOutputStderr writes log records to the standard error stream (stderr).
	LogOutputStderr LogOutput = "stderr"
)

const (
	// EnvLogOutput specifies the environment variable name for configuring the
	// [LogOutput].
	//
	// Expected values:
	//
	//  - [LogOutputStdout]
	//  - [LogOutputStderr]
	//  - A custom string (typically a file path)
	//
	// Default: [DefaultLogOutput]
	EnvLogOutput = "LOG_OUTPUT"

	// DefaultLogOutput specifies the default [LogOutput], used as the fallback when
	// [EnvLogOutput] is unset or contains an invalid value.
	DefaultLogOutput = LogOutputStdout
)

type (
	// Log defines the immutable application configuration for logging.
	Log interface {
		// Level returns the configured [LogLevel].
		Level() LogLevel

		// Format returns the configured [LogFormat].
		Format() LogFormat

		// Output returns the configured [LogOutput].
		Output() LogOutput
	}
)

type (
	log struct {
		level  LogLevel
		format LogFormat
		output LogOutput
	}
)

func newLog(level LogLevel, format LogFormat, output LogOutput) *log {
	return &log{
		level:  level,
		format: format,
		output: output,
	}
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
