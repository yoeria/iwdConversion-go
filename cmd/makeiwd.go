package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(makeIWD)
}

var makeIWD = &cobra.Command{
	Use:     "makeiwd",
	Aliases: []string{"mi"},
	Short:   "Makes iwd files from dds files",
}
