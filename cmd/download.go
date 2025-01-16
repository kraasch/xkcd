
package cmd

import (
  "xkcd/download"

  "github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
  Use:   "download",
  Short: "Download and save xkcd data",
  Run: func(cmd *cobra.Command, args []string) {
    download.PerformDownload()
  },
}

