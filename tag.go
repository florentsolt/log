package log

func (e *Event) Tag(names ...string) *Event {
	return &Event{e.Strs("tags", names)}
}

func Tag(names ...string) Wrapper {
	return Wrapper{instance.With().Strs("tags", names).Logger()}
}
