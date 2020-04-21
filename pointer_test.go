package ptr

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJsonMarshal(t *testing.T) {
	as := assert.New(t)
	model := struct {
		Name *String `json:"name,omitempty"`
		Age  *Int    `json:"age,omitempty"`
	}{
		Name: NewString("anna"),
	}
	j, err := json.Marshal(model)
	as.NoError(err)
	as.Equal(`{"name":"anna"}`, string(j))
	as.Zero(model.Age.Value())
}

func TestJsonUnMarshal(t *testing.T) {
	as := assert.New(t)
	model := struct {
		Name *String `json:"name,omitempty"`
		Age  *Int    `json:"age,omitempty"`
	}{}
	err := json.Unmarshal([]byte(`{"name":"anna"}`), &model)
	as.NoError(err)
	as.Equal(`anna`, model.Name.Value())
	as.Zero(model.Age.Value())
}

func TestString(t *testing.T) {
	as := assert.New(t)
	as.Equal("true", NewBool(true).String())
	as.Equal("false", NewBool(false).String())
	as.Equal("-1", NewInt(-1).String())
	as.Equal("0", NewInt(0).String())
	as.Equal("1", NewInt(1).String())
	as.Equal("-1", NewInt8(-1).String())
	as.Equal("0", NewInt8(0).String())
	as.Equal("1", NewInt8(1).String())
	as.Equal("-1", NewInt16(-1).String())
	as.Equal("0", NewInt16(0).String())
	as.Equal("1", NewInt16(1).String())
	as.Equal("-1", NewInt32(-1).String())
	as.Equal("0", NewInt32(0).String())
	as.Equal("1", NewInt32(1).String())
	as.Equal("-1", NewInt64(-1).String())
	as.Equal("0", NewInt64(0).String())
	as.Equal("1", NewInt64(1).String())
	as.Equal("0", NewUint(0).String())
	as.Equal("1", NewUint(1).String())
	as.Equal("0", NewUint8(0).String())
	as.Equal("1", NewUint8(1).String())
	as.Equal("0", NewUint16(0).String())
	as.Equal("1", NewUint16(1).String())
	as.Equal("0", NewUint32(0).String())
	as.Equal("1", NewUint32(1).String())
	as.Equal("0", NewUint64(0).String())
	as.Equal("1", NewUint64(1).String())
	as.Equal("1", NewFloat32(1.0).String())
	as.Equal("-1.1", NewFloat32(-1.1).String())
	as.Equal("1", NewFloat64(1.0).String())
	as.Equal("-1.1", NewFloat64(-1.1).String())
	as.Equal("(-1.1+1i)", NewComplex64(-1.1+1i).String())
	as.Equal("(1.1+0i)", NewComplex64(1.1).String())
	as.Equal("0001-01-01T00:00:00Z", NewTime(time.Time{}).String())
	as.Equal("2020-02-02T02:02:02.000000002Z", NewTime(
		time.Date(2020, 2, 2, 2, 2, 2, 2, time.UTC),
	).String())
}
