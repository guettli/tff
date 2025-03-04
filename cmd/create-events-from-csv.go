package cmd

import (
	"github.com/guettli/tff/pkg/tff"
	"github.com/spf13/cobra"
)

func init() {
	combosCmd := &cobra.Command{
		Use:   "create-events-from-csv events.csv",
		Short: "Read events from csv file, and emit the events. This does not rewrite the events like 'replay-combo-log' does.",
		RunE: func(_ *cobra.Command, args []string) error {
			path := ""
			if len(args) > 0 {
				path = args[0]
			}
			return tff.CreateEventsFromCsv(path)
		},
		Args:                  cobra.ExactArgs(1),
		DisableFlagsInUseLine: true,
	}
	RootCmd.AddCommand(combosCmd)
}
