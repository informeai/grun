package cmd

import (
	"fmt"

	"github.com/informeai/grun/lib"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update command the formula",
	Long:  `Update command saved in formula`,
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
		formula, err := lib.GetFormula(flagFormula.Value.String())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		var lstComand lib.Commands
		newCommands := []lib.Commands{}
		for _, comd := range formula.Commands {
			if comd.Name == flagName.Value.String() {
				lstComand = comd
			} else {
				newCommands = append(newCommands, comd)
			}
		}
		if lstComand.Name == flagName.Value.String() {
			lstComand.Action = flagAction.Value.String()
			lstComand.Description = flagDescription.Value.String()
			newCommands = append(newCommands, lstComand)
			formula.Commands = newCommands
			if err := lib.UpdateFormula(flagFormula.Value.String(), formula); err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("updated command -> %s\n", flagName.Value.String())
		} else {
			fmt.Println("command not found")
		}
	},
}

var (
	flagFormulaUpdate     string
	flagNameUpdate        string
	flagDescriptionUpdate string
	flagActionUpdate      string
)

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(&flagFormulaUpdate, "formula", "f", "", "name the formula")
	updateCmd.Flags().StringVarP(&flagNameUpdate, "name", "n", "", "name the command")
	updateCmd.Flags().StringVarP(&flagDescriptionUpdate, "description", "d", "", "description the command")
	updateCmd.Flags().StringVarP(&flagActionUpdate, "action", "a", "", "action command for execution")
	updateCmd.MarkFlagRequired("formula")
	updateCmd.MarkFlagRequired("name")
	updateCmd.MarkFlagRequired("action")
}
