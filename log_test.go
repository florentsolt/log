package log

import (
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	Info().Msg("Hello World!")
	Warn().Msg("Did you notice the last stack enty?")
	Debug().Heap().Msg("Ok but what about memory?")
	Stack()
	Info().Tag("APP").Msg("Support for tag")
	Info().Tag("DB", "SYS").Msg("Support for tag*s*")
	Debug().Dur("instant", 42*time.Millisecond).Dur("short", 2*time.Second).Dur("long", 6*time.Hour).Msg("Duration in seconds...")
	Trace().Msg("Lowest level trace")
}
