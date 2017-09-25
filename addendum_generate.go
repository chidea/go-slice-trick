// +build generate

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

func In_T_(slice []T, val T) bool {
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

func Iif_T_(test bool, onTrue T, onFalse T) T {
	if test {
		return onTrue
	}
	return onFalse
}
