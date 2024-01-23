package lib

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// CreateRepoDir execute creation the directory repo
func CreateRepoDir() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	repoDir := filepath.Join(home, ".grun")
	if err := os.Mkdir(repoDir, 0o700); err != nil {
		return err
	}
	return nil
}

// CreateFormula execute creation the formula file in repo
func CreateFormula(fileName string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	path := filepath.Join(home, ".grun", fileName)
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	formula := Formula{Commands: []Commands{}}
	bytesFormula, err := json.Marshal(&formula)
	if err != nil {
		return err
	}
	_, err = f.Write(bytesFormula)
	if err != nil {
		return err
	}
	return nil
}

// UpdateFormula execute update the formula file in repo
func UpdateFormula(fileName string, formula Formula) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	path := filepath.Join(home, ".grun", fileName)
	btsFormula, err := json.Marshal(&formula)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(path, btsFormula, os.ModePerm); err != nil {
		return err
	}
	return nil
}

// GetFormula return the formula from repo
func GetFormula(fileName string) (Formula, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Formula{}, err
	}
	path := filepath.Join(home, ".grun", fileName)
	formula := Formula{}
	bts, err := ioutil.ReadFile(path)
	if err != nil {
		return Formula{}, err
	}
	if err := json.Unmarshal(bts, &formula); err != nil {
		return Formula{}, err
	}
	return formula, nil
}

// ListFormula return the all formulas from repo
func ListFormula() ([]string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return []string{}, err
	}
	path := filepath.Join(home, ".grun")
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return []string{}, err
	}
	result := []string{}
	for _, fl := range files {
		if !fl.IsDir() {
			result = append(result, fl.Name())
		}
	}
	return result, nil
}

// AddFormula execute append in formula the repo
func AddFormula(fileName string, command Commands) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	path := filepath.Join(home, ".grun", fileName)
	formula := Formula{}
	bts, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bts, &formula); err != nil {
		return err
	}
	for _, comd := range formula.Commands {
		if comd.Name == command.Name {
			return ErrCommandAlreadyExist
		}
	}
	formula.Commands = append(formula.Commands, command)
	btsWriter, err := json.Marshal(formula)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(path, btsWriter, os.ModeAppend); err != nil {
		return err
	}
	return nil
}

// DeleteFormula execute deletion the formula file in repo
func DeleteFormula(fileName string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	path := filepath.Join(home, ".grun", fileName)
	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}
