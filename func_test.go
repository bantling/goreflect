package goreflect

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReflectTypeOf(t *testing.T) {
	assert.Equal(t, reflect.TypeOf(0), GetReflectTypeOf(0))
	assert.Equal(t, reflect.TypeOf(0), GetReflectTypeOf(reflect.ValueOf(0)))
	assert.Equal(t, reflect.TypeOf(0), GetReflectTypeOf(reflect.TypeOf(0)))
}

func TestTypeMatch(t *testing.T) {
	tm := NewTypeMatch(0)
	assert.Equal(t, []reflect.Kind(nil), tm.kinds)
	assert.Equal(t, []reflect.Type{reflect.TypeOf(0)}, tm.types)
	assert.Equal(t, 0, tm.minIndirection)
	assert.Equal(t, 0, tm.maxIndirection)
	assert.True(t, tm.Matches(reflect.TypeOf(0)))
	assert.False(t, tm.Matches(reflect.TypeOf((*int)(nil))))
	assert.False(t, tm.Matches(reflect.TypeOf((**int)(nil))))
	assert.Equal(t, "int", tm.String())

	ftm := FuncTypeMatch{typeMatch: tm, required: true}
	assert.Equal(t, "int", ftm.String())
	ftm.required = false
	assert.Equal(t, "[int]", ftm.String())

	tm = NewTypeMatch(0, 1)
	assert.Equal(t, []reflect.Kind(nil), tm.kinds)
	assert.Equal(t, []reflect.Type{reflect.TypeOf(0)}, tm.types)
	assert.Equal(t, 1, tm.minIndirection)
	assert.Equal(t, 1, tm.maxIndirection)
	assert.False(t, tm.Matches(reflect.TypeOf(0)))
	assert.True(t, tm.Matches(reflect.TypeOf((*int)(nil))))
	assert.False(t, tm.Matches(reflect.TypeOf((**int)(nil))))
	assert.Equal(t, "*int", tm.String())

	ftm = FuncTypeMatch{typeMatch: tm, required: true}
	assert.Equal(t, "*int", ftm.String())
	ftm.required = false
	assert.Equal(t, "[*int]", ftm.String())

	tm = NewTypeMatch(0, 1, 2)
	assert.Equal(t, []reflect.Kind(nil), tm.kinds)
	assert.Equal(t, []reflect.Type{reflect.TypeOf(0)}, tm.types)
	assert.Equal(t, 1, tm.minIndirection)
	assert.Equal(t, 2, tm.maxIndirection)
	assert.False(t, tm.Matches(reflect.TypeOf(0)))
	assert.True(t, tm.Matches(reflect.TypeOf((*int)(nil))))
	assert.True(t, tm.Matches(reflect.TypeOf((**int)(nil))))
	assert.Equal(t, "(*|**)int", tm.String())

	ftm = FuncTypeMatch{typeMatch: tm, required: true}
	assert.Equal(t, "(*|**)int", ftm.String())
	ftm.required = false
	assert.Equal(t, "[(*|**)int]", ftm.String())

	type str struct{}

	tm = NewTypeMatch(reflect.Struct)
	assert.Equal(t, []reflect.Kind{reflect.Struct}, tm.kinds)
	assert.Equal(t, []reflect.Type(nil), tm.types)
	assert.Equal(t, 0, tm.minIndirection)
	assert.Equal(t, 0, tm.maxIndirection)
	assert.True(t, tm.Matches(reflect.TypeOf(str{})))
	assert.False(t, tm.Matches(reflect.TypeOf([]int{})))
	assert.Equal(t, "struct", tm.String())

	ftm = FuncTypeMatch{typeMatch: tm, required: true}
	assert.Equal(t, "struct", ftm.String())
	ftm.required = false
	assert.Equal(t, "[struct]", ftm.String())

	tm = NewTypeMatch(reflect.Slice, 1)
	assert.Equal(t, []reflect.Kind{reflect.Slice}, tm.kinds)
	assert.Equal(t, []reflect.Type(nil), tm.types)
	assert.Equal(t, 1, tm.minIndirection)
	assert.Equal(t, 1, tm.maxIndirection)
	assert.False(t, tm.Matches(reflect.TypeOf(str{})))
	assert.True(t, tm.Matches(reflect.TypeOf(&[]int{})))
	assert.Equal(t, "*slice", tm.String())

	ftm = FuncTypeMatch{typeMatch: tm, required: true}
	assert.Equal(t, "*slice", ftm.String())
	ftm.required = false
	assert.Equal(t, "[*slice]", ftm.String())

	tm = NewMultiTypeMatch(0, 1, 0, "")
	assert.Equal(t, []reflect.Kind(nil), tm.kinds)
	assert.Equal(t, []reflect.Type{reflect.TypeOf(0), reflect.TypeOf("")}, tm.types)
	assert.Equal(t, 0, tm.minIndirection)
	assert.Equal(t, 1, tm.maxIndirection)
	assert.True(t, tm.Matches(reflect.TypeOf(0)))
	assert.True(t, tm.Matches(reflect.TypeOf((*string)(nil))))
	assert.False(t, tm.Matches(reflect.TypeOf(str{})))
	assert.Equal(t, "[*](int|string)", tm.String())

	ftm = FuncTypeMatch{typeMatch: tm, required: true}
	assert.Equal(t, "[*](int|string)", ftm.String())
	ftm.required = false
	assert.Equal(t, "[[*](int|string)]", ftm.String())

	tm = NewMultiTypeMatch(0, 2, 0, "")
	assert.Equal(t, []reflect.Kind(nil), tm.kinds)
	assert.Equal(t, []reflect.Type{reflect.TypeOf(0), reflect.TypeOf("")}, tm.types)
	assert.Equal(t, 0, tm.minIndirection)
	assert.Equal(t, 2, tm.maxIndirection)
	assert.True(t, tm.Matches(reflect.TypeOf(0)))
	assert.True(t, tm.Matches(reflect.TypeOf((*string)(nil))))
	assert.True(t, tm.Matches(reflect.TypeOf((**string)(nil))))
	assert.False(t, tm.Matches(reflect.TypeOf(str{})))
	assert.Equal(t, "[*|**](int|string)", tm.String())

	ftm = FuncTypeMatch{typeMatch: tm, required: true}
	assert.Equal(t, "[*|**](int|string)", ftm.String())
	ftm.required = false
	assert.Equal(t, "[[*|**](int|string)]", ftm.String())

	tm = NewMultiTypeMatch(1, 2, 0, str{})
	assert.Equal(t, []reflect.Kind(nil), tm.kinds)
	assert.Equal(t, []reflect.Type{reflect.TypeOf(0), reflect.TypeOf(str{})}, tm.types)
	assert.Equal(t, 1, tm.minIndirection)
	assert.Equal(t, 2, tm.maxIndirection)
	assert.True(t, tm.Matches(reflect.TypeOf((*int)(nil))))
	assert.False(t, tm.Matches(reflect.TypeOf(0)))
	assert.True(t, tm.Matches(reflect.TypeOf((**str)(nil))))
	assert.Equal(t, "(*|**)(int|goreflect.str)", tm.String())

	ftm = FuncTypeMatch{typeMatch: tm, required: true}
	assert.Equal(t, "(*|**)(int|goreflect.str)", ftm.String())
	ftm.required = false
	assert.Equal(t, "[(*|**)(int|goreflect.str)]", ftm.String())
}

func TestFuncMatcher(t *testing.T) {
	// reflect.Type matching
	matcher := NewFuncMatcher()
	assert.True(t, matcher.Matches(func() {}))
	assert.False(t, matcher.Matches(func(int) {}))
	assert.False(t, matcher.Matches(func() int { return 0 }))
	paramIndexes, returnIndexes, matches := matcher.MatchingIndexes(func() {})
	assert.Equal(t, []int{}, paramIndexes)
	assert.Equal(t, []int{}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int) {})
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func() int { return 0 })
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	assert.Equal(t, "func()", matcher.String())

	matcher = NewFuncMatcher().
		WithParamType(0)
	assert.True(t, matcher.Matches(func(int) {}))
	assert.False(t, matcher.Matches(func(string) {}))
	assert.False(t, matcher.Matches(func(int) string { return "" }))
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int) {})
	assert.Equal(t, []int{0}, paramIndexes)
	assert.Equal(t, []int{}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(string) {})
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int) string { return "" })
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	assert.Equal(t, "func(int)", matcher.String())

	matcher = NewFuncMatcher().
		WithReturnType(0)
	assert.True(t, matcher.Matches(func() int { return 0 }))
	assert.False(t, matcher.Matches(func(int) {}))
	assert.False(t, matcher.Matches(func() string { return "" }))
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func() int { return 0 })
	assert.Equal(t, []int{}, paramIndexes)
	assert.Equal(t, []int{0}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int) {})
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func() string { return "" })
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	assert.Equal(t, "func() int", matcher.String())

	matcher = NewFuncMatcher().
		WithParamType(0).
		WithReturnType(0)
	assert.True(t, matcher.Matches(func(int) int { return 0 }))
	assert.False(t, matcher.Matches(func(string) int { return 0 }))
	assert.False(t, matcher.Matches(func(int) string { return "" }))
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int) int { return 0 })
	assert.Equal(t, []int{0}, paramIndexes)
	assert.Equal(t, []int{0}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(string) int { return 0 })
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int) string { return "" })
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	assert.Equal(t, "func(int) int", matcher.String())

	matcher = NewFuncMatcher().
		WithParamType(0).
		WithParamType("").
		WithReturnType(0)
	assert.True(t, matcher.Matches(func(int, string) int { return 0 }))
	assert.False(t, matcher.Matches(func(string, int) int { return 0 }))
	assert.False(t, matcher.Matches(func(int, string) string { return "" }))
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int, string) int { return 0 })
	assert.Equal(t, []int{0, 1}, paramIndexes)
	assert.Equal(t, []int{0}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(string, int) int { return 0 })
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int, string) string { return "" })
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	assert.Equal(t, "func(int, string) int", matcher.String())

	matcher = NewFuncMatcher().
		WithParamType(0).
		WithReturnType(0).
		WithReturnType("")
	assert.True(t, matcher.Matches(func(int) (int, string) { return 0, "" }))
	assert.False(t, matcher.Matches(func(string) (int, string) { return 0, "" }))
	assert.False(t, matcher.Matches(func(int) (string, int) { return "", 0 }))
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int) (int, string) { return 0, "" })
	assert.Equal(t, []int{0}, paramIndexes)
	assert.Equal(t, []int{0, 1}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(string) (int, string) { return 0, "" })
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int) (string, int) { return "", 0 })
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	assert.Equal(t, "func(int) (int, string)", matcher.String())

	// Optional params/return values
	matcher = NewFuncMatcher().
		WithOptionalParamType(0)
	assert.True(t, matcher.Matches(func() {}))
	assert.True(t, matcher.Matches(func(int) {}))
	assert.False(t, matcher.Matches(func(string) {}))
	assert.False(t, matcher.Matches(func() int { return 0 }))
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func() {})
	assert.Equal(t, []int{}, paramIndexes)
	assert.Equal(t, []int{}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int) {})
	assert.Equal(t, []int{0}, paramIndexes)
	assert.Equal(t, []int{}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(string) {})
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func() int { return 0 })
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	assert.Equal(t, "func([int])", matcher.String())

	matcher = NewFuncMatcher().
		WithOptionalParamType(0).
		WithOptionalReturnType(0)
	assert.True(t, matcher.Matches(func() {}))
	assert.True(t, matcher.Matches(func(int) {}))
	assert.True(t, matcher.Matches(func() int { return 0 }))
	assert.True(t, matcher.Matches(func(int) int { return 0 }))
	assert.False(t, matcher.Matches(func(string) {}))
	assert.False(t, matcher.Matches(func() string { return "" }))
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func() {})
	assert.Equal(t, []int{}, paramIndexes)
	assert.Equal(t, []int{}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int) {})
	assert.Equal(t, []int{0}, paramIndexes)
	assert.Equal(t, []int{}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func() int { return 0 })
	assert.Equal(t, []int{}, paramIndexes)
	assert.Equal(t, []int{0}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int) int { return 0 })
	assert.Equal(t, []int{0}, paramIndexes)
	assert.Equal(t, []int{0}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(string) {})
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func() string { return "" })
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	assert.Equal(t, "func([int]) [int]", matcher.String())

	matcher = NewFuncMatcher().
		WithParamType(0).
		WithOptionalParamType("")
	assert.True(t, matcher.Matches(func(int) {}))
	assert.True(t, matcher.Matches(func(int, string) {}))
	assert.False(t, matcher.Matches(func(string) {}))
	assert.False(t, matcher.Matches(func(string, int) {}))
	assert.False(t, matcher.Matches(func(string) int { return 0 }))
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int) {})
	assert.Equal(t, []int{0}, paramIndexes)
	assert.Equal(t, []int{}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int, string) {})
	assert.Equal(t, []int{0, 1}, paramIndexes)
	assert.Equal(t, []int{}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(string) {})
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(string, int) {})
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(string) int { return 0 })
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	assert.Equal(t, "func(int, [string])", matcher.String())

	matcher = NewFuncMatcher().
		WithOptionalParamType(0).
		WithParamType("")
	assert.False(t, matcher.Matches(func(int) {}))
	assert.True(t, matcher.Matches(func(int, string) {}))
	assert.True(t, matcher.Matches(func(string) {}))
	assert.False(t, matcher.Matches(func(string, int) {}))
	assert.False(t, matcher.Matches(func(string) int { return 0 }))
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int) {})
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int, string) {})
	assert.Equal(t, []int{0, 1}, paramIndexes)
	assert.Equal(t, []int{}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(string) {})
	assert.Equal(t, []int{0}, paramIndexes)
	assert.Equal(t, []int{}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(string, int) {})
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(string) int { return 0 })
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	assert.Equal(t, "func([int], string)", matcher.String())

	matcher = NewFuncMatcher().
		WithOptionalReturnType(0)
	assert.True(t, matcher.Matches(func() {}))
	assert.True(t, matcher.Matches(func() int { return 0 }))
	assert.False(t, matcher.Matches(func() string { return "" }))
	assert.False(t, matcher.Matches(func() (int, string) { return 0, "" }))
	assert.False(t, matcher.Matches(func() (string, int) { return "", 0 }))
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func() {})
	assert.Equal(t, []int{}, paramIndexes)
	assert.Equal(t, []int{}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func() int { return 0 })
	assert.Equal(t, []int{}, paramIndexes)
	assert.Equal(t, []int{0}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func() string { return "" })
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func() (int, string) { return 0, "" })
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func() (string, int) { return "", 0 })
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	assert.Equal(t, "func() [int]", matcher.String())

	// reflect.Kind matching
	type str struct{}

	matcher = NewFuncMatcher().
		WithParamType(0).
		WithOptionalParamType(reflect.Struct)
	assert.True(t, matcher.Matches(func(int) {}))
	assert.True(t, matcher.Matches(func(int, str) {}))
	assert.False(t, matcher.Matches(func(str) {}))
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int) {})
	assert.Equal(t, []int{0}, paramIndexes)
	assert.Equal(t, []int{}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int, str) {})
	assert.Equal(t, []int{0, 1}, paramIndexes)
	assert.Equal(t, []int{}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(str) {})
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	assert.Equal(t, "func(int, [struct])", matcher.String())

	matcher = NewFuncMatcher().
		WithParamOfTypes(0, 0, 0, "").
		WithReturnOfTypes(0, 0, "", reflect.Struct)
	assert.True(t, matcher.Matches(func(int) string { return "" }))
	assert.True(t, matcher.Matches(func(int) str { return str{} }))
	assert.True(t, matcher.Matches(func(string) string { return "" }))
	assert.True(t, matcher.Matches(func(string) str { return str{} }))
	assert.False(t, matcher.Matches(func(str) {}))
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int) string { return "" })
	assert.Equal(t, []int{0}, paramIndexes)
	assert.Equal(t, []int{0}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int) str { return str{} })
	assert.Equal(t, []int{0}, paramIndexes)
	assert.Equal(t, []int{0}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(string) string { return "" })
	assert.Equal(t, []int{0}, paramIndexes)
	assert.Equal(t, []int{0}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(string) str { return str{} })
	assert.Equal(t, []int{0}, paramIndexes)
	assert.Equal(t, []int{0}, returnIndexes)
	assert.True(t, matches)
	assert.Equal(t, "func(int|string) string|struct", matcher.String())

	matcher = NewFuncMatcher().
		WithOptionalParamOfTypes(0, 0, 0, "").
		WithOptionalReturnOfTypes(0, 1, "", reflect.Struct)
	assert.True(t, matcher.Matches(func() {}))
	assert.True(t, matcher.Matches(func(int) string { return "" }))
	assert.True(t, matcher.Matches(func(int) str { return str{} }))
	assert.True(t, matcher.Matches(func(string) string { return "" }))
	assert.True(t, matcher.Matches(func(string) *str { return &str{} }))
	assert.False(t, matcher.Matches(func(str) {}))
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func() {})
	assert.Equal(t, []int{}, paramIndexes)
	assert.Equal(t, []int{}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int) {})
	assert.Equal(t, []int{0}, paramIndexes)
	assert.Equal(t, []int{}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(string) {})
	assert.Equal(t, []int{0}, paramIndexes)
	assert.Equal(t, []int{}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func() string { return "" })
	assert.Equal(t, []int{}, paramIndexes)
	assert.Equal(t, []int{0}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func() str { return str{} })
	assert.Equal(t, []int{}, paramIndexes)
	assert.Equal(t, []int{0}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int) string { return "" })
	assert.Equal(t, []int{0}, paramIndexes)
	assert.Equal(t, []int{0}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int) str { return str{} })
	assert.Equal(t, []int{0}, paramIndexes)
	assert.Equal(t, []int{0}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(string) string { return "" })
	assert.Equal(t, []int{0}, paramIndexes)
	assert.Equal(t, []int{0}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(string) *str { return &str{} })
	assert.Equal(t, []int{0}, paramIndexes)
	assert.Equal(t, []int{0}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(str) {})
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	assert.Equal(t, "func([int|string]) [[*](string|struct)]", matcher.String())

	matcher = NewFuncMatcher().
		WithParams(NewFuncTypeMatch(0, true)).
		WithReturns(NewFuncTypeMatch("", false))
	assert.True(t, matcher.Matches(func(int) {}))
	assert.True(t, matcher.Matches(func(int) string { return "" }))
	assert.False(t, matcher.Matches(func() {}))
	assert.False(t, matcher.Matches(func() string { return "" }))
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int) {})
	assert.Equal(t, []int{0}, paramIndexes)
	assert.Equal(t, []int{}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func(int) string { return "" })
	assert.Equal(t, []int{0}, paramIndexes)
	assert.Equal(t, []int{0}, returnIndexes)
	assert.True(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func() {})
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	paramIndexes, returnIndexes, matches = matcher.MatchingIndexes(func() string { return "" })
	assert.Equal(t, []int(nil), paramIndexes)
	assert.Equal(t, []int(nil), returnIndexes)
	assert.False(t, matches)
	assert.Equal(t, "func(int) [string]", matcher.String())
}
