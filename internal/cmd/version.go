package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string

func newVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show version of this tool",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("%s", version)
		},
	}

	return cmd
}
