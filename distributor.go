package main

import (
	"fmt"
)

type Distributor struct {
	Name string

	includes map[*Region]struct{}
	excludes map[*Region]struct{}

	parent *Distributor
}

func NewDistributor(name string, parent *Distributor) *Distributor {
	return &Distributor{
		Name:     name,
		parent:   parent,
		includes: map[*Region]struct{}{},
		excludes: map[*Region]struct{}{},
	}
}

func (d *Distributor) AddInclude(codes ...string) error {
	if len(codes) == 0 {
		return fmt.Errorf("no codes provided")
	}
	for _, c := range codes {
		region, ok := regionMap[c]
		if !ok {
			return fmt.Errorf("not a valid region")
		}

		if d.parent != nil && !d.parent.can(region) {
			return fmt.Errorf("not permissible by parent")
		}

	}

	for _, c := range codes {
		d.includes[regionMap[c]] = struct{}{}
	}

	return nil
}
func (d *Distributor) AddExclude(codes ...string) error {
	if len(codes) == 0 {
		return fmt.Errorf("no codes provided")
	}
	for _, c := range codes {
		_, ok := regionMap[c]
		if !ok {
			return fmt.Errorf("not a valid region")
		}
	}

	for _, c := range codes {
		d.excludes[regionMap[c]] = struct{}{}

	}
	return nil
}

func (d *Distributor) can(region *Region) bool {
	if d.isExcluded(region) {
		return false
	}

	if d.isIncluded(region) {
		if d.parent == nil {
			return true
		}
		return d.parent.can(region)
	}

	return false
}

func (d *Distributor) HasPermission(code string) bool {
	region, ok := regionMap[code]
	if !ok {
		return ok
	}
	return d.can(region)
}

func (d *Distributor) isIncluded(region *Region) bool {
	for r := region; r != nil; r = r.Parent {
		if _, ok := d.includes[r]; ok {
			return true
		}
	}
	return false
}

func (d *Distributor) isExcluded(region *Region) bool {
	for r := region; r != nil; r = r.Parent {
		if _, ok := d.excludes[r]; ok {
			return true
		}
	}
	return false
}
