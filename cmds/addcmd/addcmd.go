package addcmd

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
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
	baseDir string
	AddType AddType
}

func (c *AddCommand) Name() string     { return string(c.AddType) }
func (c *AddCommand) Synopsis() string { return synopsis[c.AddType] }
func (c *AddCommand) Usage() string {
	return fmt.Sprintf(`%s <some text>:
	Add a new entry to your day. If '-m' isn't provided it will open in $EDITOR
`, string(c.AddType))
}

func (c *AddCommand) SetFlags(f *flag.FlagSet) {
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

	message := strings.Join(f.Args(), " ")
	if len(message) == 0 {
		log.Printf("Must supply a message to add")
		return subcommands.ExitFailure
	}

	switch c.AddType {
	case AddCompleted:
		d.Completed.Add(message)
	case AddNext:
		d.Next.Add(message)
	case AddBlocked:
		d.Blockers.Add(message)
	}
	if err := d.WriteFile(c.baseDir); err != nil {
		log.Printf("Error: %v", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
