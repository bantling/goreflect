package goreflect

import (
	"fmt"
	"reflect"
)

// Indirection constants
const (
	Value  = 0
	Ptr    = 1
	PtrPtr = 2
)

// GetReflectValueOf takes a value and returns a reflect.Value wrapper.
// If the value is already a reflect.Value, it is returned as is.
func GetReflectValueOf(value interface{}) reflect.Value {
	if v, ok := value.(reflect.Value); ok {
		return v
	}

	return reflect.ValueOf(value)
}

// DerefdReflectValue takes a reflect.Value that may be have one or more levels of indirection, and dereferences it until it is a value type.
// If the value is Invalid, it is returned as is.
func DerefdReflectValue(value reflect.Value) reflect.Value {
	// Assume the instance is not a pointer
	derefd := value

	// Only deref a valid Value
	if derefd.IsValid() {
		for derefd.Kind() == reflect.Ptr {
			derefd = derefd.Elem()
		}
	}

	return derefd
}

// DerefdReflectType takes a reflect.Type that may be have one or more levels of indirection, and dereferences it until it is a value type.
// Only the outer type is derefd, eg *[]**string => []**string
func DerefdReflectType(typ reflect.Type) reflect.Type {
	// Assume the type is not a pointer
	derefd := typ

	// Deref the type
	for derefd.Kind() == reflect.Ptr {
		derefd = derefd.Elem()
	}

	return derefd
}

// FullyDerefdReflectType returns the fully dereferenced type, where all indirections are removed:
// T, *T, ... => T
// [1]T, *[1]T, [1]*T, ... => [1]T
// []T, *[]T, []*T, ... => []T
// chan T, *chan T, chan *T, ... => chan T
// map[K]V, *map[K]V, map[*K]V, map[K]*V, ... => map[K]V
func FullyDerefdReflectType(typ reflect.Type) reflect.Type {
	derefd := DerefdReflectType(typ)

	switch derefd.Kind() {
	case reflect.Array:
		derefd = reflect.ArrayOf(derefd.Len(), FullyDerefdReflectType(derefd.Elem()))
	case reflect.Slice:
		derefd = reflect.SliceOf(FullyDerefdReflectType(derefd.Elem()))
	case reflect.Chan:
		derefd = reflect.ChanOf(derefd.ChanDir(), FullyDerefdReflectType(derefd.Elem()))
	case reflect.Map:
		derefd = reflect.MapOf(FullyDerefdReflectType(derefd.Key()), FullyDerefdReflectType(derefd.Elem()))
	}

	return derefd
}

// NumRefs returns the number of references in the given type
func NumRefs(typ reflect.Type) int {
	numRefs := 0
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		numRefs++
	}

	return numRefs
}

// CreateRefs creates the specified number of references to the value given.
// Note that if the value already has one or more references, the specified number of references is still added.
func CreateRefs(value reflect.Value, numRefs uint) reflect.Value {
	currentValue := value
	for i := numRefs; i > 0; i-- {
		// reflect.New adds an indirection
		newPtr := reflect.New(currentValue.Type())

		// Set the new pointer to point to current value
		newPtr.Elem().Set(currentValue)

		// current value is new pointer
		currentValue = newPtr
	}

	return currentValue
}

// valueDerefer fully dereferences all components of a type.
// EG, a *[]**[]***string is dereferenced to [][]string.
// The zero value is ready to use.
type valueDerefer struct {
	derefd           reflect.Value
	composites       []reflect.Value
	currentComposite reflect.Value
	indexes          []reflect.Value
	currentIndex     reflect.Value
	valueIsMapKey    bool
}

// add a composite
func (d *valueDerefer) add(val reflect.Value) {
	d.composites = append(d.composites, val)
	d.currentComposite = val
	if len(d.composites) == 1 {
		d.derefd = val
	}
}

// add a derefd value to a composite type at current index
// if valueIsMapKey is true, then the value is a map key, otherwise it is a array/slice/map key value
func (d *valueDerefer) set(val reflect.Value) {
	switch d.currentComposite.Kind() {
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		d.currentComposite.Index(int(d.currentIndex.Int())).Set(val)

	case reflect.Map:
		if d.valueIsMapKey {
			d.currentIndex = val
			d.indexes = append(d.indexes, d.currentIndex)
		} else {
			d.currentComposite.SetMapIndex(d.currentIndex, val)
		}
	}
}

// remove a composite
func (d *valueDerefer) remove() {
	previousComposite := d.composites[len(d.composites)-1]
	d.composites = d.composites[:len(d.composites)-1]
	d.indexes = d.indexes[:len(d.indexes)-1]
	if len(d.composites) > 0 {
		d.currentComposite = d.composites[len(d.composites)-1]
		d.currentIndex = d.indexes[len(d.indexes)-1]
		d.set(previousComposite)
	} else {
		d.currentComposite = reflect.Value{}
		d.currentIndex = reflect.Value{}
	}
}

func (d *valueDerefer) Init() {
	d.derefd = reflect.Value{}
	d.composites = nil
	d.currentComposite = reflect.Value{}
	d.indexes = nil
	d.currentIndex = reflect.Value{}
	d.valueIsMapKey = false
}

func (d *valueDerefer) VisitPreArray(length int, val reflect.Value) {
	arrElemType := FullyDerefdReflectType(val.Type().Elem())
	newArray := reflect.New(reflect.ArrayOf(length, arrElemType)).Elem()
	d.add(newArray)
}

func (d *valueDerefer) VisitPreArrayIndex(_ int, idx int, _ reflect.Value) {
	idxVal := reflect.ValueOf(idx)
	d.currentIndex = idxVal
	d.indexes = append(d.indexes, d.currentIndex)
}

func (d *valueDerefer) VisitPostArray(int, reflect.Value) {
	d.remove()
}

func (d *valueDerefer) VisitPreSlice(length int, val reflect.Value) {
	sliceType := FullyDerefdReflectType(val.Type())
	newSlice := reflect.MakeSlice(sliceType, length, length)
	d.add(newSlice)
}

func (d *valueDerefer) VisitPreSliceIndex(_ int, idx int, _ reflect.Value) {
	idxVal := reflect.ValueOf(idx)
	d.currentIndex = idxVal
	d.indexes = append(d.indexes, d.currentIndex)
}

func (d *valueDerefer) VisitPostSlice(int, reflect.Value) {
	d.remove()
}

func (d *valueDerefer) VisitPreMap(length int, m reflect.Value) {
	mapType := FullyDerefdReflectType(m.Type())

	// Maps keys can be slice or map pointers, but not slice or map values
	keyType := mapType.Key()
	switch keyType.Kind() {
	case reflect.Slice:
		fallthrough
	case reflect.Map:
		panic(fmt.Errorf("Cannot fully deref a map whose key is a slice or map pointer, as their values are not legal keys"))
	}

	newMap := reflect.MakeMapWithSize(mapType, length)
	d.add(newMap)
}

func (d *valueDerefer) VisitPreMapKey(_ int, _ int, _ reflect.Value) {
	// When we visit the derefd map key, we need to store it in currentIndex
	d.valueIsMapKey = true
}

func (d *valueDerefer) VisitPreMapValue(int, int, reflect.Value) {
	// When we visit the derefd map value, we need to store it in the map using the currentIndex
	d.valueIsMapKey = false
}

func (d *valueDerefer) VisitPostMap(int, reflect.Value) {
	d.remove()
}

func (d *valueDerefer) VisitBool(val bool) {
	d.set(reflect.ValueOf(val))
}

func (d *valueDerefer) VisitInt(val int) {
	d.set(reflect.ValueOf(val))
}

func (d *valueDerefer) VisitInt8(val int8) {
	d.set(reflect.ValueOf(val))
}

func (d *valueDerefer) VisitInt16(val int) {
	d.set(reflect.ValueOf(val))
}

func (d *valueDerefer) VisitInt32(val int) {
	d.set(reflect.ValueOf(val))
}

func (d *valueDerefer) VisitInt64(val int) {
	d.set(reflect.ValueOf(val))
}

func (d *valueDerefer) VisitUint(val uint) {
	d.set(reflect.ValueOf(val))
}

func (d *valueDerefer) VisitUint8(val uint8) {
	d.set(reflect.ValueOf(val))
}

func (d *valueDerefer) VisitUint16(val uint16) {
	d.set(reflect.ValueOf(val))
}

func (d *valueDerefer) VisitUint32(val uint32) {
	d.set(reflect.ValueOf(val))
}

func (d *valueDerefer) VisitUint64(val uint64) {
	d.set(reflect.ValueOf(val))
}

func (d *valueDerefer) VisitFloat32(val float32) {
	d.set(reflect.ValueOf(val))
}

func (d *valueDerefer) VisitFloat64(val float64) {
	d.set(reflect.ValueOf(val))
}

func (d *valueDerefer) VisitComplex64(val complex64) {
	d.set(reflect.ValueOf(val))
}

func (d *valueDerefer) VisitComplex128(val complex128) {
	d.set(reflect.ValueOf(val))
}

func (d *valueDerefer) VisitString(val string) {
	d.set(reflect.ValueOf(val))
}

func (d *valueDerefer) VisitChan(val reflect.Value) {
	d.set(val)
}

func (d *valueDerefer) VisitFunc(val reflect.Value) {
	d.set(val)
}

func (d *valueDerefer) VisitPreStruct(_ int, val reflect.Value) {
	d.set(val)
}

func (d *valueDerefer) Result() reflect.Value {
	return d.derefd
}

// FullyDerefdValue returns the fully dereferenced value, where all indirections are removed.
// If the given value has no indirections, it is returned as is.
// Structs are dereferenced, but their fields are not as that would change their types to an unassignable value.
// The value may be a value to dereference or a reflect.Value wrapper that contains the value to dereference.
//
// Note that a map key can be a pointer to a slice or map, but not a slice or map value.
// As such, we cannot fully deref a map if the key is a pointer to a slice or map.
// If such a structure is encountered, a panic will occur.
func FullyDerefdValue(val interface{}) interface{} {
	// Get fully derefd top level value
	rval := DerefdReflectValue(GetReflectValueOf(val))

	// If it is a non-struct composite type, fully deref the components with a valueDerefer
	switch rval.Kind() {
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		fallthrough
	case reflect.Map:
		d := &valueDerefer{}
		NewValueDepthFirstWalker(NewValueVisitorAdapter(d)).Walk(rval)
		return d.Result().Interface()
	}

	// If it is a scalar type, just return fully derefd top level value
	return rval.Interface()
}
