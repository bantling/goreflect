package goreflect

import (
	"fmt"
	"reflect"
)

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
