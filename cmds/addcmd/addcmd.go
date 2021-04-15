package addcmd

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

type AddCommand struct {
	message string
	baseDir string
}

func (*AddCommand) Name() string     { return "add" }
func (*AddCommand) Synopsis() string { return "Add new entry" }
func (*AddCommand) Usage() string {
	return `add [-m <some text>]:
	Add a new entry to your day. If '-m' isn't provided it will open in $EDITOR
`
}

func (c *AddCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.message, "m", "", "Task to add as worked on")
	f.StringVar(&c.baseDir, "b", os.Getenv("DAYLOG_BASE"), "Base for log entries to be stored. Defaults to $DAYLOG_BASE, or ~/.daylog/")
}

func (c *AddCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if c.baseDir == "" {
		c.baseDir = path.Join(os.Getenv("HOME"), ".daylog")
	}

	d := day.NewDay(time.Now())
	d.Completed.Add(c.message)
	if err := d.WriteFile(c.baseDir); err != nil {
		log.Printf("Error: %v", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
