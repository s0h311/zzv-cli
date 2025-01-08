package main

import (
	"fmt"
	"os"
	"zzv/cli/cmd"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zzv",
	Short: "zzv-cli enables you to develop efficiently",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func main() {
	rootCmd.AddCommand(cmd.GetPullCmd())
	rootCmd.AddCommand(cmd.GetBuildCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops, error: '%s'\n", err)
		os.Exit(1)
	}
}
