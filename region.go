package main

type Region struct {
	Name   string
	Parent *Region
}

var World = &Region{
	Name:   "World",
	Parent: nil,
}

var regionMap = map[string]*Region{}

func NewRegion(name string, parent *Region) *Region {
	return &Region{
		Name:   name,
		Parent: parent,
	}
}

func (candidate *Region) isDescendant(target *Region) bool {
	for r := target; r != nil; r = r.Parent {
		if r == candidate {
			return true
		}
	}
	return false
}
