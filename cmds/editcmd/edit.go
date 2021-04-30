package editcmd

import (
	"context"
	"flag"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/google/subcommands"
	"github.com/keyneston/daylog/internal/day"
)

const (
	DEFAULT_EDITOR = "vi"
)

type EditCommand struct {
	baseDir string
}

func (*EditCommand) Name() string     { return "edit" }
func (*EditCommand) Synopsis() string { return "Edit today's entry" }
func (*EditCommand) Usage() string {
	return `edit:
	Opens the current day's log in $EDITOR. Defaults to 'vi' if no editor is set.
`
}

func (c *EditCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.baseDir, "b", os.Getenv("DAYLOG_BASE"), "Base for log entries to be stored. Defaults to $DAYLOG_BASE, or ~/.daylog/")
}

func (c *EditCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if c.baseDir == "" {
		c.baseDir = path.Join(os.Getenv("HOME"), ".daylog")
	}

	d := day.NewDay(time.Now())

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = DEFAULT_EDITOR
	}

	cmd := exec.Command(editor, d.FileName(c.baseDir))
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()

	return subcommands.ExitSuccess
}
