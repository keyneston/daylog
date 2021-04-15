package day

import "fmt"

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
