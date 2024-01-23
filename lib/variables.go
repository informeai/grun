package lib

import "errors"

var (
	ErrFormulaAlreadyExist = errors.New("formula already exist")
	ErrCommandAlreadyExist = errors.New("command already exist")
	ErrFormulaNotFound     = errors.New("formula not found")
	ErrCommandNotFound     = errors.New("command not found")
)
