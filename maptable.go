package maptable

import (
	"bytes"
	"encoding/gob"
	"sync"
)

// VT is the maptable
type VT struct {
	sync.RWMutex
	vars map[string]interface{}
}

// New initializes the maptable
func New() *VT {
	vt := VT{}
	vt.vars = make(map[string]interface{})
	return &vt
}

// Copy performs a deep copy of the maptable
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
