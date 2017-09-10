// Code generated with github.com/chidea/go-type-template DO NOT EDIT.
		

package sliceTrick

import (
	//"fmt"
	"reflect"
)

func AppendBool(a, b []bool) []bool {
	return append(a, b...)
}

func AppendInt(a, b []int) []int {
	return append(a, b...)
}

func AppendInt8(a, b []int8) []int8 {
	return append(a, b...)
}

func AppendInt16(a, b []int16) []int16 {
	return append(a, b...)
}

func AppendInt32(a, b []int32) []int32 {
	return append(a, b...)
}

func AppendInt64(a, b []int64) []int64 {
	return append(a, b...)
}

func AppendUint(a, b []uint) []uint {
	return append(a, b...)
}

func AppendUint8(a, b []uint8) []uint8 {
	return append(a, b...)
}

func AppendUint16(a, b []uint16) []uint16 {
	return append(a, b...)
}

func AppendUint32(a, b []uint32) []uint32 {
	return append(a, b...)
}

func AppendUint64(a, b []uint64) []uint64 {
	return append(a, b...)
}

func AppendUintptr(a, b []uintptr) []uintptr {
	return append(a, b...)
}

func AppendFloat32(a, b []float32) []float32 {
	return append(a, b...)
}

func AppendFloat64(a, b []float64) []float64 {
	return append(a, b...)
}

func AppendComplex64(a, b []complex64) []complex64 {
	return append(a, b...)
}

func AppendComplex128(a, b []complex128) []complex128 {
	return append(a, b...)
}

func AppendStr(a, b []string) []string {
	return append(a, b...)
}

func Append(a, b interface{}) interface{} {
	/*switch a.(type) {
	case bool:
		return interface{}(AppendBool(a.([]bool), b.([]bool)))

	case int:
		return interface{}(AppendInt(a.([]int), b.([]int)))

	case int8:
		return interface{}(AppendInt8(a.([]int8), b.([]int8)))

	case int16:
		return interface{}(AppendInt16(a.([]int16), b.([]int16)))

	case int32:
		return interface{}(AppendInt32(a.([]int32), b.([]int32)))

	case int64:
		return interface{}(AppendInt64(a.([]int64), b.([]int64)))

	case uint:
		return interface{}(AppendUint(a.([]uint), b.([]uint)))

	case uint8:
		return interface{}(AppendUint8(a.([]uint8), b.([]uint8)))

	case uint16:
		return interface{}(AppendUint16(a.([]uint16), b.([]uint16)))

	case uint32:
		return interface{}(AppendUint32(a.([]uint32), b.([]uint32)))

	case uint64:
		return interface{}(AppendUint64(a.([]uint64), b.([]uint64)))

	case uintptr:
		return interface{}(AppendUintptr(a.([]uintptr), b.([]uintptr)))

	case float32:
		return interface{}(AppendFloat32(a.([]float32), b.([]float32)))

	case float64:
		return interface{}(AppendFloat64(a.([]float64), b.([]float64)))

	case complex64:
		return interface{}(AppendComplex64(a.([]complex64), b.([]complex64)))

	case complex128:
		return interface{}(AppendComplex128(a.([]complex128), b.([]complex128)))

	case string:
		return interface{}(AppendStr(a.([]string), b.([]string)))
	} //fall back to reflect lib when it's not a precompiled type*/
	return reflect.AppendSlice(reflect.ValueOf(a), reflect.ValueOf(b)).Interface()
}
func CopyBool(a []bool) []bool {
	/*b := make([]bool, len(a))
	copy(b, a)
	return b*/
	// or
	return append([]bool(nil), a...)
}

func CopyInt(a []int) []int {
	/*b := make([]int, len(a))
	copy(b, a)
	return b*/
	// or
	return append([]int(nil), a...)
}

func CopyInt8(a []int8) []int8 {
	/*b := make([]int8, len(a))
	copy(b, a)
	return b*/
	// or
	return append([]int8(nil), a...)
}

func CopyInt16(a []int16) []int16 {
	/*b := make([]int16, len(a))
	copy(b, a)
	return b*/
	// or
	return append([]int16(nil), a...)
}

func CopyInt32(a []int32) []int32 {
	/*b := make([]int32, len(a))
	copy(b, a)
	return b*/
	// or
	return append([]int32(nil), a...)
}

func CopyInt64(a []int64) []int64 {
	/*b := make([]int64, len(a))
	copy(b, a)
	return b*/
	// or
	return append([]int64(nil), a...)
}

func CopyUint(a []uint) []uint {
	/*b := make([]uint, len(a))
	copy(b, a)
	return b*/
	// or
	return append([]uint(nil), a...)
}

func CopyUint8(a []uint8) []uint8 {
	/*b := make([]uint8, len(a))
	copy(b, a)
	return b*/
	// or
	return append([]uint8(nil), a...)
}

func CopyUint16(a []uint16) []uint16 {
	/*b := make([]uint16, len(a))
	copy(b, a)
	return b*/
	// or
	return append([]uint16(nil), a...)
}

func CopyUint32(a []uint32) []uint32 {
	/*b := make([]uint32, len(a))
	copy(b, a)
	return b*/
	// or
	return append([]uint32(nil), a...)
}

func CopyUint64(a []uint64) []uint64 {
	/*b := make([]uint64, len(a))
	copy(b, a)
	return b*/
	// or
	return append([]uint64(nil), a...)
}

func CopyUintptr(a []uintptr) []uintptr {
	/*b := make([]uintptr, len(a))
	copy(b, a)
	return b*/
	// or
	return append([]uintptr(nil), a...)
}

func CopyFloat32(a []float32) []float32 {
	/*b := make([]float32, len(a))
	copy(b, a)
	return b*/
	// or
	return append([]float32(nil), a...)
}

func CopyFloat64(a []float64) []float64 {
	/*b := make([]float64, len(a))
	copy(b, a)
	return b*/
	// or
	return append([]float64(nil), a...)
}

func CopyComplex64(a []complex64) []complex64 {
	/*b := make([]complex64, len(a))
	copy(b, a)
	return b*/
	// or
	return append([]complex64(nil), a...)
}

func CopyComplex128(a []complex128) []complex128 {
	/*b := make([]complex128, len(a))
	copy(b, a)
	return b*/
	// or
	return append([]complex128(nil), a...)
}

func CopyStr(a []string) []string {
	/*b := make([]string, len(a))
	copy(b, a)
	return b*/
	// or
	return append([]string(nil), a...)
}
func Copy(a interface{}) interface{} {
	/*switch a.(type) {
	case bool:
		return interface{}(CopyBool(a.([]bool)))

	case int:
		return interface{}(CopyInt(a.([]int)))

	case int8:
		return interface{}(CopyInt8(a.([]int8)))

	case int16:
		return interface{}(CopyInt16(a.([]int16)))

	case int32:
		return interface{}(CopyInt32(a.([]int32)))

	case int64:
		return interface{}(CopyInt64(a.([]int64)))

	case uint:
		return interface{}(CopyUint(a.([]uint)))

	case uint8:
		return interface{}(CopyUint8(a.([]uint8)))

	case uint16:
		return interface{}(CopyUint16(a.([]uint16)))

	case uint32:
		return interface{}(CopyUint32(a.([]uint32)))

	case uint64:
		return interface{}(CopyUint64(a.([]uint64)))

	case uintptr:
		return interface{}(CopyUintptr(a.([]uintptr)))

	case float32:
		return interface{}(CopyFloat32(a.([]float32)))

	case float64:
		return interface{}(CopyFloat64(a.([]float64)))

	case complex64:
		return interface{}(CopyComplex64(a.([]complex64)))

	case complex128:
		return interface{}(CopyComplex128(a.([]complex128)))

	case string:
		return interface{}(CopyStr(a.([]string)))
	}*/
	/*aval := reflect.ValueOf(a)
	alen := aval.Len()
	bval := reflect.MakeSlice(reflect.TypeOf(a), alen, alen)
	reflect.Copy(bval, aval)
	return bval.Interface()*/
	// or
	aval := reflect.ValueOf(a)
	return reflect.AppendSlice(reflect.MakeSlice(reflect.TypeOf(a), 0, aval.Len()), aval.Slice(0, aval.Len())).Interface()
}
func CutBool(a []bool, i, j int) []bool {
	return append(a[:i], a[j:]...)
}

func CutInt(a []int, i, j int) []int {
	return append(a[:i], a[j:]...)
}

func CutInt8(a []int8, i, j int) []int8 {
	return append(a[:i], a[j:]...)
}

func CutInt16(a []int16, i, j int) []int16 {
	return append(a[:i], a[j:]...)
}

func CutInt32(a []int32, i, j int) []int32 {
	return append(a[:i], a[j:]...)
}

func CutInt64(a []int64, i, j int) []int64 {
	return append(a[:i], a[j:]...)
}

func CutUint(a []uint, i, j int) []uint {
	return append(a[:i], a[j:]...)
}

func CutUint8(a []uint8, i, j int) []uint8 {
	return append(a[:i], a[j:]...)
}

func CutUint16(a []uint16, i, j int) []uint16 {
	return append(a[:i], a[j:]...)
}

func CutUint32(a []uint32, i, j int) []uint32 {
	return append(a[:i], a[j:]...)
}

func CutUint64(a []uint64, i, j int) []uint64 {
	return append(a[:i], a[j:]...)
}

func CutUintptr(a []uintptr, i, j int) []uintptr {
	return append(a[:i], a[j:]...)
}

func CutFloat32(a []float32, i, j int) []float32 {
	return append(a[:i], a[j:]...)
}

func CutFloat64(a []float64, i, j int) []float64 {
	return append(a[:i], a[j:]...)
}

func CutComplex64(a []complex64, i, j int) []complex64 {
	return append(a[:i], a[j:]...)
}

func CutComplex128(a []complex128, i, j int) []complex128 {
	return append(a[:i], a[j:]...)
}

func CutStr(a []string, i, j int) []string {
	return append(a[:i], a[j:]...)
}
func Cut(a interface{}, i, j int) interface{} {
	/*switch a.(type) {
	case bool:
		return interface{}(CutBool(a.([]bool), i, j))

	case int:
		return interface{}(CutInt(a.([]int), i, j))

	case int8:
		return interface{}(CutInt8(a.([]int8), i, j))

	case int16:
		return interface{}(CutInt16(a.([]int16), i, j))

	case int32:
		return interface{}(CutInt32(a.([]int32), i, j))

	case int64:
		return interface{}(CutInt64(a.([]int64), i, j))

	case uint:
		return interface{}(CutUint(a.([]uint), i, j))

	case uint8:
		return interface{}(CutUint8(a.([]uint8), i, j))

	case uint16:
		return interface{}(CutUint16(a.([]uint16), i, j))

	case uint32:
		return interface{}(CutUint32(a.([]uint32), i, j))

	case uint64:
		return interface{}(CutUint64(a.([]uint64), i, j))

	case uintptr:
		return interface{}(CutUintptr(a.([]uintptr), i, j))

	case float32:
		return interface{}(CutFloat32(a.([]float32), i, j))

	case float64:
		return interface{}(CutFloat64(a.([]float64), i, j))

	case complex64:
		return interface{}(CutComplex64(a.([]complex64), i, j))

	case complex128:
		return interface{}(CutComplex128(a.([]complex128), i, j))

	case string:
		return interface{}(CutStr(a.([]string), i, j))
	}*/
	aval := reflect.ValueOf(a)
	return reflect.AppendSlice(aval.Slice(0, i), aval.Slice(j, aval.Len())).Interface()
}
func DeleteBool(a []bool, i int) []bool {
	return append(a[:i], a[i+1:]...)
	// or
	//a = a[:i+copy(a[i:], a[i+1:])]
}

func DeleteInt(a []int, i int) []int {
	return append(a[:i], a[i+1:]...)
	// or
	//a = a[:i+copy(a[i:], a[i+1:])]
}

func DeleteInt8(a []int8, i int) []int8 {
	return append(a[:i], a[i+1:]...)
	// or
	//a = a[:i+copy(a[i:], a[i+1:])]
}

func DeleteInt16(a []int16, i int) []int16 {
	return append(a[:i], a[i+1:]...)
	// or
	//a = a[:i+copy(a[i:], a[i+1:])]
}

func DeleteInt32(a []int32, i int) []int32 {
	return append(a[:i], a[i+1:]...)
	// or
	//a = a[:i+copy(a[i:], a[i+1:])]
}

func DeleteInt64(a []int64, i int) []int64 {
	return append(a[:i], a[i+1:]...)
	// or
	//a = a[:i+copy(a[i:], a[i+1:])]
}

func DeleteUint(a []uint, i int) []uint {
	return append(a[:i], a[i+1:]...)
	// or
	//a = a[:i+copy(a[i:], a[i+1:])]
}

func DeleteUint8(a []uint8, i int) []uint8 {
	return append(a[:i], a[i+1:]...)
	// or
	//a = a[:i+copy(a[i:], a[i+1:])]
}

func DeleteUint16(a []uint16, i int) []uint16 {
	return append(a[:i], a[i+1:]...)
	// or
	//a = a[:i+copy(a[i:], a[i+1:])]
}

func DeleteUint32(a []uint32, i int) []uint32 {
	return append(a[:i], a[i+1:]...)
	// or
	//a = a[:i+copy(a[i:], a[i+1:])]
}

func DeleteUint64(a []uint64, i int) []uint64 {
	return append(a[:i], a[i+1:]...)
	// or
	//a = a[:i+copy(a[i:], a[i+1:])]
}

func DeleteUintptr(a []uintptr, i int) []uintptr {
	return append(a[:i], a[i+1:]...)
	// or
	//a = a[:i+copy(a[i:], a[i+1:])]
}

func DeleteFloat32(a []float32, i int) []float32 {
	return append(a[:i], a[i+1:]...)
	// or
	//a = a[:i+copy(a[i:], a[i+1:])]
}

func DeleteFloat64(a []float64, i int) []float64 {
	return append(a[:i], a[i+1:]...)
	// or
	//a = a[:i+copy(a[i:], a[i+1:])]
}

func DeleteComplex64(a []complex64, i int) []complex64 {
	return append(a[:i], a[i+1:]...)
	// or
	//a = a[:i+copy(a[i:], a[i+1:])]
}

func DeleteComplex128(a []complex128, i int) []complex128 {
	return append(a[:i], a[i+1:]...)
	// or
	//a = a[:i+copy(a[i:], a[i+1:])]
}

func DeleteStr(a []string, i int) []string {
	return append(a[:i], a[i+1:]...)
	// or
	//a = a[:i+copy(a[i:], a[i+1:])]
}
func Delete(a interface{}, i int) interface{} {
	/*switch a.(type) {
	case bool:
		return interface{}(DeleteBool(a.([]bool), i))

	case int:
		return interface{}(DeleteInt(a.([]int), i))

	case int8:
		return interface{}(DeleteInt8(a.([]int8), i))

	case int16:
		return interface{}(DeleteInt16(a.([]int16), i))

	case int32:
		return interface{}(DeleteInt32(a.([]int32), i))

	case int64:
		return interface{}(DeleteInt64(a.([]int64), i))

	case uint:
		return interface{}(DeleteUint(a.([]uint), i))

	case uint8:
		return interface{}(DeleteUint8(a.([]uint8), i))

	case uint16:
		return interface{}(DeleteUint16(a.([]uint16), i))

	case uint32:
		return interface{}(DeleteUint32(a.([]uint32), i))

	case uint64:
		return interface{}(DeleteUint64(a.([]uint64), i))

	case uintptr:
		return interface{}(DeleteUintptr(a.([]uintptr), i))

	case float32:
		return interface{}(DeleteFloat32(a.([]float32), i))

	case float64:
		return interface{}(DeleteFloat64(a.([]float64), i))

	case complex64:
		return interface{}(DeleteComplex64(a.([]complex64), i))

	case complex128:
		return interface{}(DeleteComplex128(a.([]complex128), i))

	case string:
		return interface{}(DeleteStr(a.([]string), i))
	}*/
	aval := reflect.ValueOf(a)
	return reflect.AppendSlice(aval.Slice(0, i), aval.Slice(i+1, aval.Len())).Interface()
}

func DeleteWithoutPreservingOrderBool(a []bool, i int) []bool {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func DeleteWithoutPreservingOrderInt(a []int, i int) []int {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func DeleteWithoutPreservingOrderInt8(a []int8, i int) []int8 {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func DeleteWithoutPreservingOrderInt16(a []int16, i int) []int16 {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func DeleteWithoutPreservingOrderInt32(a []int32, i int) []int32 {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func DeleteWithoutPreservingOrderInt64(a []int64, i int) []int64 {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func DeleteWithoutPreservingOrderUint(a []uint, i int) []uint {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func DeleteWithoutPreservingOrderUint8(a []uint8, i int) []uint8 {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func DeleteWithoutPreservingOrderUint16(a []uint16, i int) []uint16 {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func DeleteWithoutPreservingOrderUint32(a []uint32, i int) []uint32 {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func DeleteWithoutPreservingOrderUint64(a []uint64, i int) []uint64 {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func DeleteWithoutPreservingOrderUintptr(a []uintptr, i int) []uintptr {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func DeleteWithoutPreservingOrderFloat32(a []float32, i int) []float32 {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func DeleteWithoutPreservingOrderFloat64(a []float64, i int) []float64 {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func DeleteWithoutPreservingOrderComplex64(a []complex64, i int) []complex64 {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func DeleteWithoutPreservingOrderComplex128(a []complex128, i int) []complex128 {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func DeleteWithoutPreservingOrderStr(a []string, i int) []string {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}
func DeleteWithoutPreservingOrder(a interface{}, i int) interface{} { //interface {} {
	/*switch a.(type) {
	case bool:
		return interface{}(DeleteWithoutPreservingOrderBool(a.([]bool), i))

	case int:
		return interface{}(DeleteWithoutPreservingOrderInt(a.([]int), i))

	case int8:
		return interface{}(DeleteWithoutPreservingOrderInt8(a.([]int8), i))

	case int16:
		return interface{}(DeleteWithoutPreservingOrderInt16(a.([]int16), i))

	case int32:
		return interface{}(DeleteWithoutPreservingOrderInt32(a.([]int32), i))

	case int64:
		return interface{}(DeleteWithoutPreservingOrderInt64(a.([]int64), i))

	case uint:
		return interface{}(DeleteWithoutPreservingOrderUint(a.([]uint), i))

	case uint8:
		return interface{}(DeleteWithoutPreservingOrderUint8(a.([]uint8), i))

	case uint16:
		return interface{}(DeleteWithoutPreservingOrderUint16(a.([]uint16), i))

	case uint32:
		return interface{}(DeleteWithoutPreservingOrderUint32(a.([]uint32), i))

	case uint64:
		return interface{}(DeleteWithoutPreservingOrderUint64(a.([]uint64), i))

	case uintptr:
		return interface{}(DeleteWithoutPreservingOrderUintptr(a.([]uintptr), i))

	case float32:
		return interface{}(DeleteWithoutPreservingOrderFloat32(a.([]float32), i))

	case float64:
		return interface{}(DeleteWithoutPreservingOrderFloat64(a.([]float64), i))

	case complex64:
		return interface{}(DeleteWithoutPreservingOrderComplex64(a.([]complex64), i))

	case complex128:
		return interface{}(DeleteWithoutPreservingOrderComplex128(a.([]complex128), i))

	case string:
		return interface{}(DeleteWithoutPreservingOrderStr(a.([]string), i))
	}*/
	aval := reflect.ValueOf(a)
	l := aval.Len()
	aval.Index(i).Set(aval.Index(l - 1))
	return aval.Slice(0, l-1).Interface()
}

func ExpandBool(a []bool, i, j int) []bool {
	return append(a[:i], append(make([]bool, j), a[i:]...)...)
}

func ExpandInt(a []int, i, j int) []int {
	return append(a[:i], append(make([]int, j), a[i:]...)...)
}

func ExpandInt8(a []int8, i, j int) []int8 {
	return append(a[:i], append(make([]int8, j), a[i:]...)...)
}

func ExpandInt16(a []int16, i, j int) []int16 {
	return append(a[:i], append(make([]int16, j), a[i:]...)...)
}

func ExpandInt32(a []int32, i, j int) []int32 {
	return append(a[:i], append(make([]int32, j), a[i:]...)...)
}

func ExpandInt64(a []int64, i, j int) []int64 {
	return append(a[:i], append(make([]int64, j), a[i:]...)...)
}

func ExpandUint(a []uint, i, j int) []uint {
	return append(a[:i], append(make([]uint, j), a[i:]...)...)
}

func ExpandUint8(a []uint8, i, j int) []uint8 {
	return append(a[:i], append(make([]uint8, j), a[i:]...)...)
}

func ExpandUint16(a []uint16, i, j int) []uint16 {
	return append(a[:i], append(make([]uint16, j), a[i:]...)...)
}

func ExpandUint32(a []uint32, i, j int) []uint32 {
	return append(a[:i], append(make([]uint32, j), a[i:]...)...)
}

func ExpandUint64(a []uint64, i, j int) []uint64 {
	return append(a[:i], append(make([]uint64, j), a[i:]...)...)
}

func ExpandUintptr(a []uintptr, i, j int) []uintptr {
	return append(a[:i], append(make([]uintptr, j), a[i:]...)...)
}

func ExpandFloat32(a []float32, i, j int) []float32 {
	return append(a[:i], append(make([]float32, j), a[i:]...)...)
}

func ExpandFloat64(a []float64, i, j int) []float64 {
	return append(a[:i], append(make([]float64, j), a[i:]...)...)
}

func ExpandComplex64(a []complex64, i, j int) []complex64 {
	return append(a[:i], append(make([]complex64, j), a[i:]...)...)
}

func ExpandComplex128(a []complex128, i, j int) []complex128 {
	return append(a[:i], append(make([]complex128, j), a[i:]...)...)
}

func ExpandStr(a []string, i, j int) []string {
	return append(a[:i], append(make([]string, j), a[i:]...)...)
}
func Expand(a interface{}, i, j int) interface{} {
	/*switch a.(type) {
	case bool:
		return interface{}(ExpandBool(a.([]bool), i, j))

	case int:
		return interface{}(ExpandInt(a.([]int), i, j))

	case int8:
		return interface{}(ExpandInt8(a.([]int8), i, j))

	case int16:
		return interface{}(ExpandInt16(a.([]int16), i, j))

	case int32:
		return interface{}(ExpandInt32(a.([]int32), i, j))

	case int64:
		return interface{}(ExpandInt64(a.([]int64), i, j))

	case uint:
		return interface{}(ExpandUint(a.([]uint), i, j))

	case uint8:
		return interface{}(ExpandUint8(a.([]uint8), i, j))

	case uint16:
		return interface{}(ExpandUint16(a.([]uint16), i, j))

	case uint32:
		return interface{}(ExpandUint32(a.([]uint32), i, j))

	case uint64:
		return interface{}(ExpandUint64(a.([]uint64), i, j))

	case uintptr:
		return interface{}(ExpandUintptr(a.([]uintptr), i, j))

	case float32:
		return interface{}(ExpandFloat32(a.([]float32), i, j))

	case float64:
		return interface{}(ExpandFloat64(a.([]float64), i, j))

	case complex64:
		return interface{}(ExpandComplex64(a.([]complex64), i, j))

	case complex128:
		return interface{}(ExpandComplex128(a.([]complex128), i, j))

	case string:
		return interface{}(ExpandStr(a.([]string), i, j))
	}*/
	aval := reflect.ValueOf(a)
	return reflect.AppendSlice(aval.Slice(0, i), reflect.AppendSlice(reflect.MakeSlice(reflect.TypeOf(a), j, j), aval.Slice(i, aval.Len()))).Interface()
}
func ExtendBool(a []bool, i int) []bool {
	return append(a, make([]bool, i)...)
}

func ExtendInt(a []int, i int) []int {
	return append(a, make([]int, i)...)
}

func ExtendInt8(a []int8, i int) []int8 {
	return append(a, make([]int8, i)...)
}

func ExtendInt16(a []int16, i int) []int16 {
	return append(a, make([]int16, i)...)
}

func ExtendInt32(a []int32, i int) []int32 {
	return append(a, make([]int32, i)...)
}

func ExtendInt64(a []int64, i int) []int64 {
	return append(a, make([]int64, i)...)
}

func ExtendUint(a []uint, i int) []uint {
	return append(a, make([]uint, i)...)
}

func ExtendUint8(a []uint8, i int) []uint8 {
	return append(a, make([]uint8, i)...)
}

func ExtendUint16(a []uint16, i int) []uint16 {
	return append(a, make([]uint16, i)...)
}

func ExtendUint32(a []uint32, i int) []uint32 {
	return append(a, make([]uint32, i)...)
}

func ExtendUint64(a []uint64, i int) []uint64 {
	return append(a, make([]uint64, i)...)
}

func ExtendUintptr(a []uintptr, i int) []uintptr {
	return append(a, make([]uintptr, i)...)
}

func ExtendFloat32(a []float32, i int) []float32 {
	return append(a, make([]float32, i)...)
}

func ExtendFloat64(a []float64, i int) []float64 {
	return append(a, make([]float64, i)...)
}

func ExtendComplex64(a []complex64, i int) []complex64 {
	return append(a, make([]complex64, i)...)
}

func ExtendComplex128(a []complex128, i int) []complex128 {
	return append(a, make([]complex128, i)...)
}

func ExtendStr(a []string, i int) []string {
	return append(a, make([]string, i)...)
}
func Extend(a interface{}, i int) interface{} {
	/*switch a.(type) {
	case bool:
		return interface{}(ExtendBool(a.([]bool), i))

	case int:
		return interface{}(ExtendInt(a.([]int), i))

	case int8:
		return interface{}(ExtendInt8(a.([]int8), i))

	case int16:
		return interface{}(ExtendInt16(a.([]int16), i))

	case int32:
		return interface{}(ExtendInt32(a.([]int32), i))

	case int64:
		return interface{}(ExtendInt64(a.([]int64), i))

	case uint:
		return interface{}(ExtendUint(a.([]uint), i))

	case uint8:
		return interface{}(ExtendUint8(a.([]uint8), i))

	case uint16:
		return interface{}(ExtendUint16(a.([]uint16), i))

	case uint32:
		return interface{}(ExtendUint32(a.([]uint32), i))

	case uint64:
		return interface{}(ExtendUint64(a.([]uint64), i))

	case uintptr:
		return interface{}(ExtendUintptr(a.([]uintptr), i))

	case float32:
		return interface{}(ExtendFloat32(a.([]float32), i))

	case float64:
		return interface{}(ExtendFloat64(a.([]float64), i))

	case complex64:
		return interface{}(ExtendComplex64(a.([]complex64), i))

	case complex128:
		return interface{}(ExtendComplex128(a.([]complex128), i))

	case string:
		return interface{}(ExtendStr(a.([]string), i))
	}*/
	return reflect.AppendSlice(reflect.ValueOf(a), reflect.MakeSlice(reflect.TypeOf(a), i, i)).Interface()
}
func InsertBool(a []bool, x bool, i int) []bool {
	return append(a[:i], append([]bool{x}, a[i:]...)...)
}

func InsertInt(a []int, x int, i int) []int {
	return append(a[:i], append([]int{x}, a[i:]...)...)
}

func InsertInt8(a []int8, x int8, i int) []int8 {
	return append(a[:i], append([]int8{x}, a[i:]...)...)
}

func InsertInt16(a []int16, x int16, i int) []int16 {
	return append(a[:i], append([]int16{x}, a[i:]...)...)
}

func InsertInt32(a []int32, x int32, i int) []int32 {
	return append(a[:i], append([]int32{x}, a[i:]...)...)
}

func InsertInt64(a []int64, x int64, i int) []int64 {
	return append(a[:i], append([]int64{x}, a[i:]...)...)
}

func InsertUint(a []uint, x uint, i int) []uint {
	return append(a[:i], append([]uint{x}, a[i:]...)...)
}

func InsertUint8(a []uint8, x uint8, i int) []uint8 {
	return append(a[:i], append([]uint8{x}, a[i:]...)...)
}

func InsertUint16(a []uint16, x uint16, i int) []uint16 {
	return append(a[:i], append([]uint16{x}, a[i:]...)...)
}

func InsertUint32(a []uint32, x uint32, i int) []uint32 {
	return append(a[:i], append([]uint32{x}, a[i:]...)...)
}

func InsertUint64(a []uint64, x uint64, i int) []uint64 {
	return append(a[:i], append([]uint64{x}, a[i:]...)...)
}

func InsertUintptr(a []uintptr, x uintptr, i int) []uintptr {
	return append(a[:i], append([]uintptr{x}, a[i:]...)...)
}

func InsertFloat32(a []float32, x float32, i int) []float32 {
	return append(a[:i], append([]float32{x}, a[i:]...)...)
}

func InsertFloat64(a []float64, x float64, i int) []float64 {
	return append(a[:i], append([]float64{x}, a[i:]...)...)
}

func InsertComplex64(a []complex64, x complex64, i int) []complex64 {
	return append(a[:i], append([]complex64{x}, a[i:]...)...)
}

func InsertComplex128(a []complex128, x complex128, i int) []complex128 {
	return append(a[:i], append([]complex128{x}, a[i:]...)...)
}

func InsertStr(a []string, x string, i int) []string {
	return append(a[:i], append([]string{x}, a[i:]...)...)
}
func SetIndex(a, b reflect.Value, i int) reflect.Value {
	a.Index(i).Set(b)
	return a
}
func Insert(a, x interface{}, i int) interface{} {
	/*switch a.(type) {
	case bool:
		return interface{}(InsertBool(a.([]bool), x.(bool), i))

	case int:
		return interface{}(InsertInt(a.([]int), x.(int), i))

	case int8:
		return interface{}(InsertInt8(a.([]int8), x.(int8), i))

	case int16:
		return interface{}(InsertInt16(a.([]int16), x.(int16), i))

	case int32:
		return interface{}(InsertInt32(a.([]int32), x.(int32), i))

	case int64:
		return interface{}(InsertInt64(a.([]int64), x.(int64), i))

	case uint:
		return interface{}(InsertUint(a.([]uint), x.(uint), i))

	case uint8:
		return interface{}(InsertUint8(a.([]uint8), x.(uint8), i))

	case uint16:
		return interface{}(InsertUint16(a.([]uint16), x.(uint16), i))

	case uint32:
		return interface{}(InsertUint32(a.([]uint32), x.(uint32), i))

	case uint64:
		return interface{}(InsertUint64(a.([]uint64), x.(uint64), i))

	case uintptr:
		return interface{}(InsertUintptr(a.([]uintptr), x.(uintptr), i))

	case float32:
		return interface{}(InsertFloat32(a.([]float32), x.(float32), i))

	case float64:
		return interface{}(InsertFloat64(a.([]float64), x.(float64), i))

	case complex64:
		return interface{}(InsertComplex64(a.([]complex64), x.(complex64), i))

	case complex128:
		return interface{}(InsertComplex128(a.([]complex128), x.(complex128), i))

	case string:
		return interface{}(InsertStr(a.([]string), x.(string), i))
	}*/
	aval := reflect.ValueOf(a)
	return reflect.AppendSlice(aval.Slice(0, i), reflect.AppendSlice(SetIndex(reflect.MakeSlice(reflect.TypeOf(a), 1, 1), reflect.ValueOf(x), 0), aval.Slice(i, aval.Len()))).Interface()
}
func InsertVectorBool(a, b []bool, i int) []bool {
	return append(a[:i], append(b, a[i:]...)...)
}

func InsertVectorInt(a, b []int, i int) []int {
	return append(a[:i], append(b, a[i:]...)...)
}

func InsertVectorInt8(a, b []int8, i int) []int8 {
	return append(a[:i], append(b, a[i:]...)...)
}

func InsertVectorInt16(a, b []int16, i int) []int16 {
	return append(a[:i], append(b, a[i:]...)...)
}

func InsertVectorInt32(a, b []int32, i int) []int32 {
	return append(a[:i], append(b, a[i:]...)...)
}

func InsertVectorInt64(a, b []int64, i int) []int64 {
	return append(a[:i], append(b, a[i:]...)...)
}

func InsertVectorUint(a, b []uint, i int) []uint {
	return append(a[:i], append(b, a[i:]...)...)
}

func InsertVectorUint8(a, b []uint8, i int) []uint8 {
	return append(a[:i], append(b, a[i:]...)...)
}

func InsertVectorUint16(a, b []uint16, i int) []uint16 {
	return append(a[:i], append(b, a[i:]...)...)
}

func InsertVectorUint32(a, b []uint32, i int) []uint32 {
	return append(a[:i], append(b, a[i:]...)...)
}

func InsertVectorUint64(a, b []uint64, i int) []uint64 {
	return append(a[:i], append(b, a[i:]...)...)
}

func InsertVectorUintptr(a, b []uintptr, i int) []uintptr {
	return append(a[:i], append(b, a[i:]...)...)
}

func InsertVectorFloat32(a, b []float32, i int) []float32 {
	return append(a[:i], append(b, a[i:]...)...)
}

func InsertVectorFloat64(a, b []float64, i int) []float64 {
	return append(a[:i], append(b, a[i:]...)...)
}

func InsertVectorComplex64(a, b []complex64, i int) []complex64 {
	return append(a[:i], append(b, a[i:]...)...)
}

func InsertVectorComplex128(a, b []complex128, i int) []complex128 {
	return append(a[:i], append(b, a[i:]...)...)
}

func InsertVectorStr(a, b []string, i int) []string {
	return append(a[:i], append(b, a[i:]...)...)
}
func InsertVector(a, b interface{}, i int) interface{} {
	/*switch a.(type) {
	case bool:
		return interface{}(InsertVectorBool(a.([]bool), b.([]bool), i))

	case int:
		return interface{}(InsertVectorInt(a.([]int), b.([]int), i))

	case int8:
		return interface{}(InsertVectorInt8(a.([]int8), b.([]int8), i))

	case int16:
		return interface{}(InsertVectorInt16(a.([]int16), b.([]int16), i))

	case int32:
		return interface{}(InsertVectorInt32(a.([]int32), b.([]int32), i))

	case int64:
		return interface{}(InsertVectorInt64(a.([]int64), b.([]int64), i))

	case uint:
		return interface{}(InsertVectorUint(a.([]uint), b.([]uint), i))

	case uint8:
		return interface{}(InsertVectorUint8(a.([]uint8), b.([]uint8), i))

	case uint16:
		return interface{}(InsertVectorUint16(a.([]uint16), b.([]uint16), i))

	case uint32:
		return interface{}(InsertVectorUint32(a.([]uint32), b.([]uint32), i))

	case uint64:
		return interface{}(InsertVectorUint64(a.([]uint64), b.([]uint64), i))

	case uintptr:
		return interface{}(InsertVectorUintptr(a.([]uintptr), b.([]uintptr), i))

	case float32:
		return interface{}(InsertVectorFloat32(a.([]float32), b.([]float32), i))

	case float64:
		return interface{}(InsertVectorFloat64(a.([]float64), b.([]float64), i))

	case complex64:
		return interface{}(InsertVectorComplex64(a.([]complex64), b.([]complex64), i))

	case complex128:
		return interface{}(InsertVectorComplex128(a.([]complex128), b.([]complex128), i))

	case string:
		return interface{}(InsertVectorStr(a.([]string), b.([]string), i))
	}*/
	aval := reflect.ValueOf(a)
	return reflect.AppendSlice(aval.Slice(0, i), reflect.AppendSlice(reflect.ValueOf(b), aval.Slice(i, aval.Len()))).Interface()
}

func PopBool(a []bool) (bool, []bool) {
	return a[len(a)-1], a[:len(a)-1]
}

func PopInt(a []int) (int, []int) {
	return a[len(a)-1], a[:len(a)-1]
}

func PopInt8(a []int8) (int8, []int8) {
	return a[len(a)-1], a[:len(a)-1]
}

func PopInt16(a []int16) (int16, []int16) {
	return a[len(a)-1], a[:len(a)-1]
}

func PopInt32(a []int32) (int32, []int32) {
	return a[len(a)-1], a[:len(a)-1]
}

func PopInt64(a []int64) (int64, []int64) {
	return a[len(a)-1], a[:len(a)-1]
}

func PopUint(a []uint) (uint, []uint) {
	return a[len(a)-1], a[:len(a)-1]
}

func PopUint8(a []uint8) (uint8, []uint8) {
	return a[len(a)-1], a[:len(a)-1]
}

func PopUint16(a []uint16) (uint16, []uint16) {
	return a[len(a)-1], a[:len(a)-1]
}

func PopUint32(a []uint32) (uint32, []uint32) {
	return a[len(a)-1], a[:len(a)-1]
}

func PopUint64(a []uint64) (uint64, []uint64) {
	return a[len(a)-1], a[:len(a)-1]
}

func PopUintptr(a []uintptr) (uintptr, []uintptr) {
	return a[len(a)-1], a[:len(a)-1]
}

func PopFloat32(a []float32) (float32, []float32) {
	return a[len(a)-1], a[:len(a)-1]
}

func PopFloat64(a []float64) (float64, []float64) {
	return a[len(a)-1], a[:len(a)-1]
}

func PopComplex64(a []complex64) (complex64, []complex64) {
	return a[len(a)-1], a[:len(a)-1]
}

func PopComplex128(a []complex128) (complex128, []complex128) {
	return a[len(a)-1], a[:len(a)-1]
}

func PopStr(a []string) (string, []string) {
	return a[len(a)-1], a[:len(a)-1]
}
func Pop(a interface{}) (interface{}, interface{}) {
	/*switch a.(type) {
	case bool:
		a, b := PopBool(a.([]bool))
		return interface{}(a), interface{}(b)

	case int:
		a, b := PopInt(a.([]int))
		return interface{}(a), interface{}(b)

	case int8:
		a, b := PopInt8(a.([]int8))
		return interface{}(a), interface{}(b)

	case int16:
		a, b := PopInt16(a.([]int16))
		return interface{}(a), interface{}(b)

	case int32:
		a, b := PopInt32(a.([]int32))
		return interface{}(a), interface{}(b)

	case int64:
		a, b := PopInt64(a.([]int64))
		return interface{}(a), interface{}(b)

	case uint:
		a, b := PopUint(a.([]uint))
		return interface{}(a), interface{}(b)

	case uint8:
		a, b := PopUint8(a.([]uint8))
		return interface{}(a), interface{}(b)

	case uint16:
		a, b := PopUint16(a.([]uint16))
		return interface{}(a), interface{}(b)

	case uint32:
		a, b := PopUint32(a.([]uint32))
		return interface{}(a), interface{}(b)

	case uint64:
		a, b := PopUint64(a.([]uint64))
		return interface{}(a), interface{}(b)

	case uintptr:
		a, b := PopUintptr(a.([]uintptr))
		return interface{}(a), interface{}(b)

	case float32:
		a, b := PopFloat32(a.([]float32))
		return interface{}(a), interface{}(b)

	case float64:
		a, b := PopFloat64(a.([]float64))
		return interface{}(a), interface{}(b)

	case complex64:
		a, b := PopComplex64(a.([]complex64))
		return interface{}(a), interface{}(b)

	case complex128:
		a, b := PopComplex128(a.([]complex128))
		return interface{}(a), interface{}(b)

	case string:
		a, b := PopStr(a.([]string))
		return interface{}(a), interface{}(b)
	}*/
	aval := reflect.ValueOf(a)
	return aval.Index(aval.Len() - 1).Interface(), aval.Slice(0, aval.Len()-1).Interface()
}
func PopBackBool(a []bool) (bool, []bool) {
	return a[0], a[1:]
}

func PopBackInt(a []int) (int, []int) {
	return a[0], a[1:]
}

func PopBackInt8(a []int8) (int8, []int8) {
	return a[0], a[1:]
}

func PopBackInt16(a []int16) (int16, []int16) {
	return a[0], a[1:]
}

func PopBackInt32(a []int32) (int32, []int32) {
	return a[0], a[1:]
}

func PopBackInt64(a []int64) (int64, []int64) {
	return a[0], a[1:]
}

func PopBackUint(a []uint) (uint, []uint) {
	return a[0], a[1:]
}

func PopBackUint8(a []uint8) (uint8, []uint8) {
	return a[0], a[1:]
}

func PopBackUint16(a []uint16) (uint16, []uint16) {
	return a[0], a[1:]
}

func PopBackUint32(a []uint32) (uint32, []uint32) {
	return a[0], a[1:]
}

func PopBackUint64(a []uint64) (uint64, []uint64) {
	return a[0], a[1:]
}

func PopBackUintptr(a []uintptr) (uintptr, []uintptr) {
	return a[0], a[1:]
}

func PopBackFloat32(a []float32) (float32, []float32) {
	return a[0], a[1:]
}

func PopBackFloat64(a []float64) (float64, []float64) {
	return a[0], a[1:]
}

func PopBackComplex64(a []complex64) (complex64, []complex64) {
	return a[0], a[1:]
}

func PopBackComplex128(a []complex128) (complex128, []complex128) {
	return a[0], a[1:]
}

func PopBackStr(a []string) (string, []string) {
	return a[0], a[1:]
}
func PopBack(a interface{}) (interface{}, interface{}) {
	/*switch a.(type) {
	case bool:
		a, b := PopBackBool(a.([]bool))
		return interface{}(a), interface{}(b)

	case int:
		a, b := PopBackInt(a.([]int))
		return interface{}(a), interface{}(b)

	case int8:
		a, b := PopBackInt8(a.([]int8))
		return interface{}(a), interface{}(b)

	case int16:
		a, b := PopBackInt16(a.([]int16))
		return interface{}(a), interface{}(b)

	case int32:
		a, b := PopBackInt32(a.([]int32))
		return interface{}(a), interface{}(b)

	case int64:
		a, b := PopBackInt64(a.([]int64))
		return interface{}(a), interface{}(b)

	case uint:
		a, b := PopBackUint(a.([]uint))
		return interface{}(a), interface{}(b)

	case uint8:
		a, b := PopBackUint8(a.([]uint8))
		return interface{}(a), interface{}(b)

	case uint16:
		a, b := PopBackUint16(a.([]uint16))
		return interface{}(a), interface{}(b)

	case uint32:
		a, b := PopBackUint32(a.([]uint32))
		return interface{}(a), interface{}(b)

	case uint64:
		a, b := PopBackUint64(a.([]uint64))
		return interface{}(a), interface{}(b)

	case uintptr:
		a, b := PopBackUintptr(a.([]uintptr))
		return interface{}(a), interface{}(b)

	case float32:
		a, b := PopBackFloat32(a.([]float32))
		return interface{}(a), interface{}(b)

	case float64:
		a, b := PopBackFloat64(a.([]float64))
		return interface{}(a), interface{}(b)

	case complex64:
		a, b := PopBackComplex64(a.([]complex64))
		return interface{}(a), interface{}(b)

	case complex128:
		a, b := PopBackComplex128(a.([]complex128))
		return interface{}(a), interface{}(b)

	case string:
		a, b := PopBackStr(a.([]string))
		return interface{}(a), interface{}(b)
	}*/
	aval := reflect.ValueOf(a)
	return aval.Index(0).Interface(), aval.Slice(1, aval.Len()).Interface()
}
func PushBool(a []bool, x bool) []bool {
	return append(a, x)
}

func PushInt(a []int, x int) []int {
	return append(a, x)
}

func PushInt8(a []int8, x int8) []int8 {
	return append(a, x)
}

func PushInt16(a []int16, x int16) []int16 {
	return append(a, x)
}

func PushInt32(a []int32, x int32) []int32 {
	return append(a, x)
}

func PushInt64(a []int64, x int64) []int64 {
	return append(a, x)
}

func PushUint(a []uint, x uint) []uint {
	return append(a, x)
}

func PushUint8(a []uint8, x uint8) []uint8 {
	return append(a, x)
}

func PushUint16(a []uint16, x uint16) []uint16 {
	return append(a, x)
}

func PushUint32(a []uint32, x uint32) []uint32 {
	return append(a, x)
}

func PushUint64(a []uint64, x uint64) []uint64 {
	return append(a, x)
}

func PushUintptr(a []uintptr, x uintptr) []uintptr {
	return append(a, x)
}

func PushFloat32(a []float32, x float32) []float32 {
	return append(a, x)
}

func PushFloat64(a []float64, x float64) []float64 {
	return append(a, x)
}

func PushComplex64(a []complex64, x complex64) []complex64 {
	return append(a, x)
}

func PushComplex128(a []complex128, x complex128) []complex128 {
	return append(a, x)
}

func PushStr(a []string, x string) []string {
	return append(a, x)
}
func Push(a, x interface{}) interface{} {
	/*switch a.(type) {
	case bool:
		return interface{}(PushBool(a.([]bool), x.(bool)))

	case int:
		return interface{}(PushInt(a.([]int), x.(int)))

	case int8:
		return interface{}(PushInt8(a.([]int8), x.(int8)))

	case int16:
		return interface{}(PushInt16(a.([]int16), x.(int16)))

	case int32:
		return interface{}(PushInt32(a.([]int32), x.(int32)))

	case int64:
		return interface{}(PushInt64(a.([]int64), x.(int64)))

	case uint:
		return interface{}(PushUint(a.([]uint), x.(uint)))

	case uint8:
		return interface{}(PushUint8(a.([]uint8), x.(uint8)))

	case uint16:
		return interface{}(PushUint16(a.([]uint16), x.(uint16)))

	case uint32:
		return interface{}(PushUint32(a.([]uint32), x.(uint32)))

	case uint64:
		return interface{}(PushUint64(a.([]uint64), x.(uint64)))

	case uintptr:
		return interface{}(PushUintptr(a.([]uintptr), x.(uintptr)))

	case float32:
		return interface{}(PushFloat32(a.([]float32), x.(float32)))

	case float64:
		return interface{}(PushFloat64(a.([]float64), x.(float64)))

	case complex64:
		return interface{}(PushComplex64(a.([]complex64), x.(complex64)))

	case complex128:
		return interface{}(PushComplex128(a.([]complex128), x.(complex128)))

	case string:
		return interface{}(PushStr(a.([]string), x.(string)))
	}*/
	return reflect.Append(reflect.ValueOf(a), reflect.ValueOf(x)).Interface()
}
func PushBackBool(a []bool, x bool) []bool {
	return append([]bool{x}, a...)
}

func PushBackInt(a []int, x int) []int {
	return append([]int{x}, a...)
}

func PushBackInt8(a []int8, x int8) []int8 {
	return append([]int8{x}, a...)
}

func PushBackInt16(a []int16, x int16) []int16 {
	return append([]int16{x}, a...)
}

func PushBackInt32(a []int32, x int32) []int32 {
	return append([]int32{x}, a...)
}

func PushBackInt64(a []int64, x int64) []int64 {
	return append([]int64{x}, a...)
}

func PushBackUint(a []uint, x uint) []uint {
	return append([]uint{x}, a...)
}

func PushBackUint8(a []uint8, x uint8) []uint8 {
	return append([]uint8{x}, a...)
}

func PushBackUint16(a []uint16, x uint16) []uint16 {
	return append([]uint16{x}, a...)
}

func PushBackUint32(a []uint32, x uint32) []uint32 {
	return append([]uint32{x}, a...)
}

func PushBackUint64(a []uint64, x uint64) []uint64 {
	return append([]uint64{x}, a...)
}

func PushBackUintptr(a []uintptr, x uintptr) []uintptr {
	return append([]uintptr{x}, a...)
}

func PushBackFloat32(a []float32, x float32) []float32 {
	return append([]float32{x}, a...)
}

func PushBackFloat64(a []float64, x float64) []float64 {
	return append([]float64{x}, a...)
}

func PushBackComplex64(a []complex64, x complex64) []complex64 {
	return append([]complex64{x}, a...)
}

func PushBackComplex128(a []complex128, x complex128) []complex128 {
	return append([]complex128{x}, a...)
}

func PushBackStr(a []string, x string) []string {
	return append([]string{x}, a...)
}
func PushBack(a, x interface{}) interface{} {
	/*switch a.(type) {
	case bool:
		return interface{}(PushBackBool(a.([]bool), x.(bool)))

	case int:
		return interface{}(PushBackInt(a.([]int), x.(int)))

	case int8:
		return interface{}(PushBackInt8(a.([]int8), x.(int8)))

	case int16:
		return interface{}(PushBackInt16(a.([]int16), x.(int16)))

	case int32:
		return interface{}(PushBackInt32(a.([]int32), x.(int32)))

	case int64:
		return interface{}(PushBackInt64(a.([]int64), x.(int64)))

	case uint:
		return interface{}(PushBackUint(a.([]uint), x.(uint)))

	case uint8:
		return interface{}(PushBackUint8(a.([]uint8), x.(uint8)))

	case uint16:
		return interface{}(PushBackUint16(a.([]uint16), x.(uint16)))

	case uint32:
		return interface{}(PushBackUint32(a.([]uint32), x.(uint32)))

	case uint64:
		return interface{}(PushBackUint64(a.([]uint64), x.(uint64)))

	case uintptr:
		return interface{}(PushBackUintptr(a.([]uintptr), x.(uintptr)))

	case float32:
		return interface{}(PushBackFloat32(a.([]float32), x.(float32)))

	case float64:
		return interface{}(PushBackFloat64(a.([]float64), x.(float64)))

	case complex64:
		return interface{}(PushBackComplex64(a.([]complex64), x.(complex64)))

	case complex128:
		return interface{}(PushBackComplex128(a.([]complex128), x.(complex128)))

	case string:
		return interface{}(PushBackStr(a.([]string), x.(string)))
	}*/
	aval := reflect.ValueOf(a)
	return reflect.AppendSlice(SetIndex(reflect.MakeSlice(reflect.TypeOf(a), 1, 1+aval.Len()), reflect.ValueOf(x), 0), aval).Interface()
}

var ShiftBool = PopBackBool
var ShiftInt = PopBackInt
var ShiftInt8 = PopBackInt8
var ShiftInt16 = PopBackInt16
var ShiftInt32 = PopBackInt32
var ShiftInt64 = PopBackInt64
var ShiftUint = PopBackUint
var ShiftUint8 = PopBackUint8
var ShiftUint16 = PopBackUint16
var ShiftUint32 = PopBackUint32
var ShiftUint64 = PopBackUint64
var ShiftUintptr = PopBackUintptr
var ShiftFloat32 = PopBackFloat32
var ShiftFloat64 = PopBackFloat64
var ShiftComplex64 = PopBackComplex64
var ShiftComplex128 = PopBackComplex128
var ShiftStr = PopBackStr
var Shift = PopBack
var UnshiftBool = PushBackBool
var UnshiftInt = PushBackInt
var UnshiftInt8 = PushBackInt8
var UnshiftInt16 = PushBackInt16
var UnshiftInt32 = PushBackInt32
var UnshiftInt64 = PushBackInt64
var UnshiftUint = PushBackUint
var UnshiftUint8 = PushBackUint8
var UnshiftUint16 = PushBackUint16
var UnshiftUint32 = PushBackUint32
var UnshiftUint64 = PushBackUint64
var UnshiftUintptr = PushBackUintptr
var UnshiftFloat32 = PushBackFloat32
var UnshiftFloat64 = PushBackFloat64
var UnshiftComplex64 = PushBackComplex64
var UnshiftComplex128 = PushBackComplex128
var UnshiftStr = PushBackStr
var Unshift = PushBack

/*func Shift_T_(a []T, x T) (T, []T) {
	return a[0], a[1:]
}
func Shift(a, x interface{}) (interface{}, interface{}) {
	return reflect.ValueOf(
}
func UnshiftBool(a []bool, x bool) []bool {
	return append([]bool{x}, a...)
}

func UnshiftInt(a []int, x int) []int {
	return append([]int{x}, a...)
}

func UnshiftInt8(a []int8, x int8) []int8 {
	return append([]int8{x}, a...)
}

func UnshiftInt16(a []int16, x int16) []int16 {
	return append([]int16{x}, a...)
}

func UnshiftInt32(a []int32, x int32) []int32 {
	return append([]int32{x}, a...)
}

func UnshiftInt64(a []int64, x int64) []int64 {
	return append([]int64{x}, a...)
}

func UnshiftUint(a []uint, x uint) []uint {
	return append([]uint{x}, a...)
}

func UnshiftUint8(a []uint8, x uint8) []uint8 {
	return append([]uint8{x}, a...)
}

func UnshiftUint16(a []uint16, x uint16) []uint16 {
	return append([]uint16{x}, a...)
}

func UnshiftUint32(a []uint32, x uint32) []uint32 {
	return append([]uint32{x}, a...)
}

func UnshiftUint64(a []uint64, x uint64) []uint64 {
	return append([]uint64{x}, a...)
}

func UnshiftUintptr(a []uintptr, x uintptr) []uintptr {
	return append([]uintptr{x}, a...)
}

func UnshiftFloat32(a []float32, x float32) []float32 {
	return append([]float32{x}, a...)
}

func UnshiftFloat64(a []float64, x float64) []float64 {
	return append([]float64{x}, a...)
}

func UnshiftComplex64(a []complex64, x complex64) []complex64 {
	return append([]complex64{x}, a...)
}

func UnshiftComplex128(a []complex128, x complex128) []complex128 {
	return append([]complex128{x}, a...)
}

func UnshiftStr(a []string, x string) []string {
	return append([]string{x}, a...)
}
func Unshift(a, x interface{}) interface{} {
	aval := reflect.ValueOf(a)
	return reflect.AppendSlice(SetIndex(reflect.MakeSlice(reflect.TypeOf(a), 1, 1+aval.Len()), reflect.ValueOf(x), 0), aval)
}*/
func FilterWithoutAllocBool(a []bool, f func(bool) bool) []bool {
	var b []bool
	for _, x := range a {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}

func FilterWithoutAllocInt(a []int, f func(int) bool) []int {
	var b []int
	for _, x := range a {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}

func FilterWithoutAllocInt8(a []int8, f func(int8) bool) []int8 {
	var b []int8
	for _, x := range a {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}

func FilterWithoutAllocInt16(a []int16, f func(int16) bool) []int16 {
	var b []int16
	for _, x := range a {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}

func FilterWithoutAllocInt32(a []int32, f func(int32) bool) []int32 {
	var b []int32
	for _, x := range a {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}

func FilterWithoutAllocInt64(a []int64, f func(int64) bool) []int64 {
	var b []int64
	for _, x := range a {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}

func FilterWithoutAllocUint(a []uint, f func(uint) bool) []uint {
	var b []uint
	for _, x := range a {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}

func FilterWithoutAllocUint8(a []uint8, f func(uint8) bool) []uint8 {
	var b []uint8
	for _, x := range a {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}

func FilterWithoutAllocUint16(a []uint16, f func(uint16) bool) []uint16 {
	var b []uint16
	for _, x := range a {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}

func FilterWithoutAllocUint32(a []uint32, f func(uint32) bool) []uint32 {
	var b []uint32
	for _, x := range a {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}

func FilterWithoutAllocUint64(a []uint64, f func(uint64) bool) []uint64 {
	var b []uint64
	for _, x := range a {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}

func FilterWithoutAllocUintptr(a []uintptr, f func(uintptr) bool) []uintptr {
	var b []uintptr
	for _, x := range a {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}

func FilterWithoutAllocFloat32(a []float32, f func(float32) bool) []float32 {
	var b []float32
	for _, x := range a {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}

func FilterWithoutAllocFloat64(a []float64, f func(float64) bool) []float64 {
	var b []float64
	for _, x := range a {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}

func FilterWithoutAllocComplex64(a []complex64, f func(complex64) bool) []complex64 {
	var b []complex64
	for _, x := range a {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}

func FilterWithoutAllocComplex128(a []complex128, f func(complex128) bool) []complex128 {
	var b []complex128
	for _, x := range a {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}

func FilterWithoutAllocStr(a []string, f func(string) bool) []string {
	var b []string
	for _, x := range a {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}
func FilterWithoutAlloc(a interface{}, f interface{}) interface{} {
	/*switch a.(type) {
	case bool:
		return interface{}(FilterWithoutAllocBool(a.([]bool), f.(func(v bool) bool)))

	case int:
		return interface{}(FilterWithoutAllocInt(a.([]int), f.(func(v int) bool)))

	case int8:
		return interface{}(FilterWithoutAllocInt8(a.([]int8), f.(func(v int8) bool)))

	case int16:
		return interface{}(FilterWithoutAllocInt16(a.([]int16), f.(func(v int16) bool)))

	case int32:
		return interface{}(FilterWithoutAllocInt32(a.([]int32), f.(func(v int32) bool)))

	case int64:
		return interface{}(FilterWithoutAllocInt64(a.([]int64), f.(func(v int64) bool)))

	case uint:
		return interface{}(FilterWithoutAllocUint(a.([]uint), f.(func(v uint) bool)))

	case uint8:
		return interface{}(FilterWithoutAllocUint8(a.([]uint8), f.(func(v uint8) bool)))

	case uint16:
		return interface{}(FilterWithoutAllocUint16(a.([]uint16), f.(func(v uint16) bool)))

	case uint32:
		return interface{}(FilterWithoutAllocUint32(a.([]uint32), f.(func(v uint32) bool)))

	case uint64:
		return interface{}(FilterWithoutAllocUint64(a.([]uint64), f.(func(v uint64) bool)))

	case uintptr:
		return interface{}(FilterWithoutAllocUintptr(a.([]uintptr), f.(func(v uintptr) bool)))

	case float32:
		return interface{}(FilterWithoutAllocFloat32(a.([]float32), f.(func(v float32) bool)))

	case float64:
		return interface{}(FilterWithoutAllocFloat64(a.([]float64), f.(func(v float64) bool)))

	case complex64:
		return interface{}(FilterWithoutAllocComplex64(a.([]complex64), f.(func(v complex64) bool)))

	case complex128:
		return interface{}(FilterWithoutAllocComplex128(a.([]complex128), f.(func(v complex128) bool)))

	case string:
		return interface{}(FilterWithoutAllocStr(a.([]string), f.(func(v string) bool)))
	}*/
	bval := reflect.MakeSlice(reflect.TypeOf(a), 0, 0)
	aval := reflect.ValueOf(a)
	for i := 0; i < aval.Len(); i++ {
		if reflect.ValueOf(f).Call([]reflect.Value{aval.Index(i)})[0].Bool() {
			bval = reflect.Append(bval, aval.Index(i))
		}
	}
	return bval.Interface()
}
func ReverseBool(a *[]bool) []bool {
	/*for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a*/
	for left, right := 0, len(*a)-1; left < right; left, right = left+1, right-1 {
		(*a)[left], (*a)[right] = (*a)[right], (*a)[left]
	}
	return *a
}

func ReverseInt(a *[]int) []int {
	/*for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a*/
	for left, right := 0, len(*a)-1; left < right; left, right = left+1, right-1 {
		(*a)[left], (*a)[right] = (*a)[right], (*a)[left]
	}
	return *a
}

func ReverseInt8(a *[]int8) []int8 {
	/*for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a*/
	for left, right := 0, len(*a)-1; left < right; left, right = left+1, right-1 {
		(*a)[left], (*a)[right] = (*a)[right], (*a)[left]
	}
	return *a
}

func ReverseInt16(a *[]int16) []int16 {
	/*for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a*/
	for left, right := 0, len(*a)-1; left < right; left, right = left+1, right-1 {
		(*a)[left], (*a)[right] = (*a)[right], (*a)[left]
	}
	return *a
}

func ReverseInt32(a *[]int32) []int32 {
	/*for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a*/
	for left, right := 0, len(*a)-1; left < right; left, right = left+1, right-1 {
		(*a)[left], (*a)[right] = (*a)[right], (*a)[left]
	}
	return *a
}

func ReverseInt64(a *[]int64) []int64 {
	/*for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a*/
	for left, right := 0, len(*a)-1; left < right; left, right = left+1, right-1 {
		(*a)[left], (*a)[right] = (*a)[right], (*a)[left]
	}
	return *a
}

func ReverseUint(a *[]uint) []uint {
	/*for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a*/
	for left, right := 0, len(*a)-1; left < right; left, right = left+1, right-1 {
		(*a)[left], (*a)[right] = (*a)[right], (*a)[left]
	}
	return *a
}

func ReverseUint8(a *[]uint8) []uint8 {
	/*for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a*/
	for left, right := 0, len(*a)-1; left < right; left, right = left+1, right-1 {
		(*a)[left], (*a)[right] = (*a)[right], (*a)[left]
	}
	return *a
}

func ReverseUint16(a *[]uint16) []uint16 {
	/*for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a*/
	for left, right := 0, len(*a)-1; left < right; left, right = left+1, right-1 {
		(*a)[left], (*a)[right] = (*a)[right], (*a)[left]
	}
	return *a
}

func ReverseUint32(a *[]uint32) []uint32 {
	/*for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a*/
	for left, right := 0, len(*a)-1; left < right; left, right = left+1, right-1 {
		(*a)[left], (*a)[right] = (*a)[right], (*a)[left]
	}
	return *a
}

func ReverseUint64(a *[]uint64) []uint64 {
	/*for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a*/
	for left, right := 0, len(*a)-1; left < right; left, right = left+1, right-1 {
		(*a)[left], (*a)[right] = (*a)[right], (*a)[left]
	}
	return *a
}

func ReverseUintptr(a *[]uintptr) []uintptr {
	/*for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a*/
	for left, right := 0, len(*a)-1; left < right; left, right = left+1, right-1 {
		(*a)[left], (*a)[right] = (*a)[right], (*a)[left]
	}
	return *a
}

func ReverseFloat32(a *[]float32) []float32 {
	/*for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a*/
	for left, right := 0, len(*a)-1; left < right; left, right = left+1, right-1 {
		(*a)[left], (*a)[right] = (*a)[right], (*a)[left]
	}
	return *a
}

func ReverseFloat64(a *[]float64) []float64 {
	/*for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a*/
	for left, right := 0, len(*a)-1; left < right; left, right = left+1, right-1 {
		(*a)[left], (*a)[right] = (*a)[right], (*a)[left]
	}
	return *a
}

func ReverseComplex64(a *[]complex64) []complex64 {
	/*for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a*/
	for left, right := 0, len(*a)-1; left < right; left, right = left+1, right-1 {
		(*a)[left], (*a)[right] = (*a)[right], (*a)[left]
	}
	return *a
}

func ReverseComplex128(a *[]complex128) []complex128 {
	/*for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a*/
	for left, right := 0, len(*a)-1; left < right; left, right = left+1, right-1 {
		(*a)[left], (*a)[right] = (*a)[right], (*a)[left]
	}
	return *a
}

func ReverseStr(a *[]string) []string {
	/*for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a*/
	for left, right := 0, len(*a)-1; left < right; left, right = left+1, right-1 {
		(*a)[left], (*a)[right] = (*a)[right], (*a)[left]
	}
	return *a
}
func Reverse(a interface{}) interface{} {
	/*switch a.(type) {
	case bool:
		return interface{}(ReverseBool(a.(*[]bool)))

	case int:
		return interface{}(ReverseInt(a.(*[]int)))

	case int8:
		return interface{}(ReverseInt8(a.(*[]int8)))

	case int16:
		return interface{}(ReverseInt16(a.(*[]int16)))

	case int32:
		return interface{}(ReverseInt32(a.(*[]int32)))

	case int64:
		return interface{}(ReverseInt64(a.(*[]int64)))

	case uint:
		return interface{}(ReverseUint(a.(*[]uint)))

	case uint8:
		return interface{}(ReverseUint8(a.(*[]uint8)))

	case uint16:
		return interface{}(ReverseUint16(a.(*[]uint16)))

	case uint32:
		return interface{}(ReverseUint32(a.(*[]uint32)))

	case uint64:
		return interface{}(ReverseUint64(a.(*[]uint64)))

	case uintptr:
		return interface{}(ReverseUintptr(a.(*[]uintptr)))

	case float32:
		return interface{}(ReverseFloat32(a.(*[]float32)))

	case float64:
		return interface{}(ReverseFloat64(a.(*[]float64)))

	case complex64:
		return interface{}(ReverseComplex64(a.(*[]complex64)))

	case complex128:
		return interface{}(ReverseComplex128(a.(*[]complex128)))

	case string:
		return interface{}(ReverseStr(a.(*[]string)))
	}*/
	aval := reflect.Indirect(reflect.ValueOf(a))
	swap := reflect.Swapper(aval.Interface())
	/*for i := aval.Len()/2 - 1; i >= 0; i-- {
		swap(i, aval.Len()-1-i)
	}*/
	for left, right := 0, aval.Len()-1; left < right; left, right = left+1, right-1 {
		swap(left, right)
	}
	return aval.Interface()
}
