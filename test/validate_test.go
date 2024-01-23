package test

import (
	"log"
	"testing"

	"github.com/informeai/grun/lib"
)

func TestValidateVerifyRepoExist(t *testing.T) {
	exist := lib.VerifyRepoExist()
	log.Printf("exist: %v\n", exist)
}

func TestValidateVerifyFormula(t *testing.T) {
	name := "formula1"
	if err := lib.VerifyFormula(name); err != nil {
		t.Errorf("TestValidateVerifyFormula: expect(nil) - got(%s)\n", err.Error())
	}
}
