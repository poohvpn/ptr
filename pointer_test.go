package ptr

import (
	"encoding/json"
	"testing"

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
