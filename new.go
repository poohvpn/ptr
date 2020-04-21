package ptr

import (
	"time"
)

func NewBool(v bool) *Bool {
	return (*Bool)(&v)
}

func NewUint8(v uint8) *Uint8 {
	return (*Uint8)(&v)
}

func NewUint16(v uint16) *Uint16 {
	return (*Uint16)(&v)
}

func NewUint32(v uint32) *Uint32 {
	return (*Uint32)(&v)
}

func NewUint64(v uint64) *Uint64 {
	return (*Uint64)(&v)
}

func NewInt8(v int8) *Int8 {
	return (*Int8)(&v)
}

func NewInt16(v int16) *Int16 {
	return (*Int16)(&v)
}

func NewInt32(v int32) *Int32 {
	return (*Int32)(&v)
}

func NewInt64(v int64) *Int64 {
	return (*Int64)(&v)
}

func NewFloat32(v float32) *Float32 {
	return (*Float32)(&v)
}

func NewFloat64(v float64) *Float64 {
	return (*Float64)(&v)
}

func NewComplex64(v complex64) *Complex64 {
	return (*Complex64)(&v)
}

func NewComplex128(v complex128) *Complex128 {
	return (*Complex128)(&v)
}

func NewString(v string) *String {
	return (*String)(&v)
}

func NewInt(v int) *Int {
	return (*Int)(&v)
}

func NewUint(v uint) *Uint {
	return (*Uint)(&v)
}

func NewTime(v time.Time) *Time {
	return (*Time)(&v)
}
