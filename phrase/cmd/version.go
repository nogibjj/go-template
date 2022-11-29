package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of phrase",
	Long:  `All software has versions. This is phrases`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("phrase -- 1.0.0")
	},
}
