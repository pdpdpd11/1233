// Package logger provides a customizable logging utility using zerolog.
package logger

import (
	"os"
	"path"
	"runtime"

	"github.com/rs/zerolog"
)

// Global logger
var log zerolog.Logger

// verbose controls the verbosity of the logging output.
var verbose bool

// SetVerbose allows external packages to modify the verbosity of the logger.
func SetVerbose(v bool) {
	verbose = v
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	if !verbose {
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}
}

// GetVerbose allows external packages to get the verbosity status.
func GetVerbose() bool {
	return verbose
}

// InitLogger initializes the logger based on the current verbosity setting.
// This function should be called after the verbosity has been set via flags.
func InitLogger() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}
	log = zerolog.New(output).With().Timestamp().Logger()
}

// callerHook is a zerolog Hook designed to add source file and line number info to log messages.
type callerHook struct{}

func (h callerHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if _, file, line, ok := runtime.Caller(3); ok {
		e.Str("source", path.Base(file)+":"+string(rune(line)))
	}
}

// WithCaller adds caller info (source file and line number) to log messages.
func WithCaller() zerolog.Logger {
	return log.Hook(callerHook{})
}

// Info logs an informational message.
func Info(msg string) {
	log.Info().Msg(msg)
}

// Infof logs a formatted informational message.
func Infof(format string, args ...interface{}) {
	log.Info().Msgf(format, args...)
}

// Debug logs a debug message.
func Debug(msg string) {
	log.Debug().Msg(msg)
}

// Debugf logs a formatted debug message.
func Debugf(format string, args ...interface{}) {
	log.Debug().Msgf(format, args...)
}

// Error logs an error message.
func Error(msg string) {
	log.Error().Msg(msg)
}

// Errorf logs a formatted error message.
func Errorf(format string, args ...interface{}) {
	log.Error().Msgf(format, args...)
}

// Warn logs a warning message.
func Warn(msg string) {
	log.Warn().Msg(msg)
}

// Warnf logs a formatted warning message.
func Warnf(format string, args ...interface{}) {
	log.Warn().Msgf(format, args...)
}

// Fatal logs a fatal message and exits the program.
func Fatal(msg string) {
	log.Fatal().Msg(msg)
}

// Fatalf logs a formatted fatal message and exits the program.
func Fatalf(format string, args ...interface{}) {
	log.Fatal().Msgf(format, args...)
}

// Panic logs a panic message and panics.
func Panic(msg string) {
	log.Panic().Msg(msg)
}

// Panicf logs a formatted panic message and panics.
func Panicf(format string, args ...interface{}) {
	log.Panic().Msgf(format, args...)
}
