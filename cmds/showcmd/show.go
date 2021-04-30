package showcmd

import (
	"context"
	"flag"
	"log"
	"os"
	"path"
	"time"

	"github.com/google/subcommands"
	"github.com/keyneston/daylog/internal/day"
)

type Command struct {
	baseDir string
}

func (*Command) Name() string     { return "show" }
func (*Command) Synopsis() string { return "Display today's post" }
func (*Command) Usage() string {
	return `show:
`
}

func (c *Command) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.baseDir, "b", os.Getenv("DAYLOG_BASE"), "Base for log entries to be stored. Defaults to $DAYLOG_BASE, or ~/.daylog/")
}

func (c *Command) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if c.baseDir == "" {
		c.baseDir = path.Join(os.Getenv("HOME"), ".daylog")
	}

	d := day.NewDay(time.Now())
	if _, err := d.ReadFile(c.baseDir); err != nil {
		log.Printf("Error: %v", err)
		return subcommands.ExitFailure
	}

	if err := d.Write(os.Stdout); err != nil {
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
