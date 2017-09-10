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

func InStr(slice []string, val string) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
