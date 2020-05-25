package goreflect

import (
	"fmt"
	"reflect"
	"strings"
)

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
