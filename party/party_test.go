package party_test

import (
	"testing"

	"github.com/jub0bs/go-perf-workshop/party"
)

func TestBouncerCheck(t *testing.T) {
	type TestCase struct {
		desc   string
		guests []string
		csv    string
		want   string
		ok     bool
	}
	cases := []TestCase{
		{
			desc:   "no guests, empty csv",
			guests: []string{""},
			csv:    "",
			want:   "",
			ok:     true,
		}, {
			desc:   "one guest, empty csv",
			guests: []string{"Foo"},
			csv:    "",
			want:   "",
			ok:     true,
		}, {
			desc:   "one guest in csv",
			guests: []string{"Foo", "Bar"},
			csv:    "foo",
			want:   "foo",
			ok:     true,
		}, {
			desc:   "all guests",
			guests: []string{"Foo", "Bar"},
			csv:    "bar,foo",
			want:   "bar,foo",
			ok:     true,
		}, {
			desc:   "all guests in wrong order",
			guests: []string{"Foo", "Bar"},
			csv:    "foo,bar",
			want:   "",
			ok:     false,
		}, {
			desc:   "one non-invited name",
			guests: []string{"Foo", "Bar"},
			csv:    "bar,baz,foo",
			want:   "",
			ok:     false,
		}, {
			desc:   "one guest and one non-invited name",
			guests: []string{"Foo", "Bar"},
			csv:    "bar,baz",
			want:   "",
			ok:     false,
		}, {
			desc:   "one duplicate guest",
			guests: []string{"Foo", "Bar"},
			csv:    "bar,foo,foo",
			want:   "",
			ok:     false,
		},
	}
	for _, c := range cases {
		f := func(t *testing.T) {
			bouncer := party.NewBouncer(c.guests...)
			got, ok := bouncer.Check(c.csv)
			if ok != c.ok || got != c.want {
				const tmpl = "NewBouncer(%#v...).Check(%q): got %q, %t; want %q, %t"
				t.Errorf(tmpl, c.guests, c.csv, got, ok, c.want, c.ok)
			}
		}
		t.Run(c.desc, f)
	}
}
