package ptr

import (
	"time"
)

func NewBool(v bool) *Bool {
	p := Bool(v)
	return &p
}

func NewUint8(v uint8) *Uint8 {
	p := Uint8(v)
	return &p
}

func NewUint16(v uint16) *Uint16 {
	p := Uint16(v)
	return &p
}

func NewUint32(v uint32) *Uint32 {
	p := Uint32(v)
	return &p
}

func NewUint64(v uint64) *Uint64 {
	p := Uint64(v)
	return &p
}

func NewInt8(v int8) *Int8 {
	p := Int8(v)
	return &p
}

func NewInt16(v int16) *Int16 {
	p := Int16(v)
	return &p
}

func NewInt32(v int32) *Int32 {
	p := Int32(v)
	return &p
}

func NewInt64(v int64) *Int64 {
	p := Int64(v)
	return &p
}

func NewFloat32(v float32) *Float32 {
	p := Float32(v)
	return &p
}

func NewFloat64(v float64) *Float64 {
	p := Float64(v)
	return &p
}

func NewComplex64(v complex64) *Complex64 {
	p := Complex64(v)
	return &p
}

func NewComplex128(v complex128) *Complex128 {
	p := Complex128(v)
	return &p
}

func NewString(v string) *String {
	p := String(v)
	return &p
}

func NewInt(v int) *Int {
	p := Int(v)
	return &p
}

func NewUint(v uint) *Uint {
	p := Uint(v)
	return &p
}

func NewTime(v time.Time) *Time {
	p := Time(v)
	return &p
}
