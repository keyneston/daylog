package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
	"github.com/keyneston/daylog/cmds/addcmd"
	"github.com/keyneston/daylog/cmds/compilecmd"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&addcmd.AddCommand{AddType: addcmd.AddCompleted}, "")
	subcommands.Register(&addcmd.AddCommand{AddType: addcmd.AddNext}, "")
	subcommands.Register(&addcmd.AddCommand{AddType: addcmd.AddBlocked}, "")
	subcommands.Register(&compilecmd.CompileCommand{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))

}
