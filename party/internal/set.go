package internal

import (
	"slices"
)

type SortedSet map[string]int

func NewSet(elems ...string) SortedSet {
	sorted := slices.Clone(elems)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	set := make(SortedSet, len(sorted))
	for i, elem := range sorted {
		set[elem] = i
	}
	return set
}

func (s SortedSet) Position(elem string) int {
	pos, found := s[elem]
	if !found {
		return -1
	}
	return pos
}
