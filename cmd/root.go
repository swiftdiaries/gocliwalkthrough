package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "my-cli-app",
	Short: "A simple CLI app built using Cobra",
	Long:  `Cobra's Complete documentation is available at http://cobra.dev`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hello World !")
	},
}

// Execute forms the entrypoint for the CLI App
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
