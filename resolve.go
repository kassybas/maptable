package maptable

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antonmedv/expr"
)

// Eval evaluates the given expression using the variable stored in the maptable
func (v *VT) eval(exp string, parseNumber bool) (interface{}, error) {
	v.RLock()
	defer v.RUnlock()
	if strings.HasPrefix(exp, "$") {
		output, err := expr.Eval(exp, v.vars)
		if err != nil {
			return nil, fmt.Errorf("failed to evalute expression: %s\n%w", exp, err)
		}
		return output, nil
	}
	if parseNumber {
		if i, err := strconv.Atoi(exp); err == nil {
			return i, nil
		}
	}
	return exp, nil
}

// Resolve resolves the given expression
func (v *VT) Resolve(value interface{}) (interface{}, error) {
	switch value := value.(type) {
	case string:
		return v.eval(value, false)
	default:
		// TODO: deep eval maps and slices
		return value, nil
	}
}
