package party

import (
	"slices"
	"strings"
	"sync"

	"github.com/jub0bs/go-perf-workshop/party/internal"
)

// A Bouncer accepts guests to a party and reject everyone else.
type Bouncer struct {
	guests internal.SortedSet
	pool   *sync.Pool
}

// NewBouncer returns a new Bouncer whose list of case-insensitive guest
// names is guests.
func NewBouncer(guests ...string) Bouncer {
	guests = slices.Clone(guests)
	for i := range guests {
		guests[i] = strings.ToLower(guests[i])
	}
	set := internal.NewSet(guests...)
	pool := sync.Pool{
		New: func() any {
			bools := make([]bool, len(set))
			return &bools
		},
	}
	return Bouncer{guests: set, pool: &pool}
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
	seen := b.pool.Get().(*[]bool)
	defer func() {
		clear(*seen)
		b.pool.Put(seen)
	}()
	for {
		name, s, commaFound = strings.Cut(s, ",")
		//normalized := strings.ToLower(name)
		pos := b.guests.Position(name)
		if pos == -1 || (*seen)[pos] {
			return "", false
		}
		(*seen)[pos] = true
		if !commaFound {
			break
		}
	}
	return csv, true
}
