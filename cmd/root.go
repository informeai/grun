package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "grun",
	Short: "Execute e generate commands from formulas",
	Long: `
  Generate formulas and save them in the local repository for later use. 
  Execute commands in the terminal using pre-saved formulas.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		return
	}
}

func init() {
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
}
