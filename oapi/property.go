package oapi

import (
	"reflect"
)

type Property struct {
	Required    bool
	Name        string
	Description string
	Type        string
	SubType     string
	Enum        []string
	In          string
	Format      string
	Children    []Property
}

func (p Property) IsEmpty() bool {
	return reflect.DeepEqual(p, Property{})
}

type Properties []Property

func (ps Properties) IsEmpty() bool {
	for _, p := range ps {
		if !p.IsEmpty() {
			return false
		}
	}
	return true
}
