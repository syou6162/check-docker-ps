package checkdockerps

import (
	"os"
	"github.com/mackerelio/checkers"
)

// Do the plugin
func Do() {
	ckr := run(os.Args[1:])
	ckr.Name = "DockerPs"
	ckr.Exit()
}

func run(args []string) *checkers.Checker {
	checkSt := checkers.OK

	return checkers.NewChecker(checkSt, "Process exists")
}
