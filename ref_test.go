package goreflect

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReflectValueOf(t *testing.T) {
	assert.Equal(t, 1, GetReflectValueOf(1).Interface())
	assert.Equal(t, 1, GetReflectValueOf(reflect.ValueOf(1)).Interface())
}

func TestDerefdReflectValue(t *testing.T) {
	assert.Equal(t, 1, DerefdReflectValue(reflect.ValueOf(1)).Interface())

	i := 1
	assert.Equal(t, 1, DerefdReflectValue(reflect.ValueOf(&i)).Interface())

	p := &i
	assert.Equal(t, 1, DerefdReflectValue(reflect.ValueOf(&p)).Interface())

	pp := &p
	assert.Equal(t, 1, DerefdReflectValue(reflect.ValueOf(&pp)).Interface())
}

func TestDerefdReflectType(t *testing.T) {
	derefdType := reflect.TypeOf(0)
	assert.Equal(t, derefdType, DerefdReflectType(reflect.TypeOf(0)))
	assert.Equal(t, derefdType, DerefdReflectType(reflect.TypeOf((*int)(nil))))
	assert.Equal(t, derefdType, DerefdReflectType(reflect.TypeOf((**int)(nil))))
	assert.Equal(t, derefdType, DerefdReflectType(reflect.TypeOf((***int)(nil))))
}

func TestFullyDerefdReflectType(t *testing.T) {
	// T, *T, ... => T
	assert.Equal(t, reflect.TypeOf(""), FullyDerefdReflectType(reflect.TypeOf("")))
	assert.Equal(t, reflect.TypeOf(""), FullyDerefdReflectType(reflect.TypeOf((*string)(nil))))
	assert.Equal(t, reflect.TypeOf(""), FullyDerefdReflectType(reflect.TypeOf((**string)(nil))))
	assert.Equal(t, reflect.TypeOf(""), FullyDerefdReflectType(reflect.TypeOf((***string)(nil))))

	// [1]T, *[1]T, [1]*T, ... => [1]T
	assert.Equal(t, reflect.TypeOf([1]string{}), FullyDerefdReflectType(reflect.TypeOf([1]string{})))
	assert.Equal(t, reflect.TypeOf([1]string{}), FullyDerefdReflectType(reflect.TypeOf((*[1]string)(nil))))
	assert.Equal(t, reflect.TypeOf([1]string{}), FullyDerefdReflectType(reflect.TypeOf([1]*string{})))

	// []T, *[]T, []*T, ... => []T
	assert.Equal(t, reflect.TypeOf(([]string)(nil)), FullyDerefdReflectType(reflect.TypeOf(([]string)(nil))))
	assert.Equal(t, reflect.TypeOf(([]string)(nil)), FullyDerefdReflectType(reflect.TypeOf((*[]string)(nil))))
	assert.Equal(t, reflect.TypeOf(([]string)(nil)), FullyDerefdReflectType(reflect.TypeOf(([]*string)(nil))))

	// chan T, *chan T, chan *T, ... => chan T
	assert.Equal(t, reflect.TypeOf((chan string)(nil)), FullyDerefdReflectType(reflect.TypeOf((chan string)(nil))))
	assert.Equal(t, reflect.TypeOf((chan string)(nil)), FullyDerefdReflectType(reflect.TypeOf((*chan string)(nil))))
	assert.Equal(t, reflect.TypeOf((chan string)(nil)), FullyDerefdReflectType(reflect.TypeOf((chan *string)(nil))))

	// map[K]V, *map[K]V, map[*K]V, map[K]*V, ... => map[K]V
	assert.Equal(t, reflect.TypeOf((map[string]int)(nil)), FullyDerefdReflectType(reflect.TypeOf((map[string]int)(nil))))
	assert.Equal(t, reflect.TypeOf((map[string]int)(nil)), FullyDerefdReflectType(reflect.TypeOf((*map[string]int)(nil))))
	assert.Equal(t, reflect.TypeOf((map[string]int)(nil)), FullyDerefdReflectType(reflect.TypeOf((map[*string]int)(nil))))
	assert.Equal(t, reflect.TypeOf((map[string]int)(nil)), FullyDerefdReflectType(reflect.TypeOf((map[string]*int)(nil))))
}

func TestNumRefs(t *testing.T) {
	i := 1
	ptr := &i
	ptrptr := &ptr
	assert.Equal(t, 0, NumRefs(reflect.TypeOf(i)))
	assert.Equal(t, 1, NumRefs(reflect.TypeOf(ptr)))
	assert.Equal(t, 2, NumRefs(reflect.TypeOf(ptrptr)))
}

func TestCreateRefs(t *testing.T) {
	assert.Equal(t, 10, CreateRefs(reflect.ValueOf(10), 0).Interface())
	assert.Equal(t, 11, CreateRefs(reflect.ValueOf(11), 1).Elem().Interface())
	assert.Equal(t, 12, CreateRefs(reflect.ValueOf(12), 2).Elem().Elem().Interface())
}
