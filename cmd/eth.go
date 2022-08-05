package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ethCmd = &cobra.Command{
	Use:   "eth",
	Short: "Eth log parser",
	Long:  `Eth log parser`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("eth log parser")
	},
}
