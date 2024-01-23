package lib

import (
	"os"
	"os/exec"
	"strings"
)

// GetCommand return command from formula
func GetCommand(name string, formula Formula) (Commands, error) {
	for _, cmd := range formula.Commands {
		if cmd.Name == name {
			return cmd, nil
		}
	}
	return Commands{}, ErrCommandNotFound
}

// Run execute action command
func Run(command Commands) error {
	argCommands := strings.Split(command.Action, " ")
	cmd := exec.Command(argCommands[0], argCommands[1:]...)
  cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
