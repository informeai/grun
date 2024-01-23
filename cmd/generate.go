package cmd

import (
	"fmt"

	"github.com/informeai/grun/lib"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate formula in repo",
	Long:  `Generate formula in repository for save commands`,
	Run: func(cmd *cobra.Command, args []string) {
		flagName := cmd.Flag("name")
		if len(flagName.Value.String()) == 0 {
			fmt.Printf("invalid name")
			return
		}
		if exist := lib.VerifyRepoExist(); !exist {
			fmt.Printf("repo not exist\n")
			return
		}
		err := lib.VerifyFormula(flagName.Value.String())
		if err != nil && err == lib.ErrFormulaAlreadyExist {
			fmt.Printf(err.Error())
			return
		}
		if err != nil && err == lib.ErrFormulaNotFound {
			if errCrt := lib.CreateFormula(flagName.Value.String()); errCrt != nil {
				fmt.Printf("error in creation -> %s\n", errCrt.Error())
				return
			}
			fmt.Printf("formula -> %s created\n", flagName.Value.String())
		}
	},
}
var name string

func init() {
	rootCmd.AddCommand(generateCmd)
	
	generateCmd.Flags().StringVarP(&name, "name", "n", "", "name the formula")
	generateCmd.MarkFlagRequired("name")
}
