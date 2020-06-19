package goreflect

import (
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
	VisitPreArray(length int, v reflect.Value)
}

// PreArrayIndexVisitor previsits array value indexes
type PreArrayIndexVisitor interface {
	VisitPreArrayIndex(length int, index int, v reflect.Value)
}

// PostArrayIndexVisitor postvisits array value indexes
type PostArrayIndexVisitor interface {
	VisitPostArrayIndex(length int, index int, v reflect.Value)
}

// PostArrayVisitor postvisits array values
type PostArrayVisitor interface {
	VisitPostArray(length int, v reflect.Value)
}

// PreSliceVisitor previsits slice values
type PreSliceVisitor interface {
	VisitPreSlice(length int, v reflect.Value)
}

// PreSliceIndexVisitor previsits slice value indexes
type PreSliceIndexVisitor interface {
	VisitPreSliceIndex(length int, index int, v reflect.Value)
}

// PostSliceIndexVisitor prostvisits slice value indexes
type PostSliceIndexVisitor interface {
	VisitPostSliceIndex(length int, index int, v reflect.Value)
}

// PostSliceVisitor postvisits slice values
type PostSliceVisitor interface {
	VisitPostSlice(length int, v reflect.Value)
}

// PreMapVisitor previsits map values
type PreMapVisitor interface {
	VisitPreMap(length int, m reflect.Value)
}

// PreMapKeyValueVisitor previsits map key/value pairs
type PreMapKeyValueVisitor interface {
	VisitPreMapKeyValue(length int, index int, k reflect.Value, v reflect.Value)
}

// PreMapKeyVisitor previsits map keys
type PreMapKeyVisitor interface {
	VisitPreMapKey(length int, index int, k reflect.Value)
}

// PostMapKeyVisitor postvisits map keys
type PostMapKeyVisitor interface {
	VisitPostMapKey(length int, index int, k reflect.Value)
}

// PreMapValueVisitor previsits map key values
type PreMapValueVisitor interface {
	VisitPreMapValue(length int, index int, v reflect.Value)
}

// PostMapValueVisitor postvisits map key values
type PostMapValueVisitor interface {
	VisitPostMapValue(length int, index int, v reflect.Value)
}

// PostMapKeyValueVisitor postvisits map key/value pairs
type PostMapKeyValueVisitor interface {
	VisitPostMapKeyValue(length int, index int, k reflect.Value, v reflect.Value)
}

// PostMapVisitor postvisits map values
type PostMapVisitor interface {
	VisitPostMap(length int, m reflect.Value)
}

// PreStructVisitor previsits struct values
type PreStructVisitor interface {
	VisitPreStruct(length int, v reflect.Value)
}

// PreStructFieldValueVisitor previsits struct field value pairs
type PreStructFieldValueVisitor interface {
	VisitPreStructFieldValue(length int, index int, f reflect.StructField, v reflect.Value)
}

// PostStructFieldValueVisitor postvisits struct field value pairs
type PostStructFieldValueVisitor interface {
	VisitPostStructFieldValue(length int, index int, f reflect.StructField, v reflect.Value)
}

// PostStructVisitor postvisits struct values
type PostStructVisitor interface {
	VisitPostStruct(length int, v reflect.Value)
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
