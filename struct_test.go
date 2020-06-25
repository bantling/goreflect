package goreflect

import (
	// "fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type myGCStruct struct {
	pf1  string
	cf1  string
	gcf1 int
	gcf2 string
}

func (myGCStruct) Pm1()   {}
func (myGCStruct) Cm1()   {}
func (myGCStruct) Gcm1()  {}
func (*myGCStruct) Gcm2() {}

type myCStruct struct {
	pf1 string
	cf1 int
	cf2 string
	gc  myGCStruct
}

func (myCStruct) Pm1()  {}
func (myCStruct) Cm1()  {}
func (*myCStruct) Cm2() {}

type myPStruct struct {
	pf1 int
	pf2 string
	myCStruct
}

func (myPStruct) Pm1()  {}
func (*myPStruct) Pm2() {}

func TestFlatStruct(t *testing.T) {
	assertFields := func(
		expected []reflect.StructField,
		actual []reflect.StructField,
		actualIter func() (reflect.StructField, bool),
	) {
		// Fields are in order declared
		assert.Equal(t, len(expected), len(actual))
		for i, e := range expected {
			// fmt.Printf("%d, %v, %v\n", i, e, actual[i])
			assert.Equal(t, e, actual[i])
		}

		i := 0
		for a, hasNext := actualIter(); hasNext; a, hasNext = actualIter() {
			assert.Equal(t, expected[i], a)
			i++
		}
	}

	assertMethods := func(
		expected []reflect.Method,
		actual []reflect.Method,
		actualIterGen func() func() (reflect.Method, bool),
	) {
		// Methods are in random order
		assert.Equal(t, len(expected), len(actual))
		for _, e := range expected {
			var found bool
			for _, a := range actual {
				// fmt.Printf("%s, %s, %s, %s\n", e.Name, e.Type, a.Name, a.Type)
				if found = (e.Name == a.Name) && (e.Type == a.Type); found {
					// fmt.Println("equal")
					break
				}
			}
			assert.True(t, found)

			found = false
			actualIter := actualIterGen()
			for a, hasNext := actualIter(); hasNext; a, hasNext = actualIter() {
				// fmt.Printf("%s, %s, %s, %s\n", e.Name, e.Type, a.Name, a.Type)
				if found = (e.Name == a.Name) && (e.Type == a.Type); found {
					// fmt.Println("equal")
					break
				}
			}
			assert.True(t, found)
		}
	}

	fldByName := func(st reflect.Type, name string) reflect.StructField {
		fld, _ := st.FieldByName(name)
		return fld
	}

	mthdByName := func(st reflect.Type, name string) reflect.Method {
		mthd, _ := st.MethodByName(name)
		return mthd
	}

	//// myGCStruct

	gc := myGCStruct{}
	gcf := FlatStructOf(gc)

	gcvt := reflect.TypeOf(gc)
	assertFields(
		[]reflect.StructField{
			fldByName(gcvt, "pf1"),
			fldByName(gcvt, "cf1"),
			fldByName(gcvt, "gcf1"),
			fldByName(gcvt, "gcf2"),
		},
		gcf.fields,
		gcf.FieldIter(),
	)

	assertMethods(
		[]reflect.Method{
			mthdByName(gcvt, "Pm1"),
			mthdByName(gcvt, "Cm1"),
			mthdByName(gcvt, "Gcm1"),
		},
		gcf.valMethods,
		gcf.ValMethodsIter,
	)

	gcpt := reflect.PtrTo(gcvt)
	assertMethods(
		[]reflect.Method{
			mthdByName(gcpt, "Gcm2"),
		},
		gcf.ptrMethods,
		gcf.PtrMethodsIter,
	)

	//// myCStruct

	c := myCStruct{}
	cf := FlatStructOf(c)

	cvt := reflect.TypeOf(c)
	assertFields(
		[]reflect.StructField{
			fldByName(cvt, "pf1"),
			fldByName(cvt, "cf1"),
			fldByName(cvt, "cf2"),
			fldByName(cvt, "gc"),
			fldByName(gcvt, "gcf1"),
			fldByName(gcvt, "gcf2"),
		},
		cf.fields,
		cf.FieldIter(),
	)

	assertMethods(
		[]reflect.Method{
			mthdByName(cvt, "Pm1"),
			mthdByName(cvt, "Cm1"),
			mthdByName(gcvt, "Gcm1"),
		},
		cf.valMethods,
		cf.ValMethodsIter,
	)

	cpt := reflect.PtrTo(cvt)
	assertMethods(
		[]reflect.Method{
			mthdByName(cpt, "Cm2"),
			mthdByName(gcpt, "Gcm2"),
		},
		cf.ptrMethods,
		cf.PtrMethodsIter,
	)

	//// myPStruct

	p := myPStruct{}
	pf := FlatStructOf(p)

	pvt := reflect.TypeOf(p)
	assertFields(
		[]reflect.StructField{
			fldByName(pvt, "pf1"),
			fldByName(pvt, "pf2"),
			fldByName(pvt, "myCStruct"),
			fldByName(cvt, "cf1"),
			fldByName(cvt, "cf2"),
			fldByName(cvt, "gc"),
			fldByName(gcvt, "gcf1"),
			fldByName(gcvt, "gcf2"),
		},
		pf.fields,
		pf.FieldIter(),
	)

	assertMethods(
		[]reflect.Method{
			mthdByName(pvt, "Pm1"),
			// Embedded struct causes reflection to report its methods as declared in parent
			mthdByName(pvt, "Cm1"),
			mthdByName(gcvt, "Gcm1"),
		},
		pf.valMethods,
		pf.ValMethodsIter,
	)

	ppt := reflect.PtrTo(pvt)
	assertMethods(
		[]reflect.Method{
			mthdByName(ppt, "Pm2"),
			// Embedded struct causes reflection to report its methods as declared in parent
			mthdByName(ppt, "Cm2"),
			mthdByName(gcpt, "Gcm2"),
		},
		pf.ptrMethods,
		pf.PtrMethodsIter,
	)
}
