package log

import "github.com/rs/zerolog"

type LogLevel int8

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
	NoLevel
	Disabled
	TraceLevel LogLevel = -1
)

func (w Wrapper) Level(lvl LogLevel) Wrapper {
	return Wrapper{parent: w.parent.Level(zerolog.Level(lvl))}
}
func Level(lvl LogLevel) Wrapper {
	return instance.Level(lvl)
}
