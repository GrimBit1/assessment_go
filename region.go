package main

type Region struct {
	Name     string
	Parent   *Region
	Children map[string]*Region
}

var World = &Region{
	Name:     "World",
	Parent:   nil,
	Children: make(map[string]*Region),
}

var regionMap = map[string]*Region{}

func NewRegion(name string, parent *Region) *Region {
	return &Region{
		Name:     name,
		Parent:   parent,
		Children: make(map[string]*Region),
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
