package log

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

// Defines commons keys for logging
const ModuleKey = "module"

// ContextKey is used to store the logger in the context
var ContextKey struct{}

type Logger interface {
	Info(msg string, keyVals ...interface{})
	Error(msg string, keyVals ...interface{})
	Debug(msg string, keyVals ...interface{})
	With(keyVals ...interface{}) Logger
}

type ZeroLogWrapper struct {
	zerolog.Logger
}

func NewNopLogger() Logger {
	logger := zerolog.Nop()
	return &ZeroLogWrapper{logger}
}

func NewLogger() Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.Kitchen}
	logger := zerolog.New(output).With().Timestamp().Logger()
	return &ZeroLogWrapper{logger}
}

func NewLoggerWithKV(key, value string) Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.Kitchen}
	logger := zerolog.New(output).With().Str(key, value).Timestamp().Logger()
	return &ZeroLogWrapper{logger}
}

// Info takes a message and a set of key/value pairs and logs with level INFO.
// The key of the tuple must be a string.
func (l ZeroLogWrapper) Info(msg string, keyVals ...interface{}) {
	l.Logger.Info().Fields(keyVals).Msg(msg)
}

// Error takes a message and a set of key/value pairs and logs with level DEBUG.
// The key of the tuple must be a string.
func (l ZeroLogWrapper) Error(msg string, keyVals ...interface{}) {
	l.Logger.Error().Fields(keyVals).Msg(msg)
}

// Debug takes a message and a set of key/value pairs and logs with level ERR.
// The key of the tuple must be a string.
func (l ZeroLogWrapper) Debug(msg string, keyVals ...interface{}) {
	l.Logger.Debug().Fields(keyVals).Msg(msg)
}

// With returns a new wrapped logger with additional context provided by a set
func (l ZeroLogWrapper) With(keyVals ...interface{}) Logger {
	return ZeroLogWrapper{l.Logger.With().Fields(keyVals).Logger()}
}
