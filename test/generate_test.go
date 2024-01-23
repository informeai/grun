package test

import (
	"fmt"
	"testing"

	"github.com/informeai/grun/lib"
)

func TestCreateRepoDir(t *testing.T) {
	if err := lib.CreateRepoDir(); err != nil {
		t.Errorf("TestCreateRepoDir: expect(nil) - got(%s)\n", err.Error())
	}
}

func TestCreateFormula(t *testing.T) {
	name := "formula1"
	if err := lib.CreateFormula(name); err != nil {
		t.Errorf("TestCreateFormula: expect(nil) - got(%s)\n", err.Error())
	}
}

func TestUpdateFormula(t *testing.T) {
	name := "deploy"
	formula := lib.Formula{
		Commands: []lib.Commands{
			{
				Name:   "echo_deploy",
				Action: "echo deploy_updated",
			},
		},
	}
	if err := lib.UpdateFormula(name, formula); err != nil {
		t.Errorf("TestUpdateFormula: expect(nil) - got(%s)\n", err.Error())
	}
}

func TestGetFormula(t *testing.T) {
	name := "deploy"
	formula, err := lib.GetFormula(name)
	if err != nil {
		t.Errorf("TestGetFormula: expect(nil) - got(%s)\n", err.Error())
	}
	fmt.Printf("formula: %+v\n", formula)
}

func TestListFormula(t *testing.T) {
	formulas, err := lib.ListFormula()
	if err != nil {
		t.Errorf("TestListFormula: expect(nil) - got(%s)\n", err.Error())
	}
	fmt.Printf("formulas: %+v\n", formulas)
}

func TestAddFormula(t *testing.T) {
	name := "echo_deploy"
	description := "deploy description"
	action := "echo deploy"
	formulaName := "deploy"
	cmds := lib.Commands{Name: name, Description: description, Action: action}
	if err := lib.AddFormula(formulaName, cmds); err != nil {
		t.Errorf("TestAddFormula: expect(nil) - got(%s)\n", err.Error())
	}
}

func TestDeleteFormula(t *testing.T) {
	name := "deploy"
	if err := lib.DeleteFormula(name); err != nil {
		t.Errorf("TestDeleteFormula: expect(nil) - got(%s)\n", err.Error())
	}
}
