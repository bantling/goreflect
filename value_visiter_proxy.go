package goreflect

import (
	"fmt"
	"reflect"
)

// ValueVisitorProxy reduces all visiter interfaces to a single function that accepts a string method name and []reflect.Value of arguments
type ValueVisitorProxy struct {
	dispatcher func(method string, args []reflect.Value)
}

// NewValueVisitorProxy constructs a ValueVisitorProxy
func NewValueVisitorProxy(dispatcher ...func(method string, args []reflect.Value)) *ValueVisitorProxy {
	va := &ValueVisitorProxy{}

	if len(dispatcher) > 0 {
		va.WithDispatcher(dispatcher[0])
	}

	return va
}

// WithDispatcher sets the dispatcher to use for future calls to Walk
func (vr *ValueVisitorProxy) WithDispatcher(dispatcher func(method string, args []reflect.Value)) *ValueVisitorProxy {
	if dispatcher == nil {
		panic(fmt.Errorf("goreflect.ValueVisitorProxy.WithDispatcher: dispatcher cannot be nil"))
	}

	vr.dispatcher = dispatcher
	return vr
}

// Init does nothing
func (ValueVisitorProxy) Init() {
	//
}

// VisitBool dispatches ("VisitBool", bool)
func (vr ValueVisitorProxy) VisitBool(v bool) {
	vr.dispatcher("VisitBool", []reflect.Value{reflect.ValueOf(v)})
}

// VisitInt dispatches ("VisitInt", int)
func (vr ValueVisitorProxy) VisitInt(v int) {
	vr.dispatcher("VisitInt", []reflect.Value{reflect.ValueOf(v)})
}

// VisitInt8 dispatches ("VisitInt8", int8)
func (vr ValueVisitorProxy) VisitInt8(v int8) {
	vr.dispatcher("VisitInt8", []reflect.Value{reflect.ValueOf(v)})
}

// VisitInt16 dispatches ("VisitInt16", int16)
func (vr ValueVisitorProxy) VisitInt16(v int16) {
	vr.dispatcher("VisitInt16", []reflect.Value{reflect.ValueOf(v)})
}

// VisitInt32 dispatches ("VisitInt32", int32)
func (vr ValueVisitorProxy) VisitInt32(v int32) {
	vr.dispatcher("VisitInt32", []reflect.Value{reflect.ValueOf(v)})
}

// VisitInt64 dispatches ("VisitInt64", int64)
func (vr ValueVisitorProxy) VisitInt64(v int64) {
	vr.dispatcher("VisitInt64", []reflect.Value{reflect.ValueOf(v)})
}

// VisitUint dispatches ("VisitUint", uint)
func (vr ValueVisitorProxy) VisitUint(v uint) {
	vr.dispatcher("VisitUint", []reflect.Value{reflect.ValueOf(v)})
}

// VisitUint8 dispatches ("VisitUint8", uint8)
func (vr ValueVisitorProxy) VisitUint8(v uint8) {
	vr.dispatcher("VisitUint8", []reflect.Value{reflect.ValueOf(v)})
}

// VisitUint16 dispatches ("VisitUint16", uint16)
func (vr ValueVisitorProxy) VisitUint16(v uint16) {
	vr.dispatcher("VisitUint16", []reflect.Value{reflect.ValueOf(v)})
}

// VisitUint32 dispatches ("VisitUint32", uint32)
func (vr ValueVisitorProxy) VisitUint32(v uint32) {
	vr.dispatcher("VisitUint32", []reflect.Value{reflect.ValueOf(v)})
}

// VisitUint64 dispatches ("VisitUint64", uint64)
func (vr ValueVisitorProxy) VisitUint64(v uint64) {
	vr.dispatcher("VisitUint64", []reflect.Value{reflect.ValueOf(v)})
}

// VisitFloat32 dispatches ("VisitFloat32", float32)
func (vr ValueVisitorProxy) VisitFloat32(v float32) {
	vr.dispatcher("VisitFloat32", []reflect.Value{reflect.ValueOf(v)})
}

// VisitFloat64 dispatches ("VisitFloat64", float64)
func (vr ValueVisitorProxy) VisitFloat64(v float64) {
	vr.dispatcher("VisitFloat64", []reflect.Value{reflect.ValueOf(v)})
}

// VisitComplex64 dispatches ("VisitComplex64", complex64)
func (vr ValueVisitorProxy) VisitComplex64(v complex64) {
	vr.dispatcher("VisitComplex64", []reflect.Value{reflect.ValueOf(v)})
}

// VisitComplex128 dispatches ("VisitComplex128", complex128)
func (vr ValueVisitorProxy) VisitComplex128(v complex128) {
	vr.dispatcher("VisitComplex128", []reflect.Value{reflect.ValueOf(v)})
}

// VisitString dispatches ("VisitString", string)
func (vr ValueVisitorProxy) VisitString(v string) {
	vr.dispatcher("VisitString", []reflect.Value{reflect.ValueOf(v)})
}

// VisitChan dispatches ("VisitChan", chan)
func (vr ValueVisitorProxy) VisitChan(v reflect.Value) {
	vr.dispatcher("VisitChan", []reflect.Value{v})
}

// VisitFunc dispatches ("VisitFunc", func)
func (vr ValueVisitorProxy) VisitFunc(v reflect.Value) {
	vr.dispatcher("VisitFunc", []reflect.Value{v})
}

// VisitPrePtr dispatches ("VisitPrePtr", ptr)
func (vr ValueVisitorProxy) VisitPrePtr(v reflect.Value) {
	vr.dispatcher("VisitPrePtr", []reflect.Value{v})
}

// VisitPostPtr dispatches ("VisitPostPtr", ptr)
func (vr ValueVisitorProxy) VisitPostPtr(v reflect.Value) {
	vr.dispatcher("VisitPostPtr", []reflect.Value{v})
}

// VisitPreArray dispatches ("VisitPreArray", length, array)
func (vr ValueVisitorProxy) VisitPreArray(length int, v reflect.Value) {
	vr.dispatcher("VisitPreArray", []reflect.Value{reflect.ValueOf(length), v})
}

// VisitPreArrayIndex dispatches ("VisitPreArrayIndex", length, index, array)
func (vr ValueVisitorProxy) VisitPreArrayIndex(length int, index int, v reflect.Value) {
	vr.dispatcher("VisitPreArrayIndex", []reflect.Value{reflect.ValueOf(length), reflect.ValueOf(index), v})
}

// VisitPostArrayIndex dispatches ("VisitPostArrayIndex", length, index, array)
func (vr ValueVisitorProxy) VisitPostArrayIndex(length int, index int, v reflect.Value) {
	vr.dispatcher("VisitPostArrayIndex", []reflect.Value{reflect.ValueOf(length), reflect.ValueOf(index), v})
}

// VisitPostArray dispatches ("VisitPostArray", length, array)
func (vr ValueVisitorProxy) VisitPostArray(length int, v reflect.Value) {
	vr.dispatcher("VisitPostArray", []reflect.Value{reflect.ValueOf(length), v})
}

// VisitPreSlice dispatches ("VisitPreSlice", length, array)
func (vr ValueVisitorProxy) VisitPreSlice(length int, v reflect.Value) {
	vr.dispatcher("VisitPreSlice", []reflect.Value{reflect.ValueOf(length), v})
}

// VisitPreSliceIndex dispatches ("VisitPreSliceIndex", length, index, array)
func (vr ValueVisitorProxy) VisitPreSliceIndex(length int, index int, v reflect.Value) {
	vr.dispatcher("VisitPreSliceIndex", []reflect.Value{reflect.ValueOf(length), reflect.ValueOf(index), v})
}

// VisitPostSliceIndex dispatches ("VisitPostSliceIndex", length, index, array)
func (vr ValueVisitorProxy) VisitPostSliceIndex(length int, index int, v reflect.Value) {
	vr.dispatcher("VisitPostSliceIndex", []reflect.Value{reflect.ValueOf(length), reflect.ValueOf(index), v})
}

// VisitPostSlice dispatches ("VisitPostSlice", length, array)
func (vr ValueVisitorProxy) VisitPostSlice(length int, v reflect.Value) {
	vr.dispatcher("VisitPostSlice", []reflect.Value{reflect.ValueOf(length), v})
}

// VisitPreMap dispatches ("VisitPreMap", length, map)
func (vr ValueVisitorProxy) VisitPreMap(length int, m reflect.Value) {
	vr.dispatcher("VisitPreMap", []reflect.Value{reflect.ValueOf(length), m})
}

// VisitPreMapKeyValue dispatches ("VisitPreMapKeyValue", length, index, k, v)
func (vr ValueVisitorProxy) VisitPreMapKeyValue(length int, index int, k reflect.Value, v reflect.Value) {
	vr.dispatcher("VisitPreMapKeyValue", []reflect.Value{reflect.ValueOf(length), reflect.ValueOf(index), reflect.ValueOf(k), reflect.ValueOf(v)})
}

// VisitPreMapKey dispatches ("VisitPreMapKey", length, index, k)
func (vr ValueVisitorProxy) VisitPreMapKey(length int, index int, k reflect.Value) {
	vr.dispatcher("VisitPreMapKey", []reflect.Value{reflect.ValueOf(length), reflect.ValueOf(index), reflect.ValueOf(k)})
}

// VisitPostMapKey dispatches ("VisitPostMapKey", length, index, k)
func (vr ValueVisitorProxy) VisitPostMapKey(length int, index int, k reflect.Value) {
	vr.dispatcher("VisitPostMapKey", []reflect.Value{reflect.ValueOf(length), reflect.ValueOf(index), reflect.ValueOf(k)})
}

// VisitPreMapValue dispatches ("VisitPreMapValue", length, index, v)
func (vr ValueVisitorProxy) VisitPreMapValue(length int, index int, v reflect.Value) {
	vr.dispatcher("VisitPreMapValue", []reflect.Value{reflect.ValueOf(length), reflect.ValueOf(index), reflect.ValueOf(v)})
}

// VisitPostMapValue dispatches ("VisitPostMapValue", length, index, v)
func (vr ValueVisitorProxy) VisitPostMapValue(length int, index int, v reflect.Value) {
	vr.dispatcher("VisitPostMapValue", []reflect.Value{reflect.ValueOf(length), reflect.ValueOf(index), reflect.ValueOf(v)})
}

// VisitPostMapKeyValue dispatches ("VisitPostMapKeyValue", length, map)
func (vr ValueVisitorProxy) VisitPostMapKeyValue(length int, index int, k reflect.Value, v reflect.Value) {
	vr.dispatcher("VisitPostMapKeyValue", []reflect.Value{reflect.ValueOf(length), reflect.ValueOf(index), reflect.ValueOf(k), reflect.ValueOf(v)})
}

// VisitPostMap dispatches ("VisitPostMap", length, map)
func (vr ValueVisitorProxy) VisitPostMap(length int, m reflect.Value) {
	vr.dispatcher("VisitPostMap", []reflect.Value{reflect.ValueOf(length), m})
}

// VisitPreStruct dispatches ("VisitPreStruct", length, struct)
func (vr ValueVisitorProxy) VisitPreStruct(length int, v reflect.Value) {
	vr.dispatcher("VisitPreStruct", []reflect.Value{reflect.ValueOf(length), v})
}

// VisitPreStructFieldValue dispatches ("VisitPreStructFieldValue", length, field, value)
func (vr ValueVisitorProxy) VisitPreStructFieldValue(length int, index int, f reflect.StructField, v reflect.Value) {
	vr.dispatcher("VisitPreStructFieldValue", []reflect.Value{reflect.ValueOf(length), reflect.ValueOf(index), reflect.ValueOf(f), v})
}

// VisitPostStructFieldValue dispatches ("VisitPostStructFieldValue", length, field, value)
func (vr ValueVisitorProxy) VisitPostStructFieldValue(length int, index int, f reflect.StructField, v reflect.Value) {
	vr.dispatcher("VisitPostStructFieldValue", []reflect.Value{reflect.ValueOf(length), reflect.ValueOf(index), reflect.ValueOf(f), v})
}

// VisitPostStruct dispatches ("VisitPostStruct", length, struct)
func (vr ValueVisitorProxy) VisitPostStruct(length int, v reflect.Value) {
	vr.dispatcher("VisitPostStruct", []reflect.Value{reflect.ValueOf(length), v})
}
