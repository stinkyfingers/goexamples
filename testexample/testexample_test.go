package testexample

import "testing"

func TestFoobar(t *testing.T) {
	// create instance of Foo and assign address to f
	f := &Foo{
		Name: "The first foo",
	}
	expected := "Foo name: The first foo"

	// test logs print to stdout with -v flag
	t.Log("running test")
	resp := f.Foobar()
	// compare response with expected
	if resp != expected {
		t.Errorf("expected %s, got %s", expected, resp)
	}
}
