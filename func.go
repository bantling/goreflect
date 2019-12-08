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

// TypeMatch describes a single value type
type TypeMatch struct {
	valueType      reflect.Type
	minIndirection int
	maxIndirection int
}

func (v TypeMatch) String() string {
	return fmt.Sprintf(
		"TypeMatch: {valueType: %s, minIndirection: %d, maxIndirection: %d}",
		v.valueType,
		v.minIndirection,
		v.maxIndirection,
	)
}

// NewTypeMatch constructs a TypeMatch
// The given type cannot have more than two levels of pointer indirection
// If indirection may have up two ints, as follows:
// - 0 ints: minIndirection = maxIndirection = 0
// - 1 int:  minIndirection = maxIndirection = int
// - 2 ints: minIndirection = first int, maxIndirection = second int
// The type passed can be a value, a reflect.Value that wraps a value,
// or a reflect.Type that wraps a value type.
// In all cases, the type represented can have up to two levels of indirection.
func NewTypeMatch(val interface{}, indirection ...int) TypeMatch {
	typ := GetReflectTypeOf(val)

	// Deref the given type up to two times if neccessary to get actual type
	valueType := typ
	if valueType.Kind() == reflect.Ptr {
		valueType = valueType.Elem()
	}
	if valueType.Kind() == reflect.Ptr {
		valueType = valueType.Elem()
	}
	if valueType.Kind() == reflect.Ptr {
		panic(fmt.Errorf("Too much indirection in type %s", typ))
	}

	minIndirection := 0
	maxIndirection := 0
	if len(indirection) >= 1 {
		minIndirection = indirection[0]
		maxIndirection = minIndirection
	}
	if len(indirection) >= 2 {
		maxIndirection = indirection[1]
	}

	return TypeMatch{
		valueType:      valueType,
		minIndirection: minIndirection,
		maxIndirection: maxIndirection,
	}
}

// Matches returns true if this type matches the given reflect type
func (tm TypeMatch) Matches(t reflect.Type) bool {
	valueType := t
	indirection := 0
	if valueType.Kind() == reflect.Ptr {
		valueType = valueType.Elem()
		indirection++
	}
	if valueType.Kind() == reflect.Ptr {
		valueType = valueType.Elem()
		indirection++
	}
	if valueType.Kind() == reflect.Ptr {
		return false
	}

	return (valueType == tm.valueType) &&
		(indirection >= tm.minIndirection) &&
		(indirection <= tm.maxIndirection)
}

// FuncTypeMatch is a TypeMatch for a function parameter or return type.
// The difference is that a parameter or return type that may or may not be present in a matching function.
type FuncTypeMatch struct {
	typeMatch TypeMatch
	optional  bool
}

// NewFuncTypeMatch constructs a FuncTypeMatch
func NewFuncTypeMatch(val interface{}, optional bool, indirection ...int) FuncTypeMatch {
	return FuncTypeMatch{
		typeMatch: NewTypeMatch(val, indirection...),
		optional: optional,
	}
}

// FuncMatcher describes desired value types
type FuncMatcher struct {
	paramTypes  []FuncTypeMatch
	returnTypes []FuncTypeMatch
}

// NewFuncMatcher constructs a FuncMatcher
func NewFuncMatcher() *FuncMatcher {
	return &FuncMatcher{}
}

// WithParamType builder adds the given param type
func (f *FuncMatcher) WithParamType(val interface{}, indirection ...int) *FuncMatcher {
	f.paramTypes = append(f.paramTypes, NewFuncTypeMatch(val, false, indirection...))
	return f
}

// WithOptionalParamType builder adds the given param type
func (f *FuncMatcher) WithOptionalParamType(val interface{}, indirection ...int) *FuncMatcher {
	f.paramTypes = append(f.paramTypes, NewFuncTypeMatch(val, true, indirection...))
	return f
}

// WithReturnType builder adds the given return type
func (f *FuncMatcher) WithReturnType(val interface{}, indirection ...int) *FuncMatcher {
	f.returnTypes = append(f.returnTypes, NewFuncTypeMatch(val, false, indirection...))
	return f
}

// WithOptionalReturnType builder adds the given return type
func (f *FuncMatcher) WithOptionalReturnType(val interface{}, indirection ...int) *FuncMatcher {
	f.returnTypes = append(f.returnTypes, NewFuncTypeMatch(val, true, indirection...))
	return f
}

// Matches returns true if the given value matches the parameter and return types of this matcher.
// The value passed can be a function object, a reflect.Value that wraps a function object,
// or a reflect.Type that wraps a function type.
// If the value is not any of the above types, false is returned.
func (f *FuncMatcher) Matches(fn interface{}) bool {
	// Get a reflect.Type wrapper
	fnType := GetReflectTypeOf(fn)

	// Assume the value does not match
	matches := false
	if (fnType.Kind() == reflect.Func) &&
		(fnType.NumIn() == len(f.paramTypes)) &&
		(fnType.NumOut() == len(f.returnTypes)) {
		// Now that we know it is a func with correct number of args and return vals,
		// assume it matches unless we encounter an arg or return value that doesn't.
		matches = true

		for i, paramType := range f.paramTypes {
			if !paramType.typeMatch.Matches(fnType.In(i)) {
				return false
			}
		}

		for i, returnType := range f.returnTypes {
			if !returnType.typeMatch.Matches(fnType.Out(i)) {
				return false
			}
		}
	}

	return matches
}
