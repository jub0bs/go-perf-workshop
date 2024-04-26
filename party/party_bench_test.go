package party_test

import (
	"strings"
	"testing"

	"github.com/jub0bs/go-perf-workshop/party"
)

func BenchmarkBouncerCheck(b *testing.B) {
	type BenchCase struct {
		desc   string
		guests []string
		csv    string
	}
	cases := []BenchCase{
		{
			desc:   "all guests",
			guests: []string{"Foo", "Bar"},
			csv:    "foo,bar",
		}, {
			desc:   "all guests but duplicated many times",
			guests: []string{"Foo", "Bar"},
			csv:    strings.Repeat("foo,bar", 1024),
		}, {
			desc: "maliciously long non-invited name",
			// see https://en.wikipedia.org/wiki/Hubert_Blaine_Wolfeschlegelsteinhausenbergerdorff_Sr.
			guests: []string{"Foo", "Bar"},
			csv:    strings.Repeat("a", 1024),
		},
	}
	for _, c := range cases {
		f := func(b *testing.B) {
			bouncer := party.NewBouncer(c.guests...)
			b.ResetTimer()
			for range b.N {
				bouncer.Check(c.csv)
			}
		}
		b.Run(c.desc, f)
	}
}
