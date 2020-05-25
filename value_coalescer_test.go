package goreflect

import (
	//	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValueCoalescer(t *testing.T) {
	var (
		methodNames []string
		methodArgs  [][]interface{}
		d           = NewValueVisitorReducer(func(m string, a []reflect.Value) {
			methodNames = append(methodNames, m)
			var args []interface{}
			for _, arg := range a {
				args = append(args, arg.Interface())
			}
			methodArgs = append(methodArgs, args)
		},
		)
		c1 = NewValueCoalescer(d)
		w1 = NewValueDepthFirstWalker(NewValueVisitorAdapter(c1))
		c2 = NewValueCoalescer(d).
			WithIntCoalesceMode(IntsToString).
			WithUintCoalesceMode(UintsToString).
			WithFloatCoalesceMode(FloatsToString).
			WithComplexCoalesceMode(ComplexesToString).
			WithArrayCoalesceMode(ArraysAsIs)
		w2  = NewValueDepthFirstWalker(NewValueVisitorAdapter(c2))
		val interface{}
	)

	clear := func() {
		methodNames = []string{}
		methodArgs = [][]interface{}{}
	}

	assertOneCall := func(name string, args ...interface{}) {
		assert.Equal(t, []string{name}, methodNames)
		assert.Equal(t, 1, len(methodArgs))
		assert.Equal(t, len(args), len(methodArgs[0]))
		for i, arg := range args {
			assert.Equal(t, arg, methodArgs[0][i])
		}
	}

	// int
	clear()
	val = 0
	w1.Walk(val)
	assertOneCall("VisitInt", val)

	clear()
	w2.Walk(val)
	assertOneCall("VisitString", "0")

	// int8
	val = int8(1)
	clear()
	w1.Walk(val)
	assertOneCall("VisitInt", 1)

	clear()
	w2.Walk(val)
	assertOneCall("VisitString", "1")

	// int16
	val = int16(2)
	clear()
	w1.Walk(val)
	assertOneCall("VisitInt", 2)

	clear()
	w2.Walk(val)
	assertOneCall("VisitString", "2")

	// int32
	val = int32(3)
	clear()
	w1.Walk(val)
	assertOneCall("VisitInt", 3)

	clear()
	w2.Walk(val)
	assertOneCall("VisitString", "3")

	// int64
	val = int64(4)
	clear()
	w1.Walk(val)
	assertOneCall("VisitInt", 4)

	clear()
	w2.Walk(val)
	assertOneCall("VisitString", "4")

	// uint
	clear()
	val = uint(5)
	w1.Walk(val)
	assertOneCall("VisitInt", 5)

	clear()
	w2.Walk(val)
	assertOneCall("VisitString", "5")

	// uint8
	val = uint8(6)
	clear()
	w1.Walk(val)
	assertOneCall("VisitInt", 6)

	clear()
	w2.Walk(val)
	assertOneCall("VisitString", "6")

	// uint16
	val = uint16(7)
	clear()
	w1.Walk(val)
	assertOneCall("VisitInt", 7)

	clear()
	w2.Walk(val)
	assertOneCall("VisitString", "7")

	// uint32
	val = uint32(8)
	clear()
	w1.Walk(val)
	assertOneCall("VisitInt", 8)

	clear()
	w2.Walk(val)
	assertOneCall("VisitString", "8")

	// uint64
	val = uint64(9)
	clear()
	w1.Walk(val)
	assertOneCall("VisitInt", 9)

	clear()
	w2.Walk(val)
	assertOneCall("VisitString", "9")

	// float32
	val = float32(10.25)
	clear()
	w1.Walk(val)
	assertOneCall("VisitFloat64", 10.25)

	clear()
	w2.Walk(val)
	assertOneCall("VisitString", "10.25")

	// float64
	val = 10.5
	clear()
	w1.Walk(val)
	assertOneCall("VisitFloat64", 10.5)

	clear()
	w2.Walk(val)
	assertOneCall("VisitString", "10.5")

	// complex64
	val = complex64(11 + 12i)
	clear()
	w1.Walk(val)
	assertOneCall("VisitComplex128", (11 + 12i))

	clear()
	w2.Walk(val)
	assertOneCall("VisitString", "(11+12i)")

	// complex128
	val = 13 + 14i
	clear()
	w1.Walk(val)
	assertOneCall("VisitComplex128", (13 + 14i))

	clear()
	w2.Walk(val)
	assertOneCall("VisitString", "(13+14i)")

	// array
	val = [2]int{15, 16}
	clear()
	w1.Walk(val)
	assert.Equal(
		t,
		[]string{
			"VisitPreSlice",
			"VisitPreSliceIndex",
			"VisitInt",
			"VisitPostSliceIndex",
			"VisitPreSliceIndex",
			"VisitInt",
			"VisitPostSliceIndex",
			"VisitPostSlice",
		},
		methodNames,
	)
	assert.Equal(
		t,
		[][]interface{}{
			{2, val},
			{2, 0, 15},
			{15},
			{2, 0, 15},
			{2, 1, 16},
			{16},
			{2, 1, 16},
			{2, val},
		},
		methodArgs,
	)

	clear()
	w2.Walk(val)
	assert.Equal(
		t,
		[]string{
			"VisitPreArray",
			"VisitPreArrayIndex",
			"VisitString",
			"VisitPostArrayIndex",
			"VisitPreArrayIndex",
			"VisitString",
			"VisitPostArrayIndex",
			"VisitPostArray",
		},
		methodNames,
	)
	assert.Equal(
		t,
		[][]interface{}{
			{2, val},
			{2, 0, 15},
			{"15"},
			{2, 0, 15},
			{2, 1, 16},
			{"16"},
			{2, 1, 16},
			{2, val},
		},
		methodArgs,
	)
}
