package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/oxtoacart/bpool"
	"github.com/rs/zerolog"
)

const (
	None    = 0
	Black   = 30
	Red     = 31
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

var pool = bpool.NewBufferPool(48)

type Console struct {
	Parts   []string
	Out     io.Writer
	NoColor bool
}

func (c *Console) Write(p []byte) (n int, err error) {
	buf := pool.Get()

	decoder := json.NewDecoder(bytes.NewReader(p))
	decoder.UseNumber()
	var data map[string]interface{}
	if err := decoder.Decode(&data); err != nil {
		return n, fmt.Errorf("Cannot decode event: %s", err)
	}

	for _, part := range c.Parts {
		if data[part] != nil {
			c.WritePart(buf, data, part)
			delete(data, part)
			buf.WriteByte(' ')
		}
	}

	callers, ok := data["callers"].([]interface{})
	if callers != nil && ok {
		delete(data, "callers")
	}

	fields := make([]string, 0, len(data))
	for field := range data {
		fields = append(fields, field)
	}
	sort.Strings(fields)

	for _, field := range fields {
		buf.WriteString(c.Colorize(fmt.Sprintf("%s=", field), Cyan)) // need quote?

		if field == "heap" {
			cast, ok := data["heap"].([]interface{})
			if ok && len(cast) == 2 {
				alloc, err1 := cast[0].(json.Number).Float64()
				total, err2 := cast[0].(json.Number).Float64()
				if err1 == nil && err2 == nil {
					buf.WriteString(c.Colorize(fmt.Sprintf(
						"%.2fMB/%.2fMB",
						alloc/1000000,
						total/1000000,
					), None))
					buf.WriteByte(' ')
					continue
				}
			}
		}
		buf.WriteString(c.Colorize(fmt.Sprintf("%s", data[field]), None))
		buf.WriteByte(' ')
	}
	buf.WriteByte('\n')

	if callers != nil {
		for _, caller := range callers {
			buf.WriteByte('\t')
			buf.WriteString(fmt.Sprintf("%s", caller))
			buf.WriteByte('\n')

		}
	}

	_, err = buf.WriteTo(c.Out)
	return len(p), err
}

func (c *Console) WritePart(buf *bytes.Buffer, data map[string]interface{}, part string) {
	switch part {

	case zerolog.LevelFieldName:
		level := "   "
		if cast, ok := data[part].(string); ok {
			switch cast {
			case "trace":
				level = c.Colorize("TRC", Magenta)
			case "debug":
				level = c.Colorize("DBG", Gray)
			case "info":
				level = c.Colorize("INF", Green)
			case "warn":
				level = c.Colorize("WRN", Yellow)
			case "error":
				level = c.Colorize("ERR", Red)
			case "fatal":
				level = c.Colorize("FTL", Red)
			case "panic":
				level = c.Colorize("PNC", Red)
			}
		} else {
			level = strings.ToUpper(fmt.Sprintf("%s", data[part]))[0:3]
		}
		buf.WriteString(level)

	case zerolog.TimestampFieldName:
		ts := ""
		switch t := data[part].(type) {
		case string:
			parsed, err := time.Parse(time.RFC3339, t)
			if err != nil {
				ts = t
			} else {
				ts = parsed.Format(time.Kitchen)
			}

		case json.Number:
			i, err := t.Int64()
			if err != nil {
				ts = t.String()
			} else {
				ts = time.Unix(i, 0).UTC().Format(time.Kitchen)
			}
		}
		if ts != "" {
			buf.WriteString(c.Colorize(ts, Gray))
		}

	case zerolog.MessageFieldName:
		buf.WriteString(c.Colorize(fmt.Sprintf("%s", data[part]), None))

	case zerolog.CallerFieldName:
		if caller, ok := data[part].(string); ok {
			buf.WriteString(c.Colorize(path.Base(caller), Gray))
		}

	case zerolog.ErrorFieldName:
		buf.WriteString(c.Colorize(fmt.Sprintf("%s", data[part]), Red))

	default:
		buf.WriteString(c.Colorize(fmt.Sprintf("%s", data[part]), None))
	}
}

func (c *Console) Colorize(s interface{}, color int) string {
	if c.NoColor || color == None {
		return fmt.Sprintf("%v", s)
	}
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", color, s)
}

var Writer = &Console{
	Parts: []string{
		zerolog.TimestampFieldName,
		zerolog.LevelFieldName,
		"tags",
		zerolog.CallerFieldName,
		zerolog.MessageFieldName,
		zerolog.ErrorFieldName,
	},
	NoColor: os.Getenv(EnvNoColor) != "",
	Out:     Output(),
}

// SetOutput writer
func SetOutput(out io.Writer) {
	Writer.Out = out
}
