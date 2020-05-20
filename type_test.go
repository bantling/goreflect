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

func TestGetReflectTypeOrKindValueOf(t *testing.T) {
	var (
		typ     reflect.Type
		kind    reflect.Kind
		strType = reflect.TypeOf((*string)(nil)).Elem()
	)

	// string kind

	typ, kind = GetReflectTypeOrKindValueOf(reflect.String)
	assert.Nil(t, typ)
	assert.Equal(t, reflect.String, kind)

	// string instance, value, type

	str := ""

	typ, kind = GetReflectTypeOrKindValueOf(str)
	assert.Equal(t, strType, typ)
	assert.Equal(t, reflect.Invalid, kind)

	typ, kind = GetReflectTypeOrKindValueOf(reflect.ValueOf(str))
	assert.Equal(t, strType, typ)
	assert.Equal(t, reflect.Invalid, kind)

	typ, kind = GetReflectTypeOrKindValueOf(reflect.TypeOf(str))
	assert.Equal(t, strType, typ)
	assert.Equal(t, reflect.Invalid, kind)

	// *string instance, value, type

	pstr := &str

	typ, kind = GetReflectTypeOrKindValueOf(pstr)
	assert.Equal(t, strType, typ)
	assert.Equal(t, reflect.Invalid, kind)

	typ, kind = GetReflectTypeOrKindValueOf(reflect.ValueOf(pstr))
	assert.Equal(t, strType, typ)
	assert.Equal(t, reflect.Invalid, kind)

	typ, kind = GetReflectTypeOrKindValueOf(reflect.TypeOf(pstr))
	assert.Equal(t, strType, typ)
	assert.Equal(t, reflect.Invalid, kind)

	// **string instance, value, type

	ppstr := &pstr

	typ, kind = GetReflectTypeOrKindValueOf(ppstr)
	assert.Equal(t, strType, typ)
	assert.Equal(t, reflect.Invalid, kind)

	typ, kind = GetReflectTypeOrKindValueOf(reflect.ValueOf(ppstr))
	assert.Equal(t, strType, typ)
	assert.Equal(t, reflect.Invalid, kind)

	typ, kind = GetReflectTypeOrKindValueOf(reflect.TypeOf(ppstr))
	assert.Equal(t, strType, typ)
	assert.Equal(t, reflect.Invalid, kind)

	// ***string instance, value, type

	pppstr := &ppstr

	typ, kind = GetReflectTypeOrKindValueOf(pppstr)
	assert.Equal(t, strType, typ)
	assert.Equal(t, reflect.Invalid, kind)

	typ, kind = GetReflectTypeOrKindValueOf(reflect.ValueOf(pppstr))
	assert.Equal(t, strType, typ)
	assert.Equal(t, reflect.Invalid, kind)

	typ, kind = GetReflectTypeOrKindValueOf(reflect.TypeOf(pppstr))
	assert.Equal(t, strType, typ)
	assert.Equal(t, reflect.Invalid, kind)
}
