// Package cmd provides the command-line interface for the kubectl-envsecret plugin.
//
// This package defines the commands and options available for the kubectl-envsecret
// plugin, which simplifies the creation of Kubernetes secrets from .env files,
// including support for multiline environment variables.
package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericiooptions"
)

// NativeOptions encapsulates the configuration flags and IO streams for the command.
type RootCmdOptions struct {
	configFlags *genericclioptions.ConfigFlags // Configuration flags from kubectl CLI.

	genericclioptions.IOStreams // Input/output streams for the CLI.
}

// NewNativeOptions creates a new NativeOptions instance with the provided IO streams.
//
// Example usage:
// streams := genericiooptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
// options := NewNativeOptions(streams)
func NewRootCmdOptions(streams genericiooptions.IOStreams) *RootCmdOptions {
	return &RootCmdOptions{
		configFlags: genericclioptions.NewConfigFlags(true),

		IOStreams: streams,
	}
}

// NewCmdEnvSecret creates a new cobra command for the kubectl-envsecret plugin.
//
// Example usage:
// streams := genericiooptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
// cmd := NewCmdEnvSecret(streams)
// cmd.Execute()
func NewCmdEnvSecret(streams genericiooptions.IOStreams) *cobra.Command {
	o := NewRootCmdOptions(streams)

	// cmd represents the base command when called without any subcommands
	cmd := &cobra.Command{
		Use: "kubectl-envsecret",
		Annotations: map[string]string{
			cobra.CommandDisplayNameAnnotation: "kubectl envsecret",
		},
		Short: "Create Kubernetes secrets from .env files with multiline support.",
		Long: `kubectl-envsecret is a plugin for kubectl that simplifies the process of creating Kubernetes secrets from .env files, including support for multiline environment variables. 

  This tool reads the .env file, converts its contents into Kubernetes secret format, and applies it to your cluster. It streamlines the management of secrets, making it easier to handle configurations that include multiline values.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}

	o.configFlags.AddFlags(cmd.PersistentFlags())

	// create subcommands
	cmd.AddCommand(NewCmdCreate(streams))

	return cmd
}
