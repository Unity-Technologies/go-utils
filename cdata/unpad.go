package cdata

import (
	"reflect"
)

// Unpad returns buf with all field padding according to the memory layout of the struct data removed.
// If data is not a struct it returns buf unchanged.
func Unpad(buf []byte, data interface{}) []byte {
	return unpad(buf, data, 0)
}

func unpad(buf []byte, data interface{}, offset uintptr) []byte {
	v := reflect.Indirect(reflect.ValueOf(data))
	t := v.Type()
	if t.Kind() != reflect.Struct {
		return buf
	}

	tmp := make([]byte, 0, t.Size())
	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		if f.Type.Kind() == reflect.Struct {
			tmp = append(tmp, unpad(buf, v.Field(i).Interface(), offset+f.Offset)...)
		} else {
			tmp = append(tmp, buf[offset+f.Offset:offset+f.Offset+f.Type.Size()]...)
		}
	}

	return tmp
}
