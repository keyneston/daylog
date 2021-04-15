package day

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"
)

type Day struct {
	Date time.Time
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
