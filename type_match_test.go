package goreflect

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeMatch(t *testing.T) {
	tm := NewTypeMatch(0)
	assert.Equal(t, []reflect.Kind(nil), tm.kinds)
	assert.Equal(t, []reflect.Type{reflect.TypeOf(0)}, tm.types)
	assert.Equal(t, Value, tm.minIndirection)
	assert.Equal(t, Value, tm.maxIndirection)
	assert.True(t, tm.Matches(reflect.TypeOf(0)))
	assert.False(t, tm.Matches(reflect.TypeOf((*int)(nil))))
	assert.False(t, tm.Matches(reflect.TypeOf((**int)(nil))))
	assert.Equal(t, "int", tm.String())

	tm = NewTypeMatch(0, Ptr)
	assert.Equal(t, []reflect.Kind(nil), tm.kinds)
	assert.Equal(t, []reflect.Type{reflect.TypeOf(0)}, tm.types)
	assert.Equal(t, Ptr, tm.minIndirection)
	assert.Equal(t, Ptr, tm.maxIndirection)
	assert.False(t, tm.Matches(reflect.TypeOf(0)))
	assert.True(t, tm.Matches(reflect.TypeOf((*int)(nil))))
	assert.False(t, tm.Matches(reflect.TypeOf((**int)(nil))))
	assert.Equal(t, "*int", tm.String())

	tm = NewTypeMatch(0, Ptr, PtrPtr)
	assert.Equal(t, []reflect.Kind(nil), tm.kinds)
	assert.Equal(t, []reflect.Type{reflect.TypeOf(0)}, tm.types)
	assert.Equal(t, Ptr, tm.minIndirection)
	assert.Equal(t, PtrPtr, tm.maxIndirection)
	assert.False(t, tm.Matches(reflect.TypeOf(0)))
	assert.True(t, tm.Matches(reflect.TypeOf((*int)(nil))))
	assert.True(t, tm.Matches(reflect.TypeOf((**int)(nil))))
	assert.Equal(t, "(*|**)int", tm.String())

	type str struct{}

	tm = NewTypeMatch(reflect.Struct)
	assert.Equal(t, []reflect.Kind{reflect.Struct}, tm.kinds)
	assert.Equal(t, []reflect.Type(nil), tm.types)
	assert.Equal(t, Value, tm.minIndirection)
	assert.Equal(t, Value, tm.maxIndirection)
	assert.True(t, tm.Matches(reflect.TypeOf(str{})))
	assert.False(t, tm.Matches(reflect.TypeOf(map[int]bool{})))
	assert.Equal(t, "struct", tm.String())

	tm = NewTypeMatch(reflect.Slice, Ptr)
	assert.Equal(t, []reflect.Kind{reflect.Slice}, tm.kinds)
	assert.Equal(t, []reflect.Type(nil), tm.types)
	assert.Equal(t, Ptr, tm.minIndirection)
	assert.Equal(t, Ptr, tm.maxIndirection)
	assert.False(t, tm.Matches(reflect.TypeOf(str{})))
	assert.True(t, tm.Matches(reflect.TypeOf(&[]int{})))
	assert.Equal(t, "*slice", tm.String())

	tm = NewMultiTypeMatch(Value, Ptr, 0, "")
	assert.Equal(t, []reflect.Kind(nil), tm.kinds)
	assert.Equal(t, []reflect.Type{reflect.TypeOf(0), reflect.TypeOf("")}, tm.types)
	assert.Equal(t, Value, tm.minIndirection)
	assert.Equal(t, Ptr, tm.maxIndirection)
	assert.True(t, tm.Matches(reflect.TypeOf(0)))
	assert.True(t, tm.Matches(reflect.TypeOf((*string)(nil))))
	assert.False(t, tm.Matches(reflect.TypeOf(str{})))
	assert.Equal(t, "[*](int|string)", tm.String())

	tm = NewMultiTypeMatch(Value, PtrPtr, 0, "")
	assert.Equal(t, []reflect.Kind(nil), tm.kinds)
	assert.Equal(t, []reflect.Type{reflect.TypeOf(0), reflect.TypeOf("")}, tm.types)
	assert.Equal(t, Value, tm.minIndirection)
	assert.Equal(t, PtrPtr, tm.maxIndirection)
	assert.True(t, tm.Matches(reflect.TypeOf(0)))
	assert.True(t, tm.Matches(reflect.TypeOf((*string)(nil))))
	assert.True(t, tm.Matches(reflect.TypeOf((**string)(nil))))
	assert.False(t, tm.Matches(reflect.TypeOf(str{})))
	assert.Equal(t, "[*|**](int|string)", tm.String())

	tm = NewMultiTypeMatch(Ptr, PtrPtr, 0, str{})
	assert.Equal(t, []reflect.Kind(nil), tm.kinds)
	assert.Equal(t, []reflect.Type{reflect.TypeOf(0), reflect.TypeOf(str{})}, tm.types)
	assert.Equal(t, Ptr, tm.minIndirection)
	assert.Equal(t, PtrPtr, tm.maxIndirection)
	assert.True(t, tm.Matches(reflect.TypeOf((*int)(nil))))
	assert.False(t, tm.Matches(reflect.TypeOf(0)))
	assert.True(t, tm.Matches(reflect.TypeOf((**str)(nil))))
	assert.Equal(t, "(*|**)(int|goreflect.str)", tm.String())
}
