package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dumpCmd = &cobra.Command{
	Use:     "dump",
	Aliases: []string{"display"},
	Short:   "Dump the header of the gba rom",
	Long:    "Dump the header of the gba rom",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		err, res := dump(args[0])
		if err != nil {
			return err
		}
		fmt.Printf("%s header:\n%s\n", args[0], res)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(dumpCmd)
}
