// Package typist gets a reflect.Type for a type specified by a string.
package typist

import (
	"fmt"
	"reflect"
	"unsafe"
)

//go:linkname typesByString reflect.typesByString
func typesByString(s string) []unsafe.Pointer

/*
TypeByString tries to find a reflect.Type corresponding to the type specified by
s.

It calls the unexported `reflect.typesByString` to do so. It will fail if
the type can't be found or if more than one type with the given name exist.

This relies on the following assumptions:
    * The signature of `reflect.typesByString` must not change
    * The value returned by `reflect.TypeOf(0)` is a `*reflect.rtype`
    * The `reflect.Value` struct contains a `ptr` field of type `unsafe.Pointer`
*/
func TypeByString(s string) (reflect.Type, error) {
	types := typesByString(s)

	if len(types) == 0 {
		return nil, fmt.Errorf("Type '%s' not found", s)
	}
	if len(types) > 1 {
		return nil, fmt.Errorf("Type '%s' is ambiguous", s)
	}

	t := types[0]

	pRtypeType := reflect.ValueOf(reflect.TypeOf(0)).Type()
	rtype := reflect.New(pRtypeType).Elem()

	ptr := unsafe.Pointer(reflect.ValueOf(rtype).FieldByName("ptr").Pointer())
	*(*unsafe.Pointer)(ptr) = t

	typ := rtype.Interface().(reflect.Type)
	return typ, nil
}
