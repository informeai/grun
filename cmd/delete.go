package cmd

import (
	"fmt"

	"github.com/informeai/grun/lib"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete formula the repo",
	Long:  `delete formula the repository and actions in formula.`,
	Run: func(cmd *cobra.Command, args []string) {
		flagName := cmd.Flag("name")
		err := lib.VerifyFormula(flagName.Value.String())
		if err != nil && err == lib.ErrFormulaAlreadyExist {
			if errDlt := lib.DeleteFormula(flagName.Value.String()); errDlt != nil {
				fmt.Printf("failed in delete formula -> %s\n", flagName.Value.String())
				return
			}
			fmt.Printf("deleted formula -> %s\n", flagName.Value.String())
			return
		}
		if err != nil && err == lib.ErrFormulaNotFound {
			fmt.Printf("formula not found -> %s\n", flagName.Value.String())
			return
		}
	},
}
var nameFlagDel string

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVarP(&nameFlagDel, "name", "n", "", "name the formula")
	deleteCmd.MarkFlagRequired("name")
}
