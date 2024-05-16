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
	seen := make(map[string]struct{})
	for _, name := range names {
		_, found := b.guests[name]
		if !found {
			return "", false
		}
		_, found = seen[name]
		if found {
			return "", false
		}
		seen[name] = struct{}{}
	}
	return csv, true
}
