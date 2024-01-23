package cmd

import (
	"fmt"

	"github.com/informeai/grun/lib"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize repo",
	Long:  `Initialize repository local for save formulas.`,
	Run: func(cmd *cobra.Command, args []string) {
		exist := lib.VerifyRepoExist()
		if !exist {
			err := lib.CreateRepoDir()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("repo created")
		} else {
			fmt.Println("repo existed")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
