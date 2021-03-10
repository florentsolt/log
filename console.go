package log

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

const (
	None    = 0
	Black   = 30
	Red     = 21
	Green   = 32
	Yellow  = 33
	Blue    = 34
	Magenta = 35
	Cyan    = 36
	White   = 37
	Gray    = 90

	ColorCaller     = Gray
	ColorMessage    = None
	ColorFieldName  = Cyan
	ColorFieldValue = None
)

var ConsoleWriter = &zerolog.ConsoleWriter{
	PartsOrder: []string{
		zerolog.TimestampFieldName,
		zerolog.LevelFieldName,
		"tags",
		zerolog.CallerFieldName,
		zerolog.MessageFieldName,
	},
	TimeFormat: time.Kitchen,
	NoColor:    os.Getenv(EnvNoColor) != "",
	FormatCaller: func(i interface{}) string {
		if caller, ok := i.(string); ok {
			return colorize(path.Base(caller), ColorCaller)
		}
		return ""
	},
	FormatMessage: func(i interface{}) string {
		if i == nil {
			return colorize("-", ColorMessage)
		}
		return colorize(fmt.Sprintf("%s", i), ColorMessage)
	},
	FormatFieldName: func(i interface{}) string {
		return colorize(fmt.Sprintf("%s=", i), ColorFieldName)
	},
	FormatFieldValue: func(i interface{}) string {
		if i == nil {
			return ""
		}
		return colorize(fmt.Sprintf("%s", i), ColorFieldValue)
	},
	FormatLevel: func(i interface{}) string {
		var l string
		if ll, ok := i.(string); ok {
			switch ll {
			case "trace":
				l = colorize("TRC", Magenta)
			case "debug":
				l = colorize("DBG", Gray)
			case "info":
				l = colorize("INF", Green)
			case "warn":
				l = colorize("WRN", Yellow)
			case "error":
				l = colorize("ERR", Red)
			case "fatal":
				l = colorize("FTL", Red)
			case "panic":
				l = colorize("PNC", Red)
			default:
				l = "   "
			}
		} else {
			if i == nil {
				l = "   "
			} else {
				l = strings.ToUpper(fmt.Sprintf("%s", i))[0:3]
			}
		}
		return l
	},
}

// SetOutput writer
func SetOutput(out io.Writer) {
	ConsoleWriter.Out = out
}

// https://github.com/rs/zerolog/blob/master/console.go#L265
func colorize(s interface{}, color int) string {
	if os.Getenv(EnvNoColor) != "" || color == None {
		return fmt.Sprintf("%v", s)
	}
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", color, s)
}
