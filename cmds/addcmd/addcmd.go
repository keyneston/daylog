package addcmd

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/google/subcommands"
	"github.com/keyneston/daylog/internal/day"
)

type AddType string

const (
	AddCompleted AddType = "add"
	AddBlocked   AddType = "blocked"
	AddNext      AddType = "next"
)

var synopsis = map[AddType]string{
	AddCompleted: "Add new entry to list of completed tasks",
	AddNext:      "Add new entry to list of next tasks",
	AddBlocked:   "Add new entry to list of blockers",
}

type AddCommand struct {
	message string
	baseDir string
	AddType AddType
}

func (c *AddCommand) Name() string     { return string(c.AddType) }
func (c *AddCommand) Synopsis() string { return synopsis[c.AddType] }
func (c *AddCommand) Usage() string {
	return fmt.Sprintf(`%s [-m <some text>]:
	Add a new entry to your day. If '-m' isn't provided it will open in $EDITOR
`, string(c.AddType))
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

	if _, err := d.ReadFile(c.baseDir); err != nil {
		log.Printf("Error: %v", err)
		return subcommands.ExitFailure
	}

	switch c.AddType {
	case AddCompleted:
		log.Printf("Adding to completed/%s", c.AddType)
		d.Completed.Add(c.message)
	case AddNext:
		log.Printf("Adding to next/%s", c.AddType)
		d.Next.Add(c.message)
	case AddBlocked:
		log.Printf("Adding to blocked/%s", c.AddType)
		d.Blockers.Add(c.message)
	}
	if err := d.WriteFile(c.baseDir); err != nil {
		log.Printf("Error: %v", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
