package log

import (
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

	if os.Getenv(EnvTCP) != "" {
		tcp := &TcpInterceptor{}
		go func() {
			err := tcp.ListenAndAccept(os.Getenv(EnvTCP))
			if err != nil {
				log.Error().Err(err).Msg("Unable to tcp listen")
			}
		}()
		log = zerolog.New(tcp).With().Timestamp().Caller().Logger()
	} else {
		log = zerolog.New(ConsoleWriter).With().Timestamp().Caller().Logger()
	}

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
func Debug() *zerolog.Event { return log.Debug() }

// Info starts a new message with info level.
func Info() *zerolog.Event { return log.Info() }

// Warn starts a new message with warn level.
func Warn() *zerolog.Event { return log.Warn() }

// Error starts a new message with error level.
func Error() *zerolog.Event { return log.Error() }

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
func Fatal() *zerolog.Event { return log.Fatal() }

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
func Panic() *zerolog.Event { return log.Panic() }

// Print sends a log event using debug level and no extra field.
func Print(v ...interface{}) { log.Print(v...) }

// Printf sends a log event using debug level and no extra field.
func Printf(format string, v ...interface{}) { log.Printf(format, v...) }

// Stack dump the current stack in the log
func Stack() {
	stack := make([]byte, 10*1024)
	runtime.Stack(stack, false)
	_, _ = log.Write(stack)
}
