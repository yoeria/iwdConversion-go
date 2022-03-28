package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(makeDDS)
}

var(
	Dir string
	File string
)

var makeDDS = &cobra.Command{
	Use: "makedds",
	Aliases: []string{"md"},
	Short: "Makes dds files from iwd files",
}
