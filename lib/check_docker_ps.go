package checkdockerps

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/mackerelio/checkers"
	"os"
)

// Do the plugin
func Do() {
	ckr := run(os.Args[1:])
	ckr.Name = "DockerPs"
	ckr.Exit()
}

func run(args []string) *checkers.Checker {
	if len(args) != 1 {
		return checkers.NewChecker(checkers.UNKNOWN, "No container name specified")
	}
	containerName := args[0]

	ctx := context.Background()
	client, err := client.NewEnvClient()
	if err != nil {
		return checkers.NewChecker(checkers.UNKNOWN, err.Error())
	}

	containers, err := client.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return checkers.NewChecker(checkers.UNKNOWN, err.Error())
	}

	checkSt := checkers.CRITICAL
	for _, container := range containers {
		if container.Names[0] == "/"+containerName {
			return checkers.NewChecker(checkers.OK, fmt.Sprintf("Process %s exists", containerName))
		}
	}
	return checkers.NewChecker(checkSt, fmt.Sprintf("Process %s does not exists", containerName))
}
