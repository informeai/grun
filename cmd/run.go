package cmd

import (
	"fmt"
	"strings"

	"github.com/informeai/grun/lib"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "running command",
	Long: `Running the command in formula
  usage: grun run FORMULA:COMMAND_NAME`,
	Run: func(cmd *cobra.Command, args []string) {
		if exist := lib.VerifyRepoExist(); !exist {
			fmt.Printf("repo not exist\n")
			return
		}
		if len(args) != 1 {
			fmt.Printf("invalid format command\n")
			return
		}
		if isOk := strings.Contains(args[0], ":"); isOk {
			formulaAndCommand := strings.Split(args[0], ":")
			formula, errF := lib.GetFormula(formulaAndCommand[0])
			if errF != nil {
				fmt.Println(errF.Error())
				return
			}
			var exComand lib.Commands
			for _, comd := range formula.Commands {
				if comd.Name == formulaAndCommand[1] {
					exComand = comd
				}
			}
			if err := lib.Run(exComand); err != nil {
				fmt.Println(err.Error())
				return
			}
		} else {
			fmt.Println("command not found")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
