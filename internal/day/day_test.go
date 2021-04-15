package day

import (
	"testing"
	"time"
)

func newDay(year, month, day int) Day {
	return Day{
		Date: time.Date(year, time.Month(month), day, 15, 35, 5, 23, time.UTC),
	}
}

func TestFileName(t *testing.T) {
	type testCase struct {
		base     string
		input    Day
		expected string
	}

	cases := []testCase{
		{"foo", newDay(2020, 2, 1), "foo/2020/02/01.md"},
	}

	for _, c := range cases {
		out := c.input.FileName(c.base)
		if out != c.expected {
			t.Errorf("Day.FileName(%q) = %q; Want %q", c.base, out, c.expected)
		}
	}

}
