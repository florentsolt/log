package log

import (
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/rs/zerolog"
)

var log zerolog.Logger

var output = os.Stdout

// Return the output used by the logger
func Output() io.Writer {
	return output
}

func init() {
	if os.Getenv(EnvOutput) == "stderr" {
		output = os.Stderr
	}
	ConsoleWriter.Out = output

	log = zerolog.New(ConsoleWriter).With().Timestamp().Caller().Logger()

	if os.Getenv(EnvLevel) != "" {
		if os.Getenv(EnvLevel) == "disabled" {
			zerolog.SetGlobalLevel(zerolog.Disabled)
		} else {
			if err := SetLevel(os.Getenv(EnvLevel)); err != nil {
				log.Error().Err(err).Msg("Unable to set error level")
			}
		}
	}
}

// SetLevel set global level of the logger: debug, info, warn, error, fatal, panic.
func SetLevel(level string) error {
	l, err := zerolog.ParseLevel(level)
	if err != nil {
		return err
	}
	zerolog.SetGlobalLevel(l)
	return nil
}

// Return the logger object
func Logger() zerolog.Logger {
	return log
}

// from https://github.com/rs/zerolog/blob/master/log/log.go

// Debug starts a new message with debug level.
func Debug() *Event { return &Event{log.Debug()} }

// Info starts a new message with info level.
func Info() *Event { return &Event{log.Info()} }

// Warn starts a new message with warn level.
func Warn() *Event { return &Event{log.Warn()} }

// Error starts a new message with error level.
func Error() *Event { return &Event{log.Error()} }

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
func Fatal() *Event { return &Event{log.Fatal()} }

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
func Panic() *Event { return &Event{log.Panic()} }

// Print sends a log event using debug level and no extra field.
func Print(v ...interface{}) {
	log := log.With().CallerWithSkipFrameCount(3).Logger()
	log.Debug().Msg(fmt.Sprint(v...))
}

// Printf sends a log event using debug level and no extra field.
func Printf(format string, v ...interface{}) {
	log := log.With().CallerWithSkipFrameCount(3).Logger()
	log.Debug().Msg(fmt.Sprintf(format, v...))
}

// Stack dump the current stack in the log
func Stack() {
	stack := make([]byte, 10*1024)
	runtime.Stack(stack, false)
	_, _ = log.Write(stack)
}

// Write implements the io.Writer interface
func Write(p []byte) (n int, err error) {
	return log.Write(p)
}
