package log

import (
	"fmt"
	"net"
	"time"

	"github.com/rs/zerolog"
)

type Event struct {
	*zerolog.Event
}

func (e *Event) Dict(key string, dict *Event) *Event {
	return &Event{e.Event.Dict(key, dict.Event)}
}

func (e *Event) Array(key string, arr zerolog.LogArrayMarshaler) *Event {
	return &Event{e.Event.Array(key, arr)}
}

func (e *Event) Object(key string, obj zerolog.LogObjectMarshaler) *Event {
	return &Event{e.Event.Object(key, obj)}
}

func (e *Event) EmbedObject(obj zerolog.LogObjectMarshaler) *Event {
	return &Event{e.Event.EmbedObject(obj)}
}

func (e *Event) Str(key, val string) *Event {
	return &Event{e.Event.Str(key, val)}
}

func (e *Event) Strs(key string, vals []string) *Event {
	return &Event{e.Event.Strs(key, vals)}
}

func (e *Event) Stringer(key string, val fmt.Stringer) *Event {
	return &Event{e.Event.Stringer(key, val)}
}

func (e *Event) Bytes(key string, val []byte) *Event {
	return &Event{e.Event.Bytes(key, val)}
}

func (e *Event) Hex(key string, val []byte) *Event {
	return &Event{e.Event.Hex(key, val)}
}

func (e *Event) RawJSON(key string, b []byte) *Event {
	return &Event{e.Event.RawJSON(key, b)}
}

func (e *Event) AnErr(key string, err error) *Event {
	return &Event{e.Event.AnErr(key, err)}
}

func (e *Event) Errs(key string, errs []error) *Event {
	return &Event{e.Event.Errs(key, errs)}
}

func (e *Event) Err(err error) *Event {
	return &Event{e.Event.Err(err)}
}

func (e *Event) Bool(key string, b bool) *Event {
	return &Event{e.Event.Bool(key, b)}
}

func (e *Event) Bools(key string, b []bool) *Event {
	return &Event{e.Event.Bools(key, b)}
}

func (e *Event) Int(key string, i int) *Event {
	return &Event{e.Event.Int(key, i)}
}

func (e *Event) Ints(key string, i []int) *Event {
	return &Event{e.Event.Ints(key, i)}
}

func (e *Event) Int8(key string, i int8) *Event {
	return &Event{e.Event.Int8(key, i)}
}

func (e *Event) Ints8(key string, i []int8) *Event {
	return &Event{e.Event.Ints8(key, i)}
}

func (e *Event) Int16(key string, i int16) *Event {
	return &Event{e.Event.Int16(key, i)}
}

func (e *Event) Ints16(key string, i []int16) *Event {
	return &Event{e.Event.Ints16(key, i)}
}

func (e *Event) Int32(key string, i int32) *Event {
	return &Event{e.Event.Int32(key, i)}
}

func (e *Event) Ints32(key string, i []int32) *Event {
	return &Event{e.Event.Ints32(key, i)}
}

func (e *Event) Int64(key string, i int64) *Event {
	return &Event{e.Event.Int64(key, i)}
}

func (e *Event) Ints64(key string, i []int64) *Event {
	return &Event{e.Event.Ints64(key, i)}
}

func (e *Event) Uint(key string, i uint) *Event {
	return &Event{e.Event.Uint(key, i)}
}

func (e *Event) Uints(key string, i []uint) *Event {
	return &Event{e.Event.Uints(key, i)}
}

func (e *Event) Uint8(key string, i uint8) *Event {
	return &Event{e.Event.Uint8(key, i)}
}

func (e *Event) Uints8(key string, i []uint8) *Event {
	return &Event{e.Event.Uints8(key, i)}
}

func (e *Event) Uint16(key string, i uint16) *Event {
	return &Event{e.Event.Uint16(key, i)}
}

func (e *Event) Uints16(key string, i []uint16) *Event {
	return &Event{e.Event.Uints16(key, i)}
}

func (e *Event) Uint32(key string, i uint32) *Event {
	return &Event{e.Event.Uint32(key, i)}
}

func (e *Event) Uints32(key string, i []uint32) *Event {
	return &Event{e.Event.Uints32(key, i)}
}

func (e *Event) Uint64(key string, i uint64) *Event {
	return &Event{e.Event.Uint64(key, i)}
}

func (e *Event) Uints64(key string, i []uint64) *Event {
	return &Event{e.Event.Uints64(key, i)}
}

func (e *Event) Float32(key string, f float32) *Event {
	return &Event{e.Event.Float32(key, f)}
}

func (e *Event) Floats32(key string, f []float32) *Event {
	return &Event{e.Event.Floats32(key, f)}
}

func (e *Event) Float64(key string, f float64) *Event {
	return &Event{e.Event.Float64(key, f)}
}

func (e *Event) Floats64(key string, f []float64) *Event {
	return &Event{e.Event.Floats64(key, f)}
}

func (e *Event) Timestamp() *Event {
	return &Event{e.Event.Timestamp()}
}

func (e *Event) Time(key string, t time.Time) *Event {
	return &Event{e.Event.Time(key, t)}
}

func (e *Event) Times(key string, t []time.Time) *Event {
	return &Event{e.Event.Times(key, t)}
}

func (e *Event) Dur(key string, d time.Duration) *Event {
	return &Event{e.Event.Dur(key, d)}
}

func (e *Event) Durs(key string, d []time.Duration) *Event {
	return &Event{e.Event.Durs(key, d)}
}

func (e *Event) TimeDiff(key string, t time.Time, start time.Time) *Event {
	return &Event{e.Event.TimeDiff(key, t, start)}
}

func (e *Event) Interface(key string, i interface{}) *Event {
	return &Event{e.Event.Interface(key, i)}
}

func (e *Event) Caller(skip ...int) *Event {
	return &Event{e.Event.Caller(skip...)}
}

func (e *Event) IPAddr(key string, ip net.IP) *Event {
	return &Event{e.Event.IPAddr(key, ip)}
}

func (e *Event) IPPrefix(key string, pfx net.IPNet) *Event {
	return &Event{e.Event.IPPrefix(key, pfx)}
}

func (e *Event) MACAddr(key string, ha net.HardwareAddr) *Event {
	return &Event{e.Event.MACAddr(key, ha)}
}
