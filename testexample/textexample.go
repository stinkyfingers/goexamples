package testexample

import "fmt"

// Foo represents an object.
// The capitalized name "Foo" means it can be used as
// an import in other packages.
type Foo struct {
	Name string
}

// Foobar prints a message build from Foo's Name
func (f *Foo) Foobar() string {
	return fmt.Sprintf("Foo name: %s", f.Name)
}
