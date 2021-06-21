package log

import (
	"fmt"
	"net"
	"time"

	"github.com/rs/zerolog"
)

type Event struct {
	parent *zerolog.Event
}

func (e *Event) Send() {
	e.parent.Send()
}

func (e *Event) Msg(msg string) {
	e.parent.Msg(msg)
}

func (e *Event) Msgf(format string, v ...interface{}) {
	e.parent.Msgf(format, v...)
}

func (e *Event) Dict(key string, dict *Event) *Event {
	return &Event{e.parent.Dict(key, dict.parent)}
}

func (e *Event) Array(key string, arr zerolog.LogArrayMarshaler) *Event {
	return &Event{e.parent.Array(key, arr)}
}

func (e *Event) Object(key string, obj zerolog.LogObjectMarshaler) *Event {
	return &Event{e.parent.Object(key, obj)}
}

func (e *Event) EmbedObject(obj zerolog.LogObjectMarshaler) *Event {
	return &Event{e.parent.EmbedObject(obj)}
}

func (e *Event) Str(key, val string) *Event {
	return &Event{e.parent.Str(key, val)}
}

func (e *Event) Strs(key string, vals []string) *Event {
	return &Event{e.parent.Strs(key, vals)}
}

func (e *Event) Stringer(key string, val fmt.Stringer) *Event {
	return &Event{e.parent.Stringer(key, val)}
}

func (e *Event) Bytes(key string, val []byte) *Event {
	return &Event{e.parent.Bytes(key, val)}
}

func (e *Event) Hex(key string, val []byte) *Event {
	return &Event{e.parent.Hex(key, val)}
}

func (e *Event) RawJSON(key string, b []byte) *Event {
	return &Event{e.parent.RawJSON(key, b)}
}

func (e *Event) AnErr(key string, err error) *Event {
	return &Event{e.parent.AnErr(key, err)}
}

func (e *Event) Errs(key string, errs []error) *Event {
	return &Event{e.parent.Errs(key, errs)}
}

func (e *Event) Err(err error) *Event {
	return &Event{e.parent.Err(err)}
}

func (e *Event) Bool(key string, b bool) *Event {
	return &Event{e.parent.Bool(key, b)}
}

func (e *Event) Bools(key string, b []bool) *Event {
	return &Event{e.parent.Bools(key, b)}
}

func (e *Event) Int(key string, i int) *Event {
	return &Event{e.parent.Int(key, i)}
}

func (e *Event) Ints(key string, i []int) *Event {
	return &Event{e.parent.Ints(key, i)}
}

func (e *Event) Int8(key string, i int8) *Event {
	return &Event{e.parent.Int8(key, i)}
}

func (e *Event) Ints8(key string, i []int8) *Event {
	return &Event{e.parent.Ints8(key, i)}
}

func (e *Event) Int16(key string, i int16) *Event {
	return &Event{e.parent.Int16(key, i)}
}

func (e *Event) Ints16(key string, i []int16) *Event {
	return &Event{e.parent.Ints16(key, i)}
}

func (e *Event) Int32(key string, i int32) *Event {
	return &Event{e.parent.Int32(key, i)}
}

func (e *Event) Ints32(key string, i []int32) *Event {
	return &Event{e.parent.Ints32(key, i)}
}

func (e *Event) Int64(key string, i int64) *Event {
	return &Event{e.parent.Int64(key, i)}
}

func (e *Event) Ints64(key string, i []int64) *Event {
	return &Event{e.parent.Ints64(key, i)}
}

func (e *Event) Uint(key string, i uint) *Event {
	return &Event{e.parent.Uint(key, i)}
}

func (e *Event) Uints(key string, i []uint) *Event {
	return &Event{e.parent.Uints(key, i)}
}

func (e *Event) Uint8(key string, i uint8) *Event {
	return &Event{e.parent.Uint8(key, i)}
}

func (e *Event) Uints8(key string, i []uint8) *Event {
	return &Event{e.parent.Uints8(key, i)}
}

func (e *Event) Uint16(key string, i uint16) *Event {
	return &Event{e.parent.Uint16(key, i)}
}

func (e *Event) Uints16(key string, i []uint16) *Event {
	return &Event{e.parent.Uints16(key, i)}
}

func (e *Event) Uint32(key string, i uint32) *Event {
	return &Event{e.parent.Uint32(key, i)}
}

func (e *Event) Uints32(key string, i []uint32) *Event {
	return &Event{e.parent.Uints32(key, i)}
}

func (e *Event) Uint64(key string, i uint64) *Event {
	return &Event{e.parent.Uint64(key, i)}
}

func (e *Event) Uints64(key string, i []uint64) *Event {
	return &Event{e.parent.Uints64(key, i)}
}

func (e *Event) Float32(key string, f float32) *Event {
	return &Event{e.parent.Float32(key, f)}
}

func (e *Event) Floats32(key string, f []float32) *Event {
	return &Event{e.parent.Floats32(key, f)}
}

func (e *Event) Float64(key string, f float64) *Event {
	return &Event{e.parent.Float64(key, f)}
}

func (e *Event) Floats64(key string, f []float64) *Event {
	return &Event{e.parent.Floats64(key, f)}
}

func (e *Event) Timestamp() *Event {
	return &Event{e.parent.Timestamp()}
}

func (e *Event) Time(key string, t time.Time) *Event {
	return &Event{e.parent.Time(key, t)}
}

func (e *Event) Times(key string, t []time.Time) *Event {
	return &Event{e.parent.Times(key, t)}
}

func (e *Event) Dur(key string, d time.Duration) *Event {
	return &Event{e.parent.Dur(key, d)}
}

func (e *Event) Durs(key string, d []time.Duration) *Event {
	return &Event{e.parent.Durs(key, d)}
}

func (e *Event) TimeDiff(key string, t time.Time, start time.Time) *Event {
	return &Event{e.parent.TimeDiff(key, t, start)}
}

func (e *Event) Interface(key string, i interface{}) *Event {
	return &Event{e.parent.Str(key, fmt.Sprintf("%#v", i))}
}

func (e *Event) Caller(skip ...int) *Event {
	return &Event{e.parent.Caller(skip...)}
}

func (e *Event) IPAddr(key string, ip net.IP) *Event {
	return &Event{e.parent.IPAddr(key, ip)}
}

func (e *Event) IPPrefix(key string, pfx net.IPNet) *Event {
	return &Event{e.parent.IPPrefix(key, pfx)}
}

func (e *Event) MACAddr(key string, ha net.HardwareAddr) *Event {
	return &Event{e.parent.MACAddr(key, ha)}
}
