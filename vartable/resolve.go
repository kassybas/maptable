package vartable

import (
	"fmt"
	"strings"

	"github.com/antonmedv/expr"
)

// Eval evaluates the given expression using the variable stored in the vartable
func (v *VT) Eval(exp string) (interface{}, error) {
	v.RLock()
	defer v.RUnlock()
	if !strings.HasPrefix(exp, "$") {
		return exp, nil
	}
	output, err := expr.Eval(exp, expr.Env(v.vars))
	if err != nil {
		return nil, fmt.Errorf("failed to evalute expression: %s\n%w", exp, err)
	}
	return output, nil
}
