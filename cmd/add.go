package cmd

import (
	"fmt"

	"github.com/informeai/grun/lib"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add command to formula",
	Long:  `add command to formula for use.`,
	Run: func(cmd *cobra.Command, args []string) {
		flagFormula := cmd.Flag("formula")
		flagName := cmd.Flag("name")
		flagDescription := cmd.Flag("description")
		flagAction := cmd.Flag("action")
		if exist := lib.VerifyRepoExist(); !exist {
			fmt.Printf("repo not exist\n")
			return
		}
		err := lib.VerifyFormula(flagFormula.Value.String())
		if err != nil && err == lib.ErrFormulaNotFound {
			fmt.Printf(err.Error())
			return
		}
		addCmd := lib.Commands{Name: flagName.Value.String(), Description: flagDescription.Value.String(), Action: flagAction.Value.String()}
		if errCrt := lib.AddFormula(flagFormula.Value.String(), addCmd); errCrt != nil {
			fmt.Println(errCrt.Error())
			return
		}
    fmt.Println("added command")
	},
}

var (
	flagFormula     string
	flagName        string
	flagDescription string
	flagAction      string
)

func init() {
	rootCmd.AddCommand(addCmd)
	
	addCmd.Flags().StringVarP(&flagFormula, "formula", "f", "", "name the formula")
	addCmd.Flags().StringVarP(&flagName, "name", "n", "", "name the command")
	addCmd.Flags().StringVarP(&flagDescription, "description", "d", "", "description the command")
	addCmd.Flags().StringVarP(&flagAction, "action", "a", "", "action command for execution")
	addCmd.MarkFlagRequired("formula")
	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("action")
}
