package goreflect

import (
	"fmt"
	"reflect"
)

// ValueVisitorAdapter composes any subset of interfaces defined in ValueVisitor into a full ValueVisitor implementation.
// Unimplemented interfaces are filled in with empty implementations.
type ValueVisitorAdapter struct {
	initVisitor                 func()
	boolVisitor                 func(bool)
	intVisitor                  func(int)
	int8Visitor                 func(int8)
	int16Visitor                func(int16)
	int32Visitor                func(int32)
	int64Visitor                func(int64)
	uintVisitor                 func(uint)
	uint8Visitor                func(uint8)
	uint16Visitor               func(uint16)
	uint32Visitor               func(uint32)
	uint64Visitor               func(uint64)
	float32Visitor              func(float32)
	float64Visitor              func(float64)
	complex64Visitor            func(complex64)
	complex128Visitor           func(complex128)
	stringVisitor               func(string)
	chanVisitor                 func(reflect.Value)
	funcVisitor                 func(reflect.Value)
	prePtrVisitor               func(reflect.Value)
	postPtrVisitor              func(reflect.Value)
	preArrayVisitor             func(int, reflect.Value)
	preArrayIndexVisitor        func(int, int, reflect.Value)
	postArrayIndexVisitor       func(int, int, reflect.Value)
	postArrayVisitor            func(int, reflect.Value)
	preSliceVisitor             func(int, reflect.Value)
	preSliceIndexVisitor        func(int, int, reflect.Value)
	postSliceIndexVisitor       func(int, int, reflect.Value)
	postSliceVisitor            func(int, reflect.Value)
	preMapVisitor               func(int, reflect.Value)
	preMapKeyValueVisitor       func(int, int, reflect.Value, reflect.Value)
	preMapKeyVisitor            func(int, int, reflect.Value)
	postMapKeyVisitor           func(int, int, reflect.Value)
	preMapValueVisitor          func(int, int, reflect.Value)
	postMapValueVisitor         func(int, int, reflect.Value)
	postMapKeyValueVisitor      func(int, int, reflect.Value, reflect.Value)
	postMapVisitor              func(int, reflect.Value)
	preStructVisitor            func(int, reflect.Value)
	preStructFieldValueVisitor  func(int, int, reflect.StructField, reflect.Value)
	postStructFieldValueVisitor func(int, int, reflect.StructField, reflect.Value)
	postStructVisitor           func(int, reflect.Value)
}

// NewValueVisitorAdapter constructs a ValueVisitorAdapter
func NewValueVisitorAdapter(visitor ...interface{}) *ValueVisitorAdapter {
	va := &ValueVisitorAdapter{}

	if len(visitor) > 0 {
		va.WithVisitor(visitor[0])
	}

	return va
}

// WithVisitor sets the visitor to use for future calls to Walk
func (va *ValueVisitorAdapter) WithVisitor(visitor interface{}) *ValueVisitorAdapter {
	if visitor == nil {
		panic(fmt.Errorf("goreflect.ValueVisitorAdapter.WithVisitor: visitor cannot be nil"))
	}

	// Compose ValueVisitor with a combination of interfaces implemented by the given visitor,
	// and empty functions for unimplemented interfaces.

	va.initVisitor = func() {}
	if initv, ok := visitor.(InitVisitor); ok {
		va.initVisitor = initv.Init
	}

	va.boolVisitor = func(bool) {}
	if bv, ok := visitor.(BoolVisitor); ok {
		va.boolVisitor = bv.VisitBool
	}

	va.intVisitor = func(int) {}
	if iv, ok := visitor.(IntVisitor); ok {
		va.intVisitor = iv.VisitInt
	}

	va.int8Visitor = func(int8) {}
	if iv8, ok := visitor.(Int8Visitor); ok {
		va.int8Visitor = iv8.VisitInt8
	}

	va.int16Visitor = func(int16) {}
	if iv16, ok := visitor.(Int16Visitor); ok {
		va.int16Visitor = iv16.VisitInt16
	}

	va.int32Visitor = func(int32) {}
	if iv32, ok := visitor.(Int32Visitor); ok {
		va.int32Visitor = iv32.VisitInt32
	}

	va.int64Visitor = func(int64) {}
	if iv64, ok := visitor.(Int64Visitor); ok {
		va.int64Visitor = iv64.VisitInt64
	}

	va.uintVisitor = func(uint) {}
	if uiv, ok := visitor.(UintVisitor); ok {
		va.uintVisitor = uiv.VisitUint
	}

	va.uint8Visitor = func(uint8) {}
	if uiv8, ok := visitor.(Uint8Visitor); ok {
		va.uint8Visitor = uiv8.VisitUint8
	}

	va.uint16Visitor = func(uint16) {}
	if uiv16, ok := visitor.(Uint16Visitor); ok {
		va.uint16Visitor = uiv16.VisitUint16
	}

	va.uint32Visitor = func(uint32) {}
	if uiv32, ok := visitor.(Uint32Visitor); ok {
		va.uint32Visitor = uiv32.VisitUint32
	}

	va.uint64Visitor = func(uint64) {}
	if uiv64, ok := visitor.(Uint64Visitor); ok {
		va.uint64Visitor = uiv64.VisitUint64
	}

	va.float32Visitor = func(float32) {}
	if f32v, ok := visitor.(Float32Visitor); ok {
		va.float32Visitor = f32v.VisitFloat32
	}

	va.float64Visitor = func(float64) {}
	if f64v, ok := visitor.(Float64Visitor); ok {
		va.float64Visitor = f64v.VisitFloat64
	}

	va.complex64Visitor = func(complex64) {}
	if c64v, ok := visitor.(Complex64Visitor); ok {
		va.complex64Visitor = c64v.VisitComplex64
	}

	va.complex128Visitor = func(complex128) {}
	if c128v, ok := visitor.(Complex128Visitor); ok {
		va.complex128Visitor = c128v.VisitComplex128
	}

	va.stringVisitor = func(string) {}
	if sv, ok := visitor.(StringVisitor); ok {
		va.stringVisitor = sv.VisitString
	}

	va.chanVisitor = func(reflect.Value) {}
	if chanv, ok := visitor.(ChanVisitor); ok {
		va.chanVisitor = chanv.VisitChan
	}

	va.funcVisitor = func(reflect.Value) {}
	if funcv, ok := visitor.(FuncVisitor); ok {
		va.funcVisitor = funcv.VisitFunc
	}

	va.prePtrVisitor = func(reflect.Value) {}
	if prePtrv, ok := visitor.(PrePtrVisitor); ok {
		va.prePtrVisitor = prePtrv.VisitPrePtr
	}

	va.postPtrVisitor = func(reflect.Value) {}
	if postPtrv, ok := visitor.(PostPtrVisitor); ok {
		va.postPtrVisitor = postPtrv.VisitPostPtr
	}

	va.preArrayVisitor = func(int, reflect.Value) {}
	if preArrayv, ok := visitor.(PreArrayVisitor); ok {
		va.preArrayVisitor = preArrayv.VisitPreArray
	}

	va.preArrayIndexVisitor = func(int, int, reflect.Value) {}
	if preArrayIndexv, ok := visitor.(PreArrayIndexVisitor); ok {
		va.preArrayIndexVisitor = preArrayIndexv.VisitPreArrayIndex
	}

	va.postArrayIndexVisitor = func(int, int, reflect.Value) {}
	if postArrayIndexv, ok := visitor.(PostArrayIndexVisitor); ok {
		va.postArrayIndexVisitor = postArrayIndexv.VisitPostArrayIndex
	}

	va.postArrayVisitor = func(int, reflect.Value) {}
	if postArrayv, ok := visitor.(PostArrayVisitor); ok {
		va.postArrayVisitor = postArrayv.VisitPostArray
	}

	va.preSliceVisitor = func(int, reflect.Value) {}
	if preSlicev, ok := visitor.(PreSliceVisitor); ok {
		va.preSliceVisitor = preSlicev.VisitPreSlice
	}

	va.preSliceIndexVisitor = func(int, int, reflect.Value) {}
	if preSliceIndexv, ok := visitor.(PreSliceIndexVisitor); ok {
		va.preSliceIndexVisitor = preSliceIndexv.VisitPreSliceIndex
	}

	va.postSliceIndexVisitor = func(int, int, reflect.Value) {}
	if postSliceIndexv, ok := visitor.(PostSliceIndexVisitor); ok {
		va.postSliceIndexVisitor = postSliceIndexv.VisitPostSliceIndex
	}

	va.postSliceVisitor = func(int, reflect.Value) {}
	if postSlicev, ok := visitor.(PostSliceVisitor); ok {
		va.postSliceVisitor = postSlicev.VisitPostSlice
	}

	va.preMapVisitor = func(int, reflect.Value) {}
	if preMapv, ok := visitor.(PreMapVisitor); ok {
		va.preMapVisitor = preMapv.VisitPreMap
	}

	va.preMapKeyValueVisitor = func(int, int, reflect.Value, reflect.Value) {}
	if preMapkv, ok := visitor.(PreMapKeyValueVisitor); ok {
		va.preMapKeyValueVisitor = preMapkv.VisitPreMapKeyValue
	}

	va.preMapKeyVisitor = func(int, int, reflect.Value) {}
	if preMapk, ok := visitor.(PreMapKeyVisitor); ok {
		va.preMapKeyVisitor = preMapk.VisitPreMapKey
	}

	va.postMapKeyVisitor = func(int, int, reflect.Value) {}
	if postMapk, ok := visitor.(PostMapKeyVisitor); ok {
		va.postMapKeyVisitor = postMapk.VisitPostMapKey
	}

	va.preMapValueVisitor = func(int, int, reflect.Value) {}
	if preMapv, ok := visitor.(PreMapValueVisitor); ok {
		va.preMapValueVisitor = preMapv.VisitPreMapValue
	}

	va.postMapValueVisitor = func(int, int, reflect.Value) {}
	if postMapv, ok := visitor.(PostMapValueVisitor); ok {
		va.postMapValueVisitor = postMapv.VisitPostMapValue
	}

	va.postMapKeyValueVisitor = func(int, int, reflect.Value, reflect.Value) {}
	if postMapkv, ok := visitor.(PostMapKeyValueVisitor); ok {
		va.postMapKeyValueVisitor = postMapkv.VisitPostMapKeyValue
	}

	va.postMapVisitor = func(int, reflect.Value) {}
	if postMapv, ok := visitor.(PostMapVisitor); ok {
		va.postMapVisitor = postMapv.VisitPostMap
	}

	va.preStructVisitor = func(int, reflect.Value) {}
	if preStructv, ok := visitor.(PreStructVisitor); ok {
		va.preStructVisitor = preStructv.VisitPreStruct
	}

	va.preStructFieldValueVisitor = func(int, int, reflect.StructField, reflect.Value) {}
	if preStructfv, ok := visitor.(PreStructFieldValueVisitor); ok {
		va.preStructFieldValueVisitor = preStructfv.VisitPreStructFieldValue
	}

	va.postStructFieldValueVisitor = func(int, int, reflect.StructField, reflect.Value) {}
	if postStructfv, ok := visitor.(PostStructFieldValueVisitor); ok {
		va.postStructFieldValueVisitor = postStructfv.VisitPostStructFieldValue
	}

	va.postStructVisitor = func(int, reflect.Value) {}
	if postStructv, ok := visitor.(PostStructVisitor); ok {
		va.postStructVisitor = postStructv.VisitPostStruct
	}

	return va
}

// Init delegates to given InitVisitor
func (va ValueVisitorAdapter) Init() {
	va.initVisitor()
}

// VisitBool delegates to composed BoolVisitor
func (va ValueVisitorAdapter) VisitBool(v bool) {
	va.boolVisitor(v)
}

// VisitInt delegates to composed IntVisitor
func (va ValueVisitorAdapter) VisitInt(v int) {
	va.intVisitor(v)
}

// VisitInt8 delegates to composed Int8Visitor
func (va ValueVisitorAdapter) VisitInt8(v int8) {
	va.int8Visitor(v)
}

// VisitInt16 delegates to composed Int16Visitor
func (va ValueVisitorAdapter) VisitInt16(v int16) {
	va.int16Visitor(v)
}

// VisitInt32 delegates to composed Int32Visitor
func (va ValueVisitorAdapter) VisitInt32(v int32) {
	va.int32Visitor(v)
}

// VisitInt64 delegates to composed Int64Visitor
func (va ValueVisitorAdapter) VisitInt64(v int64) {
	va.int64Visitor(v)
}

// VisitUint delegates to composed UintVisitor
func (va ValueVisitorAdapter) VisitUint(v uint) {
	va.uintVisitor(v)
}

// VisitUint8 delegates to composed Uint8Visitor
func (va ValueVisitorAdapter) VisitUint8(v uint8) {
	va.uint8Visitor(v)
}

// VisitUint16 delegates to composed Uint16Visitor
func (va ValueVisitorAdapter) VisitUint16(v uint16) {
	va.uint16Visitor(v)
}

// VisitUint32 delegates to composed Uint32Visitor
func (va ValueVisitorAdapter) VisitUint32(v uint32) {
	va.uint32Visitor(v)
}

// VisitUint64 delegates to composed Uint64Visitor
func (va ValueVisitorAdapter) VisitUint64(v uint64) {
	va.uint64Visitor(v)
}

// VisitFloat32 delegates to composed Float32Visitor
func (va ValueVisitorAdapter) VisitFloat32(v float32) {
	va.float32Visitor(v)
}

// VisitFloat64 delegates to composed Float64Visitor
func (va ValueVisitorAdapter) VisitFloat64(v float64) {
	va.float64Visitor(v)
}

// VisitComplex64 delegates to composed Complex64Visitor
func (va ValueVisitorAdapter) VisitComplex64(v complex64) {
	va.complex64Visitor(v)
}

// VisitComplex128 delegates to composed Complex128Visitor
func (va ValueVisitorAdapter) VisitComplex128(v complex128) {
	va.complex128Visitor(v)
}

// VisitString delegates to composed StringVisitor
func (va ValueVisitorAdapter) VisitString(v string) {
	va.stringVisitor(v)
}

// VisitChan delegates to composed ChanVisitor
func (va ValueVisitorAdapter) VisitChan(v reflect.Value) {
	va.chanVisitor(v)
}

// VisitFunc delegates to composed FuncVisitor
func (va ValueVisitorAdapter) VisitFunc(v reflect.Value) {
	va.funcVisitor(v)
}

// VisitPrePtr delegates to composed PrePtrVisitor
func (va ValueVisitorAdapter) VisitPrePtr(v reflect.Value) {
	va.prePtrVisitor(v)
}

// VisitPostPtr delegates to composed PostPtrVisitor
func (va ValueVisitorAdapter) VisitPostPtr(v reflect.Value) {
	va.postPtrVisitor(v)
}

// VisitPreArray delegates to composed PreArrayVisitor
func (va ValueVisitorAdapter) VisitPreArray(len int, v reflect.Value) {
	va.preArrayVisitor(len, v)
}

// VisitPreArrayIndex delegates to composed PreArrayIndexVisitor
func (va ValueVisitorAdapter) VisitPreArrayIndex(len int, idx int, v reflect.Value) {
	va.preArrayIndexVisitor(len, idx, v)
}

// VisitPostArrayIndex delegates to composed PostArrayIndexVisitor
func (va ValueVisitorAdapter) VisitPostArrayIndex(len int, idx int, v reflect.Value) {
	va.postArrayIndexVisitor(len, idx, v)
}

// VisitPostArray delegates to composed PostArrayVisitor
func (va ValueVisitorAdapter) VisitPostArray(len int, v reflect.Value) {
	va.postArrayVisitor(len, v)
}

// VisitPreSlice delegates to composed PreSliceVisitor
func (va ValueVisitorAdapter) VisitPreSlice(len int, v reflect.Value) {
	va.preSliceVisitor(len, v)
}

// VisitPreSliceIndex delegates to composed PreSliceIndexVisitor
func (va ValueVisitorAdapter) VisitPreSliceIndex(len int, idx int, v reflect.Value) {
	va.preSliceIndexVisitor(len, idx, v)
}

// VisitPostSliceIndex delegates to composed PostSliceIndexVisitor
func (va ValueVisitorAdapter) VisitPostSliceIndex(len int, idx int, v reflect.Value) {
	va.postSliceIndexVisitor(len, idx, v)
}

// VisitPostSlice delegates to composed PostSliceVisitor
func (va ValueVisitorAdapter) VisitPostSlice(len int, v reflect.Value) {
	va.postSliceVisitor(len, v)
}

// VisitPreMap delegates to composed PreMapVisitor
func (va ValueVisitorAdapter) VisitPreMap(len int, m reflect.Value) {
	va.preMapVisitor(len, m)
}

// VisitPreMapKeyValue delegates to composed PreMapKeyValueVisitor
func (va ValueVisitorAdapter) VisitPreMapKeyValue(len int, idx int, k reflect.Value, v reflect.Value) {
	va.preMapKeyValueVisitor(len, idx, k, v)
}

// VisitPreMapKey delegates to composed PreMapKeyVisitor
func (va ValueVisitorAdapter) VisitPreMapKey(len int, idx int, k reflect.Value) {
	va.preMapKeyVisitor(len, idx, k)
}

// VisitPostMapKey delegates to composed PostMapKeyVisitor
func (va ValueVisitorAdapter) VisitPostMapKey(len int, idx int, k reflect.Value) {
	va.postMapKeyVisitor(len, idx, k)
}

// VisitPreMapValue delegates to composed PreMapValueVisitor
func (va ValueVisitorAdapter) VisitPreMapValue(len int, idx int, v reflect.Value) {
	va.preMapValueVisitor(len, idx, v)
}

// VisitPostMapValue delegates to composed PostMapValueVisitor
func (va ValueVisitorAdapter) VisitPostMapValue(len int, idx int, v reflect.Value) {
	va.postMapValueVisitor(len, idx, v)
}

// VisitPostMapKeyValue delegates to composed PostMapKeyValueVisitor
func (va ValueVisitorAdapter) VisitPostMapKeyValue(len int, idx int, k reflect.Value, v reflect.Value) {
	va.postMapKeyValueVisitor(len, idx, k, v)
}

// VisitPostMap delegates to composed PostMapVisitor
func (va ValueVisitorAdapter) VisitPostMap(len int, m reflect.Value) {
	va.postMapVisitor(len, m)
}

// VisitPreStruct delegates to composed PreStructVisitor
func (va ValueVisitorAdapter) VisitPreStruct(len int, v reflect.Value) {
	va.preStructVisitor(len, v)
}

// VisitPreStructFieldValue delegates to composed PreStructFieldValueVisitor
func (va ValueVisitorAdapter) VisitPreStructFieldValue(len int, idx int, f reflect.StructField, v reflect.Value) {
	va.preStructFieldValueVisitor(len, idx, f, v)
}

// VisitPostStructFieldValue delegates to composed PostStructFieldValueVisitor
func (va ValueVisitorAdapter) VisitPostStructFieldValue(len int, idx int, f reflect.StructField, v reflect.Value) {
	va.postStructFieldValueVisitor(len, idx, f, v)
}

// VisitPostStruct delegates to composed PostStructVisitor
func (va ValueVisitorAdapter) VisitPostStruct(len int, v reflect.Value) {
	va.postStructVisitor(len, v)
}
