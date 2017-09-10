//go:generate go build go-type-template/template.go
//go:generate ./template -d bool int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr float32 float64 complex64 complex128 string
//go:generate go build

package sliceTrick
