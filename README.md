# Generalized Golang (Slice trick)[https://github.com/golang/go/wiki/SliceTricks] with reflect
### Are you ready to exchange runtime speed with ease of coding?


### Benchmark result

##### with string only precompiled

Function Name                  StringSpecified  Reflection
===========================================================

##### with all default types precompiled


Function Name                  StringSpecified  Reflection
===========================================================



### How to change precompiled types list
 - Open `generate.go` and edit type list (line number 2). Below is default value specifying all of the go value types.
```
//go:generate ./template -d bool int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr float32 float64 complex64 complex128 string
```
 - Type `go generate` to generate and build.
