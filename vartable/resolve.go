package vartable

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antonmedv/expr"
)

// Eval evaluates the given expression using the variable stored in the vartable
func (v *VT) Eval(exp string) (interface{}, error) {
	v.RLock()
	defer v.RUnlock()
	if strings.HasPrefix(exp, "$") {
		output, err := expr.Eval(exp, v.vars)
		if err != nil {
			return nil, fmt.Errorf("failed to evalute expression: %s\n%w", exp, err)
		}
		return output, nil
	}
	if i, err := strconv.Atoi(exp); err == nil {
		return i, nil
	}
	return exp, nil
}
