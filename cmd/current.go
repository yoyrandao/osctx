package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yoyrandao/osctx/internal/utils"
)

var currentCmd = &cobra.Command{
	Use:           "current",
	Aliases:       []string{"c"},
	Short:         "Print the current cloud",
	Args:          cobra.NoArgs,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Fprintln(os.Stderr, utils.GetOSCloud())
		return nil
	},
}
