package goreflect

import (
	"fmt"
	"reflect"
	"strings"
)

// TypeMatch describes a single match by any one of multiple kinds and/or types
type TypeMatch struct {
	types          []reflect.Type
	kinds          []reflect.Kind
	minIndirection int
	maxIndirection int
}

// String returns a signature for the types matched, use vertical bars to separate multiple choices.
// If pointer indirections are allowed, they occur once at the beginning.
// If there are multiple pointer indirections, they are in parantheses.
// If there are multiple type choices, they are in parantheses only if there is at least one pointer indirection
// Examples:
// "string"
// "*string"
// "(*|**)string"
// "string|slice"
// "*(string|slice)"
// "(*|**)(string|slice)"
func (tm TypeMatch) String() string {
	var str strings.Builder

	// Add pointer indirection(s)
	var numIndirections int
	if tm.maxIndirection > 0 {
		numIndirections = tm.maxIndirection - tm.minIndirection + 1
	}
	multipleIndirections := numIndirections > 1
	useBrackets := tm.minIndirection == 0

	if multipleIndirections {
		if useBrackets {
			str.WriteRune('[')
		} else {
			str.WriteRune('(')
		}
	}

	for i := 1; i <= int(tm.maxIndirection); i++ {
		if i > 1 {
			str.WriteRune('|')
		}

		str.WriteString(strings.Repeat("*", i))
	}

	if multipleIndirections {
		if useBrackets {
			str.WriteRune(']')
		} else {
			str.WriteRune(')')
		}
	}

	// Add type(s), then kind(s)
	needTypeParens := (numIndirections > 1) && ((len(tm.types) + len(tm.kinds)) > 1)
	if needTypeParens {
		str.WriteRune('(')
	}

	firstType := true
	for _, typ := range tm.types {
		if !firstType {
			str.WriteRune('|')
		}
		firstType = false

		str.WriteString(typ.String())
	}

	for _, kind := range tm.kinds {
		if !firstType {
			str.WriteRune('|')
		}
		firstType = false

		str.WriteString(kind.String())
	}

	if needTypeParens {
		str.WriteRune(')')
	}

	return str.String()
}

// NewTypeMatch constructs a TypeMatch
// The value passed can be a value, reflect.Value, reflect.Type, or reflect.Kind.
// Indirection may have up two ints, as follows:
// - 0 ints: minIndirection = maxIndirection = 0
// - 1 int:  minIndirection = maxIndirection = int
// - 2 ints: minIndirection = first int, maxIndirection = second int
// Panics if maxIndirection < minIndirection
func NewTypeMatch(val interface{}, indirection ...int) TypeMatch {
	var types []reflect.Type
	var kinds []reflect.Kind

	typ, kind := GetReflectTypeOrKindValueOf(val)
	if kind == reflect.Invalid {
		types = []reflect.Type{typ}
	} else {
		kinds = []reflect.Kind{kind}
	}

	minIndirection := Value
	maxIndirection := Value
	if len(indirection) >= 1 {
		minIndirection = indirection[0]
		maxIndirection = minIndirection
	}
	if len(indirection) >= 2 {
		maxIndirection = indirection[1]
	}

	if maxIndirection < minIndirection {
		panic(fmt.Errorf("NewTypeMatch: maxIndirection %d < minIndirection %d", maxIndirection, minIndirection))
	}

	return TypeMatch{
		types:          types,
		kinds:          kinds,
		minIndirection: minIndirection,
		maxIndirection: maxIndirection,
	}
}

// NewMultiTypeMatch constructs a TypeMatch that can match against any of several choices.
// Each choice is the same as for NewTypeMatch.
func NewMultiTypeMatch(
	minIndirection int,
	maxIndirection int,
	vals ...interface{},
) TypeMatch {
	var types []reflect.Type
	var kinds []reflect.Kind

	for _, val := range vals {
		typ, kind := GetReflectTypeOrKindValueOf(val)
		if kind == reflect.Invalid {
			types = append(types, typ)
		} else {
			kinds = append(kinds, kind)
		}
	}

	return TypeMatch{
		types:          types,
		kinds:          kinds,
		minIndirection: minIndirection,
		maxIndirection: maxIndirection,
	}
}

// Matches returns true if this type matches the given reflect type.
// If the given type is nil, false is returned.
func (tm TypeMatch) Matches(t reflect.Type) bool {
	if t == nil {
		return false
	}

	// Get the given type as a zero indirection value type, counting indirections
	valueType := DerefdReflectType(t)
	indirection := NumRefs(t)

	// Check indirection levels first
	if (indirection < tm.minIndirection) || (indirection > tm.maxIndirection) {
		return false
	}

	// Check if any type matches the zero indirection value type
	for _, typ := range tm.types {
		if typ == valueType {
			return true
		}
	}

	// Check if any kind matches the zero indirection value kind
	valueKind := valueType.Kind()
	for _, kind := range tm.kinds {
		if kind == valueKind {
			return true
		}
	}

	return false
}

// Types is the types accessor
func (tm TypeMatch) Types() []reflect.Type {
	typesCopy := make([]reflect.Type, len(tm.types))
	copy(typesCopy, tm.types)
	return typesCopy
}

// Kinds is the kinds accessor
func (tm TypeMatch) Kinds() []reflect.Kind {
	kindsCopy := make([]reflect.Kind, len(tm.kinds))
	copy(kindsCopy, tm.kinds)
	return kindsCopy
}

// MinIndirection is the minIndirection accessor
func (tm TypeMatch) MinIndirection() int {
	return tm.minIndirection
}

// MaxIndirection is the maxIndirection accessor
func (tm TypeMatch) MaxIndirection() int {
	return tm.maxIndirection
}
