package script

import (
	"testing"
)

type person struct {
	name string
	age  int
}

var personMethods = map[string]interface{}{
	"get_name": getPersonName,
	"set_age":  setPersonName,
	"get_age":  getPersonAge,
}

func newPerson(name string) *person {
	return &person{
		name: name,
		age:  0,
	}
}

func getPersonName(p *person) string {
	return p.name
}

func setPersonName(p *person, age int) {
	p.age = age
}

func getPersonAge(p *person) int {
	return p.age
}

// ----
// Test
// ----

func TestCallLuaFunction(t *testing.T) {
	s := NewLuaScript()
	s.LoadString(`
		function add(a, b)
			return a + b
		end
	`)

	want := 7
	got, err := s.CallFunction("add", 5, 2)

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCallGoFunction(t *testing.T) {
	s := NewLuaScript()
	s.RegisterFunction("sub", func(a int, b int) int {
		return a - b
	})

	want := 10
	got, err := s.CallFunction("sub", 20, 10)

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestRegisterInt(t *testing.T) {
	s := NewLuaScript()
	s.RegisterInt("global_value", 1234567890)
	s.LoadString(`
		function get_glbal_value_plus_one()
			return global_value + 1
		end
	`)

	want := 1234567891
	got, err := s.CallFunction("get_glbal_value_plus_one")

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestRegisterString(t *testing.T) {
	s := NewLuaScript()
	s.RegisterString("global_value", "abc")
	s.LoadString(`
		function get_glbal_value_plus_char()
			return global_value .. "d"
		end
	`)

	want := "abcd"
	got, err := s.CallFunction("get_glbal_value_plus_char")

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestRegisterTypeAndCallGetter(t *testing.T) {
	s := NewLuaScript()
	s.RegisterType("person", newPerson, personMethods)
	s.LoadString(`
		function main()
			p = person.new("Batman")
			return p:get_name()
		end
	`)

	want := "Batman"
	got, err := s.CallFunction("main")

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestRegisterTypeAndCallSetter(t *testing.T) {
	s := NewLuaScript()
	s.RegisterType("person", newPerson, personMethods)
	s.LoadString(`
		function main()
			p = person.new("Batman")
			p:set_age(33)
			return p:get_age()
		end
	`)

	want := 33
	got, err := s.CallFunction("main")

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestLoadFile(t *testing.T) {
	s := NewLuaScript()
	err := s.LoadFile("app_test.lua")

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	want := 1000
	got, err := s.CallFunction("main")

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

// ---------
// Benchmark
// ---------

func BenchmarkDoWork(b *testing.B) {
	myScript := NewLuaScript()

	myScript.LoadFile("mul_test.lua")

	myScript.RegisterFunction("mul", func(n int, f int) int {
		return n * f
	})

	myScript.CallFunction("test")
}
