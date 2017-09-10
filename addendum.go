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
