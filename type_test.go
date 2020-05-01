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

func TestGetReflectKindOrTypeValueOf(t *testing.T) {
	var (
		kind    reflect.Kind
		typ     reflect.Type
		strType = reflect.TypeOf((*string)(nil)).Elem()
	)

	// string kind

	kind, typ = GetReflectKindOrTypeValueOf(reflect.String)
	assert.Equal(t, reflect.String, kind)
	assert.Nil(t, typ)

	// string instance, value, type

	str := ""

	kind, typ = GetReflectKindOrTypeValueOf(str)
	assert.Equal(t, reflect.Invalid, kind)
	assert.Equal(t, strType, typ)

	kind, typ = GetReflectKindOrTypeValueOf(reflect.ValueOf(str))
	assert.Equal(t, reflect.Invalid, kind)
	assert.Equal(t, strType, typ)

	kind, typ = GetReflectKindOrTypeValueOf(reflect.TypeOf(str))
	assert.Equal(t, reflect.Invalid, kind)
	assert.Equal(t, strType, typ)

	// *string instance, value, type

	pstr := &str

	kind, typ = GetReflectKindOrTypeValueOf(pstr)
	assert.Equal(t, reflect.Invalid, kind)
	assert.Equal(t, strType, typ)

	kind, typ = GetReflectKindOrTypeValueOf(reflect.ValueOf(pstr))
	assert.Equal(t, reflect.Invalid, kind)
	assert.Equal(t, strType, typ)

	kind, typ = GetReflectKindOrTypeValueOf(reflect.TypeOf(pstr))
	assert.Equal(t, reflect.Invalid, kind)
	assert.Equal(t, strType, typ)

	// **string instance, value, type

	ppstr := &pstr

	kind, typ = GetReflectKindOrTypeValueOf(ppstr)
	assert.Equal(t, reflect.Invalid, kind)
	assert.Equal(t, strType, typ)

	kind, typ = GetReflectKindOrTypeValueOf(reflect.ValueOf(ppstr))
	assert.Equal(t, reflect.Invalid, kind)
	assert.Equal(t, strType, typ)

	kind, typ = GetReflectKindOrTypeValueOf(reflect.TypeOf(ppstr))
	assert.Equal(t, reflect.Invalid, kind)
	assert.Equal(t, strType, typ)
}
