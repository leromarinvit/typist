package typist

import (
	"os"
	"reflect"
	"testing"
)

func TestRtypeAssumption(t *testing.T) {
	pRtypeType := reflect.ValueOf(reflect.TypeOf(0)).Type()
	if pRtypeType.String() != "*reflect.rtype" {
		t.Fatal("Expected *reflect.rtype, got", pRtypeType.String())
	}
}

func TestValueAssumption(t *testing.T) {
	field, ok := reflect.TypeOf(reflect.Value{}).FieldByName("ptr")
	if !ok {
		t.Fatal("reflect.Value has no 'ptr' field")
	}
	if field.Type.String() != "unsafe.Pointer" {
		t.Fatalf("reflect.Value.ptr has wrong type %s (expected unsafe.Pointer)", field.Type.String())
	}
}

func TestTypeByString(t *testing.T) {
	typeName := "*os.File"
	var value *os.File

	typ, err := TypeByString(typeName)

	if err != nil {
		t.Fatal(err)
	}

	if typ.String() != typeName {
		t.Fatalf("Expected %s, got %s", typeName, typ.String())
	}

	if !typ.AssignableTo(reflect.TypeOf(value)) {
		t.Fatalf("%s not assignable to %s", typ.String(), reflect.TypeOf(value).String())
	}
}
