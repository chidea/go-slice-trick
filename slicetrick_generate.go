// +build generate

package sliceTrick

import (
	//"fmt"
	"reflect"
)

func Append_T_(a, b []T) []T {
	return append(a, b...)
}

func Append(a, b interface{}) interface{} {
	/*switch a.(type) {
	case T:
		return interface{}(Append_T_(a.([]T), b.([]T)))
	} //fall back to reflect lib when it's not a precompiled type*/
	return reflect.AppendSlice(reflect.ValueOf(a), reflect.ValueOf(b)).Interface()
}
func Copy_T_(a []T) []T {
	/*b := make([]T, len(a))
	copy(b, a)
	return b*/
	// or
	return append([]T(nil), a...)
}
func Copy(a interface{}) interface{} {
	/*switch a.(type) {
	case T:
		return interface{}(Copy_T_(a.([]T)))
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
func Cut_T_(a []T, i, j int) []T {
	return append(a[:i], a[j:]...)
}
func Cut(a interface{}, i, j int) interface{} {
	/*switch a.(type) {
	case T:
		return interface{}(Cut_T_(a.([]T), i, j))
	}*/
	aval := reflect.ValueOf(a)
	return reflect.AppendSlice(aval.Slice(0, i), aval.Slice(j, aval.Len())).Interface()
}
func Delete_T_(a []T, i int) []T {
	if len(a)-1 == i {
		return a[:i]
	}
	return append(a[:i], a[i+1:]...)
	// or
	//a = a[:i+copy(a[i:], a[i+1:])]
}
func Delete(a interface{}, i int) interface{} {
	/*switch a.(type) {
	case T:
		return interface{}(Delete_T_(a.([]T), i))
	}*/
	aval := reflect.ValueOf(a)
	return reflect.AppendSlice(aval.Slice(0, i), aval.Slice(i+1, aval.Len())).Interface()
}

func DeleteWithoutPreservingOrder_T_(a []T, i int) []T {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}
func DeleteWithoutPreservingOrder(a interface{}, i int) interface{} { //interface {} {
	/*switch a.(type) {
	case T:
		return interface{}(DeleteWithoutPreservingOrder_T_(a.([]T), i))
	}*/
	aval := reflect.ValueOf(a)
	l := aval.Len()
	aval.Index(i).Set(aval.Index(l - 1))
	return aval.Slice(0, l-1).Interface()
}

func Expand_T_(a []T, i, j int) []T {
	return append(a[:i], append(make([]T, j), a[i:]...)...)
}
func Expand(a interface{}, i, j int) interface{} {
	/*switch a.(type) {
	case T:
		return interface{}(Expand_T_(a.([]T), i, j))
	}*/
	aval := reflect.ValueOf(a)
	return reflect.AppendSlice(aval.Slice(0, i), reflect.AppendSlice(reflect.MakeSlice(reflect.TypeOf(a), j, j), aval.Slice(i, aval.Len()))).Interface()
}
func Extend_T_(a []T, i int) []T {
	return append(a, make([]T, i)...)
}
func Extend(a interface{}, i int) interface{} {
	/*switch a.(type) {
	case T:
		return interface{}(Extend_T_(a.([]T), i))
	}*/
	return reflect.AppendSlice(reflect.ValueOf(a), reflect.MakeSlice(reflect.TypeOf(a), i, i)).Interface()
}
func Insert_T_(a []T, x T, i int) []T {
	return append(a[:i], append([]T{x}, a[i:]...)...)
}
func SetIndex(a, b reflect.Value, i int) reflect.Value {
	a.Index(i).Set(b)
	return a
}
func Insert(a, x interface{}, i int) interface{} {
	/*switch a.(type) {
	case T:
		return interface{}(Insert_T_(a.([]T), x.(T), i))
	}*/
	aval := reflect.ValueOf(a)
	return reflect.AppendSlice(aval.Slice(0, i), reflect.AppendSlice(SetIndex(reflect.MakeSlice(reflect.TypeOf(a), 1, 1), reflect.ValueOf(x), 0), aval.Slice(i, aval.Len()))).Interface()
}
func InsertVector_T_(a, b []T, i int) []T {
	return append(a[:i], append(b, a[i:]...)...)
}
func InsertVector(a, b interface{}, i int) interface{} {
	/*switch a.(type) {
	case T:
		return interface{}(InsertVector_T_(a.([]T), b.([]T), i))
	}*/
	aval := reflect.ValueOf(a)
	return reflect.AppendSlice(aval.Slice(0, i), reflect.AppendSlice(reflect.ValueOf(b), aval.Slice(i, aval.Len()))).Interface()
}

func Pop_T_(a []T) (T, []T) {
	return a[len(a)-1], a[:len(a)-1]
}
func Pop(a interface{}) (interface{}, interface{}) {
	/*switch a.(type) {
	case T:
		a, b := Pop_T_(a.([]T))
		return interface{}(a), interface{}(b)
	}*/
	aval := reflect.ValueOf(a)
	return aval.Index(aval.Len() - 1).Interface(), aval.Slice(0, aval.Len()-1).Interface()
}
func PopBack_T_(a []T) (T, []T) {
	return a[0], a[1:]
}
func PopBack(a interface{}) (interface{}, interface{}) {
	/*switch a.(type) {
	case T:
		a, b := PopBack_T_(a.([]T))
		return interface{}(a), interface{}(b)
	}*/
	aval := reflect.ValueOf(a)
	return aval.Index(0).Interface(), aval.Slice(1, aval.Len()).Interface()
}
func Push_T_(a []T, x T) []T {
	return append(a, x)
}
func Push(a, x interface{}) interface{} {
	/*switch a.(type) {
	case T:
		return interface{}(Push_T_(a.([]T), x.(T)))
	}*/
	return reflect.Append(reflect.ValueOf(a), reflect.ValueOf(x)).Interface()
}
func PushBack_T_(a []T, x T) []T {
	return append([]T{x}, a...)
}
func PushBack(a, x interface{}) interface{} {
	/*switch a.(type) {
	case T:
		return interface{}(PushBack_T_(a.([]T), x.(T)))
	}*/
	aval := reflect.ValueOf(a)
	return reflect.AppendSlice(SetIndex(reflect.MakeSlice(reflect.TypeOf(a), 1, 1+aval.Len()), reflect.ValueOf(x), 0), aval).Interface()
}

var Shift_T_ = PopBack_T_
var Shift = PopBack
var Unshift_T_ = PushBack_T_
var Unshift = PushBack

/*func Shift_T_(a []T, x T) (T, []T) {
	return a[0], a[1:]
}
func Shift(a, x interface{}) (interface{}, interface{}) {
	return reflect.ValueOf(
}
func Unshift_T_(a []T, x T) []T {
	return append([]T{x}, a...)
}
func Unshift(a, x interface{}) interface{} {
	aval := reflect.ValueOf(a)
	return reflect.AppendSlice(SetIndex(reflect.MakeSlice(reflect.TypeOf(a), 1, 1+aval.Len()), reflect.ValueOf(x), 0), aval)
}*/
func FilterWithoutAlloc_T_(a []T, f func(T) bool) []T {
	var b []T
	for _, x := range a {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}
func FilterWithoutAlloc(a interface{}, f interface{}) interface{} {
	/*switch a.(type) {
	case T:
		return interface{}(FilterWithoutAlloc_T_(a.([]T), f.(func(v T) bool)))
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
func Reverse_T_(a *[]T) []T {
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
	case T:
		return interface{}(Reverse_T_(a.(*[]T)))
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
