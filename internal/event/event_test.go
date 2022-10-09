package event

import (
	"testing"
)

func TestEventSetGetField(t *testing.T) {
	e := New()
	e.SetField("name", "Batman")

	want := "Batman"
	got := e.GetField("name")

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestEventHasField(t *testing.T) {
	e := New()
	e.SetField("name", "Batman")

	want := true
	got := e.HasField("name")

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestEventDoesNotHaveField(t *testing.T) {
	e := New()
	e.SetField("name", "Batman")

	want := false
	got := e.HasField("age")

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

// --

func TestDef(t *testing.T) {
	f1 := NewFieldDefinition("name", FIELD_STRING, make(map[string]interface{}))
	f2 := NewFieldDefinition("age", FIELD_NUMBER, map[string]interface{}{})

	ed := NewDefinition([]*FieldDefinition{f1, f2})

	ageDef, err := ed.GetFieldDefinition("age")
	if err != nil {
		panic(err)
	}

	if ageDef.GetFieldType() != FIELD_NUMBER {
		t.Error("zzz")
	}
}
