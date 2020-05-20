package goreflect

import (
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

// GetReflectTypeOrKindValueOf returns either a reflect.Type or a reflect.Kind that represents
// a zero indirect version of the value or type provided.
// Passing any of the following returns (reflect.TypeOf(string), reflect.Invalid):
// - "str"
// - &"str"
// - &&"str"
// - reflect.ValueOf("")
// - reflect.ValueOf(&"")
// - reflect.ValueOf(&&"")
// - reflect.TypeOf("")
// - reflect.TypeOf(&"")
// - reflect.TypeOf(&&"")
// Passing reflect.String returns (nil, reflect.String)
func GetReflectTypeOrKindValueOf(val interface{}) (reflect.Type, reflect.Kind) {
	if kind, ok := val.(reflect.Kind); ok {
		return nil, kind
	}

	// Deref the given type to get actual type
	return DerefdReflectType(GetReflectTypeOf(val)), reflect.Invalid
}
