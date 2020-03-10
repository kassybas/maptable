package vartable

import (
	"fmt"

	"github.com/imdario/mergo"
)

// AddPath adds a new value to the vartable based on dot/bracket notation
// accepted path fields can be separated by . or []
//   eg. foo.bar or foo["bar"]
// if variable field is used, only the [] notation is accepted with unqoted field
//   eg. foo[varbar] where varbar must be a existing variable in the vartable
// Adding a new field to an existing map is possible
func (v *VT) AddPath(path string, value interface{}) error {
	newVar, err := unflattenPath(path, value)
	if err != nil {
		return err
	}
	if err := mergo.Merge(&v.vars, &newVar); err != nil {
		golog.Fatal(err)
	}

	fmt.Println(newVar)
	// TODO
	return nil
}
