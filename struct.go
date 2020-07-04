package goreflect

import (
	"fmt"
	"reflect"
)

// FlatStructItemMode describes the mode of a FlatStructItem
type FlatStructItemMode uint

const (
	Field FlatStructItemMode = iota // an ordinary (non-function) field
	Func // a field that is defined to be a function
	Method // a method
)

// FlatStructItem contains a single struct item (field or method), and provides a consistent view of both.
//
// If a struct item represents a field whose type is NOT a function:
// - Mode is Field
// - Name is the field name
// - Type is the field type
// - Func is a getter/setter
//   The first arg is a pointer to the top level struct instance
//   If there is no second arg, the call is a getter and returns the current value
//   If there is a second arg, the call is a setter and returns the previous value
// - Tag contains any tags
//
// If a struct item represents a field whose type IS a function:
// - Mode is Func
// - Name is the field name
// - Type is the declared function type
// - Func is a wrapper for the underlying func
//   The first arg is a pointer to the top level struct instance
//   The remaining args are unwrapped to call the underlying function, and the results are wrapped
//   The (un)wrapping occurs even if underlying func accepts and/or returns reflect.Value
// - Tag contains any tags
//
// If a struct item represents a method:
// - Mode is Method
// - Name is the method name
// - Type is the method type
// - Func is a wrapper for the underlying method 
//   The first arg is a pointer to the top level struct instance
//   The remaining args are unwrapped to call the underlying method, and the results are wrapped
//   The (un)wrapping occurs even if underlying method accepts and/or returns reflect.Value
// - Tag is a zero value
//
// Since Go only allows struct tags on fields, the only way to have an annotated method in Go is to have a func field.
// The purpose of the Func mode is to allow the caller to recognize such a use case, and use a tag to generate an implementation.  
type FlatStructItem {
	Mode FlatStructItemMode
	Name string
	Type reflect.Type
	Func func([]reflect.Value) []reflect.Value
	Tag  StructTag
}

// FlatStructType provides a flat view of the fields and methods of a struct.
// The flat view contains the set of fields and methods of the struct itself, as well as those of child structs it contains, recursively.
// For a set of fields or methods that have the same name, the first one encountered in bread first traversal is selected.
// Effectively, the view provided is like the view provided by the compiler for embedded structs.
// However, the view does not distinguish between embedded child structs and ordinary child structs.
// The method objects require a receiver as the first argument.
type FlatStructType struct {
	items []FlatStructItem
}

// flatStructTypeCache caches details of a struct type.
// Multiple calls can be made to analyze the same struct, and only the first call actually performs the analysis,
// the other calls just get the cached information.
var flatStructTypeCache = map[reflect.Type]FlatStructType
	
// A generator for a field accessor (getter/setter) func
// topLevelValueType is a type that describes the top level struct instance as a value type.
// path is a list of string field names to dig down from top level struct to get the struct instance that contains the target field.
// fieldName is the target field name.
// If the target field is in the top level struct, then the path is empty.
func fieldAccessorFuncGen(topLevelValueType reflect.Type, path []string, fieldName string) func([]reflect.Value) []reflect.Value {
	return func(args []reflect.Value) []reflect.Value {
		// Must have first arg of top level struct instance ptr
		n := len(args)
		if (n == 0) || (args[0].Kind() != reflect.Ptr) || (!args[0].Elem().Type() != topLevelValueType) {
			panic(fmt.Errorf("FlatStructType: first arg for %s.%s accessor must be an instance pointer", structValueType, fldName)) 
		}
		
		// Dig down to the value of the struct that contains the target field
		target := args[0].Elem()
		for _, name := range path {
			targetStruct = targetStruct.FieldByName(name)
		}
		
		// Get target field
		target := target.FieldByName(fieldName)
		
		// Result to return
		var result reflect.Value
		
		switch n {
			case 1: // getter operation
				result = target
			case 2: // setter operation
				// Second arg must be correct type
				if !args[1].IsAssignableTo(target.Type()) {
					panic(fmt.Errorf("FlatStructType: second arg for %s.%s accessor must be assignable to type %s", structValueType, fldName, target.Type()))
				}
				
				// Result is previous value
				result = reflect.ValueOf(target.Interface())
				target.Set(args[1])
				
			default: // invalid
				panic("FlatStructType: the accessor for %s.%s accepts at most two args, an instance pointer and new value", structValueType, fldName)  
		}
		
		return []reflect.value{result}
	}
}

// FlatStructOf flattens a struct instance by visiting the struct fields and methods breadth first.
// If a given field or method name has already been encountered, it is simply ignored, else the field or method is added.
// If reflect.Type passed is not a struct type or ptr to one, a panic will occur.
// The first time this function encounters a given struct type, it analyzes it, further calls will return cached information.
func FlatStructTypeOf(typ reflect.Type) FlatStruct {
	// Get value as a reflect.Value wrapper
	dtyp := DerefdReflectType(typ)
	if (dtyp.Kind() != Struct) {
		panic(fmt.Errorf("FlatStructTypeOf: type %s is not a struct type or ptr to a struct type")) 
	}
	
	// Shortcut: if the top level type is cached, just return it
	if cachedInfo, exists := flatStructTypeCache[dtyp]; exists {
		return cachedInfo
	}
	
	// The unique set of item names for a struct type
	itemNames := map[string]interface{}
	
	// The actual items for a struct type
	items := []FlatStructItem

	// Breadth first recursive traversal of structs
	// Ignore any field whose name is already mapped
	// Only analyze structs that do not exist in the cache
	// Define recursive function on a type
	recurse := func(rtyp reflect.Type) {
		// Check cache first
		if cachedInfo, cached := flatStructTypeCache[rtyp]; cached {
			// Copy the cached data whose item names haven't been encountered yet - this may occur at any level of the tree
			for _, item := range cachedInfo.items {
				if _, exists := itemNames[item.Name]; !exists {
					itemNames[item.Name] = nil
					items = append(items, item)
				}
			}
		} else {
			// Breadth first; collect unique fields and any sub structs
			subStructs := []reflect.StructField{}
			for i, n := 0, vt.NumField(); i < n; i++ {
				fld := vt.Field(i)
				if _, exists := itemNames[fld.Name]; !exists {
					itemNames[fld.Name] = nil
					
					if fld.Type.Kind() != reflect.Func
					
					items = append(items, FlatStructItem{
						Mode: Field,
						Name: fld.Name,
						Type: fld.Type,
						Func: 
						Tag: fld.Tag,
					})
	
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
