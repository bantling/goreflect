package goreflect

import (
	"fmt"
	"reflect"
	"strings"
)

// TypeMatch describes a single match by any one of multiple kinds and/or types
type TypeMatch struct {
	types          []reflect.Type
	kinds          []reflect.Kind
	minIndirection int
	maxIndirection int
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
func (tm TypeMatch) String() string {
	var str strings.Builder

	// Add pointer indirection(s)
	var numIndirections int
	if tm.maxIndirection > 0 {
		numIndirections = tm.maxIndirection - tm.minIndirection + 1
	}
	multipleIndirections := numIndirections > 1
	useBrackets := tm.minIndirection == 0

	if multipleIndirections {
		if useBrackets {
			str.WriteRune('[')
		} else {
			str.WriteRune('(')
		}
	}

	for i := 1; i <= int(tm.maxIndirection); i++ {
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
	needTypeParens := (numIndirections > 1) && ((len(tm.types) + len(tm.kinds)) > 1)
	if needTypeParens {
		str.WriteRune('(')
	}

	firstType := true
	for _, typ := range tm.types {
		if !firstType {
			str.WriteRune('|')
		}
		firstType = false

		str.WriteString(typ.String())
	}

	for _, kind := range tm.kinds {
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
// The value passed can be a value, reflect.Value, reflect.Type, or reflect.Kind.
// Indirection may have up two ints, as follows:
// - 0 ints: minIndirection = maxIndirection = 0
// - 1 int:  minIndirection = maxIndirection = int
// - 2 ints: minIndirection = first int, maxIndirection = second int
// Panics if maxIndirection < minIndirection
func NewTypeMatch(val interface{}, indirection ...int) TypeMatch {
	var types []reflect.Type
	var kinds []reflect.Kind

	typ, kind := GetReflectTypeOrKindValueOf(val)
	if kind == reflect.Invalid {
		types = []reflect.Type{typ}
	} else {
		kinds = []reflect.Kind{kind}
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
		panic(fmt.Errorf("NewTypeMatch: maxIndirection %d < minIndirection %d", maxIndirection, minIndirection))
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
	minIndirection int,
	maxIndirection int,
	vals ...interface{},
) TypeMatch {
	var types []reflect.Type
	var kinds []reflect.Kind

	for _, val := range vals {
		typ, kind := GetReflectTypeOrKindValueOf(val)
		if kind == reflect.Invalid {
			types = append(types, typ)
		} else {
			kinds = append(kinds, kind)
		}
	}

	return TypeMatch{
		types:          types,
		kinds:          kinds,
		minIndirection: minIndirection,
		maxIndirection: maxIndirection,
	}
}

// Matches returns true if this type matches the given reflect type.
// If the given type is nil, false is returned.
func (tm TypeMatch) Matches(t reflect.Type) bool {
	if t == nil {
		return false
	}

	// Get the given type as a zero indirection value type, counting indirections
	valueType := DerefdReflectType(t)
	indirection := NumRefs(t)

	// Check indirection levels first
	if (indirection < tm.minIndirection) || (indirection > tm.maxIndirection) {
		return false
	}

	// Check if any type matches the zero indirection value type
	for _, typ := range tm.types {
		if typ == valueType {
			return true
		}
	}

	// Check if any kind matches the zero indirection value kind
	valueKind := valueType.Kind()
	for _, kind := range tm.kinds {
		if kind == valueKind {
			return true
		}
	}

	return false
}

// Types is the types accessor
func (tm TypeMatch) Types() []reflect.Type {
	typesCopy := make([]reflect.Type, len(tm.types))
	copy(typesCopy, tm.types)
	return typesCopy
}

// Kinds is the kinds accessor
func (tm TypeMatch) Kinds() []reflect.Kind {
	kindsCopy := make([]reflect.Kind, len(tm.kinds))
	copy(kindsCopy, tm.kinds)
	return kindsCopy
}

// MinIndirection is the minIndirection accessor
func (tm TypeMatch) MinIndirection() int {
	return tm.minIndirection
}

// MaxIndirection is the maxIndirection accessor
func (tm TypeMatch) MaxIndirection() int {
	return tm.maxIndirection
}

// Optionality describes whether a type is required or optional
type Optionality bool

// Optionality constants
const (
	Required Optionality = true
	Optional Optionality = false

	requiredString = "required"
	optionalString = "optional"
)

func (o Optionality) String() string {
	if o {
		return requiredString
	}

	return optionalString
}

// FuncTypeMatch is a TypeMatch for a function parameter or return type.
// The difference is that a parameter or return type that may or may not be present in a matching function.
type FuncTypeMatch struct {
	typeMatch TypeMatch
	required  Optionality
}

// String returns type signature, with [] around it if it is optional
func (ftm FuncTypeMatch) String() string {
	typeString := ftm.typeMatch.String()
	if ftm.required {
		return typeString
	}

	return fmt.Sprintf("[%s]", typeString)
}

// NewFuncTypeMatch constructs a FuncTypeMatch
// See NewTypeMatch
func NewFuncTypeMatch(val interface{}, required Optionality, indirection ...int) FuncTypeMatch {
	return FuncTypeMatch{
		typeMatch: NewTypeMatch(val, indirection...),
		required:  required,
	}
}

// NewFuncMultiTypeMatch constructs a FuncTypeMatch
// See NewMultiTypeMatch
func NewFuncMultiTypeMatch(
	minIndirection int,
	maxIndirection int,
	required Optionality,
	vals ...interface{},
) FuncTypeMatch {
	return FuncTypeMatch{
		typeMatch: NewMultiTypeMatch(minIndirection, maxIndirection, vals...),
		required:  required,
	}
}

// TypeMatch accessor
func (ftm FuncTypeMatch) TypeMatch() TypeMatch {
	return ftm.typeMatch
}

// Required accessor
func (ftm FuncTypeMatch) Required() Optionality {
	return ftm.required
}

// FuncMatcher describes desired value types
type FuncMatcher struct {
	paramTypes  []FuncTypeMatch
	returnTypes []FuncTypeMatch
}

// String returns signature of matching functions
func (fm FuncMatcher) String() string {
	var bldr strings.Builder

	bldr.WriteString("func(")
	first := true
	for _, pt := range fm.paramTypes {
		if !first {
			bldr.WriteString(", ")
		}
		first = false

		bldr.WriteString(pt.String())
	}
	bldr.WriteRune(')')

	if len(fm.returnTypes) > 0 {
		bldr.WriteRune(' ')
		multiTypes := len(fm.returnTypes) > 1
		if multiTypes {
			bldr.WriteRune('(')
		}

		first = true
		for _, rt := range fm.returnTypes {
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
func (fm *FuncMatcher) WithParamType(val interface{}, indirection ...int) *FuncMatcher {
	fm.paramTypes = append(fm.paramTypes, NewFuncTypeMatch(val, Required, indirection...))
	return fm
}

// WithOptionalParamType builder adds the given param type
func (fm *FuncMatcher) WithOptionalParamType(val interface{}, indirection ...int) *FuncMatcher {
	fm.paramTypes = append(fm.paramTypes, NewFuncTypeMatch(val, Optional, indirection...))
	return fm
}

// WithParamOfTypes builder adds a single param that be any one of multiple types
func (fm *FuncMatcher) WithParamOfTypes(
	minIndirection int,
	maxIndirection int,
	vals ...interface{},
) *FuncMatcher {
	fm.paramTypes = append(
		fm.paramTypes,
		NewFuncMultiTypeMatch(minIndirection, maxIndirection, Required, vals...),
	)

	return fm
}

// WithOptionalParamOfTypes builder adds a single optional param that be any one of multiple types
func (fm *FuncMatcher) WithOptionalParamOfTypes(
	minIndirection int,
	maxIndirection int,
	vals ...interface{},
) *FuncMatcher {
	fm.paramTypes = append(
		fm.paramTypes,
		NewFuncMultiTypeMatch(minIndirection, maxIndirection, Optional, vals...),
	)

	return fm
}

// WithParams builder adds the given FuncTypeMatch objects to the parameters
func (fm *FuncMatcher) WithParams(
	funcTypeMatches ...FuncTypeMatch,
) *FuncMatcher {
	fm.paramTypes = append(fm.paramTypes, funcTypeMatches...)

	return fm
}

// WithReturnType builder adds the given return type
func (fm *FuncMatcher) WithReturnType(val interface{}, indirection ...int) *FuncMatcher {
	fm.returnTypes = append(fm.returnTypes, NewFuncTypeMatch(val, Required, indirection...))
	return fm
}

// WithOptionalReturnType builder adds the given return type
func (fm *FuncMatcher) WithOptionalReturnType(val interface{}, indirection ...int) *FuncMatcher {
	fm.returnTypes = append(fm.returnTypes, NewFuncTypeMatch(val, Optional, indirection...))
	return fm
}

// WithReturnOfTypes builder adds a single return that be any one of multiple types
func (fm *FuncMatcher) WithReturnOfTypes(
	minIndirection int,
	maxIndirection int,
	vals ...interface{},
) *FuncMatcher {
	fm.returnTypes = append(
		fm.returnTypes,
		NewFuncMultiTypeMatch(minIndirection, maxIndirection, Required, vals...),
	)

	return fm
}

// WithOptionalReturnOfTypes builder adds a single optional return that be any one of multiple types
func (fm *FuncMatcher) WithOptionalReturnOfTypes(
	minIndirection int,
	maxIndirection int,
	vals ...interface{},
) *FuncMatcher {
	fm.returnTypes = append(
		fm.returnTypes,
		NewFuncMultiTypeMatch(minIndirection, maxIndirection, Optional, vals...),
	)

	return fm
}

// WithReturns builder adds the given FuncTypeMatch objects to the returns
func (fm *FuncMatcher) WithReturns(
	funcTypeMatch FuncTypeMatch,
	funcTypeMatches ...FuncTypeMatch,
) *FuncMatcher {
	fm.returnTypes = append(fm.returnTypes, funcTypeMatch)
	fm.returnTypes = append(fm.returnTypes, funcTypeMatches...)

	return fm
}

// MatchingTypes returns a set of matched param types, a set of matche returned types,
// and true if the given value matches the parameter and return types of this matcher.
// The value passed can be a function object, a reflect.Value that wraps a function object,
// or a reflect.Type that wraps a function type.
// If the value is not any of the above types, false is returned.
// If the value is a matching function, the types of the matching parameter and return types are
// also returned. If there are optional parameter and/or return types, this allows the caller to
// determine which particular parameter and return types were actually used by the function, as
// optional types indexs that were not matched will be mapped to nil.
// Note that if a matching func has no parameters and/or return types, the related index map(s) will be zero length.
// By constrast, if the func does not match, both index maps will be nil.
func (fm FuncMatcher) MatchingTypes(fn interface{}) (map[int]reflect.Type, map[int]reflect.Type, bool) {
	var (
		paramTypes  = map[int]reflect.Type{}
		returnTypes = map[int]reflect.Type{}
		matches     = true
	)

	// Get a reflect.Type wrapper for fn
	fnType := GetReflectTypeOf(fn)

	if fnType.Kind() != reflect.Func {
		// Any non-func value can't be a match
		return nil, nil, false
	}

	// Iterate function params
	numParams := fnType.NumIn()
	var (
		paramIndex   int
		paramType    FuncTypeMatch
		fnParamIndex int
	)
	// If we have no param types to match, then the func must accept no params
	if len(fm.paramTypes) == 0 {
		if numParams != 0 {
			return nil, nil, false
		}
	} else {
		// See if our params match that of the function
		for paramIndex, paramType = range fm.paramTypes {
			// Advance to next loop if we have a matching param
			if fnParamIndex < numParams {
				actualParamType := fnType.In(fnParamIndex)
				if paramType.typeMatch.Matches(actualParamType) {
					paramTypes[paramIndex] = actualParamType
					fnParamIndex++
					continue
				}
			}

			// Required params must match
			if paramType.required {
				return nil, nil, false
			}
		}
	}

	// If there are still parameters we haven't matched, it's not a match
	if fnParamIndex < numParams {
		return nil, nil, false
	}

	// Iterate return values
	numReturns := fnType.NumOut()
	var (
		returnIndex   int
		returnType    FuncTypeMatch
		fnReturnIndex int
	)
	// If we have no return types to match, then the func must return no values
	if len(fm.returnTypes) == 0 {
		if numReturns != 0 {
			return nil, nil, false
		}
	} else {
		// See if our returns match that of the function
		for returnIndex, returnType = range fm.returnTypes {
			// Advance to next loop if we have a matching return
			if fnReturnIndex < numReturns {
				actualReturnType := fnType.Out(fnReturnIndex)
				if returnType.typeMatch.Matches(actualReturnType) {
					returnTypes[returnIndex] = actualReturnType
					fnReturnIndex++
					continue
				}
			}

			// Required returns must match
			if returnType.required {
				return nil, nil, false
			}
		}
	}

	// If there are still returns we haven't matched, it's not a match
	if fnReturnIndex < numReturns {
		return nil, nil, false
	}

	return paramTypes, returnTypes, matches
}

// MatchingIndexes is like MatchingTypes, but returns true or false for matching indexes rather than type or nil.
func (fm FuncMatcher) MatchingIndexes(fn interface{}) (map[int]bool, map[int]bool, bool) {
	// Leverage MatchingTypes
	var (
		paramIndexes  = map[int]bool{}
		returnIndexes = map[int]bool{}
	)
	paramTypes, returnTypes, matches := fm.MatchingTypes(fn)

	// Convert return results
	if matches {
		for i, paramType := range paramTypes {
			if paramType != nil {
				paramIndexes[i] = true
			}
		}

		for i, returnType := range returnTypes {
			if returnType != nil {
				returnIndexes[i] = true
			}
		}

		return paramIndexes, returnIndexes, matches
	}

	return nil, nil, false
}

// Matches simply calls MatchingTypes and returns only the bool result, for simple yes/no matching
func (fm FuncMatcher) Matches(fn interface{}) bool {
	_, _, matches := fm.MatchingTypes(fn)

	return matches
}

// ParamTypes is the paramTypes accessor
func (fm FuncMatcher) ParamTypes() []FuncTypeMatch {
	paramTypesCopy := make([]FuncTypeMatch, len(fm.paramTypes))
	copy(paramTypesCopy, fm.paramTypes)
	return paramTypesCopy
}

// ReturnTypes is the returnTypes accessor
func (fm FuncMatcher) ReturnTypes() []FuncTypeMatch {
	returnTypesCopy := make([]FuncTypeMatch, len(fm.returnTypes))
	copy(returnTypesCopy, fm.returnTypes)
	return returnTypesCopy
}
