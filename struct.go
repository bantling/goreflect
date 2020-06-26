package goreflect

import (
	"fmt"
	"reflect"
)

// FlatStruct provides a flat view of the fields and methods of a struct.
// The flat view contains the set of fields and methods of the struct itself, as well as those of child structs it contains, recursively.
// For a set of fields or methods that have the same name, the first one encountered in bread first traversal is selected.
// Effectively, the view provided is like the view provided by the compiler for embedded structs.
// However, the view does not distinguish between embedded child structs and ordinary child structs.
type FlatStruct struct {
	fields     []reflect.StructField
	valMethods []reflect.Method
	ptrMethods []reflect.Method
}

// FlatStructOf flattens a struct instance by visiting the struct fields and methods breadth first.
// If a given field or method name has already been encountered, it is simply ignored, else the field or method is added to the map.
// The value passed may be any of the following:
// - a struct instance/ptr
// - a reflect.Value wrapper of a struct instance/ptr
// - a reflect.Type wrapper of a struct type/ptr
// Note that while a struct instance may be passed, a FlatStruct is a view of a type
func FlatStructOf(val interface{}) FlatStruct {
	// Get type of struct value
	valType, isType := val.(reflect.Type)
	if !isType {
		valType = GetReflectValueOf(val).Type()
	}
	valType = DerefdReflectType(valType)

	if valType.Kind() != reflect.Struct {
		panic(fmt.Errorf("FlatStructOf: %T is not a struct instance, a reflect.Value wrapper of a struct instance, or a reflect.Type wrapper of a struct type", val))
	}

	// Breadth first recursive traversal of structs
	// Ignore any field whose name is already mapped
	var (
		fieldNames     = map[string]interface{}{}
		valMethodNames = map[string]interface{}{}
		ptrMethodNames = map[string]interface{}{}
		fields         = []reflect.StructField{}
		valMethods     = []reflect.Method{}
		ptrMethods     = []reflect.Method{}
		recurse        func(reflect.Type)
	)

	// Define recursive function on value type
	recurse = func(vt reflect.Type) {
		// Breadth first; collect fields and sub structs
		subStructs := []reflect.StructField{}
		for i, n := 0, vt.NumField(); i < n; i++ {
			fld := vt.Field(i)
			if _, exists := fieldNames[fld.Name]; !exists {
				// Fields
				fieldNames[fld.Name] = nil
				fields = append(fields, fld)

				// Sub structs
				if derefdType := DerefdReflectType(fld.Type); derefdType.Kind() == reflect.Struct {
					subStructs = append(subStructs, fld)
				}
			}
		}

		// Breadth first; collect value methods
		for i, n := 0, vt.NumMethod(); i < n; i++ {
			mthd := vt.Method(i)
			if _, exists := valMethodNames[mthd.Name]; !exists {
				valMethodNames[mthd.Name] = nil
				valMethods = append(valMethods, mthd)
			}
		}

		// Breadth first; collect pointer methods
		// All methods can be called with pointers.
		// If a method has a value receiver and you look for it in a reflect.Type that is a pointer,
		// the type will report it as taking a pointer receiver.
		// So all we can do is filter out any methods that are known to be pointer methods,
		// then further filter out known value methods.
		pt := reflect.PtrTo(vt)
		for i, n := 0, pt.NumMethod(); i < n; i++ {
			mthd := pt.Method(i)
			if _, exists := ptrMethodNames[mthd.Name]; !exists {
				if _, exists = valMethodNames[mthd.Name]; !exists {
					ptrMethodNames[mthd.Name] = nil
					ptrMethods = append(ptrMethods, mthd)
				}
			}
		}

		// Recurse sub structs
		for _, fld := range subStructs {
			// Ensure we pass value type to recursive call
			recurse(DerefdReflectType(fld.Type))
		}
	}

	// Call recursive function on value type
	recurse(valType)

	// Return a FlatStruct
	return FlatStruct{
		fields:     fields,
		valMethods: valMethods,
		ptrMethods: ptrMethods,
	}
}

// Fields returns a copy of all fields
func (f FlatStruct) Fields() []reflect.StructField {
	fieldsCopy := make([]reflect.StructField, len(f.fields))
	copy(fieldsCopy, f.fields)
	return fieldsCopy
}

// FieldIter returns an iterator function for the fields of the struct hierarchy that a FlatStruct describes.
// For each field, the iterator returns (reflect.StructField instance, true).
// After the last field has been iterated, all further calls return (reflect.StructField zero value, false).
func (f FlatStruct) FieldIter() func() (reflect.StructField, bool) {
	var (
		i = 0
		n = len(f.fields)
	)

	return func() (reflect.StructField, bool) {
		if i < n {
			j := i
			i++
			return f.fields[j], true
		}

		return reflect.StructField{}, false
	}
}

// ValMethods returns a copy of all value methods
func (f FlatStruct) ValMethods() []reflect.Method {
	valMethodsCopy := make([]reflect.Method, len(f.valMethods))
	copy(valMethodsCopy, f.valMethods)
	return valMethodsCopy
}

// ValMethodsIter returns an iterator function for the methods of the struct hierarchy that have a value receiver.
// For each method, the iterator returns (reflect.Method instance, true).
// After the last method has been iterated, all further calls return (reflect.Method{}, false).
func (f FlatStruct) ValMethodsIter() func() (method reflect.Method, hasNext bool) {
	var (
		i = 0
		n = len(f.valMethods)
	)

	return func() (reflect.Method, bool) {
		if i < n {
			j := i
			i++
			return f.valMethods[j], true
		}

		return reflect.Method{}, false
	}
}

// PtrMethods returns a copy of all pointer methods
func (f FlatStruct) PtrMethods() []reflect.Method {
	ptrMethodsCopy := make([]reflect.Method, len(f.ptrMethods))
	copy(ptrMethodsCopy, f.ptrMethods)
	return ptrMethodsCopy
}

// PtrMethodsIter returns an iterator function for the methods of the struct hierarchy that have a pointer receiver.
// For each method, the iterator returns (reflect.Method instance, true).
// After the last method has been iterated, all further calls return (reflect.Method{}, false).
func (f FlatStruct) PtrMethodsIter() func() (method reflect.Method, hasNext bool) {
	var (
		i = 0
		n = len(f.ptrMethods)
	)

	return func() (reflect.Method, bool) {
		if i < n {
			j := i
			i++
			return f.ptrMethods[j], true
		}

		return reflect.Method{}, false
	}
}

// AllMethods returns a copy of all value methods followed by all pointer methods
func (f FlatStruct) AllMethods() []reflect.Method {
	allMethodsCopy := make([]reflect.Method, len(f.valMethods)+len(f.ptrMethods))
	copy(allMethodsCopy, f.valMethods)
	copy(allMethodsCopy[len(f.valMethods):], f.ptrMethods)
	return allMethodsCopy
}

// AllMethodsIter returns an iterator function for all methods of the struct hierarchy.
// For each method, the iterator returns (reflect.Method instance, true for val receiver | false for pointer receiver, true).
// After the last method has been iterated, all further calls return (reflect.Method{}, false, false).
func (f FlatStruct) AllMethodsIter() func() (method reflect.Method, isVal bool, hasNext bool) {
	var (
		vi = 0
		vn = len(f.valMethods)
		pi = 0
		pn = len(f.ptrMethods)
	)

	return func() (reflect.Method, bool, bool) {
		if vi < vn {
			j := vi
			vi++
			return f.valMethods[j], true, true
		} else if pi < pn {
			j := pi
			pi++
			return f.ptrMethods[j], false, true
		}

		return reflect.Method{}, false, false
	}
}
