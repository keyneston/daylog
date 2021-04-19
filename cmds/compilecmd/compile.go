package compilecmd

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

type CompileCommand struct {
	baseDir string
}

func (*CompileCommand) Name() string     { return "compile" }
func (*CompileCommand) Synopsis() string { return "Compile entries for posting" }
func (*CompileCommand) Usage() string {
	return `compile:
`
}

func (c *CompileCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.baseDir, "b", os.Getenv("DAYLOG_BASE"), "Base for log entries to be stored. Defaults to $DAYLOG_BASE, or ~/.daylog/")
}

func (c *CompileCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if c.baseDir == "" {
		c.baseDir = path.Join(os.Getenv("HOME"), ".daylog")
	}

	d := day.NewDay(time.Now())

	if _, err := d.ReadFile(c.baseDir); err != nil {
		log.Printf("Error: %v", err)
		return subcommands.ExitFailure
	}

	if err := d.Compile(os.Stdout, c.baseDir); err != nil {
		log.Printf("Error: %v", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
