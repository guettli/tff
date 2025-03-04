package cmd

import (
	"github.com/guettli/tff/pkg/tff"
	"github.com/spf13/cobra"
)

func init() {
	validateCmd := &cobra.Command{
		Use:   "validate combos.yaml",
		Short: "Validate the config file.",
		RunE: func(_ *cobra.Command, args []string) error {
			return tff.ValidateMain(args[0])
		},
		Args:                  cobra.ExactArgs(1),
		DisableFlagsInUseLine: true,
		SilenceUsage:          true,
	}
	RootCmd.AddCommand(validateCmd)
}
