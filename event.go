package log

import (
	"github.com/rs/zerolog"
)

type Event struct {
	*zerolog.Event
}
