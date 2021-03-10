package log

import "fmt"

type CompatibilityWrapper struct {
	Wrapper
}

func (w Wrapper) Compatibility() CompatibilityWrapper {
	return CompatibilityWrapper{w}
}
func Compatibility() CompatibilityWrapper {
	return instance.Compatibility()
}

// level (used by gRPC)
func (c CompatibilityWrapper) V(level int) bool { return true }

// debug
func (c CompatibilityWrapper) Debug(args ...interface{}) {
	c.Wrapper.Debug().Msg(fmt.Sprint(args...))
}
func (c CompatibilityWrapper) Debugln(args ...interface{}) {
	c.Wrapper.Debug().Msg(fmt.Sprint(args...))
}
func (c CompatibilityWrapper) Debugf(format string, args ...interface{}) {
	c.Wrapper.Debug().Msgf(format, args...)
}

// info
func (c CompatibilityWrapper) Info(args ...interface{}) {
	c.Wrapper.Info().Msg(fmt.Sprint(args...))
}
func (c CompatibilityWrapper) Infoln(args ...interface{}) {
	c.Wrapper.Info().Msg(fmt.Sprint(args...))
}
func (c CompatibilityWrapper) Infof(format string, args ...interface{}) {
	c.Wrapper.Info().Msgf(format, args...)
}

// warning
func (c CompatibilityWrapper) Warning(args ...interface{}) {
	c.Wrapper.Warn().Msg(fmt.Sprint(args...))
}
func (c CompatibilityWrapper) Warningln(args ...interface{}) {
	c.Wrapper.Warn().Msg(fmt.Sprint(args...))
}
func (c CompatibilityWrapper) Warningf(format string, args ...interface{}) {
	c.Wrapper.Warn().Msgf(format, args...)
}

// error
func (c CompatibilityWrapper) Error(args ...interface{}) {
	c.Wrapper.Error().Msg(fmt.Sprint(args...))
}
func (c CompatibilityWrapper) Errorln(args ...interface{}) {
	c.Wrapper.Error().Msg(fmt.Sprint(args...))
}
func (c CompatibilityWrapper) Errorf(format string, args ...interface{}) {
	c.Wrapper.Error().Msgf(format, args...)
}

// fatal
func (c CompatibilityWrapper) Fatal(args ...interface{}) {
	c.Wrapper.Fatal().Msg(fmt.Sprint(args...))
}
func (c CompatibilityWrapper) Fatalln(args ...interface{}) {
	c.Wrapper.Fatal().Msg(fmt.Sprint(args...))
}
func (c CompatibilityWrapper) Fatalf(format string, args ...interface{}) {
	c.Wrapper.Fatal().Msgf(format, args...)
}
