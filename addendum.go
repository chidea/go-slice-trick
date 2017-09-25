// Code generated with github.com/chidea/go-type-template DO NOT EDIT.
		
package sliceTrick

import (
	//"log"
	"reflect"
)

var Equal = reflect.DeepEqual

func Len(a interface{}) int {
	return reflect.ValueOf(a).Len()
}

type T struct{ Val interface{} }

func (self T) In(slice interface{}) bool {
	return In(slice, self.Val)
}
func In(slice, val interface{}) bool {
	vov := reflect.ValueOf(val)
	sliceval := reflect.ValueOf(slice)
	for i := 0; i < sliceval.Len(); i++ {
		if Equal(sliceval.Index(i).Interface(), vov.Interface()) {
			return true
		}
	}
	return false
}

func InBool(slice []bool, val bool) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}


func InInt(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}


func InInt8(slice []int8, val int8) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}


func InInt16(slice []int16, val int16) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}


func InInt32(slice []int32, val int32) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}


func InInt64(slice []int64, val int64) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}


func InUint(slice []uint, val uint) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}


func InUint8(slice []uint8, val uint8) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}


func InUint16(slice []uint16, val uint16) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}


func InUint32(slice []uint32, val uint32) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}


func InUint64(slice []uint64, val uint64) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}


func InUintptr(slice []uintptr, val uintptr) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}


func InFloat32(slice []float32, val float32) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}


func InFloat64(slice []float64, val float64) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}


func InComplex64(slice []complex64, val complex64) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}


func InComplex128(slice []complex128, val complex128) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}


func InStr(slice []string, val string) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func Iif(test bool, onTrue interface{}, onFalse interface{}) interface{} {
	if test {
		return onTrue
	}
	return onFalse
}

func IifBool(test bool, onTrue bool, onFalse bool) bool {
	if test {
		return onTrue
	}
	return onFalse
}


func IifInt(test bool, onTrue int, onFalse int) int {
	if test {
		return onTrue
	}
	return onFalse
}


func IifInt8(test bool, onTrue int8, onFalse int8) int8 {
	if test {
		return onTrue
	}
	return onFalse
}


func IifInt16(test bool, onTrue int16, onFalse int16) int16 {
	if test {
		return onTrue
	}
	return onFalse
}


func IifInt32(test bool, onTrue int32, onFalse int32) int32 {
	if test {
		return onTrue
	}
	return onFalse
}


func IifInt64(test bool, onTrue int64, onFalse int64) int64 {
	if test {
		return onTrue
	}
	return onFalse
}


func IifUint(test bool, onTrue uint, onFalse uint) uint {
	if test {
		return onTrue
	}
	return onFalse
}


func IifUint8(test bool, onTrue uint8, onFalse uint8) uint8 {
	if test {
		return onTrue
	}
	return onFalse
}


func IifUint16(test bool, onTrue uint16, onFalse uint16) uint16 {
	if test {
		return onTrue
	}
	return onFalse
}


func IifUint32(test bool, onTrue uint32, onFalse uint32) uint32 {
	if test {
		return onTrue
	}
	return onFalse
}


func IifUint64(test bool, onTrue uint64, onFalse uint64) uint64 {
	if test {
		return onTrue
	}
	return onFalse
}


func IifUintptr(test bool, onTrue uintptr, onFalse uintptr) uintptr {
	if test {
		return onTrue
	}
	return onFalse
}


func IifFloat32(test bool, onTrue float32, onFalse float32) float32 {
	if test {
		return onTrue
	}
	return onFalse
}


func IifFloat64(test bool, onTrue float64, onFalse float64) float64 {
	if test {
		return onTrue
	}
	return onFalse
}


func IifComplex64(test bool, onTrue complex64, onFalse complex64) complex64 {
	if test {
		return onTrue
	}
	return onFalse
}


func IifComplex128(test bool, onTrue complex128, onFalse complex128) complex128 {
	if test {
		return onTrue
	}
	return onFalse
}


func IifStr(test bool, onTrue string, onFalse string) string {
	if test {
		return onTrue
	}
	return onFalse
}
