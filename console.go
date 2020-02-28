package log

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

const (
	ColorCaller     = "gray"
	ColorMessage    = "none"
	ColorFieldName  = "cyan"
	ColorFieldValue = "none"
)

var colors = map[string]int{
	"black":   30,
	"red":     31,
	"green":   32,
	"yellow":  33,
	"blue":    34,
	"magenta": 35,
	"cyan":    36,
	"white":   37,
	"gray":    90,
}

var ConsoleWriter = zerolog.ConsoleWriter{
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
		return colorize(fmt.Sprintf("%s", i), ColorFieldValue)
	},
	FormatLevel: func(i interface{}) string {
		var l string
		if ll, ok := i.(string); ok {
			switch ll {
			case "debug":
				l = colorize("DBG", "gray")
			case "info":
				l = colorize("INF", "green")
			case "warn":
				l = colorize("WRN", "yellow")
			case "error":
				l = colorize("ERR", "red")
			case "fatal":
				l = colorize("FTL", "red")
			case "panic":
				l = colorize("PNC", "red")
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

// https://github.com/rs/zerolog/blob/master/console.go#L265
func colorize(s interface{}, color string) string {
	if os.Getenv(EnvNoColor) != "" || color == "none" {
		return fmt.Sprintf("%v", s)
	}
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", colors[color], s)
}

// Write implements the io.Writer interface
// In case of Console(), it transforms the JSON input with formatters
func Write(p []byte) (n int, err error) {
	return log.Write(p)
}
