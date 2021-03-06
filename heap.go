package log

import (
	"runtime"
)

func (e *Event) Heap() *Event {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	return e.Uints64("heap", []uint64{
		mem.HeapAlloc, mem.TotalAlloc,
	})
}
