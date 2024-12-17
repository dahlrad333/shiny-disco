package server

import (
	"fmt"
	"reflect"
)

func PrintTypeInfo(i interface{}, indent string) {
	val := reflect.ValueOf(i)
	typ := reflect.TypeOf(i)

	for val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	if val.Kind() == reflect.Struct {
		fmt.Printf("%sStruct: %s\n", indent, typ.Name())

		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			fieldValue := val.Field(i)

			fmt.Printf("%s- Field: %s, Type: %s\n", indent, field.Name, field.Type)

			if fieldValue.Kind() == reflect.Struct || fieldValue.Kind() == reflect.Ptr {
				PrintTypeInfo(fieldValue.Interface(), indent+"  ")
			}
		}
	} else {
		fmt.Printf("%sType: %s, Value: %v\n", indent, typ, val)
	}
}