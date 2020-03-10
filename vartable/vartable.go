package vartable

import (
	"bytes"
	"encoding/gob"
	"sync"
)

// VT is the vartable
type VT struct {
	sync.RWMutex
	vars map[string]interface{}
}

// New initializes the vartable
func New() *VT {
	vt := VT{}
	vt.vars = make(map[string]interface{})
	return &vt
}

// Copy performs a deep copy of the vartable
func (v *VT) Copy() (map[string]interface{}, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	v.RLock()
	err := enc.Encode(v.vars)
	if err != nil {
		return nil, err
	}
	v.RUnlock()
	var copy map[string]interface{}
	err = dec.Decode(&copy)
	if err != nil {
		return nil, err
	}
	return copy, nil
}
