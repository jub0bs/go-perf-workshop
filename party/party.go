package party

import (
	"strings"
)

// A Bouncer accepts guests to a party and reject everyone else.
type Bouncer struct {
	guests map[string]int
}

// NewBouncer returns a new Bouncer whose list of case-insensitive guest
// names is guests.
func NewBouncer(guests ...string) Bouncer {
	set := make(map[string]int)
	for i, guest := range guests {
		set[strings.ToLower(guest)] = i
	}
	return Bouncer{guests: set}
}

// Check verifies whether csv is a list of unique, lowercase, comma-separated
// names of guests;
// if so, it returns that list and true;
// otherwise, it returns the empty string and false.
func (b Bouncer) Check(csv string) (string, bool) {
	s := csv
	var name string
	//seen := make(map[string]struct{})
	seen := make([]bool, len(b.guests))
	for len(s) > 0 {
		name, s, _ = strings.Cut(s, ",")
		pos, found := b.guests[name]
		if !found {
			return "", false
		}
		if seen[pos] {
			return "", false
		}
		seen[pos] = true
	}
	return csv, true
}
