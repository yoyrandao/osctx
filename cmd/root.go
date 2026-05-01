package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yoyrandao/osctx/internal/clouds"
	"github.com/yoyrandao/osctx/internal/selector"
)

var rootCmd = &cobra.Command{
	Use:   "osctx",
	Short: "Interactive switcher for OpenStack clouds (clouds.yaml)",
	Long: `osctx selects an OpenStack cloud from clouds.yaml and emits the
shell command to set OS_CLOUD on stdout. Pair with this shell wrapper:
 
  osctx() { eval "$(command osctx "$@")"; }`,
	Args:          cobra.NoArgs,
	RunE:          runRoot,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute is the single entry-point called from main.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetOut(os.Stderr)
	rootCmd.SetErr(os.Stderr)
	rootCmd.AddCommand(lsCmd, currentCmd, unsetCmd, versionCmd)
}

func runRoot(cmd *cobra.Command, args []string) error {
	names, err := clouds.List()
	if err != nil {
		return err
	}

	cloud, err := selector.Select(names, os.Stderr)
	if err != nil {
		return err
	}

	fmt.Printf("export OS_CLOUD=%s\n", cloud)
	fmt.Fprintf(os.Stderr, "Switched to cloud: %s\n", cloud)
	return nil
}
