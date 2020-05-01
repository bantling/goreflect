package goreflect

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValuePrinter(t *testing.T) {
	var (
		p   = &ValuePrinter{}
		w   = NewValueDepthFirstWalker(NewValueVisitorAdapter(p))
		pp  = &ValuePrinter{WithAddress: true}
		wp  = NewValueDepthFirstWalker(NewValueVisitorAdapter(pp))
		ci  chan int
		fn  = func(int) string { return "" }
		i   = 5
		ptr = &i
		arr = [2]int{3, 4}
		sl  = []int{5, 6, 7}
		mp  = map[int]string{8: "9"}
		st  = struct {
			Foo string
			Bar int
		}{
			Foo: "fooish",
			Bar: 10,
		}

		ptrptr = &ptr
		arrarr = [2][3]int{[3]int{11, 12, 13}, [3]int{14, 15, 16}}
		slsl   = [][]int{[]int{17, 18, 19}, []int{20, 21, 22}}
	)

	// Plain bool
	w.Walk(false)
	assert.Equal(t, "false", p.Result())
	wp.Walk(true)
	assert.Equal(t, "true", pp.Result())

	// reflect.Value wrapped bool
	w.Walk(reflect.ValueOf(true))
	assert.Equal(t, "true", p.Result())
	wp.Walk(reflect.ValueOf(false))
	assert.Equal(t, "false", pp.Result())

	// Int types
	w.Walk(int(1))
	assert.Equal(t, "1", p.Result())
	wp.Walk(int(2))
	assert.Equal(t, "2", pp.Result())

	w.Walk(int8(2))
	assert.Equal(t, "2", p.Result())
	wp.Walk(int8(3))
	assert.Equal(t, "3", pp.Result())

	w.Walk(int16(3))
	assert.Equal(t, "3", p.Result())
	wp.Walk(int16(4))
	assert.Equal(t, "4", pp.Result())

	w.Walk(int32(4))
	assert.Equal(t, "4", p.Result())
	wp.Walk(int32(5))
	assert.Equal(t, "5", pp.Result())

	w.Walk(int64(5))
	assert.Equal(t, "5", p.Result())
	wp.Walk(int64(6))
	assert.Equal(t, "6", pp.Result())

	// Uint types
	w.Walk(uint(10))
	assert.Equal(t, "10", p.Result())
	wp.Walk(uint(11))
	assert.Equal(t, "11", pp.Result())

	w.Walk(uint8(11))
	assert.Equal(t, "11", p.Result())
	wp.Walk(uint8(12))
	assert.Equal(t, "12", pp.Result())

	w.Walk(uint16(12))
	assert.Equal(t, "12", p.Result())
	wp.Walk(uint16(13))
	assert.Equal(t, "13", pp.Result())

	w.Walk(uint32(13))
	assert.Equal(t, "13", p.Result())
	wp.Walk(uint32(14))
	assert.Equal(t, "14", pp.Result())

	w.Walk(uint64(14))
	assert.Equal(t, "14", p.Result())
	wp.Walk(uint64(15))
	assert.Equal(t, "15", pp.Result())

	// Float types
	w.Walk(float32(1.25))
	assert.Equal(t, "1.25", p.Result())
	wp.Walk(float32(1.5))
	assert.Equal(t, "1.5", pp.Result())

	w.Walk(float64(2.5))
	assert.Equal(t, "2.5", p.Result())
	wp.Walk(float64(2.75))
	assert.Equal(t, "2.75", pp.Result())

	// Complex types
	w.Walk(complex64((1 + 2i)))
	assert.Equal(t, "(1+2i)", p.Result())
	wp.Walk(complex64((2 + 4i)))
	assert.Equal(t, "(2+4i)", pp.Result())

	w.Walk(complex128((3 + 4i)))
	assert.Equal(t, "(3+4i)", p.Result())
	wp.Walk(complex128((6 + 8i)))
	assert.Equal(t, "(6+8i)", pp.Result())

	// String
	w.Walk("foo")
	assert.Equal(t, `"foo"`, p.Result())
	wp.Walk("bar")
	assert.Equal(t, `"bar"`, pp.Result())

	// Chan
	w.Walk(ci)
	assert.Equal(t, "chan int", p.Result())
	wp.Walk(ci)
	assert.Equal(t, fmt.Sprintf("chan int @[%p]", ci), pp.Result())

	// Func
	w.Walk(fn)
	assert.Equal(t, "func(int) string", p.Result())
	wp.Walk(fn)
	assert.Equal(t, fmt.Sprintf("func(int) string @[%p]", fn), pp.Result())

	// Ptr
	w.Walk(ptr)
	assert.Equal(t, "&5", p.Result())
	wp.Walk(ptr)
	assert.Equal(t, fmt.Sprintf("&@[%p]5", ptr), pp.Result())

	// Array
	w.Walk(arr)
	assert.Equal(t, "[2]int{3, 4}", p.Result())
	wp.Walk(arr)
	assert.Equal(t, "[2]int{3, 4}", pp.Result())

	// Slice
	w.Walk(sl)
	assert.Equal(t, "[]int{5, 6, 7}", p.Result())
	wp.Walk(sl)
	assert.Equal(t, fmt.Sprintf("[]int@[%p]{5, 6, 7}", sl), pp.Result())

	// Map
	w.Walk(mp)
	assert.Equal(t, `map[int]string{8: "9"}`, p.Result())
	wp.Walk(mp)
	assert.Equal(t, fmt.Sprintf(`map[int]string@[%p]{8: "9"}`, mp), pp.Result())

	// Struct
	w.Walk(st)
	assert.Equal(t, `struct { Foo string; Bar int }{Foo: "fooish", Bar: 10}`, p.Result())
	wp.Walk(st)
	assert.Equal(t, `struct { Foo string; Bar int }{Foo: "fooish", Bar: 10}`, pp.Result())

	// PtrPtr
	w.Walk(ptrptr)
	assert.Equal(t, "&&5", p.Result())
	wp.Walk(ptrptr)
	assert.Equal(t, fmt.Sprintf("&@[%p]&@[%p]5", ptrptr, ptr), pp.Result())

	// ArrayArray
	w.Walk(arrarr)
	assert.Equal(t, "[2][3]int{[3]int{11, 12, 13}, [3]int{14, 15, 16}}", p.Result())
	wp.Walk(arrarr)
	assert.Equal(t, "[2][3]int{[3]int{11, 12, 13}, [3]int{14, 15, 16}}", p.Result())

	// SliceSlice
	w.Walk(slsl)
	assert.Equal(t, "[][]int{[]int{17, 18, 19}, []int{20, 21, 22}}", p.Result())
	wp.Walk(slsl)
	assert.Equal(t, fmt.Sprintf("[][]int@[%p]{[]int@[%p]{17, 18, 19}, []int@[%p]{20, 21, 22}}", slsl, slsl[0], slsl[1]), pp.Result())
}
