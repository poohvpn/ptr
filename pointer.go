package ptr

import (
	"fmt"
	"time"
)

type (
	Bool       bool
	Uint8      uint8
	Uint16     uint16
	Uint32     uint32
	Uint64     uint64
	Int8       int8
	Int16      int16
	Int32      int32
	Int64      int64
	Float32    float32
	Float64    float64
	Complex64  complex64
	Complex128 complex128
	String     string
	Int        int
	Uint       uint
	Time       time.Time
)

func (p *Bool) Value() (res bool) {
	if p != nil {
		return bool(*p)
	}
	return
}

func (p *Bool) Clone() *Bool {
	return NewBool(p.Value())
}

func (p *Bool) String() string {
	return fmt.Sprint(p.Value())
}

func (p *Uint8) Value() (res uint8) {
	if p != nil {
		return uint8(*p)
	}
	return
}

func (p *Uint8) Clone() *Uint8 {
	return NewUint8(p.Value())
}

func (p *Uint8) String() string {
	return fmt.Sprint(p.Value())
}

func (p *Uint16) Value() (res uint16) {
	if p != nil {
		return uint16(*p)
	}
	return
}

func (p *Uint16) Clone() *Uint16 {
	return NewUint16(p.Value())
}

func (p *Uint16) String() string {
	return fmt.Sprint(p.Value())
}

func (p *Uint32) Value() (res uint32) {
	if p != nil {
		return uint32(*p)
	}
	return
}

func (p *Uint32) Clone() *Uint32 {
	return NewUint32(p.Value())
}

func (p *Uint32) String() string {
	return fmt.Sprint(p.Value())
}

func (p *Uint64) Value() (res uint64) {
	if p != nil {
		return uint64(*p)
	}
	return
}

func (p *Uint64) Clone() *Uint64 {
	return NewUint64(p.Value())
}

func (p *Uint64) String() string {
	return fmt.Sprint(p.Value())
}

func (p *Int8) Value() (res int8) {
	if p != nil {
		return int8(*p)
	}
	return
}

func (p *Int8) Clone() *Int8 {
	return NewInt8(p.Value())
}

func (p *Int8) String() string {
	return fmt.Sprint(p.Value())
}

func (p *Int16) Value() (res int16) {
	if p != nil {
		return int16(*p)
	}
	return
}

func (p *Int16) Clone() *Int16 {
	return NewInt16(p.Value())
}

func (p *Int16) String() string {
	return fmt.Sprint(p.Value())
}

func (p *Int32) Value() (res int32) {
	if p != nil {
		return int32(*p)
	}
	return
}

func (p *Int32) Clone() *Int32 {
	return NewInt32(p.Value())
}

func (p *Int32) String() string {
	return fmt.Sprint(p.Value())
}

func (p *Int64) Value() (res int64) {
	if p != nil {
		return int64(*p)
	}
	return
}

func (p *Int64) Clone() *Int64 {
	return NewInt64(p.Value())
}

func (p *Int64) String() string {
	return fmt.Sprint(p.Value())
}

func (p *Float32) Value() (res float32) {
	if p != nil {
		return float32(*p)
	}
	return
}

func (p *Float32) Clone() *Float32 {
	return NewFloat32(p.Value())
}

func (p *Float32) String() string {
	return fmt.Sprint(p.Value())
}

func (p *Float64) Value() (res float64) {
	if p != nil {
		return float64(*p)
	}
	return
}

func (p *Float64) Clone() *Float64 {
	return NewFloat64(p.Value())
}

func (p *Float64) String() string {
	return fmt.Sprint(p.Value())
}

func (p *Complex64) Value() (res complex64) {
	if p != nil {
		return complex64(*p)
	}
	return
}

func (p *Complex64) Clone() *Complex64 {
	return NewComplex64(p.Value())
}

func (p *Complex64) String() string {
	return fmt.Sprint(p.Value())
}

func (p *Complex128) Value() (res complex128) {
	if p != nil {
		return complex128(*p)
	}
	return
}

func (p *Complex128) Clone() *Complex128 {
	return NewComplex128(p.Value())
}

func (p *Complex128) String() string {
	return fmt.Sprint(p.Value())
}

func (p *Int) Value() (res int) {
	if p != nil {
		return int(*p)
	}
	return
}

func (p *Int) Clone() *Int {
	return NewInt(p.Value())
}

func (p *Int) String() string {
	return fmt.Sprint(p.Value())
}

func (p *Uint) Value() (res uint) {
	if p != nil {
		return uint(*p)
	}
	return
}

func (p *Uint) Clone() *Uint {
	return NewUint(p.Value())
}

func (p *Uint) String() string {
	return fmt.Sprint(p.Value())
}

func (p *Time) Value() (res time.Time) {
	if p != nil {
		return time.Time(*p)
	}
	return
}

func (p *Time) Clone() *Time {
	return NewTime(p.Value())
}

func (p *Time) String() string {
	return string(p.Value().AppendFormat(make([]byte, 0, len(time.RFC3339Nano)), time.RFC3339Nano))
}

func (p *String) Value() (res string) {
	if p != nil {
		return string(*p)
	}
	return
}

func (p *String) Clone() *String {
	return NewString(p.Value())
}

func (p *String) String() string {
	return p.Value()
}
