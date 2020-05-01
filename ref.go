package goreflect

import (
	"reflect"
)

// Indirection constants
const (
	Value  = 0
	Ptr    = 1
	PtrPtr = 2
)

// GetReflectValueOf takes a value and returns a reflect.Value wrapper.
// If the value is already a reflect.Value, it is returned as is.
func GetReflectValueOf(value interface{}) reflect.Value {
	if v, ok := value.(reflect.Value); ok {
		return v
	}

	return reflect.ValueOf(value)
}

// DerefdReflectValue takes a reflect.Value that may be have one or more levels of indirection, and dereferences it until it is a value type.
// If the value is Invalid, it is returned as is.
func DerefdReflectValue(value reflect.Value) reflect.Value {
	// Assume the instance is not a pointer
	derefd := value

	// Only deref a valid Value
	if derefd.IsValid() {
		for derefd.Kind() == reflect.Ptr {
			derefd = derefd.Elem()
		}
	}

	return derefd
}

// DerefdReflectType takes a reflect.Type that may be have one or more levels of indirection, and dereferences it until it is a value type.
// Only the outer type is derefd, eg *[]**string => []**string
func DerefdReflectType(typ reflect.Type) reflect.Type {
	// Assume the type is not a pointer
	derefd := typ

	// Deref the type
	for derefd.Kind() == reflect.Ptr {
		derefd = derefd.Elem()
	}

	return derefd
}

// FullyDerefdReflectType returns the fully dereferenced type, where all indirections are removed:
// T, *T, ... => T
// [1]T, *[1]T, [1]*T, ... => [1]T
// []T, *[]T, []*T, ... => []T
// chan T, *chan T, chan *T, ... => chan T
// map[K]V, *map[K]V, map[*K]V, map[K]*V, ... => map[K]V
func FullyDerefdReflectType(typ reflect.Type) reflect.Type {
	derefd := DerefdReflectType(typ)

	switch derefd.Kind() {
	case reflect.Array:
		derefd = reflect.ArrayOf(derefd.Len(), FullyDerefdReflectType(derefd.Elem()))
	case reflect.Slice:
		derefd = reflect.SliceOf(FullyDerefdReflectType(derefd.Elem()))
	case reflect.Chan:
		derefd = reflect.ChanOf(derefd.ChanDir(), FullyDerefdReflectType(derefd.Elem()))
	case reflect.Map:
		derefd = reflect.MapOf(FullyDerefdReflectType(derefd.Key()), FullyDerefdReflectType(derefd.Elem()))
	}

	return derefd
}

// NumRefs returns the number of references in the given type
func NumRefs(typ reflect.Type) int {
	numRefs := 0
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		numRefs++
	}

	return numRefs
}

// CreateRefs creates the specified number of references to the value given.
// Note that if the value already has one or more references, the specified number of references is still added.
func CreateRefs(value reflect.Value, numRefs uint) reflect.Value {
	currentValue := value
	for i := numRefs; i > 0; i-- {
		// reflect.New adds an indirection
		newPtr := reflect.New(currentValue.Type())

		// Set the new pointer to point to current value
		newPtr.Elem().Set(currentValue)

		// current value is new pointer
		currentValue = newPtr
	}

	return currentValue
}
