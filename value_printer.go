package goreflect

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// ValuePrinter prints out values similar to fmt.Sprintf("%+v"), with some exceptions:
// - strings are double quoted
// - chan and func values are printed as their type
// - pointer values are printed with a leading & for each indirection
// - array, slice, map, and struct values are printed with same format as inline initialization
// If desired, the address can also be printed for chan, func, pointer, slice, and map values.
// The address is inside "@[]" in hex form, and is printed after the type.
// In the case of multiple pointer indirections, each indirection shows the address after the &.
// The zero value is ready to use, and will not print pointer addresses.
type ValuePrinter struct {
	bldr        *strings.Builder
	WithAddress bool
}

// Init initializes printer with empty string
func (p *ValuePrinter) Init() {
	if p.bldr == nil {
		p.bldr = &strings.Builder{}
	} else {
		p.bldr.Reset()
	}
}

// VisitBool prints a boolean
func (p *ValuePrinter) VisitBool(val bool) {
	p.bldr.WriteString(strconv.FormatBool(val))
}

// VisitInt prints an int
func (p *ValuePrinter) VisitInt(val int) {
	p.bldr.WriteString(strconv.FormatInt(int64(val), 10))
}

// VisitInt8 prints an int8
func (p *ValuePrinter) VisitInt8(val int8) {
	p.bldr.WriteString(strconv.FormatInt(int64(val), 10))
}

// VisitInt16 prints an int16
func (p *ValuePrinter) VisitInt16(val int16) {
	p.bldr.WriteString(strconv.FormatInt(int64(val), 10))
}

// VisitInt32 prints an int32
func (p *ValuePrinter) VisitInt32(val int32) {
	p.bldr.WriteString(strconv.FormatInt(int64(val), 10))
}

// VisitInt64 prints an int64
func (p *ValuePrinter) VisitInt64(val int64) {
	p.bldr.WriteString(strconv.FormatInt(val, 10))
}

// VisitUint prints a uint
func (p *ValuePrinter) VisitUint(val uint) {
	p.bldr.WriteString(strconv.FormatUint(uint64(val), 10))
}

// VisitUint8 prints a uint8
func (p *ValuePrinter) VisitUint8(val uint8) {
	p.bldr.WriteString(strconv.FormatUint(uint64(val), 10))
}

// VisitUint16 prints a uint16
func (p *ValuePrinter) VisitUint16(val uint16) {
	p.bldr.WriteString(strconv.FormatUint(uint64(val), 10))
}

// VisitUint32 prints a uint32
func (p *ValuePrinter) VisitUint32(val uint32) {
	p.bldr.WriteString(strconv.FormatUint(uint64(val), 10))
}

// VisitUint64 prints a uint64
func (p *ValuePrinter) VisitUint64(val uint64) {
	p.bldr.WriteString(strconv.FormatUint(val, 10))
}

// VisitFloat32 prints a float32
func (p *ValuePrinter) VisitFloat32(val float32) {
	p.bldr.WriteString(strconv.FormatFloat(float64(val), 'g', -1, 32))
}

// VisitFloat64 prints a float64
func (p *ValuePrinter) VisitFloat64(val float64) {
	p.bldr.WriteString(strconv.FormatFloat(val, 'g', -1, 64))
}

// VisitComplex64 prints a complex64
func (p *ValuePrinter) VisitComplex64(val complex64) {
	p.bldr.WriteString(fmt.Sprint(val))
}

// VisitComplex128 prints a complex128
func (p *ValuePrinter) VisitComplex128(val complex128) {
	p.bldr.WriteString(fmt.Sprint(val))
}

// VisitString prints a string
func (p *ValuePrinter) VisitString(val string) {
	p.bldr.WriteRune('"')
	p.bldr.WriteString(val)
	p.bldr.WriteRune('"')
}

// VisitChan prints a chan
func (p *ValuePrinter) VisitChan(val reflect.Value) {
	p.bldr.WriteString(val.Type().String())
	if p.WithAddress {
		p.bldr.WriteString(fmt.Sprintf(" @[%p]", val.Interface()))
	}
}

// VisitFunc prints a func
func (p *ValuePrinter) VisitFunc(val reflect.Value) {
	p.bldr.WriteString(val.Type().String())
	if p.WithAddress {
		p.bldr.WriteString(fmt.Sprintf(" @[%p]", val.Interface()))
	}
}

// VisitPrePtr prints a ptr
func (p *ValuePrinter) VisitPrePtr(val reflect.Value) {
	p.bldr.WriteRune('&')
	if p.WithAddress {
		p.bldr.WriteString(fmt.Sprintf("@[%p]", val.Interface()))
	}
}

// VisitPreArray prints an array
func (p *ValuePrinter) VisitPreArray(_ int, val reflect.Value) {
	p.bldr.WriteString(val.Type().String())
	p.bldr.WriteRune('{')
}

// VisitPreArrayIndex prints a value of a array
func (p *ValuePrinter) VisitPreArrayIndex(_ int, idx int, _ reflect.Value) {
	if idx > 0 {
		p.bldr.WriteString(", ")
	}
}

// VisitPostArray prints an array
func (p *ValuePrinter) VisitPostArray(_ int, _ reflect.Value) {
	p.bldr.WriteRune('}')
}

// VisitPreSlice prints a slice
func (p *ValuePrinter) VisitPreSlice(_ int, val reflect.Value) {
	p.bldr.WriteString(val.Type().String())
	if p.WithAddress {
		p.bldr.WriteString(fmt.Sprintf("@[%p]", val.Interface()))
	}
	p.bldr.WriteRune('{')
}

// VisitPreSliceIndex prints a value of a slice
func (p *ValuePrinter) VisitPreSliceIndex(_ int, idx int, _ reflect.Value) {
	if idx > 0 {
		p.bldr.WriteString(", ")
	}
}

// VisitPostSlice prints a slice
func (p *ValuePrinter) VisitPostSlice(_ int, _ reflect.Value) {
	p.bldr.WriteRune('}')
}

// VisitPreMap prints a map
func (p *ValuePrinter) VisitPreMap(_ int, val reflect.Value) {
	p.bldr.WriteString(val.Type().String())
	if p.WithAddress {
		p.bldr.WriteString(fmt.Sprintf("@[%p]", val.Interface()))
	}
	p.bldr.WriteRune('{')
}

// VisitPreMapKey prints a key of a map
func (p *ValuePrinter) VisitPreMapKey(_ int, idx int, _ reflect.Value) {
	if idx > 0 {
		p.bldr.WriteString(", ")
	}
}

// VisitPreMapValue prints a value of a map key
func (p *ValuePrinter) VisitPreMapValue(_ int, _ int, _ reflect.Value) {
	p.bldr.WriteString(": ")
}

// VisitPostMap prints a map
func (p *ValuePrinter) VisitPostMap(_ int, _ reflect.Value) {
	p.bldr.WriteRune('}')
}

// VisitPreStruct prints a struct
func (p *ValuePrinter) VisitPreStruct(_ int, val reflect.Value) {
	p.bldr.WriteString(val.Type().String())
	p.bldr.WriteRune('{')
}

// VisitPreStructFieldValue prints the value of a struct field
func (p *ValuePrinter) VisitPreStructFieldValue(_ int, idx int, fld reflect.StructField, _ reflect.Value) {
	if idx > 0 {
		p.bldr.WriteString(", ")
	}
	p.bldr.WriteString(fld.Name)
	p.bldr.WriteString(": ")
}

// VisitPostStruct prints a struct
func (p *ValuePrinter) VisitPostStruct(_ int, _ reflect.Value) {
	p.bldr.WriteRune('}')
}

// Result returns the generated string
func (p *ValuePrinter) Result() string {
	return p.bldr.String()
}
