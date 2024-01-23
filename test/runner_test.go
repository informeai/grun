package test

import (
	"fmt"
	"testing"

	"github.com/informeai/grun/lib"
)

func TestGetCommand(t *testing.T) {
	formula := lib.Formula{
		Commands: []lib.Commands{
			{
				Name:   "echo_deploy",
				Action: "echo deploy",
			},
		},
	}
	cmd, err := lib.GetCommand("echo_deploy", formula)
	if err != nil {
		t.Errorf("TestGetCommand: expect(nil) - got(%s)\n", err.Error())
	}
	fmt.Printf("cmd: %+v\n", cmd)
}

func TestRun(t *testing.T) {
	cmd := lib.Commands{
		Name:   "echo_deploy",
		Action: "echo deploy",
	}
	if err := lib.Run(cmd); err != nil {
		t.Errorf("TestRun: expect(nil) - got(%s)\n", err.Error())
	}
}
