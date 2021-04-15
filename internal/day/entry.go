package day

import (
	"fmt"
	"io"
)

type Entries struct {
	Entries []string
}

func (e *Entries) Add(in string) {
	e.Entries = append(e.Entries, in)
}

func (e *Entries) Len() int {
	return len(e.Entries)
}

func (e *Entries) String() string {
	return fmt.Sprintf("Entry[%d]", e.Len())
}

func (e *Entries) Write(w io.Writer, header string) error {
	if e.Len() == 0 {
		return nil
	}

	if _, err := fmt.Fprintln(w, header); err != nil {
		return err
	}

	for _, e := range e.Entries {
		if _, err := fmt.Fprintln(w, e); err != nil {
			return err
		}
	}

	return nil
}
