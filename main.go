package main

import (
	"os"

	"github.com/AtleSkaanes/gosnake/cli"
	"github.com/AtleSkaanes/gosnake/game"
	"github.com/AtleSkaanes/gosnake/tui"
)

func main() {
	cliArgs := cli.ParseArgs(os.Args)
	cli.CheckArgs(*cliArgs)

	game.Init(cliArgs.Width, cliArgs.Height, cliArgs.Loop)
	tui.Init(cliArgs.Speed, cliArgs.Color)

	if cliArgs.ExitCode {
		os.Exit(game.GetScore())
	}
}
