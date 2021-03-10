package log

import (
	"errors"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	Info().Msg("Hello World!")
	Error().Err(errors.New("Bad things")).Msg("Oops")
	Warn().Msg("Did you notice the last stack enty?")
	Debug().Heap().Msg("Ok but what about memory?")
	Stack()
	Info().Tag("APP").Msg("Support for tag")
	Info().Tag("DB", "SYS").Msg("Support for tag*s*")
	Debug().Dur("instant", 42*time.Millisecond).Dur("short", 2*time.Second).Dur("long", 6*time.Hour).Msg("Duration in seconds...")
	Trace().Msg("Lowest level trace")
	ctx := Logger().With().Strs("tags", []string{"DB", "SYS"}).Logger()
	ctx.Info().Msg("Tags from context")
}
