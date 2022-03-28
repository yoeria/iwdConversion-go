package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "iwdgo",
	Short: "iwdgo extracts dds files from iwd files",
	Long: `iwdgo extracts dds files from iwd files
	It can also pack all dds files in a directory into iwd files`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&Dir, "dir", "d", "", "Path (relative/absolute)")
	rootCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "Path (relative/absolute)")
}
