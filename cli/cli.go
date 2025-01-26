package cli

import (
	"fmt"
	"os"

	"github.com/fred1268/go-clap/clap"
)

type CliArgs struct {
	Width    uint8 `clap:"--width"`
	Height   uint8 `clap:"--height"`
	Loop     bool  `clap:"--loop"`
	Speed    int   `clap:"--speed,-s"`
	Help     bool  `clap:"--help,-h"`
	Version  bool  `clap:"--version,-v"`
	ExitCode bool  `clap:"--exit-with-code,-e"`
}

func ParseArgs(args []string) *CliArgs {
	var err error

	cliArgs := &CliArgs{Width: 16, Height: 16, Speed: 100, Loop: true, Help: false, Version: false, ExitCode: false}
	if _, err = clap.Parse(args, cliArgs); err != nil {
		os.Exit(1)
	}

	if cliArgs.Help {
		PrintHelp()
		os.Exit(0)
	}
	if cliArgs.Version {
		PrintVersion()
		os.Exit(0)
	}

	return cliArgs
}

func CheckArgs(args CliArgs) {
	if args.Width < 8 {
		fmt.Printf("Invalid Width. Has to be atleast 8. Got %d\n", args.Width)
		os.Exit(1)
	}
	if args.Height < 8 {
		fmt.Printf("Invalid Height. Has to be atleast 8. Got %d\n", args.Height)
		os.Exit(1)
	}
	if args.Speed < 50 {
		fmt.Printf("Invalid Speed. Has to be atleast 50ms. Got %d\n", args.Speed)
		os.Exit(1)
	}
}

func PrintHelp() {
	fmt.Println("GoSnake!")
	fmt.Println("A snake game for the terminal, made in Go (By someone who doesn't know Go)")

	fmt.Println("\n  Options:")
	fmt.Println("    --width <int>          Sets the window width. 16 by default")
	fmt.Println("    --height <int>         Sets the window height. 16 by default")
	fmt.Println("    --loop                 Enables the snake to loop around. Enabled by default")
	fmt.Println("    --no-loop              Disables the snakes ability to loop around")
	fmt.Println("    -s, --speed            Sets the time between frames, in ms. 100ms by default")
	fmt.Println("    -e, --exit-with-code   Exits the program with the score as the exitcode")
	fmt.Println("    -h, --help             Prints this help page")
	fmt.Println("    -v, --version          Prints the version number")
}

func PrintVersion() {
	fmt.Println("Version: 1.0.0")
}
