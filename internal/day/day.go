package day

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"time"

	"github.com/fatih/color"
)

var (
	commentLine = regexp.MustCompile(`^\W*#.*$`)
)

const (
	HeaderCompleted = "# What I worked on today"
	HeaderNext      = "# What I’m working on next"
	HeaderBlockers  = "# What’s blocking me"

	HeaderCompiledCompleted = ":: What I worked on today"
	HeaderCompiledNext      = ":: What I’m working on next"
	HeaderCompiledBlockers  = ":: What’s blocking me"

	PrefixDiskEntry     = " - "
	PrefixCompiledEntry = " • "

	DefaultDirPerm  = 0755
	DefaultFilePerm = 0644
)

type state string

const (
	stateUnknown   state = "unknown"
	stateCompleted state = "completed"
	stateBlockers  state = "blockers"
	stateNext      state = "next"
)

type Day struct {
	Date time.Time

	Completed *Entries
	Next      *Entries
	Blockers  *Entries

	state state
}

func NewDay(date time.Time) *Day {
	return &Day{
		Date: date,

		Completed: &Entries{},
		Next:      &Entries{},
		Blockers:  &Entries{},

		state: stateUnknown,
	}
}

func (d *Day) Write(w io.Writer) error {
	d.Completed.Write(w, PrefixDiskEntry, HeaderCompleted)
	d.Next.Write(w, PrefixDiskEntry, HeaderNext)
	d.Blockers.Write(w, PrefixDiskEntry, HeaderBlockers)

	return nil
}

// ReadFile returns true if the file was found and read, otherwise it returns
// false.
func (d *Day) ReadFile(base string) (bool, error) {
	fName := d.FileName(base)

	f, err := os.Open(fName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}

		return false, err
	}
	defer f.Close()

	return true, d.Parse(f)
}

func (d *Day) WriteFile(base string) error {
	fName := d.FileName(base)
	if err := os.MkdirAll(filepath.Dir(fName), DefaultDirPerm); err != nil {
		return err
	}

	f, err := os.OpenFile(fName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, DefaultFilePerm)
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

		err := d.parseLine(line)
		if err != nil {
			return err
		}
	}
}

func (d *Day) parseLine(line []byte) error {
	line = bytes.TrimSpace(line)
	line = bytes.TrimLeft(line, PrefixDiskEntry)

	if len(line) == 0 {
		return nil
	}

	if commentLine.Match(line) {
		switch string(line) {
		case HeaderCompleted:
			d.state = stateCompleted
		case HeaderNext:
			d.state = stateNext
		case HeaderCompiledBlockers:
			d.state = stateBlockers
		}
		return nil
	}

	switch d.state {
	case stateCompleted:
		d.Completed.Add(string(line))
	case stateBlockers:
		d.Blockers.Add(string(line))
	case stateNext:
		d.Next.Add(string(line))
	}

	return nil
}

func (d *Day) Compile(w io.Writer, baseDir string) error {
	bold := color.New(color.Bold).SprintFunc()

	d.Completed.Write(w, PrefixCompiledEntry, bold(HeaderCompiledCompleted))
	d.Next.Write(w, PrefixCompiledEntry, bold(HeaderCompiledNext))
	d.Blockers.Write(w, PrefixCompiledEntry, bold(HeaderCompiledBlockers))

	return nil
}
