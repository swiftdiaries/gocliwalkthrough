package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(printCmd)
}

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print the input argument",
	Long:  `A simple print command from the CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			fmt.Printf(arg)
		}
	},
}
