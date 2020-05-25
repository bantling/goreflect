package goreflect

import (
	"fmt"
	"reflect"
	"strconv"
)

// IntCoalesceMode is an enum of int coalesce modes
type IntCoalesceMode uint

// Int coalesce modes
const (
	IntsToInt = iota
	IntsToInt64
	IntsToUint
	IntsToUint64
	IntsToString
	IntsAsIs
)

// UintCoalesceMode is an enum of uint coalesce modes
type UintCoalesceMode uint

// Uint coalesce modes
const (
	UintsToInt = iota
	UintsToInt64
	UintsToUint
	UintsToUint64
	UintsToString
	UintsAsIs
)

// FloatCoalesceMode is an enum of float coalesce modes
type FloatCoalesceMode uint

// Float coalesce modes
const (
	FloatsToFloat64 = iota
	FloatsToFloat32
	FloatsToString
	FloatsAsIs
)

// ComplexCoalesceMode is an enum of complex coalesce modes
type ComplexCoalesceMode uint

// Complex coalesce modes
const (
	ComplexesToComplex128 = iota
	ComplexesToComplex64
	ComplexesToString
	ComplexesAsIs
)

// ArrayCoalesceMode is an enum of array coalesce modes
type ArrayCoalesceMode bool

// Array coalesce modes
const (
	ArraysToSlice ArrayCoalesceMode = false
	ArraysAsIs    ArrayCoalesceMode = true
)

// ValueCoalescer coalesces numeric types as follows:
// All signed integer types may be coalesced into int, int64, uint, uint64 or string
// All unsigned integer types may be coalesced into int, int64, uint, uint64, or string
// float32 and float64 may be coaleseced into float32, float64, or string
// complex64 and complex128 may be coalesced into complex64, complex128, or string
// Arrays may be coalesced into slices
// All other types are passed through as is
type ValueCoalescer struct {
	visitor             ValueVisitor
	intCoalesceMode     IntCoalesceMode
	uintCoalesceMode    UintCoalesceMode
	floatCoalesceMode   FloatCoalesceMode
	complexCoalesceMode ComplexCoalesceMode
	arrayCoalesceMode   ArrayCoalesceMode
}

// NewValueCoalescer creates a ValueCoalescer with the following defaults:
// - signed and unsigned integers are coalesced into int
// - float32 and float64 are coalesced into float64
// - complex64 and complex128 are coaleseced into  complex128
// - arrays are coalesced into slices
func NewValueCoalescer(visitor ValueVisitor) *ValueCoalescer {
	if visitor == nil {
		panic("NewValueCoalescer: visitor cannor cannot be nil")
	}
	return &ValueCoalescer{visitor: visitor}
}

// WithIntCoalesceMode sets the int coalesce mode
func (c *ValueCoalescer) WithIntCoalesceMode(intCoalesceMode IntCoalesceMode) *ValueCoalescer {
	c.intCoalesceMode = intCoalesceMode
	return c
}

// WithUintCoalesceMode sets the uint coalesce mode
func (c *ValueCoalescer) WithUintCoalesceMode(uintCoalesceMode UintCoalesceMode) *ValueCoalescer {
	c.uintCoalesceMode = uintCoalesceMode
	return c
}

// WithFloatCoalesceMode sets the float coalesce mode
func (c *ValueCoalescer) WithFloatCoalesceMode(floatCoalesceMode FloatCoalesceMode) *ValueCoalescer {
	c.floatCoalesceMode = floatCoalesceMode
	return c
}

// WithComplexCoalesceMode sets the int coalesce mode
func (c *ValueCoalescer) WithComplexCoalesceMode(complexCoalesceMode ComplexCoalesceMode) *ValueCoalescer {
	c.complexCoalesceMode = complexCoalesceMode
	return c
}

// WithArrayCoalesceMode sets the array coalesce mode
func (c *ValueCoalescer) WithArrayCoalesceMode(arrayCoalesceMode ArrayCoalesceMode) *ValueCoalescer {
	c.arrayCoalesceMode = arrayCoalesceMode
	return c
}

// VisitInt coalesces an int
func (c ValueCoalescer) VisitInt(val int) {
	switch c.intCoalesceMode {
	case IntsToInt:
		c.visitor.VisitInt(val)
	case IntsToInt64:
		c.visitor.VisitInt64(int64(val))
	case IntsToUint:
		c.visitor.VisitUint(uint(val))
	case IntsToUint64:
		c.visitor.VisitUint64(uint64(val))
	case IntsToString:
		c.visitor.VisitString(strconv.FormatInt(int64(val), 10))
	case IntsAsIs:
		c.visitor.VisitInt(val)
	}
}

// VisitInt8 coalesces an int8
func (c ValueCoalescer) VisitInt8(val int8) {
	switch c.intCoalesceMode {
	case IntsToInt:
		c.visitor.VisitInt(int(val))
	case IntsToInt64:
		c.visitor.VisitInt64(int64(val))
	case IntsToUint:
		c.visitor.VisitUint(uint(val))
	case IntsToUint64:
		c.visitor.VisitUint64(uint64(val))
	case IntsToString:
		c.visitor.VisitString(strconv.FormatInt(int64(val), 10))
	case IntsAsIs:
		c.visitor.VisitInt8(val)
	}
}

// VisitInt16 coalesces an int16
func (c ValueCoalescer) VisitInt16(val int16) {
	switch c.intCoalesceMode {
	case IntsToInt:
		c.visitor.VisitInt(int(val))
	case IntsToInt64:
		c.visitor.VisitInt64(int64(val))
	case IntsToUint:
		c.visitor.VisitUint(uint(val))
	case IntsToUint64:
		c.visitor.VisitUint64(uint64(val))
	case IntsToString:
		c.visitor.VisitString(strconv.FormatInt(int64(val), 10))
	case IntsAsIs:
		c.visitor.VisitInt16(val)
	}
}

// VisitInt32 coalesces an int32
func (c ValueCoalescer) VisitInt32(val int32) {
	switch c.intCoalesceMode {
	case IntsToInt:
		c.visitor.VisitInt(int(val))
	case IntsToInt64:
		c.visitor.VisitInt64(int64(val))
	case IntsToUint:
		c.visitor.VisitUint(uint(val))
	case IntsToUint64:
		c.visitor.VisitUint64(uint64(val))
	case IntsToString:
		c.visitor.VisitString(strconv.FormatInt(int64(val), 10))
	case IntsAsIs:
		c.visitor.VisitInt32(val)
	}
}

// VisitInt64 coalesces an int64
func (c ValueCoalescer) VisitInt64(val int64) {
	switch c.intCoalesceMode {
	case IntsToInt:
		c.visitor.VisitInt(int(val))
	case IntsToInt64:
		c.visitor.VisitInt64(val)
	case IntsToUint:
		c.visitor.VisitUint(uint(val))
	case IntsToUint64:
		c.visitor.VisitUint64(uint64(val))
	case IntsToString:
		c.visitor.VisitString(strconv.FormatInt(int64(val), 10))
	case IntsAsIs:
		c.visitor.VisitInt64(val)
	}
}

// VisitUint coalesces a uint
func (c ValueCoalescer) VisitUint(val uint) {
	switch c.uintCoalesceMode {
	case UintsToInt:
		c.visitor.VisitInt(int(val))
	case IntsToInt64:
		c.visitor.VisitInt64(int64(val))
	case IntsToUint:
		c.visitor.VisitUint(val)
	case IntsToUint64:
		c.visitor.VisitUint64(uint64(val))
	case IntsToString:
		c.visitor.VisitString(strconv.FormatUint(uint64(val), 10))
	case IntsAsIs:
		c.visitor.VisitUint(val)
	}
}

// VisitUint8 coalesces a uint8
func (c ValueCoalescer) VisitUint8(val uint8) {
	switch c.uintCoalesceMode {
	case UintsToInt:
		c.visitor.VisitInt(int(val))
	case IntsToInt64:
		c.visitor.VisitInt64(int64(val))
	case IntsToUint:
		c.visitor.VisitUint(uint(val))
	case IntsToUint64:
		c.visitor.VisitUint64(uint64(val))
	case IntsToString:
		c.visitor.VisitString(strconv.FormatUint(uint64(val), 10))
	case IntsAsIs:
		c.visitor.VisitUint8(val)
	}
}

// VisitUint16 coalesces a uint16
func (c ValueCoalescer) VisitUint16(val uint16) {
	switch c.uintCoalesceMode {
	case UintsToInt:
		c.visitor.VisitInt(int(val))
	case IntsToInt64:
		c.visitor.VisitInt64(int64(val))
	case IntsToUint:
		c.visitor.VisitUint(uint(val))
	case IntsToUint64:
		c.visitor.VisitUint64(uint64(val))
	case IntsToString:
		c.visitor.VisitString(strconv.FormatUint(uint64(val), 10))
	case IntsAsIs:
		c.visitor.VisitUint16(val)
	}
}

// VisitUint32 coalesces a uint32
func (c ValueCoalescer) VisitUint32(val uint32) {
	switch c.uintCoalesceMode {
	case UintsToInt:
		c.visitor.VisitInt(int(val))
	case IntsToInt64:
		c.visitor.VisitInt64(int64(val))
	case IntsToUint:
		c.visitor.VisitUint(uint(val))
	case IntsToUint64:
		c.visitor.VisitUint64(uint64(val))
	case IntsToString:
		c.visitor.VisitString(strconv.FormatUint(uint64(val), 10))
	case IntsAsIs:
		c.visitor.VisitUint32(val)
	}
}

// VisitUint64 coalesces a uint64
func (c ValueCoalescer) VisitUint64(val uint64) {
	switch c.uintCoalesceMode {
	case UintsToInt:
		c.visitor.VisitInt(int(val))
	case IntsToInt64:
		c.visitor.VisitInt64(int64(val))
	case IntsToUint:
		c.visitor.VisitUint(uint(val))
	case IntsToUint64:
		c.visitor.VisitUint64(val)
	case IntsToString:
		c.visitor.VisitString(strconv.FormatUint(val, 10))
	case IntsAsIs:
		c.visitor.VisitUint64(val)
	}
}

// VisitFloat32 coalesces a float32
func (c ValueCoalescer) VisitFloat32(val float32) {
	switch c.floatCoalesceMode {
	case FloatsToFloat32:
		c.visitor.VisitFloat32(val)
	case FloatsToFloat64:
		c.visitor.VisitFloat64(float64(val))
	case FloatsToString:
		c.visitor.VisitString(strconv.FormatFloat(float64(val), 'g', -1, 32))
	case FloatsAsIs:
		c.visitor.VisitFloat32(val)
	}
}

// VisitFloat64 coalesces a float64
func (c ValueCoalescer) VisitFloat64(val float64) {
	switch c.floatCoalesceMode {
	case FloatsToFloat32:
		c.visitor.VisitFloat32(float32(val))
	case FloatsToFloat64:
		c.visitor.VisitFloat64(val)
	case FloatsToString:
		c.visitor.VisitString(strconv.FormatFloat(val, 'g', -1, 64))
	case FloatsAsIs:
		c.visitor.VisitFloat64(val)
	}
}

// VisitComplex64 coalesces a complex64
func (c ValueCoalescer) VisitComplex64(val complex64) {
	switch c.complexCoalesceMode {
	case ComplexesToComplex64:
		c.visitor.VisitComplex64(val)
	case ComplexesToComplex128:
		c.visitor.VisitComplex128(complex128(val))
	case ComplexesToString:
		c.visitor.VisitString(fmt.Sprint(val))
	case ComplexesAsIs:
		c.visitor.VisitComplex64(val)
	}
}

// VisitComplex128 coalesces a complex128
func (c ValueCoalescer) VisitComplex128(val complex128) {
	switch c.complexCoalesceMode {
	case ComplexesToComplex64:
		c.visitor.VisitComplex64(complex64(val))
	case ComplexesToComplex128:
		c.visitor.VisitComplex128(val)
	case ComplexesToString:
		c.visitor.VisitString(fmt.Sprint(val))
	case ComplexesAsIs:
		c.visitor.VisitComplex128(val)
	}
}

// VisitPreArray coalesces an array
func (c ValueCoalescer) VisitPreArray(len int, val reflect.Value) {
	switch c.arrayCoalesceMode {
	case ArraysToSlice:
		c.visitor.VisitPreSlice(len, val)
	case ArraysAsIs:
		c.visitor.VisitPreArray(len, val)
	}
}

// VisitPreArrayIndex coalesces a value of a array
func (c ValueCoalescer) VisitPreArrayIndex(len int, idx int, val reflect.Value) {
	switch c.arrayCoalesceMode {
	case ArraysToSlice:
		c.visitor.VisitPreSliceIndex(len, idx, val)
	case ArraysAsIs:
		c.visitor.VisitPreArrayIndex(len, idx, val)
	}
}

// VisitPostArrayIndex coalesces a value of a array
func (c ValueCoalescer) VisitPostArrayIndex(len int, idx int, val reflect.Value) {
	switch c.arrayCoalesceMode {
	case ArraysToSlice:
		c.visitor.VisitPostSliceIndex(len, idx, val)
	case ArraysAsIs:
		c.visitor.VisitPostArrayIndex(len, idx, val)
	}
}

// VisitPostArray coalesces an array
func (c ValueCoalescer) VisitPostArray(len int, val reflect.Value) {
	switch c.arrayCoalesceMode {
	case ArraysToSlice:
		c.visitor.VisitPostSlice(len, val)
	case ArraysAsIs:
		c.visitor.VisitPostArray(len, val)
	}
}
