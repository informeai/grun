package cmd

import (
	"fmt"

	"github.com/informeai/grun/lib"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove command the formula",
	Long:  `remove command the formula in repository local.`,
	Run: func(cmd *cobra.Command, args []string) {
		flagName := cmd.Flag("name")
		flagFormula := cmd.Flag("formula")
		err := lib.VerifyFormula(flagFormula.Value.String())
		if err != nil && err == lib.ErrFormulaAlreadyExist {
			fml, errGet := lib.GetFormula(flagFormula.Value.String())
			if errGet != nil {
				fmt.Printf("failed in get formula -> %s\n", flagFormula.Value.String())
				return
			}
			newCommands := []lib.Commands{}
			for _, cmd := range fml.Commands {
				if cmd.Name != flagName.Value.String() {
					newCommands = append(newCommands, cmd)
				}
			}
			fml.Commands = newCommands
			if err := lib.UpdateFormula(flagFormula.Value.String(), fml); err != nil {
				fmt.Printf("failed in remove command -> %s\n", flagName.Value.String())
				return
			}
			fmt.Printf("deleted command -> %s\n", flagName.Value.String())
			return
		}
		if err != nil && err == lib.ErrFormulaNotFound {
			fmt.Printf("formula not found -> %s\n", flagName.Value.String())
			return
		}
	},
}

var (
	nameFlag    string
	formulaFlag string
)

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().StringVarP(&nameFlag, "name", "n", "", "name the command")
	removeCmd.Flags().StringVarP(&formulaFlag, "formula", "f", "", "name the formula")
	removeCmd.MarkFlagRequired("name")
	removeCmd.MarkFlagRequired("formula")
}
