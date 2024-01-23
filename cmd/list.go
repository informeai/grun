package cmd

import (
	"fmt"

	"github.com/informeai/grun/lib"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all formulas in repo",
	Long: `List all formulas saved in repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		if exist := lib.VerifyRepoExist(); !exist {
			fmt.Printf("repo not exist\n")
			return
		}
    formulas, err := lib.ListFormula()
    if err != nil{
      fmt.Printf("failed in list formulas\n")
      return
    }
    fmt.Println("FORMULAS:")
    for _, fml := range formulas{
      fmt.Println(fml)
    }
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
