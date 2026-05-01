package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yoyrandao/osctx/internal/clouds"
)

var lsCmd = &cobra.Command{
	Use:           "ls",
	Aliases:       []string{"list"},
	Short:         "List all available clouds",
	Args:          cobra.NoArgs,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		names, err := clouds.List()
		if err != nil {
			return err
		}
		for _, name := range names {
			fmt.Fprintln(os.Stderr, name)
		}
		return nil
	},
}
