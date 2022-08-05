package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gapipaser",
	Long:  `All software has versions. This is gapipaser's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Gapipaser v0.1 -- HEAD")
	},
}
