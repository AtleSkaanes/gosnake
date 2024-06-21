package cli

import (
	"github.com/fred1268/go-clap/clap"
)

type CliArgs struct {
	Width  uint8 `clap:"--width,-w"`
	Height uint8 `clap:"--height,-h"`
	Loop   bool  `clap:"--loop"`
}

func ParseArgs(args []string) *CliArgs {
	var err error

	cliArgs := &CliArgs{Width: 16, Height: 16, Loop: true}
	if _, err = clap.Parse(args, cliArgs); err != nil {
		return nil
	}

	return cliArgs
}
