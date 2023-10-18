package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// New returns a new cobra.Command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "ysr",
		Short:         "ysh replaces template file according to config file.",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.AddCommand(newTemplateCommand())
	cmd.AddCommand(newVersionCommand())

	return cmd
}

func checkArgsLength(argsReceived int, requiredArgs ...string) error {
	expectedNum := len(requiredArgs)
	if argsReceived != expectedNum {
		arg := "arguments"
		if expectedNum == 1 {
			arg = "argument"
		}
		return fmt.Errorf("this command needs %v %s: %s", expectedNum, arg, strings.Join(requiredArgs, ", "))
	}
	return nil
}
