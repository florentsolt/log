package log

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog"
)

type CompatibilityWrapper struct {
	wrapper zerolog.Logger
}

func (w Wrapper) Compatibility() CompatibilityWrapper {
	return CompatibilityWrapper{w.With().CallerWithSkipFrameCount(3).Logger()}
}
func Compatibility() CompatibilityWrapper {
	return instance.Compatibility()
}

// level (used by gRPC)
func (c CompatibilityWrapper) V(level int) bool { return true }

// debug
func (c CompatibilityWrapper) Debug(args ...interface{}) {
	c.wrapper.Debug().Msg(strings.TrimSpace(fmt.Sprint(args...)))
}
func (c CompatibilityWrapper) Debugln(args ...interface{}) {
	c.wrapper.Debug().Msg(strings.TrimSpace(fmt.Sprint(args...)))
}
func (c CompatibilityWrapper) Debugf(format string, args ...interface{}) {
	c.wrapper.Debug().Msgf(strings.TrimSpace(fmt.Sprintf(format, args...)))
}

// info
func (c CompatibilityWrapper) Info(args ...interface{}) {
	c.wrapper.Info().Msg(strings.TrimSpace(fmt.Sprint(args...)))
}
func (c CompatibilityWrapper) Infoln(args ...interface{}) {
	c.wrapper.Info().Msg(strings.TrimSpace(fmt.Sprint(args...)))
}
func (c CompatibilityWrapper) Infof(format string, args ...interface{}) {
	c.wrapper.Info().Msgf(strings.TrimSpace(fmt.Sprintf(format, args...)))
}

// warning
func (c CompatibilityWrapper) Warning(args ...interface{}) {
	c.wrapper.Warn().Msg(strings.TrimSpace(fmt.Sprint(args...)))
}
func (c CompatibilityWrapper) Warningln(args ...interface{}) {
	c.wrapper.Warn().Msg(strings.TrimSpace(fmt.Sprint(args...)))
}
func (c CompatibilityWrapper) Warningf(format string, args ...interface{}) {
	c.wrapper.Warn().Msgf(strings.TrimSpace(fmt.Sprintf(format, args...)))
}

// error
func (c CompatibilityWrapper) Error(args ...interface{}) {
	c.wrapper.Error().Msg(strings.TrimSpace(fmt.Sprint(args...)))
}
func (c CompatibilityWrapper) Errorln(args ...interface{}) {
	c.wrapper.Error().Msg(strings.TrimSpace(fmt.Sprint(args...)))
}
func (c CompatibilityWrapper) Errorf(format string, args ...interface{}) {
	c.wrapper.Error().Msgf(strings.TrimSpace(fmt.Sprintf(format, args...)))
}

// fatal
func (c CompatibilityWrapper) Fatal(args ...interface{}) {
	c.wrapper.Fatal().Msg(strings.TrimSpace(fmt.Sprint(args...)))
}
func (c CompatibilityWrapper) Fatalln(args ...interface{}) {
	c.wrapper.Fatal().Msg(strings.TrimSpace(fmt.Sprint(args...)))
}
func (c CompatibilityWrapper) Fatalf(format string, args ...interface{}) {
	c.wrapper.Fatal().Msgf(strings.TrimSpace(fmt.Sprintf(format, args...)))
}
