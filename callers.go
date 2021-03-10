package log

import (
	"fmt"
	"runtime"
)

// Callers dump the current stack of callers
func (e *Event) Callers() *Event {
	callers := make([]uintptr, 50)
	lines := []string{}
	runtime.Callers(3, callers)
	frames := runtime.CallersFrames(callers)
	for {
		frame, more := frames.Next()
		if !more {
			break
		}
		lines = append(lines, fmt.Sprintf("%s:%d\t%s()", frame.File, frame.Line, frame.Function))
	}
	return e.Strs("callers", lines)
}
