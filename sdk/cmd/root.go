package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sdk",
	Short: "sdk is a cli tool to help the realisation of the gba user group",
	Long:  "sdk is a cli tool to help the realisation of the gba user group - dump and generation of header gba header.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing sdk '%s'\n", err)
		os.Exit(1)
	}
}
