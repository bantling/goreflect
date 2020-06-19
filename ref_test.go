package goreflect

import (
	"fmt"
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
	assert.Equal(t, Value, NumRefs(reflect.TypeOf(i)))
	assert.Equal(t, Ptr, NumRefs(reflect.TypeOf(ptr)))
	assert.Equal(t, PtrPtr, NumRefs(reflect.TypeOf(ptrptr)))
}

func TestCreateRefs(t *testing.T) {
	assert.Equal(t, 10, CreateRefs(reflect.ValueOf(10), Value).Interface())
	assert.Equal(t, 11, CreateRefs(reflect.ValueOf(11), Ptr).Elem().Interface())
	assert.Equal(t, 12, CreateRefs(reflect.ValueOf(12), PtrPtr).Elem().Elem().Interface())
}

func TestFullyDerefdValue(t *testing.T) {
	//// Scalars
	assert.Equal(t, 1, FullyDerefdValue(1))
	assert.Equal(t, int8(2), FullyDerefdValue(int8(2)))
	assert.Equal(t, int16(3), FullyDerefdValue(int16(3)))
	assert.Equal(t, int32(4), FullyDerefdValue(int32(4)))
	assert.Equal(t, int64(5), FullyDerefdValue(int64(5)))
	assert.Equal(t, uint(6), FullyDerefdValue(uint(6)))
	assert.Equal(t, uint8(7), FullyDerefdValue(uint8(7)))
	assert.Equal(t, uint16(8), FullyDerefdValue(uint16(8)))
	assert.Equal(t, uint32(9), FullyDerefdValue(uint32(9)))
	assert.Equal(t, uint64(10), FullyDerefdValue(uint64(10)))
	assert.Equal(t, float32(11), FullyDerefdValue(float32(11)))
	assert.Equal(t, float64(12), FullyDerefdValue(float64(12)))
	assert.Equal(t, complex64((13 + 14i)), FullyDerefdValue(complex64((13 + 14i))))
	assert.Equal(t, complex128((15 + 16i)), FullyDerefdValue(complex128((15 + 16i))))
	assert.Equal(t, "17", FullyDerefdValue("17"))

	c := make(chan int)
	assert.Equal(t, c, FullyDerefdValue(c))

	f := func() {}
	assert.Equal(t, fmt.Sprintf("%p", f), fmt.Sprintf("%p", FullyDerefdValue(f)))

	//// Pointer
	i := 1
	assert.Equal(t, i, FullyDerefdValue(&i))

	//// Array of scalars
	assert.Equal(t, [1]int{1}, FullyDerefdValue([1]int{i}))
	assert.Equal(t, [1]int{1}, FullyDerefdValue(&[1]int{i}))
	assert.Equal(t, [1]int{1}, FullyDerefdValue([1]*int{&i}))
	assert.Equal(t, [1]int{1}, FullyDerefdValue(&[1]*int{&i}))

	//// Slice of scalars
	assert.Equal(t, []int{1}, FullyDerefdValue([]int{i}))
	assert.Equal(t, []int{1}, FullyDerefdValue(&[]int{i}))
	assert.Equal(t, []int{1}, FullyDerefdValue([]*int{&i}))
	assert.Equal(t, []int{1}, FullyDerefdValue(&[]*int{&i}))

	//// Map of scalars
	str := "1"
	assert.Equal(t, map[int]string{1: "1"}, FullyDerefdValue(map[int]string{i: str}))
	assert.Equal(t, map[int]string{1: "1"}, FullyDerefdValue(&map[int]string{i: str}))
	assert.Equal(t, map[int]string{1: "1"}, FullyDerefdValue(map[*int]string{&i: str}))
	assert.Equal(t, map[int]string{1: "1"}, FullyDerefdValue(map[int]*string{i: &str}))
	assert.Equal(t, map[int]string{1: "1"}, FullyDerefdValue(&map[*int]*string{&i: &str}))

	//// Array of array
	assert.Equal(t, [1][1]int{{1}}, FullyDerefdValue([1][1]int{{i}}))
	assert.Equal(t, [1][1]int{{1}}, FullyDerefdValue(&[1][1]int{{i}}))
	assert.Equal(t, [1][1]int{{1}}, FullyDerefdValue([1][1]*int{{&i}}))
	assert.Equal(t, [1][1]int{{1}}, FullyDerefdValue(&[1]*[1]*int{{&i}}))

	//// Array of slice
	assert.Equal(t, [1][]int{{1}}, FullyDerefdValue([1][]int{{i}}))
	assert.Equal(t, [1][]int{{1}}, FullyDerefdValue(&[1][]int{{i}}))
	assert.Equal(t, [1][]int{{1}}, FullyDerefdValue([1][]*int{{&i}}))
	assert.Equal(t, [1][]int{{1}}, FullyDerefdValue(&[1]*[]*int{{&i}}))

	//// Array of map
	assert.Equal(t, [1]map[int]string{{i: str}}, FullyDerefdValue([1]map[int]string{{i: str}}))
	assert.Equal(t, [1]map[int]string{{i: str}}, FullyDerefdValue(&[1]map[int]string{{i: str}}))
	assert.Equal(t, [1]map[int]string{{i: str}}, FullyDerefdValue([1]*map[int]string{{i: str}}))
	assert.Equal(t, [1]map[int]string{{i: str}}, FullyDerefdValue([1]map[*int]string{{&i: str}}))
	assert.Equal(t, [1]map[int]string{{i: str}}, FullyDerefdValue([1]map[int]*string{{i: &str}}))
	assert.Equal(t, [1]map[int]string{{i: str}}, FullyDerefdValue(&[1]*map[*int]*string{{&i: &str}}))

	//// Slice of array
	assert.Equal(t, [][1]int{{1}}, FullyDerefdValue([][1]int{{i}}))
	assert.Equal(t, [][1]int{{1}}, FullyDerefdValue(&[][1]int{{i}}))
	assert.Equal(t, [][1]int{{1}}, FullyDerefdValue([][1]*int{{&i}}))
	assert.Equal(t, [][1]int{{1}}, FullyDerefdValue(&[]*[1]*int{{&i}}))

	//// Slice of slice
	assert.Equal(t, [][]int{{1}}, FullyDerefdValue([][]int{{i}}))
	assert.Equal(t, [][]int{{1}}, FullyDerefdValue(&[][]int{{i}}))
	assert.Equal(t, [][]int{{1}}, FullyDerefdValue([][]*int{{&i}}))
	assert.Equal(t, [][]int{{1}}, FullyDerefdValue(&[]*[]*int{{&i}}))

	//// Slice of map
	assert.Equal(t, []map[int]string{{i: str}}, FullyDerefdValue([]map[int]string{{i: str}}))
	assert.Equal(t, []map[int]string{{i: str}}, FullyDerefdValue(&[]map[int]string{{i: str}}))
	assert.Equal(t, []map[int]string{{i: str}}, FullyDerefdValue([]*map[int]string{{i: str}}))
	assert.Equal(t, []map[int]string{{i: str}}, FullyDerefdValue([]map[*int]string{{&i: str}}))
	assert.Equal(t, []map[int]string{{i: str}}, FullyDerefdValue([]map[int]*string{{i: &str}}))
	assert.Equal(t, []map[int]string{{i: str}}, FullyDerefdValue(&[]*map[*int]*string{{&i: &str}}))

	//// Map of array
	j := 2
	assert.Equal(t, map[int][1]int{1: {2}}, FullyDerefdValue(map[int][1]int{i: {j}}))
	assert.Equal(t, map[int][1]int{1: {2}}, FullyDerefdValue(&map[int][1]int{i: {j}}))
	assert.Equal(t, map[int][1]int{1: {2}}, FullyDerefdValue(map[*int][1]int{&i: {j}}))
	assert.Equal(t, map[int][1]int{1: {2}}, FullyDerefdValue(map[int]*[1]int{i: {j}}))
	assert.Equal(t, map[int][1]int{1: {2}}, FullyDerefdValue(map[int][1]*int{i: {&j}}))
	assert.Equal(t, map[int][1]int{1: {2}}, FullyDerefdValue(&map[*int]*[1]*int{&i: {&j}}))

	//// Map of slice
	assert.Equal(t, map[int][]int{1: {2}}, FullyDerefdValue(map[int][]int{i: {j}}))
	assert.Equal(t, map[int][]int{1: {2}}, FullyDerefdValue(&map[int][]int{i: {j}}))
	assert.Equal(t, map[int][]int{1: {2}}, FullyDerefdValue(map[*int][]int{&i: {j}}))
	assert.Equal(t, map[int][]int{1: {2}}, FullyDerefdValue(map[int]*[]int{i: {j}}))
	assert.Equal(t, map[int][]int{1: {2}}, FullyDerefdValue(map[int][]*int{i: {&j}}))
	assert.Equal(t, map[int][]int{1: {2}}, FullyDerefdValue(&map[*int]*[]*int{&i: {&j}}))

	//// Map of map
	jstr := "2"
	assert.Equal(t, map[int]map[int]string{1: {2: "2"}}, FullyDerefdValue(map[int]map[int]string{i: {j: jstr}}))
	assert.Equal(t, map[int]map[int]string{1: {2: "2"}}, FullyDerefdValue(&map[int]map[int]string{i: {j: jstr}}))
	assert.Equal(t, map[int]map[int]string{1: {2: "2"}}, FullyDerefdValue(map[*int]map[int]string{&i: {j: jstr}}))
	assert.Equal(t, map[int]map[int]string{1: {2: "2"}}, FullyDerefdValue(map[int]*map[int]string{i: {j: jstr}}))
	assert.Equal(t, map[int]map[int]string{1: {2: "2"}}, FullyDerefdValue(map[int]map[*int]string{i: {&j: jstr}}))
	assert.Equal(t, map[int]map[int]string{1: {2: "2"}}, FullyDerefdValue(map[int]map[int]*string{i: {j: &jstr}}))
	assert.Equal(t, map[int]map[int]string{1: {2: "2"}}, FullyDerefdValue(&map[*int]*map[*int]*string{&i: {&j: &jstr}}))
}
