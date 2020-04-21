# Pointer Helper for Go
This library can help to handle null value with golang primitive types, which is implemented by pointer.

## Installing
```shell script
go get -u github.com/poohvpn/ptr
```

## Usage
```go
package main

import (
    "github.com/poohvpn/ptr"
)

func main() {
	model := struct {
		Name *ptr.String `json:"name,omitempty"`
		Age  *ptr.Int    `json:"age,omitempty"`
	}{
		Name: ptr.NewString("anna"),
	}

	j, err := json.Marshal(model)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(j), model.Age.Value())
	// {"name":"anna"} 0
}
```