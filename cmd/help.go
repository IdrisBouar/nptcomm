package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:              "help",
	Long:             `Field of the Help menu []`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Help menu: On Work !")
	},
}
