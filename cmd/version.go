package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var (
	version    = "0.1.0"
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Long:  `Print the version number, which is hard-coded into the source code.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("version=%s\n", version)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
