package cmd

import (
	"fmt"

	"github.com/informeai/grun/lib"
	"github.com/spf13/cobra"
)

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "describe commands in formula",
	Long:  `Describe commands saved in formulas.`,
	Run: func(cmd *cobra.Command, args []string) {
    formulaFlag := cmd.Flag("formula")
		if exist := lib.VerifyRepoExist(); !exist {
			fmt.Printf("repo not exist\n")
			return
		}
    formula, err := lib.GetFormula(formulaFlag.Value.String())
    if err != nil{
      fmt.Printf("failed describe formula -> %s\n",formulaFlag.Value.String())
      return
    }
    fmt.Println("DESCRIBE:")
    for _, comd := range formula.Commands{
      fmt.Printf("-----------\nName: %s\nDescription: %s\nAction: %s\n-----------\n",comd.Name,comd.Description,comd.Action)
    }
	},
}
var flagFormulaDesc string

func init() {
	rootCmd.AddCommand(describeCmd)

	describeCmd.Flags().StringVarP(&flagFormulaDesc, "formula", "f", "", "formula name")
	describeCmd.MarkFlagRequired("formula")
}
