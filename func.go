package goreflect

import (
	"fmt"
	"reflect"
	"strings"
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
// a zero indirect version of the type provided.
// Passing reflect.Struct returns (reflect.Struct, nil)
// Passing string, *string, **string returns (reflect.Invalid, reflect.TypeOf(string))
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

type Indirection uint

const (
	Value Indirection = iota
	Ptr
	PtrToPtr
)

var (
	indirectionToString = map[Indirection]string{
		Value: "Value",
		Ptr: "Ptr",
		PtrToPtr: "PtrToPtr",
	}
)

// String returns Indirection as a string
func (i Indirection) String() string {
	return indirectionToString[i]
}

// TypeMatch describes a single match by any one of multiple kinds and/or types
type TypeMatch struct {
	types          []reflect.Type
	kinds          []reflect.Kind
	minIndirection Indirection
	maxIndirection Indirection
}

// String returns a signature for the types matched, use vertical bars to separate multiple choices.
// If pointer indirections are allowed, they occur once at the beginning.
// If there are multiple pointer indirections, they are in parantheses.
// If there are multiple type choices, they are in parantheses only if there is at least one pointer indirection
// Examples:
// "string"
// "*string"
// "(*|**)string"
// "string|slice"
// "*(string|slice)"
// "(*|**)(string|slice)"
func (v TypeMatch) String() string {
	var str strings.Builder

	// Add pointer indirection(s)
	var numIndirections Indirection = 0
	if v.maxIndirection > 0 {
		numIndirections = v.maxIndirection - v.minIndirection + 1
	}
	multipleIndirections := numIndirections > 1
	useBrackets := v.minIndirection == 0

	if multipleIndirections {
		if useBrackets {
			str.WriteRune('[')
		} else {
			str.WriteRune('(')
		}
	}

	for i := 1; i <= int(v.maxIndirection); i++ {
		if i > 1 {
			str.WriteRune('|')
		}

		str.WriteString(strings.Repeat("*", i))
	}

	if multipleIndirections {
		if useBrackets {
			str.WriteRune(']')
		} else {
			str.WriteRune(')')
		}
	}

	// Add type(s), then kind(s)
	needTypeParens := (numIndirections > 1) && ((len(v.types) + len(v.kinds)) > 1)
	if needTypeParens {
		str.WriteRune('(')
	}

	firstType := true
	for _, typ := range v.types {
		if !firstType {
			str.WriteRune('|')
		}
		firstType = false

		str.WriteString(typ.String())
	}

	for _, kind := range v.kinds {
		if !firstType {
			str.WriteRune('|')
		}
		firstType = false

		str.WriteString(kind.String())
	}

	if needTypeParens {
		str.WriteRune(')')
	}

	return str.String()
}

// NewTypeMatch constructs a TypeMatch
// The value passed can be a value, a reflect.Value that wraps a value,
// a reflect.Type that wraps a value type, or a reflect.Kind.
// If the value is not a reflect.Kind, then it cannot have more than two levels of pointer indirection.
// Indirection may have up two ints, as follows:
// - 0 ints: minIndirection = maxIndirection = 0
// - 1 int:  minIndirection = maxIndirection = int
// - 2 ints: minIndirection = first int, maxIndirection = second int
// Panics if maxIndirection < minIndirection
func NewTypeMatch(val interface{}, indirection ...Indirection) TypeMatch {
	var kinds []reflect.Kind
	var types []reflect.Type

	kind, typ := GetReflectKindOrTypeValueOf(val)
	if kind != reflect.Invalid {
		kinds = []reflect.Kind{kind}
	} else {
		types = []reflect.Type{typ}
	}

	minIndirection := Value
	maxIndirection := Value
	if len(indirection) >= 1 {
		minIndirection = indirection[0]
		maxIndirection = minIndirection
	}
	if len(indirection) >= 2 {
		maxIndirection = indirection[1]
	}

	if maxIndirection < minIndirection {
		panic(fmt.Errorf("NewTypeMatch: maxIndirection %s < minIndirection %s", maxIndirection, minIndirection))
	}

	return TypeMatch{
		types:          types,
		kinds:          kinds,
		minIndirection: minIndirection,
		maxIndirection: maxIndirection,
	}
}

// NewMultiTypeMatch constructs a TypeMatch that can match against any of several choices.
// Each choice is the same as for NewTypeMatch.
func NewMultiTypeMatch(
	minIndirection Indirection,
	maxIndirection Indirection,
	vals ...interface{},
) TypeMatch {
	var types []reflect.Type
	var kinds []reflect.Kind

	for _, val := range vals {
		kind, typ := GetReflectKindOrTypeValueOf(val)
		if kind != reflect.Invalid {
			kinds = append(kinds, kind)
		} else {
			types = append(types, typ)
		}
	}

	return TypeMatch{
		kinds:          kinds,
		types:          types,
		minIndirection: minIndirection,
		maxIndirection: maxIndirection,
	}
}

// Matches returns true if this type matches the given reflect type
func (tm TypeMatch) Matches(t reflect.Type) bool {
	// Get the given type as a zero indirection value type, counting indirections
	valueType := t
	indirection := Value

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

	// Check indirection levels first
	if (indirection < tm.minIndirection) || (indirection > tm.maxIndirection) {
		return false
	}

	// Check if any kind matches the zero indirection value kind
	valueKind := valueType.Kind()
	for _, kind := range tm.kinds {
		if kind == valueKind {
			return true
		}
	}

	// Check if any type matches the zero indirection value type
	for _, typ := range tm.types {
		if typ == valueType {
			return true
		}
	}

	return false
}

// FuncTypeMatch is a TypeMatch for a function parameter or return type.
// The difference is that a parameter or return type that may or may not be present in a matching function.
type FuncTypeMatch struct {
	typeMatch TypeMatch
	required  bool
}

// String returns type signature, with [] around it if it is optional
func (f FuncTypeMatch) String() string {
	typeString := f.typeMatch.String()
	if f.required {
		return typeString
	}

	return fmt.Sprintf("[%s]", typeString)
}

// NewFuncTypeMatch constructs a FuncTypeMatch
// See NewTypeMatch
func NewFuncTypeMatch(val interface{}, required bool, indirection ...Indirection) FuncTypeMatch {
	return FuncTypeMatch{
		typeMatch: NewTypeMatch(val, indirection...),
		required:  required,
	}
}

// NewFuncMultiTypeMatch constructs a FuncTypeMatch
// See NewMultiTypeMatch
func NewFuncMultiTypeMatch(
	minIndirection Indirection,
	maxIndirection Indirection,
	required bool,
	vals ...interface{},
) FuncTypeMatch {
	return FuncTypeMatch{
		typeMatch: NewMultiTypeMatch(minIndirection, maxIndirection, vals...),
		required:  required,
	}
}

// FuncMatcher describes desired value types
type FuncMatcher struct {
	paramTypes  []FuncTypeMatch
	returnTypes []FuncTypeMatch
}

// String returns signature of matching functions
func (m FuncMatcher) String() string {
	var bldr strings.Builder

	bldr.WriteString("func(")
	first := true
	for _, pt := range m.paramTypes {
		if !first {
			bldr.WriteString(", ")
		}
		first = false

		bldr.WriteString(pt.String())
	}
	bldr.WriteRune(')')

	if len(m.returnTypes) > 0 {
		bldr.WriteRune(' ')
		multiTypes := len(m.returnTypes) > 1
		if multiTypes {
			bldr.WriteRune('(')
		}

		first = true
		for _, rt := range m.returnTypes {
			if !first {
				bldr.WriteString(", ")
			}
			first = false

			bldr.WriteString(rt.String())
		}

		if multiTypes {
			bldr.WriteRune(')')
		}
	}

	return bldr.String()
}

// NewFuncMatcher constructs a FuncMatcher
func NewFuncMatcher() *FuncMatcher {
	return &FuncMatcher{}
}

// WithParamType builder adds the given param type
func (f *FuncMatcher) WithParamType(val interface{}, indirection ...Indirection) *FuncMatcher {
	f.paramTypes = append(f.paramTypes, NewFuncTypeMatch(val, true, indirection...))
	return f
}

// WithOptionalParamType builder adds the given param type
func (f *FuncMatcher) WithOptionalParamType(val interface{}, indirection ...Indirection) *FuncMatcher {
	f.paramTypes = append(f.paramTypes, NewFuncTypeMatch(val, false, indirection...))
	return f
}

// WithParamOfTypes builder adds a single param that be any one of multiple types
func (f *FuncMatcher) WithParamOfTypes(
	minIndirection Indirection,
	maxIndirection Indirection,
	vals ...interface{},
) *FuncMatcher {
	f.paramTypes = append(
		f.paramTypes,
		NewFuncMultiTypeMatch(minIndirection, maxIndirection, true, vals...),
	)

	return f
}

// WithOptionalParamOfTypes builder adds a single optional param that be any one of multiple types
func (f *FuncMatcher) WithOptionalParamOfTypes(
	minIndirection Indirection,
	maxIndirection Indirection,
	vals ...interface{},
) *FuncMatcher {
	f.paramTypes = append(
		f.paramTypes,
		NewFuncMultiTypeMatch(minIndirection, maxIndirection, false, vals...),
	)

	return f
}

// WithParams builder adds the given FuncTypeMatch objects to the parameters
func (f *FuncMatcher) WithParams(
	funcTypeMatches ...FuncTypeMatch,
) *FuncMatcher {
	f.paramTypes = append(f.paramTypes, funcTypeMatches...)

	return f
}

// WithReturnType builder adds the given return type
func (f *FuncMatcher) WithReturnType(val interface{}, indirection ...Indirection) *FuncMatcher {
	f.returnTypes = append(f.returnTypes, NewFuncTypeMatch(val, true, indirection...))
	return f
}

// WithOptionalReturnType builder adds the given return type
func (f *FuncMatcher) WithOptionalReturnType(val interface{}, indirection ...Indirection) *FuncMatcher {
	f.returnTypes = append(f.returnTypes, NewFuncTypeMatch(val, false, indirection...))
	return f
}

// WithReturnOfTypes builder adds a single return that be any one of multiple types
func (f *FuncMatcher) WithReturnOfTypes(
	minIndirection Indirection,
	maxIndirection Indirection,
	vals ...interface{},
) *FuncMatcher {
	f.returnTypes = append(
		f.returnTypes,
		NewFuncMultiTypeMatch(minIndirection, maxIndirection, true, vals...),
	)

	return f
}

// WithOptionalReturnOfTypes builder adds a single optional return that be any one of multiple types
func (f *FuncMatcher) WithOptionalReturnOfTypes(
	minIndirection Indirection,
	maxIndirection Indirection,
	vals ...interface{},
) *FuncMatcher {
	f.returnTypes = append(
		f.returnTypes,
		NewFuncMultiTypeMatch(minIndirection, maxIndirection, false, vals...),
	)

	return f
}

// WithReturns builder adds the given FuncTypeMatch objects to the returns
func (f *FuncMatcher) WithReturns(
	funcTypeMatch FuncTypeMatch,
	funcTypeMatches ...FuncTypeMatch,
) *FuncMatcher {
	f.returnTypes = append(f.returnTypes, funcTypeMatch)
	f.returnTypes = append(f.returnTypes, funcTypeMatches...)

	return f
}

// MatchingIndexes returns true if the given value matches the parameter and return types of this matcher.
// The value passed can be a function object, a reflect.Value that wraps a function object,
// or a reflect.Type that wraps a function type.
// If the value is not any of the above types, false is returned.
// If the value is a matching function, the indexes of the matching parameter and return types are
// also returned. If there are optional parameter and/or return types, this allows the caller to
// determine which particular parameter and return types were actually used by the function.
// Note that if a matching func has no parameters and/or return types, the index array(s) will be
// zero length. By constrast, if the func does not match, both index arrays will be nil.
func (f FuncMatcher) MatchingIndexes(fn interface{}) (params []int, returns []int, matches bool) {
	// Indexes to return
	paramIndexes := []int{}
	returnIndexes := []int{}

	// Get a reflect.Type wrapper
	fnType := GetReflectTypeOf(fn)

	if fnType.Kind() == reflect.Func {
		// Iterate function params
		paramIndex := 0
		numParams := fnType.NumIn()
		// If we have no param types to match, then the func must accept no params
		if len(f.paramTypes) == 0 {
			if numParams != 0 {
				return nil, nil, false
			}
		} else {
			// See if our params match that of the function
			for _, paramType := range f.paramTypes {
				// Advance to next loop if we have a matching param
				if (paramIndex < numParams) && paramType.typeMatch.Matches(fnType.In(paramIndex)) {
					paramIndexes = append(paramIndexes, paramIndex)
					paramIndex++
					continue
				}

				// Required params must match
				if paramType.required {
					return nil, nil, false
				}
			}
		}

		// If there are still parameters we haven't matched, it's not a match
		if paramIndex < numParams {
			return nil, nil, false
		}

		// Iterate return values
		returnIndex := 0
		numReturns := fnType.NumOut()
		// If we have no return types to match, then the func must return no values
		if len(f.returnTypes) == 0 {
			if numReturns != 0 {
				return nil, nil, false
			}
		} else {
			// See if our returns match that of the function
			for _, returnType := range f.returnTypes {
				// Advance to next loop if we have a matching return
				if (returnIndex < numReturns) && returnType.typeMatch.Matches(fnType.Out(returnIndex)) {
					returnIndexes = append(returnIndexes, returnIndex)
					returnIndex++
					continue
				}

				// Required returns must match
				if returnType.required {
					return nil, nil, false
				}
			}
		}

		// If there are still returns we haven't matched, it's not a match
		if returnIndex < numReturns {
			return nil, nil, false
		}
	} else {
		return nil, nil, false
	}

	return paramIndexes, returnIndexes, true
}

// Matches simply calls MatchingIndexes and returns only the bool result, for simple yes/no matching
func (f FuncMatcher) Matches(fn interface{}) bool {
	_, _, matches := f.MatchingIndexes(fn)
	return matches
}

// Signature returns a string representing the possible signature(s) this matcher will accept.
// The string looks like a Go anonymous function declaration, except that if a given parameter or return type
// has multiple choices, they are separated by vertical bars.
// EG, "func (string|int, slice|struct) (string|int, error)"
// func (f FuncMatcher) Signature() string {
// 	var signature strings.Builder
// 	signature.WriteString("func (")

// 	// Add params
// 	for _, paramType := range f.paramTypes {

// 	}
// }
