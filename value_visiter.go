package goreflect

import (
	"fmt"
	"reflect"
)

// InitVisitor initializes the visiting process
type InitVisitor interface {
	Init()
}

// BoolVisitor visits bool values
type BoolVisitor interface {
	VisitBool(bool)
}

// IntVisitor visits int values
type IntVisitor interface {
	VisitInt(int)
}

// Int8Visitor visits int8 values
type Int8Visitor interface {
	VisitInt8(int8)
}

// Int16Visitor visits int16 values
type Int16Visitor interface {
	VisitInt16(int16)
}

// Int32Visitor visits int32 values
type Int32Visitor interface {
	VisitInt32(int32)
}

// Int64Visitor visits int64 values
type Int64Visitor interface {
	VisitInt64(int64)
}

// UintVisitor visits uint values
type UintVisitor interface {
	VisitUint(uint)
}

// Uint8Visitor visits uint8 values
type Uint8Visitor interface {
	VisitUint8(uint8)
}

// Uint16Visitor visits uint16 values
type Uint16Visitor interface {
	VisitUint16(uint16)
}

// Uint32Visitor visits uint32 values
type Uint32Visitor interface {
	VisitUint32(uint32)
}

// Uint64Visitor visits uint64 values
type Uint64Visitor interface {
	VisitUint64(uint64)
}

// Float32Visitor visits float32 values
type Float32Visitor interface {
	VisitFloat32(float32)
}

// Float64Visitor visits float64 values
type Float64Visitor interface {
	VisitFloat64(float64)
}

// Complex64Visitor visits complex64 values
type Complex64Visitor interface {
	VisitComplex64(complex64)
}

// Complex128Visitor visits complex128 values
type Complex128Visitor interface {
	VisitComplex128(complex128)
}

// StringVisitor visits string values
type StringVisitor interface {
	VisitString(string)
}

// ChanVisitor visits chan values
type ChanVisitor interface {
	VisitChan(reflect.Value)
}

// FuncVisitor visits func values
type FuncVisitor interface {
	VisitFunc(reflect.Value)
}

// PrePtrVisitor previsits ptr values
type PrePtrVisitor interface {
	VisitPrePtr(reflect.Value)
}

// PostPtrVisitor postvisits ptr values
type PostPtrVisitor interface {
	VisitPostPtr(reflect.Value)
}

// PreArrayVisitor previsits array values
type PreArrayVisitor interface {
	VisitPreArray(len int, v reflect.Value)
}

// PreArrayIndexVisitor previsits array value indexes
type PreArrayIndexVisitor interface {
	VisitPreArrayIndex(len int, idx int, v reflect.Value)
}

// PostArrayIndexVisitor postvisits array value indexes
type PostArrayIndexVisitor interface {
	VisitPostArrayIndex(len int, idx int, v reflect.Value)
}

// PostArrayVisitor postvisits array values
type PostArrayVisitor interface {
	VisitPostArray(len int, v reflect.Value)
}

// PreSliceVisitor previsits slice values
type PreSliceVisitor interface {
	VisitPreSlice(len int, v reflect.Value)
}

// PreSliceIndexVisitor previsits slice value indexes
type PreSliceIndexVisitor interface {
	VisitPreSliceIndex(len int, idx int, v reflect.Value)
}

// PostSliceIndexVisitor prostvisits slice value indexes
type PostSliceIndexVisitor interface {
	VisitPostSliceIndex(len int, idx int, v reflect.Value)
}

// PostSliceVisitor postvisits slice values
type PostSliceVisitor interface {
	VisitPostSlice(len int, v reflect.Value)
}

// PreMapVisitor previsits map values
type PreMapVisitor interface {
	VisitPreMap(len int, m reflect.Value)
}

// PreMapKeyValueVisitor previsits map key/value pairs
type PreMapKeyValueVisitor interface {
	VisitPreMapKeyValue(len int, idx int, k reflect.Value, v reflect.Value)
}

// PreMapKeyVisitor previsits map keys
type PreMapKeyVisitor interface {
	VisitPreMapKey(len int, idx int, k reflect.Value)
}

// PostMapKeyVisitor postvisits map keys
type PostMapKeyVisitor interface {
	VisitPostMapKey(len int, idx int, k reflect.Value)
}

// PreMapValueVisitor previsits map key values
type PreMapValueVisitor interface {
	VisitPreMapValue(len int, idx int, v reflect.Value)
}

// PostMapValueVisitor postvisits map key values
type PostMapValueVisitor interface {
	VisitPostMapValue(len int, idx int, v reflect.Value)
}

// PostMapKeyValueVisitor postvisits map key/value pairs
type PostMapKeyValueVisitor interface {
	VisitPostMapKeyValue(len int, idx int, k reflect.Value, v reflect.Value)
}

// PostMapVisitor postvisits map values
type PostMapVisitor interface {
	VisitPostMap(len int, m reflect.Value)
}

// PreStructVisitor previsits struct values
type PreStructVisitor interface {
	VisitPreStruct(len int, v reflect.Value)
}

// PreStructFieldValueVisitor previsits struct field value pairs
type PreStructFieldValueVisitor interface {
	VisitPreStructFieldValue(len int, idx int, f reflect.StructField, v reflect.Value)
}

// PostStructFieldValueVisitor postvisits struct field value pairs
type PostStructFieldValueVisitor interface {
	VisitPostStructFieldValue(len int, idx int, f reflect.StructField, v reflect.Value)
}

// PostStructVisitor postvisits struct values
type PostStructVisitor interface {
	VisitPostStruct(len int, v reflect.Value)
}

// ValueVisitor combines all above interfaces into one
type ValueVisitor interface {
	InitVisitor
	BoolVisitor
	IntVisitor
	Int8Visitor
	Int16Visitor
	Int32Visitor
	Int64Visitor
	UintVisitor
	Uint8Visitor
	Uint16Visitor
	Uint32Visitor
	Uint64Visitor
	Float32Visitor
	Float64Visitor
	Complex64Visitor
	Complex128Visitor
	StringVisitor
	ChanVisitor
	FuncVisitor
	PrePtrVisitor
	PostPtrVisitor
	PreArrayVisitor
	PreArrayIndexVisitor
	PostArrayIndexVisitor
	PostArrayVisitor
	PreSliceVisitor
	PreSliceIndexVisitor
	PostSliceIndexVisitor
	PostSliceVisitor
	PreMapVisitor
	PreMapKeyValueVisitor
	PreMapKeyVisitor
	PostMapKeyVisitor
	PreMapValueVisitor
	PostMapValueVisitor
	PostMapKeyValueVisitor
	PostMapVisitor
	PreStructVisitor
	PreStructFieldValueVisitor
	PostStructFieldValueVisitor
	PostStructVisitor
}

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

// ValueVisitorReducer reduces all visiter interfaces to a single function that accepts the string method name and []reflect.Value of arguments
type ValueVisitorReducer struct {
	dispatcher func(method string, args []reflect.Value)
}

// NewValueVisitorReducer constructs a ValueVisitorReducer
func NewValueVisitorReducer(dispatcher ...func(method string, args []reflect.Value)) *ValueVisitorReducer {
	va := &ValueVisitorReducer{}

	if len(dispatcher) > 0 {
		va.WithDispatcher(dispatcher[0])
	}

	return va
}

// WithDispatcher sets the dispatcher to use for future calls to Walk
func (vr *ValueVisitorReducer) WithDispatcher(dispatcher func(method string, args []reflect.Value)) *ValueVisitorReducer {
	if dispatcher == nil {
		panic(fmt.Errorf("goreflect.ValueVisitorReducer.WithDispatcher: dispatcher cannot be nil"))
	}

	vr.dispatcher = dispatcher
	return vr
}

// Init does nothing
func (ValueVisitorReducer) Init() {
	//
}

// VisitBool dispatches ("VisitBool", bool)
func (vr ValueVisitorReducer) VisitBool(v bool) {
	vr.dispatcher("VisitBool", []reflect.Value{reflect.ValueOf(v)})
}

// VisitInt dispatches ("VisitInt", int)
func (vr ValueVisitorReducer) VisitInt(v int) {
	vr.dispatcher("VisitInt", []reflect.Value{reflect.ValueOf(v)})
}

// VisitInt8 dispatches ("VisitInt8", int8)
func (vr ValueVisitorReducer) VisitInt8(v int8) {
	vr.dispatcher("VisitInt8", []reflect.Value{reflect.ValueOf(v)})
}

// VisitInt16 dispatches ("VisitInt16", int16)
func (vr ValueVisitorReducer) VisitInt16(v int16) {
	vr.dispatcher("VisitInt16", []reflect.Value{reflect.ValueOf(v)})
}

// VisitInt32 dispatches ("VisitInt32", int32)
func (vr ValueVisitorReducer) VisitInt32(v int32) {
	vr.dispatcher("VisitInt32", []reflect.Value{reflect.ValueOf(v)})
}

// VisitInt64 dispatches ("VisitInt64", int64)
func (vr ValueVisitorReducer) VisitInt64(v int64) {
	vr.dispatcher("VisitInt64", []reflect.Value{reflect.ValueOf(v)})
}

// VisitUint dispatches ("VisitUint", uint)
func (vr ValueVisitorReducer) VisitUint(v uint) {
	vr.dispatcher("VisitUint", []reflect.Value{reflect.ValueOf(v)})
}

// VisitUint8 dispatches ("VisitUint8", uint8)
func (vr ValueVisitorReducer) VisitUint8(v uint8) {
	vr.dispatcher("VisitUint8", []reflect.Value{reflect.ValueOf(v)})
}

// VisitUint16 dispatches ("VisitUint16", uint16)
func (vr ValueVisitorReducer) VisitUint16(v uint16) {
	vr.dispatcher("VisitUint16", []reflect.Value{reflect.ValueOf(v)})
}

// VisitUint32 dispatches ("VisitUint32", uint32)
func (vr ValueVisitorReducer) VisitUint32(v uint32) {
	vr.dispatcher("VisitUint32", []reflect.Value{reflect.ValueOf(v)})
}

// VisitUint64 dispatches ("VisitUint64", uint64)
func (vr ValueVisitorReducer) VisitUint64(v uint64) {
	vr.dispatcher("VisitUint64", []reflect.Value{reflect.ValueOf(v)})
}

// VisitFloat32 dispatches ("VisitFloat32", float32)
func (vr ValueVisitorReducer) VisitFloat32(v float32) {
	vr.dispatcher("VisitFloat32", []reflect.Value{reflect.ValueOf(v)})
}

// VisitFloat64 dispatches ("VisitFloat64", float64)
func (vr ValueVisitorReducer) VisitFloat64(v float64) {
	vr.dispatcher("VisitFloat64", []reflect.Value{reflect.ValueOf(v)})
}

// VisitComplex64 dispatches ("VisitComplex64", complex64)
func (vr ValueVisitorReducer) VisitComplex64(v complex64) {
	vr.dispatcher("VisitComplex64", []reflect.Value{reflect.ValueOf(v)})
}

// VisitComplex128 dispatches ("VisitComplex128", complex128)
func (vr ValueVisitorReducer) VisitComplex128(v complex128) {
	vr.dispatcher("VisitComplex128", []reflect.Value{reflect.ValueOf(v)})
}

// VisitString dispatches ("VisitString", string)
func (vr ValueVisitorReducer) VisitString(v string) {
	vr.dispatcher("VisitString", []reflect.Value{reflect.ValueOf(v)})
}

// VisitChan dispatches ("VisitChan", chan)
func (vr ValueVisitorReducer) VisitChan(v reflect.Value) {
	vr.dispatcher("VisitChan", []reflect.Value{v})
}

// VisitFunc dispatches ("VisitFunc", func)
func (vr ValueVisitorReducer) VisitFunc(v reflect.Value) {
	vr.dispatcher("VisitFunc", []reflect.Value{v})
}

// VisitPrePtr dispatches ("VisitPrePtr", ptr)
func (vr ValueVisitorReducer) VisitPrePtr(v reflect.Value) {
	vr.dispatcher("VisitPrePtr", []reflect.Value{v})
}

// VisitPostPtr dispatches ("VisitPostPtr", ptr)
func (vr ValueVisitorReducer) VisitPostPtr(v reflect.Value) {
	vr.dispatcher("VisitPostPtr", []reflect.Value{v})
}

// VisitPreArray dispatches ("VisitPreArray", len, array)
func (vr ValueVisitorReducer) VisitPreArray(len int, v reflect.Value) {
	vr.dispatcher("VisitPreArray", []reflect.Value{reflect.ValueOf(len), v})
}

// VisitPreArrayIndex dispatches ("VisitPreArrayIndex", len, idx, array)
func (vr ValueVisitorReducer) VisitPreArrayIndex(len int, idx int, v reflect.Value) {
	vr.dispatcher("VisitPreArrayIndex", []reflect.Value{reflect.ValueOf(len), reflect.ValueOf(idx), v})
}

// VisitPostArrayIndex dispatches ("VisitPostArrayIndex", len, idx, array)
func (vr ValueVisitorReducer) VisitPostArrayIndex(len int, idx int, v reflect.Value) {
	vr.dispatcher("VisitPostArrayIndex", []reflect.Value{reflect.ValueOf(len), reflect.ValueOf(idx), v})
}

// VisitPostArray dispatches ("VisitPostArray", len, array)
func (vr ValueVisitorReducer) VisitPostArray(len int, v reflect.Value) {
	vr.dispatcher("VisitPostArray", []reflect.Value{reflect.ValueOf(len), v})
}

// VisitPreSlice dispatches ("VisitPreSlice", len, array)
func (vr ValueVisitorReducer) VisitPreSlice(len int, v reflect.Value) {
	vr.dispatcher("VisitPreSlice", []reflect.Value{reflect.ValueOf(len), v})
}

// VisitPreSliceIndex dispatches ("VisitPreSliceIndex", len, idx, array)
func (vr ValueVisitorReducer) VisitPreSliceIndex(len int, idx int, v reflect.Value) {
	vr.dispatcher("VisitPreSliceIndex", []reflect.Value{reflect.ValueOf(len), reflect.ValueOf(idx), v})
}

// VisitPostSliceIndex dispatches ("VisitPostSliceIndex", len, idx, array)
func (vr ValueVisitorReducer) VisitPostSliceIndex(len int, idx int, v reflect.Value) {
	vr.dispatcher("VisitPostSliceIndex", []reflect.Value{reflect.ValueOf(len), reflect.ValueOf(idx), v})
}

// VisitPostSlice dispatches ("VisitPostSlice", len, array)
func (vr ValueVisitorReducer) VisitPostSlice(len int, v reflect.Value) {
	vr.dispatcher("VisitPostSlice", []reflect.Value{reflect.ValueOf(len), v})
}

// VisitPreMap dispatches ("VisitPreMap", len, map)
func (vr ValueVisitorReducer) VisitPreMap(len int, m reflect.Value) {
	vr.dispatcher("VisitPreMap", []reflect.Value{reflect.ValueOf(len), m})
}

// VisitPreMapKeyValue dispatches ("VisitPreMapKeyValue", len, idx, k, v)
func (vr ValueVisitorReducer) VisitPreMapKeyValue(len int, idx int, k reflect.Value, v reflect.Value) {
	vr.dispatcher("VisitPreMapKeyValue", []reflect.Value{reflect.ValueOf(len), reflect.ValueOf(idx), reflect.ValueOf(k), reflect.ValueOf(v)})
}

// VisitPreMapKey dispatches ("VisitPreMapKey", len, idx, k)
func (vr ValueVisitorReducer) VisitPreMapKey(len int, idx int, k reflect.Value) {
	vr.dispatcher("VisitPreMapKey", []reflect.Value{reflect.ValueOf(len), reflect.ValueOf(idx), reflect.ValueOf(k)})
}

// VisitPostMapKey dispatches ("VisitPostMapKey", len, idx, k)
func (vr ValueVisitorReducer) VisitPostMapKey(len int, idx int, k reflect.Value) {
	vr.dispatcher("VisitPostMapKey", []reflect.Value{reflect.ValueOf(len), reflect.ValueOf(idx), reflect.ValueOf(k)})
}

// VisitPreMapValue dispatches ("VisitPreMapValue", len, idx, v)
func (vr ValueVisitorReducer) VisitPreMapValue(len int, idx int, v reflect.Value) {
	vr.dispatcher("VisitPreMapValue", []reflect.Value{reflect.ValueOf(len), reflect.ValueOf(idx), reflect.ValueOf(v)})
}

// VisitPostMapValue dispatches ("VisitPostMapValue", len, idx, v)
func (vr ValueVisitorReducer) VisitPostMapValue(len int, idx int, v reflect.Value) {
	vr.dispatcher("VisitPostMapValue", []reflect.Value{reflect.ValueOf(len), reflect.ValueOf(idx), reflect.ValueOf(v)})
}

// VisitPostMapKeyValue dispatches ("VisitPostMapKeyValue", len, map)
func (vr ValueVisitorReducer) VisitPostMapKeyValue(len int, idx int, k reflect.Value, v reflect.Value) {
	vr.dispatcher("VisitPostMapKeyValue", []reflect.Value{reflect.ValueOf(len), reflect.ValueOf(idx), reflect.ValueOf(k), reflect.ValueOf(v)})
}

// VisitPostMap dispatches ("VisitPostMap", len, map)
func (vr ValueVisitorReducer) VisitPostMap(len int, m reflect.Value) {
	vr.dispatcher("VisitPostMap", []reflect.Value{reflect.ValueOf(len), m})
}

// VisitPreStruct dispatches ("VisitPreStruct", len, struct)
func (vr ValueVisitorReducer) VisitPreStruct(len int, v reflect.Value) {
	vr.dispatcher("VisitPreStruct", []reflect.Value{reflect.ValueOf(len), v})
}

// VisitPreStructFieldValue dispatches ("VisitPreStructFieldValue", len, field, value)
func (vr ValueVisitorReducer) VisitPreStructFieldValue(len int, idx int, f reflect.StructField, v reflect.Value) {
	vr.dispatcher("VisitPreStructFieldValue", []reflect.Value{reflect.ValueOf(len), reflect.ValueOf(idx), reflect.ValueOf(f), v})
}

// VisitPostStructFieldValue dispatches ("VisitPostStructFieldValue", len, field, value)
func (vr ValueVisitorReducer) VisitPostStructFieldValue(len int, idx int, f reflect.StructField, v reflect.Value) {
	vr.dispatcher("VisitPostStructFieldValue", []reflect.Value{reflect.ValueOf(len), reflect.ValueOf(idx), reflect.ValueOf(f), v})
}

// VisitPostStruct dispatches ("VisitPostStruct", len, struct)
func (vr ValueVisitorReducer) VisitPostStruct(len int, v reflect.Value) {
	vr.dispatcher("VisitPostStruct", []reflect.Value{reflect.ValueOf(len), v})
}

// ValueDepthFirstWalker visits a value in a depth first traversal.
// A walk begins with the Walk method, which accepts a value to walk and a ValueVisitor to call on each part of the value.
// A walker instance is reuseable, it can be called many times wth different values and/or different ValueVisitors.
//
// Example call sequence to visit a *map[string]int = &map[string]int{"foo": 1, "bar": 2}:
// Walk(&map[string]int{"foo": 1, "bar": 2}, rv)
// VisitPrePtr(&map[string]int{"foo": 1, "bar": 2})
// VisitPreMap(2, map[string]int{"foo": 1, "bar": 2})
// VisitPreMapKeyValue(2, 0, "foo", 1)
// VisitPreMapKey(2, 0, "foo")
// VisitString("foo")
// VisitPostMapKey(2, 0, "foo")
// VisitPreMapValue(2, 0, 1)
// VisitInt(1)
// VisitPostMapValue(2, 0, 1)
// VisitPostMapKeyValue(2, 0, "foo", 1)
// VisitPreMapKeyValue(2, 1, "bar", 2)
// VisitPreMapKey(2, 1, "bar")
// VisitString("bar")
// VisitPostMapKey(2, 1, "bar")
// VisitPreMapValue(2, 1, 2)
// VisitInt(2)
// VisitPostMapValue(2, 1, 2)
// VisitPostMapKeyValue(2, 1, "bar", 2)
// VisitPostMap(2, map[string]int{"foo": 1, "bar": 2})
// VisitPostPtr(&map[string]int{"foo": 1, "bar": 2})
//
// Note that the map will iterate the key/value pairs in whatever order go.reflect provides.
type ValueDepthFirstWalker struct {
	visitor ValueVisitor
}

// NewValueDepthFirstWalker constructs a ValueDepthFirstWalker with a ValueVisitor
func NewValueDepthFirstWalker(visitor ValueVisitor) ValueDepthFirstWalker {
	return ValueDepthFirstWalker{visitor: visitor}
}

// Walk walks the given value in a depth-first traversal.
// The value passed can be a reflect.Value wrapper or a plain value.
// The Walk can be invoked multiple times with different values, as each walk begins by calling the Init() method the visitor given in the costructor.
// There is no return result from the walk. Instead, the visitor is expected to have a Result() method that returns the appropriate type.
func (w ValueDepthFirstWalker) Walk(val interface{}) {
	w.visitor.Init()
	w.Dispatch(GetReflectValueOf(val))
}

// Dispatch executes the appropriate visitor methods for a value based on the type
func (w ValueDepthFirstWalker) Dispatch(v reflect.Value) {
	switch v.Kind() {
	case reflect.Bool:
		w.visitor.VisitBool(v.Bool())

	case reflect.Int:
		w.visitor.VisitInt(int(v.Int()))

	case reflect.Int8:
		w.visitor.VisitInt8(int8(v.Int()))

	case reflect.Int16:
		w.visitor.VisitInt16(int16(v.Int()))

	case reflect.Int32:
		w.visitor.VisitInt32(int32(v.Int()))

	case reflect.Int64:
		w.visitor.VisitInt64(v.Int())

	case reflect.Uint:
		w.visitor.VisitUint(uint(v.Uint()))

	case reflect.Uint8:
		w.visitor.VisitUint8(uint8(v.Uint()))

	case reflect.Uint16:
		w.visitor.VisitUint16(uint16(v.Uint()))

	case reflect.Uint32:
		w.visitor.VisitUint32(uint32(v.Uint()))

	case reflect.Uint64:
		w.visitor.VisitUint64(v.Uint())

	case reflect.Float32:
		w.visitor.VisitFloat32(float32(v.Float()))

	case reflect.Float64:
		w.visitor.VisitFloat64(v.Float())

	case reflect.Complex64:
		w.visitor.VisitComplex64(complex64(v.Complex()))

	case reflect.Complex128:
		w.visitor.VisitComplex128(v.Complex())

	case reflect.String:
		w.visitor.VisitString(v.String())

	case reflect.Chan:
		w.visitor.VisitChan(v)

	case reflect.Func:
		w.visitor.VisitFunc(v)

	case reflect.Ptr:
		w.visitor.VisitPrePtr(v)
		w.Dispatch(v.Elem())
		w.visitor.VisitPostPtr(v)

	case reflect.Array:
		{
			n := v.Len()
			w.visitor.VisitPreArray(n, v)

			for i := 0; i < n; i++ {
				e := v.Index(i)

				w.visitor.VisitPreArrayIndex(n, i, e)
				w.Dispatch(e)
				w.visitor.VisitPostArrayIndex(n, i, e)
			}

			w.visitor.VisitPostArray(n, v)
		}

	case reflect.Slice:
		{
			n := v.Len()
			w.visitor.VisitPreSlice(n, v)

			for i := 0; i < n; i++ {
				e := v.Index(i)

				w.visitor.VisitPreSliceIndex(n, i, e)
				w.Dispatch(e)
				w.visitor.VisitPostSliceIndex(n, i, e)
			}

			w.visitor.VisitPostSlice(n, v)
		}

	case reflect.Map:
		{
			n := v.Len()
			w.visitor.VisitPreMap(n, v)

			for i, iter := 0, v.MapRange(); iter.Next(); i++ {
				mk, mv := iter.Key(), iter.Value()
				w.visitor.VisitPreMapKeyValue(n, i, mk, mv)

				w.visitor.VisitPreMapKey(n, i, mk)
				w.Dispatch(mk)
				w.visitor.VisitPostMapKey(n, i, mk)

				w.visitor.VisitPreMapValue(n, i, mv)
				w.Dispatch(mv)
				w.visitor.VisitPostMapValue(n, i, mv)

				w.visitor.VisitPostMapKeyValue(n, i, mk, mv)
			}

			w.visitor.VisitPostMap(n, v)
		}

	case reflect.Struct:
		{
			n := v.NumField()
			w.visitor.VisitPreStruct(n, v)

			for i := 0; i < n; i++ {
				sf, sv := v.Type().Field(i), v.Field(i)

				w.visitor.VisitPreStructFieldValue(n, i, sf, sv)
				w.Dispatch(sv)
				w.visitor.VisitPostStructFieldValue(n, i, sf, sv)
			}

			w.visitor.VisitPostStruct(n, v)
		}

	default:
		panic(fmt.Errorf("goreflect.ValueWalker.dispatch: value of kind %s cannot be visited", v.Kind()))
	}
}
