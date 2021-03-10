package log

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var instance Wrapper

// Return the output used by the logger
func Output() io.Writer {
	if os.Getenv(EnvOutput) == "stderr" {
		return os.Stderr
	} else {
		return os.Stdout
	}
}

func init() {
	w := Output()
	if os.Getenv(EnvJson) == "" {
		w = Writer
	}
	instance = Wrapper{zerolog.New(w).With().Timestamp().Caller().Logger()}
	zerolog.DurationFieldUnit = time.Second
	zerolog.DurationFieldInteger = false

	if os.Getenv(EnvLevel) != "" {
		if os.Getenv(EnvLevel) == "disabled" {
			zerolog.SetGlobalLevel(zerolog.Disabled)
		} else {
			if err := SetLevel(os.Getenv(EnvLevel)); err != nil {
				instance.Error().Err(err).Msg("Unable to set error level")
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

type Wrapper struct {
	zerolog.Logger
}

// Return the logger object
func Logger() Wrapper {
	return instance
}

// from https://github.com/rs/zerolog/blob/master/log/log.go

// Trace starts a new message with debug level.
func Trace() *Event { return &Event{instance.Trace()} }

// Debug starts a new message with debug level.
func Debug() *Event { return &Event{instance.Debug()} }

// Info starts a new message with info level.
func Info() *Event { return &Event{instance.Info()} }

// Warn starts a new message with warn level.
func Warn() *Event { return &Event{instance.Warn()} }

// Error starts a new message with error level.
func Error() *Event { return &Event{instance.Error()} }

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
func Fatal() *Event { return &Event{instance.Fatal()} }

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
func Panic() *Event { return &Event{instance.Panic()} }

// Print sends a log event using debug level and no extra field.
func (w Wrapper) Print(v ...interface{}) {
	l := w.With().CallerWithSkipFrameCount(3).Logger()
	l.Debug().Msg(fmt.Sprint(v...))
}
func Print(v ...interface{}) {
	instance.Print(v...)
}

// Printf sends a log event using debug level and no extra field.
func (w Wrapper) Printf(format string, v ...interface{}) {
	l := w.With().CallerWithSkipFrameCount(3).Logger()
	l.Debug().Msg(fmt.Sprintf(format, v...))
}
func Printf(format string, v ...interface{}) {
	instance.Printf(format, v...)
}

// Write implements the io.Writer interface
func Write(p []byte) (n int, err error) {
	return instance.Write(p)
}
