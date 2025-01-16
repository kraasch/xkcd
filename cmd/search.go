
package cmd

import (
  "xkcd/search"

  "github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
  Use:   "search",
  Short: "Search through xkcd data",
  Run: func(cmd *cobra.Command, args []string) {
    search.PerformSearch()
  },
}

