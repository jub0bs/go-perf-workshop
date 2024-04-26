package party

import (
	"maps"
	"strings"

	"github.com/jub0bs/go-perf-workshop/party/internal"
)

// A Bouncer accepts guests to a party and reject everyone else.
type Bouncer struct {
	guests internal.Set
}

// NewBouncer returns a new Bouncer whose list of case-insensitive guest
// names is guests.
func NewBouncer(guests ...string) Bouncer {
	set := internal.NewSet()
	for _, guest := range guests {
		set.Add(strings.ToLower(guest))
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
	var (
		name       string
		commaFound bool
	)
	s := csv
	guests := maps.Clone(b.guests)
	for {
		name, s, commaFound = strings.Cut(s, ",")
		normalized := strings.ToLower(name)
		if !guests.Contains(normalized) {
			return "", false
		}
		delete(guests, normalized)
		if !commaFound {
			break
		}
	}
	return csv, true
}
