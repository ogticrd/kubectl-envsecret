package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericiooptions"
)

// VersionOptions contains the options for the version command.
type VersionOptions struct {
	genericiooptions.IOStreams // Input/output streams for the CLI.
	version                    string
}

// NewVersionOptions initializes VersionOptions with the provided IO streams.
//
// Example usage:
// streams := genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
// options := NewVersionOptions(streams)
func NewVersionOptions(streams genericiooptions.IOStreams) *VersionOptions {
	return &VersionOptions{
		version:   AppVersion,
		IOStreams: streams,
	}
}

// NewCmdCreate creates a new cobra command for printing kubectl-envsecret version.
//
// Example usage:
// streams := genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
// cmd := NewCmdVersion(streams)
// cmd.Execute()
func NewCmdVersion(streams genericiooptions.IOStreams) *cobra.Command {
	// versionCmd represents the version command
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print kubectl-envsecret version.",
		Long:  "Print kubectl-envsecret version.",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(AppVersion)
			return nil
		},
	}

	return versionCmd
}
