package controller

import (
	"reflect"
	"strconv"
	"time"
)

func StringToUint(s string) uint64 {
	val, _ := strconv.ParseUint(s, 10, 64)
	return val
}

func AssignNonNullFields(dest, src interface{}) {
	destValue := reflect.ValueOf(dest).Elem()
	srcValue := reflect.ValueOf(src).Elem()

	for i := 0; i < destValue.NumField(); i++ {
		destField := destValue.Field(i)
		srcField := srcValue.Field(i)

		switch srcField.Kind() {
		case reflect.String:
			if srcField.Len() > 0 {
				destField.SetString(srcField.String())
			}
		case reflect.Uint8:
			if srcField.Uint() > 0 {
				destField.SetUint(srcField.Uint())
			}
		case reflect.Bool:
			destField.SetBool(srcField.Bool())
		case reflect.Struct:
			if !srcField.Interface().(time.Time).IsZero() {
				destField.Set(reflect.ValueOf(srcField.Interface()))
			}
		}
	}
}
