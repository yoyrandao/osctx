package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	Version = "dev"
	Commit  = ""
)

var versionCmd = &cobra.Command{
	Use:           "version",
	Short:         "Display version information",
	Args:          cobra.NoArgs,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Fprintf(os.Stderr, "Version=%s\nCommit=%s\n", Version, Commit)
		return nil
	},
}
