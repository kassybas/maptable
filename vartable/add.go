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
	var err error
	fields, err := splitFields(path)
	if err != nil {
		return fmt.Errorf("failed to parse path: %s\n%w", path, err)
	}
	interFields := make([]interface{}, len(fields))
	interFields[0] = fields[0]
	// Resolve fields except for the first field
	// which is the name of the variable
	for i, field := range fields[1:] {
		interFields[i+1], err = v.Eval(field)
		if err != nil {
			return fmt.Errorf("incorrect variable path: %s\n%w", path, err)
		}
	}
	newVar := unflattenPath(interFields, value)
	v.Lock()
	defer v.Unlock()
	// recover mergo Merge panic
	defer recover()
	defer func() {
		if recover() != nil {
			err = fmt.Errorf("cannot add to variable: %s\n%w", path, err)
		}
	}()
	if err := mergo.Merge(&v.vars, &newVar, mergo.WithOverride); err != nil {
		return fmt.Errorf("cannot add to variable: %s\n%w", path, err)
	}
	return err
}
