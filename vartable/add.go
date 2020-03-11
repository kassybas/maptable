package vartable

import (
	"fmt"
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
	// Resolve fields except for the first field
	// which is the name of the variable
	fieldsInter := make([]interface{}, len(fields))
	fieldsInter[0] = fields[0]
	for i, field := range fields[1:] {
		fieldsInter[i+1], err = v.eval(field, true)
		if err != nil {
			return fmt.Errorf("incorrect variable path: %s\n%w", path, err)
		}
	}
	v.RLock()
	_, exists := v.vars[fields[0]]
	v.RUnlock()
	if !exists {
		return fmt.Errorf("refered variable does not exist: %s (in path: %s)", fields[0], path)
	}
	v.Lock()
	v.vars[fields[0]], err = setValueByFields(fieldsInter[1:], value, v.vars[fields[0]])
	v.Unlock()
	if err != nil {
		return err
	}
	return err
}

func (v *VT) AddMultiplePath(paths []string, values []interface{}, allowLessPath true) error {
	if (!allowLessPath && len(paths) < len(values)) || len(values) > len(paths) {
		return fmt.Errorf("not matching names and values: %d (%s) != %d (%s)", len(values), values, len(names), names)
	}
	for i := range paths {
		err := v.AddPath(paths[i], values[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func setValueByFields(fieldsInter []interface{}, value interface{}, curVar interface{}) (interface{}, error) {
	if len(fieldsInter) == 0 {
		return value, nil
	}
	field := fieldsInter[0]
	var err error
	var exists bool
	switch curVar.(type) {
	case map[interface{}]interface{}:
		{
			_, exists = curVar.(map[interface{}]interface{})[field]
			if !exists {
				return nil, fmt.Errorf("field does not exist: %s", field)
			}
			curVar.(map[interface{}]interface{})[field], err =
				setValueByFields(
					fieldsInter[1:],
					value,
					curVar.(map[interface{}]interface{})[field],
				)
			return curVar, err

		}
	case []interface{}:
		{
			index, ok := field.(int)
			if !ok {
				return nil, fmt.Errorf("non-integer index on list: %s", field)
			}
			if index >= len(curVar.([]interface{})) {
				return nil, fmt.Errorf("index out of range: %d", index)
			}
			curVar.([]interface{})[index], err = setValueByFields(
				fieldsInter[1:],
				value,
				curVar.([]interface{})[index],
			)
			return curVar, err
		}
	default:
		return nil, fmt.Errorf("indexing or reference on scalar value: %T[%s]", curVar, field)
	}
}
