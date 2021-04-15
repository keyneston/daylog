package day

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"time"
)

var (
	commentLine = regexp.MustCompile(`^\W*#.*$`)
)

type Day struct {
	Date time.Time

	Completed *Entries
	TODO      *Entries
	Blockers  *Entries
}

func NewDay(date time.Time) *Day {
	return &Day{
		Date: date,

		Completed: &Entries{},
		TODO:      &Entries{},
		Blockers:  &Entries{},
	}
}

func (d *Day) Write(w io.Writer) error {
	return nil
}

func (d *Day) WriteFile(base string) error {
	f, err := os.OpenFile(d.FileName(base), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return d.Write(f)
}

func (d *Day) FileName(base string) string {
	return path.Join(base,
		fmt.Sprintf("%d/%02d/%02d.md", d.Date.Year(), d.Date.Month(), d.Date.Day()))
}

func (d *Day) Parse(input io.Reader) error {
	r := bufio.NewReader(input)

	var line []byte
	var err error

	for {
		line, err = r.ReadBytes('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}

			return err
		}

		if err := d.parseLine(line); err != nil {
			return err
		}
	}
}

func (d *Day) parseLine(line []byte) error {
	line = bytes.TrimSpace(line)

	if len(line) == 0 {
		return nil
	}

	if commentLine.Match(line) {
		return nil
	}

	d.Completed.Add(string(line))
	return nil
}
