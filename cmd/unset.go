package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var unsetCmd = &cobra.Command{
	Use:           "unset",
	Short:         "Clear the current cloud",
	Args:          cobra.NoArgs,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if os.Getenv("OS_CLOUD") == "" {
			fmt.Fprintln(os.Stderr, "no cloud set")
			return nil
		}
		fmt.Println("unset OS_CLOUD")
		fmt.Fprintln(os.Stderr, "OS_CLOUD cleared")
		return nil
	},
}
