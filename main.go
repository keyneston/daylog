package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
	"github.com/keyneston/daylog/cmds/addcmd"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&addcmd.AddCommand{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))

}
