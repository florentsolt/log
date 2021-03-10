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

func (w Wrapper) Trace() *Event { return &Event{w.Logger.Trace()} }
func Trace() *Event             { return instance.Trace() }

func (w Wrapper) Debug() *Event { return &Event{w.Logger.Debug()} }
func Debug() *Event             { return instance.Debug() }

func (w Wrapper) Info() *Event { return &Event{w.Logger.Info()} }
func Info() *Event             { return instance.Info() }

func (w Wrapper) Warn() *Event { return &Event{w.Logger.Warn()} }
func Warn() *Event             { return instance.Warn() }

func (w Wrapper) Error() *Event { return &Event{w.Logger.Error()} }
func Error() *Event             { return instance.Error() }

func (w Wrapper) Fatal() *Event { return &Event{w.Logger.Fatal()} }
func Fatal() *Event             { return instance.Fatal() }

func (w Wrapper) Panic() *Event { return &Event{w.Logger.Panic()} }
func Panic() *Event             { return instance.Panic() }

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
