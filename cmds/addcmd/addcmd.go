package addcmd

import (
	"context"
	"flag"
	"log"

	"github.com/google/subcommands"
)

type AddCommand struct {
	message string
}

func (*AddCommand) Name() string     { return "add" }
func (*AddCommand) Synopsis() string { return "Add new entry" }
func (*AddCommand) Usage() string {
	return `add [-m <some text>]:
	Add a new entry to your day. If '-m' isn't provided it will open in $EDITOR
`
}

func (c *AddCommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.message, "m", "", "Message")
}

func (c *AddCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	log.Printf("msg:\n%s", c.message)

	return subcommands.ExitSuccess
}
