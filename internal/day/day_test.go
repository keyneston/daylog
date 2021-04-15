package day

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/go-test/deep"
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
		{"bar", newDay(2020, 12, 23), "bar/2020/12/23.md"},
	}

	for _, c := range cases {
		out := c.input.FileName(c.base)
		if out != c.expected {
			t.Errorf("Day.FileName(%q) = %q; Want %q", c.base, out, c.expected)
		}
	}
}

func TestParseFile(t *testing.T) {
	type testCase struct {
		name        string
		input       string
		expected    *Day
		expectedErr error
	}

	cases := []testCase{
		{name: "Single Complete", input: "# What are you doing today?\n - Foo Bar\n", expected: NewDay(time.Time{})},
	}

	for _, c := range cases {
		d := NewDay(time.Time{})

		err := d.Parse(bytes.NewBufferString(c.input))
		if c.expectedErr != err {
			t.Errorf("Day.Parse(%q) = %v; want %v", c.name, err, c.expectedErr)
			continue
		}
		if diff := deep.Equal(c.expected, d); diff != nil {
			t.Errorf("Day.Parse(%q) =\n%v", c.name, strings.Join(diff, "\n"))
			continue
		}
		if d.Completed.Len() != c.expected.Completed.Len() {
			t.Errorf("Day.Parse(%q).Completed.Len() = %v; want %v", c.name, d.Completed.Len(), c.expected.Completed.Len())
		}
	}
}
