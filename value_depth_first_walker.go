package goreflect

import (
	"fmt"
	"reflect"
)

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
