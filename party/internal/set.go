package internal

import (
	"slices"
)

type SortedSet struct {
	m      map[string]int
	maxLen int
}

func NewSet(elems ...string) SortedSet {
	sorted := slices.Clone(elems)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	m := make(map[string]int, len(sorted))
	var maxLen int
	for i, elem := range sorted {
		m[elem] = i
		maxLen = max(maxLen, len(elem))
	}
	return SortedSet{
		m:      m,
		maxLen: maxLen,
	}
}

func (s SortedSet) MaxLen() int {
	return s.maxLen
}

func (s SortedSet) Position(elem string) int {
	pos, found := s.m[elem]
	if !found {
		return -1
	}
	return pos
}

func (s SortedSet) Size() int {
	return len(s.m)
}
