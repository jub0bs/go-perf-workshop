package internal

type Set map[string]struct{}

func NewSet(elems ...string) Set {
	return make(map[string]struct{}, len(elems))
}

func (s Set) Add(elem string) {
	s[elem] = struct{}{}
}

func (s Set) Contains(elem string) bool {
	_, found := s[elem]
	return found
}
