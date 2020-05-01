package goreflect

import (
	"fmt"
	"reflect"
)

// GetReflectTypeOf takes a value, reflect.Value wrapper, or reflect.Type wrapper and returns a reflect.Type wrapper.
func GetReflectTypeOf(val interface{}) reflect.Type {
	// Get a reflect.Type wrapper
	var valType reflect.Type
	if theType, ok := val.(reflect.Type); ok {
		valType = theType
	} else if theVal, ok := val.(reflect.Value); ok {
		valType = theVal.Type()
	} else {
		valType = reflect.TypeOf(val)
	}

	return valType
}

// GetReflectKindOrTypeValueOf returns either a reflect.Kind or a reflect.Type that repesents
// a zero indirect version of the value or type provided.
// Passing reflect.String returns (reflect.String, nil)
// Passing ay of the following returns (reflect.Invalid, reflect.TypeOf(string)):
// - "str"
// - &"str"
// - &&"str"
// - reflect.ValueOf("")
// - reflect.ValueOf(&"")
// - reflect.ValueOf(&&"")
// - reflect.TypeOf("")
// - reflect.TypeOf(&"")
// - reflect.TypeOf(&&"")
func GetReflectKindOrTypeValueOf(val interface{}) (reflect.Kind, reflect.Type) {
	if kind, ok := val.(reflect.Kind); ok {
		return kind, nil
	}
	typ := GetReflectTypeOf(val)

	// Deref the given type up to two times if neccessary to get actual type
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() == reflect.Ptr {
		panic(fmt.Errorf("Too much indirection in type %s", typ))
	}

	return reflect.Invalid, typ
}
