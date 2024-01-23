package lib

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// VerifyRepoExist return of repository exist in home
func VerifyRepoExist() bool {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return false
	}
	path := filepath.Join(home, ".grun")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// VerifyFormula execute verification the formula already exist in repo
func VerifyFormula(fileName string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	path := filepath.Join(home, ".grun", fileName)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return ErrFormulaNotFound
	} else {
		return ErrFormulaAlreadyExist
	}
}
