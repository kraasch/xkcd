
package cmd

import (
  "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xkcd",
	Short: "xkcd cli tool",
	Long:  "A command line tool to download and search xkcd meta data.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(searchCmd)
}

