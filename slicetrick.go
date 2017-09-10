// Code generated with github.com/chidea/go-type-template DO NOT EDIT.
		

package sliceTrick

import (
	//"fmt"
	"reflect"
)

func AppendStr(a, b []string) []string {
	return append(a, b...)
}

func Append(a, b interface{}) interface{} {
	switch a.(type) { // switch clause is usually a litter faster than reflect lib
	case string:
		return interface{}(AppendStr(a.([]string), b.([]string)))
	} //fall back to reflect lib when it's not a precompiled type
	return reflect.AppendSlice(reflect.ValueOf(a), reflect.ValueOf(b)).Interface()
}
func CopyStr(a []string) []string {
	/*b := make([]string, len(a))
	copy(b, a)
	return b*/
	// or
	return append([]string(nil), a...)
}
func Copy(a interface{}) interface{} {
	switch a.(type) {
	case string:
		return interface{}(CopyStr(a.([]string)))
	}
	/*aval := reflect.ValueOf(a)
	alen := aval.Len()
	bval := reflect.MakeSlice(reflect.TypeOf(a), alen, alen)
	reflect.Copy(bval, aval)
	return bval.Interface()*/
	// or
	aval := reflect.ValueOf(a)
	return reflect.AppendSlice(reflect.MakeSlice(reflect.TypeOf(a), 0, aval.Len()), aval.Slice(0, aval.Len())).Interface()
}
func CutStr(a []string, i, j int) []string {
	return append(a[:i], a[j:]...)
}
func Cut(a interface{}, i, j int) interface{} {
	switch a.(type) {
	case string:
		return interface{}(CutStr(a.([]string), i, j))
	}
	aval := reflect.ValueOf(a)
	return reflect.AppendSlice(aval.Slice(0, i), aval.Slice(j, aval.Len())).Interface()
}
func DeleteStr(a []string, i int) []string {
	return append(a[:i], a[i+1:]...)
	// or
	//a = a[:i+copy(a[i:], a[i+1:])]
}
func Delete(a interface{}, i int) interface{} {
	switch a.(type) {
	case string:
		return interface{}(DeleteStr(a.([]string), i))
	}
	aval := reflect.ValueOf(a)
	return reflect.AppendSlice(aval.Slice(0, i), aval.Slice(i+1, aval.Len())).Interface()
}

func DeleteWithoutPreservingOrderStr(a []string, i int) []string {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}
func DeleteWithoutPreservingOrder(a interface{}, i int) interface{} { //interface {} {
	switch a.(type) {
	case string:
		return interface{}(DeleteWithoutPreservingOrderStr(a.([]string), i))
	}
	aval := reflect.ValueOf(a)
	l := aval.Len()
	aval.Index(i).Set(aval.Index(l - 1))
	return aval.Slice(0, l-1).Interface()
}

func ExpandStr(a []string, i, j int) []string {
	return append(a[:i], append(make([]string, j), a[i:]...)...)
}
func Expand(a interface{}, i, j int) interface{} {
	switch a.(type) {
	case string:
		return interface{}(ExpandStr(a.([]string), i, j))
	}
	aval := reflect.ValueOf(a)
	return reflect.AppendSlice(aval.Slice(0, i), reflect.AppendSlice(reflect.MakeSlice(reflect.TypeOf(a), j, j), aval.Slice(i, aval.Len()))).Interface()
}
func ExtendStr(a []string, i int) []string {
	return append(a, make([]string, i)...)
}
func Extend(a interface{}, i int) interface{} {
	switch a.(type) {
	case string:
		return interface{}(ExtendStr(a.([]string), i))
	}
	return reflect.AppendSlice(reflect.ValueOf(a), reflect.MakeSlice(reflect.TypeOf(a), i, i)).Interface()
}
func InsertStr(a []string, x string, i int) []string {
	return append(a[:i], append([]string{x}, a[i:]...)...)
}
func SetIndex(a, b reflect.Value, i int) reflect.Value {
	a.Index(i).Set(b)
	return a
}
func Insert(a, x interface{}, i int) interface{} {
	switch a.(type) {
	case string:
		return interface{}(InsertStr(a.([]string), x.(string), i))
	}
	aval := reflect.ValueOf(a)
	return reflect.AppendSlice(aval.Slice(0, i), reflect.AppendSlice(SetIndex(reflect.MakeSlice(reflect.TypeOf(a), 1, 1), reflect.ValueOf(x), 0), aval.Slice(i, aval.Len()))).Interface()
}
func InsertVectorStr(a, b []string, i int) []string {
	return append(a[:i], append(b, a[i:]...)...)
}
func InsertVector(a, b interface{}, i int) interface{} {
	switch a.(type) {
	case string:
		return interface{}(InsertVectorStr(a.([]string), b.([]string), i))
	}
	aval := reflect.ValueOf(a)
	return reflect.AppendSlice(aval.Slice(0, i), reflect.AppendSlice(reflect.ValueOf(b), aval.Slice(i, aval.Len()))).Interface()
}

func PopStr(a []string) (string, []string) {
	return a[len(a)-1], a[:len(a)-1]
}
func Pop(a interface{}) (interface{}, interface{}) {
	switch a.(type) {
	case string:
		a, b := PopStr(a.([]string))
		return interface{}(a), interface{}(b)
	}
	aval := reflect.ValueOf(a)
	return aval.Index(aval.Len() - 1).Interface(), aval.Slice(0, aval.Len()-1).Interface()
}
func PopBackStr(a []string) (string, []string) {
	return a[0], a[1:]
}
func PopBack(a interface{}) (interface{}, interface{}) {
	switch a.(type) {
	case string:
		a, b := PopBackStr(a.([]string))
		return interface{}(a), interface{}(b)
	}
	aval := reflect.ValueOf(a)
	return aval.Index(0).Interface(), aval.Slice(1, aval.Len()).Interface()
}
func PushStr(a []string, x string) []string {
	return append(a, x)
}
func Push(a, x interface{}) interface{} {
	switch a.(type) {
	case string:
		return interface{}(PushStr(a.([]string), x.(string)))
	}
	return reflect.Append(reflect.ValueOf(a), reflect.ValueOf(x)).Interface()
}
func PushBackStr(a []string, x string) []string {
	return append([]string{x}, a...)
}
func PushBack(a, x interface{}) interface{} {
	switch a.(type) {
	case string:
		return interface{}(PushBackStr(a.([]string), x.(string)))
	}
	aval := reflect.ValueOf(a)
	return reflect.AppendSlice(SetIndex(reflect.MakeSlice(reflect.TypeOf(a), 1, 1+aval.Len()), reflect.ValueOf(x), 0), aval).Interface()
}

var ShiftStr = PopBackStr
var Shift = PopBack
var UnshiftStr = PushBackStr
var Unshift = PushBack

/*func Shift_T_(a []T, x T) (T, []T) {
	return a[0], a[1:]
}
func Shift(a, x interface{}) (interface{}, interface{}) {
	return reflect.ValueOf(
}
func UnshiftStr(a []string, x string) []string {
	return append([]string{x}, a...)
}
func Unshift(a, x interface{}) interface{} {
	aval := reflect.ValueOf(a)
	return reflect.AppendSlice(SetIndex(reflect.MakeSlice(reflect.TypeOf(a), 1, 1+aval.Len()), reflect.ValueOf(x), 0), aval)
}*/
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
	switch a.(type) {
	case string:
		return interface{}(FilterWithoutAllocStr(a.([]string), func() func(string) bool {
			return f.(func(v string) bool)
		}()))
	}
	bval := reflect.MakeSlice(reflect.TypeOf(a), 0, 0)
	aval := reflect.ValueOf(a)
	for i := 0; i < aval.Len(); i++ {
		if reflect.ValueOf(f).Call([]reflect.Value{aval.Index(i)})[0].Bool() {
			bval = reflect.Append(bval, aval.Index(i))
		}
	}
	return bval.Interface()
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
	switch a.(type) {
	case string:
		return interface{}(ReverseStr(a.(*[]string)))
	}
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
