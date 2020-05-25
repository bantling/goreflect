package goreflect

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncTypeMatch(t *testing.T) {
	tm := NewTypeMatch(0)
	ftm := FuncTypeMatch{typeMatch: tm, required: Required}
	assert.Equal(t, "int", ftm.String())
	ftm.required = false
	assert.Equal(t, "[int]", ftm.String())

	tm = NewTypeMatch(0, Ptr)
	ftm = FuncTypeMatch{typeMatch: tm, required: Required}
	assert.Equal(t, "*int", ftm.String())
	ftm.required = false
	assert.Equal(t, "[*int]", ftm.String())

	tm = NewTypeMatch(0, Ptr, PtrPtr)
	ftm = FuncTypeMatch{typeMatch: tm, required: Required}
	assert.Equal(t, "(*|**)int", ftm.String())
	ftm.required = false
	assert.Equal(t, "[(*|**)int]", ftm.String())

	type str struct{}

	tm = NewTypeMatch(reflect.Struct)
	ftm = FuncTypeMatch{typeMatch: tm, required: Required}
	assert.Equal(t, "struct", ftm.String())
	ftm.required = false
	assert.Equal(t, "[struct]", ftm.String())

	tm = NewTypeMatch(reflect.Slice, Ptr)
	ftm = FuncTypeMatch{typeMatch: tm, required: Required}
	assert.Equal(t, "*slice", ftm.String())
	ftm.required = false
	assert.Equal(t, "[*slice]", ftm.String())

	tm = NewMultiTypeMatch(Value, Ptr, 0, "")
	ftm = FuncTypeMatch{typeMatch: tm, required: Required}
	assert.Equal(t, "[*](int|string)", ftm.String())
	ftm.required = Optional
	assert.Equal(t, "[[*](int|string)]", ftm.String())

	tm = NewMultiTypeMatch(Value, PtrPtr, 0, "")
	ftm = FuncTypeMatch{typeMatch: tm, required: Required}
	assert.Equal(t, "[*|**](int|string)", ftm.String())
	ftm.required = Optional
	assert.Equal(t, "[[*|**](int|string)]", ftm.String())

	tm = NewMultiTypeMatch(Ptr, PtrPtr, 0, str{})
	ftm = FuncTypeMatch{typeMatch: tm, required: Required}
	assert.Equal(t, "(*|**)(int|goreflect.str)", ftm.String())
	ftm.required = Optional
	assert.Equal(t, "[(*|**)(int|goreflect.str)]", ftm.String())
}

func TestFuncMatcher(t *testing.T) {
	falseTester := func(matcher *FuncMatcher, fn interface{}) {
		paramTypes, returnTypes, matches := matcher.MatchingTypes(fn)
		assert.Equal(t, map[int]reflect.Type(nil), paramTypes)
		assert.Equal(t, map[int]reflect.Type(nil), returnTypes)
		assert.False(t, matches)
		paramIndexes, returnIndexes, matches := matcher.MatchingIndexes(fn)
		assert.Equal(t, map[int]bool(nil), paramIndexes)
		assert.Equal(t, map[int]bool(nil), returnIndexes)
		assert.False(t, matches)
		matches = matcher.Matches(fn)
		assert.False(t, matches)
	}

	matcher := NewFuncMatcher()
	var fn interface{} = func() {}
	paramTypes, returnTypes, matches := matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{}, returnTypes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches := matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{}, paramIndexes)
	assert.Equal(t, map[int]bool{}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)
	assert.Equal(t, "func()", matcher.String())

	falseTester(matcher, func(int) {})
	falseTester(matcher, func() int { return 0 })
	falseTester(matcher, func(int) int { return 0 })

	matcher = NewFuncMatcher().
		WithParamType(0)
	fn = func(int) {}
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(0)}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{}, returnTypes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{0: true}, paramIndexes)
	assert.Equal(t, map[int]bool{}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)
	assert.Equal(t, "func(int)", matcher.String())

	falseTester(matcher, func() {})
	falseTester(matcher, func() int { return 0 })
	falseTester(matcher, func(int) int { return 0 })

	matcher = NewFuncMatcher().
		WithReturnType(0)
	fn = func() int { return 0 }
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(0)}, returnTypes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{}, paramIndexes)
	assert.Equal(t, map[int]bool{0: true}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)
	assert.Equal(t, "func() int", matcher.String())

	falseTester(matcher, func() {})
	falseTester(matcher, func(int) {})
	falseTester(matcher, func(int) int { return 0 })

	matcher = NewFuncMatcher().
		WithParamType(0).
		WithReturnType(0)
	fn = func(int) int { return 0 }
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(0)}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(0)}, returnTypes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{0: true}, paramIndexes)
	assert.Equal(t, map[int]bool{0: true}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)
	assert.Equal(t, "func(int) int", matcher.String())

	falseTester(matcher, func() {})
	falseTester(matcher, func(int) {})
	falseTester(matcher, func() int { return 0 })

	matcher = NewFuncMatcher().
		WithParamType(0).
		WithParamType("").
		WithReturnType(0)
	fn = func(int, string) int { return 0 }
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(0), 1: reflect.TypeOf("")}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(0)}, returnTypes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{0: true, 1: true}, paramIndexes)
	assert.Equal(t, map[int]bool{0: true}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)
	assert.Equal(t, "func(int, string) int", matcher.String())

	falseTester(matcher, func() {})
	falseTester(matcher, func(int) {})
	falseTester(matcher, func() int { return 0 })
	falseTester(matcher, func(int) int { return 0 })

	matcher = NewFuncMatcher().
		WithParamType(0).
		WithReturnType(0).
		WithReturnType("")
	fn = func(int) (int, string) { return 0, "" }
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(0)}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(0), 1: reflect.TypeOf("")}, returnTypes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{0: true}, paramIndexes)
	assert.Equal(t, map[int]bool{0: true, 1: true}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)
	assert.Equal(t, "func(int) (int, string)", matcher.String())

	falseTester(matcher, func() {})
	falseTester(matcher, func(int) {})
	falseTester(matcher, func() int { return 0 })
	falseTester(matcher, func(int) int { return 0 })

	// Optional params/return values
	matcher = NewFuncMatcher().
		WithOptionalParamType(0)
	fn = func() {}
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{}, returnTypes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{}, paramIndexes)
	assert.Equal(t, map[int]bool{}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)
	assert.Equal(t, "func([int])", matcher.String())

	fn = func(int) {}
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(0)}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{}, returnTypes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{0: true}, paramIndexes)
	assert.Equal(t, map[int]bool{}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)

	falseTester(matcher, func(string) {})
	falseTester(matcher, func() int { return 0 })
	falseTester(matcher, func(int) int { return 0 })

	matcher = NewFuncMatcher().
		WithOptionalReturnType(0)
	fn = func() {}
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{}, returnTypes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{}, paramIndexes)
	assert.Equal(t, map[int]bool{}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)
	assert.Equal(t, "func() [int]", matcher.String())

	fn = func() int { return 0 }
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(0)}, returnTypes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{}, paramIndexes)
	assert.Equal(t, map[int]bool{0: true}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)

	falseTester(matcher, func(int) {})
	falseTester(matcher, func() string { return "" })

	matcher = NewFuncMatcher().
		WithOptionalParamType(0).
		WithOptionalReturnType(0)
	fn = func() {}
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{}, returnTypes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{}, paramIndexes)
	assert.Equal(t, map[int]bool{}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)
	assert.Equal(t, "func([int]) [int]", matcher.String())

	fn = func(int) {}
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(0)}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{}, returnTypes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{0: true}, paramIndexes)
	assert.Equal(t, map[int]bool{}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)

	fn = func() int { return 0 }
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(0)}, returnTypes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{}, paramIndexes)
	assert.Equal(t, map[int]bool{0: true}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)

	fn = func(int) int { return 0 }
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(0)}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(0)}, returnTypes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{0: true}, paramIndexes)
	assert.Equal(t, map[int]bool{0: true}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)

	falseTester(matcher, func(string) {})
	falseTester(matcher, func() string { return "" })

	matcher = NewFuncMatcher().
		WithParamType(0).
		WithOptionalParamType("")
	fn = func(int) {}
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(0)}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{}, returnTypes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{0: true}, paramIndexes)
	assert.Equal(t, map[int]bool{}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)
	assert.Equal(t, "func(int, [string])", matcher.String())

	fn = func(int, string) {}
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(0), 1: reflect.TypeOf("")}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{}, returnTypes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{0: true, 1: true}, paramIndexes)
	assert.Equal(t, map[int]bool{}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)

	falseTester(matcher, func() {})
	falseTester(matcher, func(string) {})
	falseTester(matcher, func(string, int) {})
	falseTester(matcher, func(int, string) int { return 0 })

	matcher = NewFuncMatcher().
		WithOptionalParamType(0).
		WithParamType("")
	fn = func(string) {}
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{1: reflect.TypeOf("")}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{}, returnTypes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{1: true}, paramIndexes)
	assert.Equal(t, map[int]bool{}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)
	assert.Equal(t, "func([int], string)", matcher.String())

	fn = func(int, string) {}
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(0), 1: reflect.TypeOf("")}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{}, returnTypes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{0: true, 1: true}, paramIndexes)
	assert.Equal(t, map[int]bool{}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)

	falseTester(matcher, func() {})
	falseTester(matcher, func(int) {})
	falseTester(matcher, func(string, int) {})
	falseTester(matcher, func(int, string) int { return 0 })

	matcher = NewFuncMatcher().
		WithOptionalReturnType(0)
	fn = func() {}
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{}, returnTypes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{}, paramIndexes)
	assert.Equal(t, map[int]bool{}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)
	assert.Equal(t, "func() [int]", matcher.String())

	fn = func() int { return 0 }
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(0)}, returnTypes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{}, paramIndexes)
	assert.Equal(t, map[int]bool{0: true}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)

	falseTester(matcher, func(int) {})
	falseTester(matcher, func() string { return "" })

	// reflect.Kind matching
	type str struct{}

	matcher = NewFuncMatcher().
		WithParamType(0).
		WithOptionalParamType(reflect.Struct)
	fn = func(int) {}
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(0)}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{}, returnTypes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{0: true}, paramIndexes)
	assert.Equal(t, map[int]bool{}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)
	assert.Equal(t, "func(int, [struct])", matcher.String())

	fn = func(int, str) {}
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(0), 1: reflect.TypeOf(str{})}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{}, returnTypes)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{0: true, 1: true}, paramIndexes)
	assert.Equal(t, map[int]bool{}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)

	falseTester(matcher, func() {})
	falseTester(matcher, func(int, string) {})

	matcher = NewFuncMatcher().
		WithParamOfTypes(Value, Value, 0, "").
		WithReturnOfTypes(Value, Value, "", reflect.Struct)
	fn = func(int) string { return "" }
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(0)}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf("")}, returnTypes)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{0: true}, paramIndexes)
	assert.Equal(t, map[int]bool{0: true}, returnIndexes)
	assert.True(t, matcher.Matches(fn))
	matches = matcher.Matches(fn)
	assert.True(t, matches)
	assert.Equal(t, "func(int|string) string|struct", matcher.String())

	fn = func(string) str { return str{} }
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf("")}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(str{})}, returnTypes)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{0: true}, paramIndexes)
	assert.Equal(t, map[int]bool{0: true}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)

	falseTester(matcher, func() {})
	falseTester(matcher, func(int, string) {})

	matcher = NewFuncMatcher().
		WithOptionalParamOfTypes(Value, Value, 0, "").
		WithOptionalReturnOfTypes(Value, Ptr, "", reflect.Struct)
	fn = func() {}
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{}, returnTypes)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{}, paramIndexes)
	assert.Equal(t, map[int]bool{}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)
	assert.Equal(t, "func([int|string]) [[*](string|struct)]", matcher.String())

	fn = func(string) *str { return &str{} }
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf("")}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf((*str)(nil))}, returnTypes)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{0: true}, paramIndexes)
	assert.Equal(t, map[int]bool{0: true}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)

	falseTester(matcher, func(str) {})
	falseTester(matcher, func() int { return 0 })

	matcher = NewFuncMatcher().
		WithParams(NewFuncTypeMatch(0, Required)).
		WithReturns(NewFuncTypeMatch("", Optional))
	fn = func(int) {}
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(0)}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{}, returnTypes)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{0: true}, paramIndexes)
	assert.Equal(t, map[int]bool{}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)
	assert.Equal(t, "func(int) [string]", matcher.String())

	fn = func(int) string { return "" }
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf(0)}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{0: reflect.TypeOf("")}, returnTypes)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{0: true}, paramIndexes)
	assert.Equal(t, map[int]bool{0: true}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)

	falseTester(matcher, func(str) {})
	falseTester(matcher, func() int { return 0 })

	matcher = NewFuncMatcher().
		WithOptionalParamType(0).
		WithOptionalParamType("").
		WithOptionalReturnType(0).
		WithOptionalReturnType("")
	fn = func(string) string { return "" }
	paramTypes, returnTypes, matches = matcher.MatchingTypes(fn)
	assert.Equal(t, map[int]reflect.Type{1: reflect.TypeOf("")}, paramTypes)
	assert.Equal(t, map[int]reflect.Type{1: reflect.TypeOf("")}, returnTypes)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(fn)
	assert.Equal(t, map[int]bool{1: true}, paramIndexes)
	assert.Equal(t, map[int]bool{1: true}, returnIndexes)
	assert.True(t, matches)
	matches = matcher.Matches(fn)
	assert.True(t, matches)

	falseTester(matcher, func(str) {})
}
