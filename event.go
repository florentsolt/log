package log

import (
	"fmt"
	"runtime"

	"github.com/rs/zerolog"
)

type Event struct {
	*zerolog.Event
}

func (e *Event) Heap() *Event {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	return &Event{
		e.Str("heap", fmt.Sprintf("%.2fMB/%.2fMB",
			float64(mem.HeapAlloc)/float64(1000000),
			float64(mem.TotalAlloc)/float64(1000000),
		)),
	}
}

func (e *Event) Tag(names ...string) *Event {
	return &Event{e.Strs("tags", names)}
}
