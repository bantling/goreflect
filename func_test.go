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
	assert.Equal(t, reflect.Int, tm.valueType.Kind())
	assert.Equal(t, 0, tm.minIndirection)
	assert.Equal(t, 0, tm.maxIndirection)
	assert.True(t, tm.Matches(reflect.TypeOf(0)))
	assert.False(t, tm.Matches(reflect.TypeOf((*int)(nil))))
	assert.False(t, tm.Matches(reflect.TypeOf((**int)(nil))))

	tm = NewTypeMatch(0, 1)
	assert.Equal(t, reflect.Int, tm.valueType.Kind())
	assert.Equal(t, 1, tm.minIndirection)
	assert.Equal(t, 1, tm.maxIndirection)
	assert.False(t, tm.Matches(reflect.TypeOf(0)))
	assert.True(t, tm.Matches(reflect.TypeOf((*int)(nil))))
	assert.False(t, tm.Matches(reflect.TypeOf((**int)(nil))))

	tm = NewTypeMatch(0, 1, 2)
	assert.Equal(t, reflect.Int, tm.valueType.Kind())
	assert.Equal(t, 1, tm.minIndirection)
	assert.Equal(t, 2, tm.maxIndirection)
	assert.False(t, tm.Matches(reflect.TypeOf(0)))
	assert.True(t, tm.Matches(reflect.TypeOf((*int)(nil))))
	assert.True(t, tm.Matches(reflect.TypeOf((**int)(nil))))
}

func TestFuncMatcher(t *testing.T) {
	// Required params/return values
	matcher := NewFuncMatcher()
	assert.True(t, matcher.Matches(func (){}))
	assert.False(t, matcher.Matches(func(int){}))
	assert.False(t, matcher.Matches(func() int {return 0}))

	matcher = NewFuncMatcher().
		WithParamType(0)
	assert.True(t, matcher.Matches(func (int){}))
	assert.False(t, matcher.Matches(func (string){}))
	assert.False(t, matcher.Matches(func (int) string {return ""}))

	matcher = NewFuncMatcher().
		WithReturnType(0)
	assert.True(t, matcher.Matches(func () int {return 0}))
	assert.False(t, matcher.Matches(func (int) {}))
	assert.False(t, matcher.Matches(func () string {return ""}))

	matcher = NewFuncMatcher().
		WithParamType(0).
		WithReturnType(0)
	assert.True(t, matcher.Matches(func (int) int {return 0}))
	assert.False(t, matcher.Matches(func (string) int {return 0}))
	assert.False(t, matcher.Matches(func (int) string {return ""}))

	matcher = NewFuncMatcher().
		WithParamType(0).
		WithParamType("").
		WithReturnType(0)
	assert.True(t, matcher.Matches(func (int, string) int {return 0}))
	assert.False(t, matcher.Matches(func (string, int) int {return 0}))
	assert.False(t, matcher.Matches(func (int, string) string {return ""}))

	matcher = NewFuncMatcher().
		WithParamType(0).
		WithReturnType(0).
		WithReturnType("")
	assert.True(t, matcher.Matches(func (int) (int, string) {return 0, ""}))
	assert.False(t, matcher.Matches(func (string) (int, string) {return 0, ""}))
	assert.False(t, matcher.Matches(func (int) (string, int) {return "", 0}))

	// Optional params/return values
	matcher = NewFuncMatcher().
		WithOptionalParamType(0)
	assert.True(t, matcher.Matches(func() {}))
	assert.True(t, matcher.Matches(func(int) {}))
	assert.False(t, matcher.Matches(func(string) {}))
	assert.False(t, matcher.Matches(func() int {return 0}))

	matcher = NewFuncMatcher().
		WithOptionalParamType(0).
		WithOptionalReturnType(0)
	assert.True(t, matcher.Matches(func() {}))
	assert.True(t, matcher.Matches(func(int) {}))
	assert.True(t, matcher.Matches(func() int {return 0}))
	assert.True(t, matcher.Matches(func(int) int {return 0}))
	assert.False(t, matcher.Matches(func(string) {}))
	assert.False(t, matcher.Matches(func() string {return ""}))

	matcher = NewFuncMatcher().
		WithParamType(0).
		WithOptionalParamType("")
	assert.True(t, matcher.Matches(func(int) {}))
	assert.True(t, matcher.Matches(func(int, string) {}))
	assert.False(t, matcher.Matches(func(string) {}))
	assert.False(t, matcher.Matches(func(string, int) {}))
	assert.False(t, matcher.Matches(func(string) int {return 0}))

	matcher = NewFuncMatcher().
		WithOptionalParamType(0).
		WithParamType("")
	assert.False(t, matcher.Matches(func(int) {}))
	assert.True(t, matcher.Matches(func(int, string) {}))
	assert.True(t, matcher.Matches(func(string) {}))
	assert.False(t, matcher.Matches(func(string, int) {}))
	assert.False(t, matcher.Matches(func(string) int {return 0}))

	matcher = NewFuncMatcher().
		WithOptionalReturnType(0)
	assert.True(t, matcher.Matches(func() {}))
	assert.True(t, matcher.Matches(func() int {return 0}))
	assert.False(t, matcher.Matches(func() string {return ""}))
	assert.False(t, matcher.Matches(func() (int, string) {return 0, ""}))
	assert.False(t, matcher.Matches(func() (string, int) {return "", 0}))
}
