package party

import (
	"strings"
)

// A Bouncer accepts guests to a party and reject everyone else.
type Bouncer struct {
	guests map[string]struct{}
}

// NewBouncer returns a new Bouncer whose list of case-insensitive guest
// names is guests.
func NewBouncer(guests ...string) Bouncer {
	set := make(map[string]struct{})
	for _, guest := range guests {
		set[strings.ToLower(guest)] = struct{}{}
	}
	return Bouncer{guests: set}
}

// Check verifies whether csv is a list of unique, lowercase, comma-separated
// names of guests;
// if so, it returns that list and true;
// otherwise, it returns the empty string and false.
func (b Bouncer) Check(csv string) (string, bool) {
	if csv == "" {
		return "", true
	}
	names := strings.Split(csv, ",")
	last := names[0]
	for _, name := range names[1:] {
		if _, ok := b.guests[name]; !ok {
			return "", false
		}
		if name <= last {
			return "", false
		}
		last = name
	}
	return csv, true
}
