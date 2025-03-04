package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "tff",
	Short: "tff is tool to modify Linux evdev keyboard events. You can build custom shortcuts (combos) to get 'ten flying fingers'.",
	Long:  `tff (ten flying fingers). Most commands need root permissions (sudo). https://github.com/guettli/tff`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
