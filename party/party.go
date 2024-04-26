package party

import (
	"slices"
	"strings"
)

// A Bouncer accepts guests to a party and reject everyone else.
type Bouncer struct {
	guests []string
}

// NewBouncer returns a new Bouncer whose list of case-insensitive guest
// names is guests.
func NewBouncer(guests ...string) Bouncer {
	set := make(map[string]struct{})
	for _, guest := range guests {
		set[strings.ToLower(guest)] = struct{}{}
	}
	normalized := make([]string, 0, len(set))
	for guest := range set {
		normalized = append(normalized, guest)
	}
	return Bouncer{guests: normalized}
}

// Check verifies whether csv is a list of unique, lowercase, comma-separated
// names of guests;
// if so, it returns that list and true;
// otherwise, it returns the empty string and false.
func (b Bouncer) Check(csv string) (string, bool) {
	var accepted []string
	if csv == "" {
		return "", true
	}
	names := strings.Split(csv, ",")
	for _, name := range names {
		var ok bool
		for _, guest := range b.guests {
			normalized := strings.ToLower(guest)
			if name == normalized {
				accepted = append(accepted, normalized)
				ok = true
				break
			}
		}
		if !ok {
			return "", false
		}
	}
	deduped := slices.Clone(accepted)
	slices.Sort(deduped)
	deduped = slices.Compact(deduped)
	if len(deduped) < len(accepted) {
		return "", false
	}
	return strings.Join(accepted, ","), true
}
