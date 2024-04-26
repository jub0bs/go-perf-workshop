package party

import (
	"slices"
	"strings"

	"github.com/jub0bs/go-perf-workshop/party/internal"
)

// A Bouncer accepts guests to a party and reject everyone else.
type Bouncer struct {
	guests internal.SortedSet
}

// NewBouncer returns a new Bouncer whose list of case-insensitive guest
// names is guests.
func NewBouncer(guests ...string) Bouncer {
	guests = slices.Clone(guests)
	for i := range guests {
		guests[i] = strings.ToLower(guests[i])
	}
	return Bouncer{guests: internal.NewSet(guests...)}
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
		name          string
		commaPos      int
		posOfLastSeen = -1
	)
	s := csv
	for {
		end := min(b.guests.MaxLen()+1, len(s))
		commaPos = strings.IndexByte(s[:end], ',')
		if commaPos == -1 {
			name = s
		} else {
			name = s[:commaPos]
		}
		pos := b.guests.Position(name)
		if pos == -1 || pos <= posOfLastSeen {
			return "", false
		}
		posOfLastSeen = pos
		if commaPos == -1 {
			break
		}
		s = s[commaPos+1:]
	}
	return csv, true
}
