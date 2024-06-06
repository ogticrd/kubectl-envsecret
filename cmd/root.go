/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericiooptions"
)

// NativeOptions encapsulates the configuration flags and IO streams for the command.
type NativeOptions struct {
	configFlags *genericclioptions.ConfigFlags // Configuration flags for the CLI.

	genericclioptions.IOStreams // Input/output streams for the CLI.
}

// NewNativeOptions creates a new NativeOptions instance with the provided IO streams.
//
// Example usage:
// streams := genericiooptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
// options := NewNativeOptions(streams)
func NewNativeOptions(streams genericiooptions.IOStreams) *NativeOptions {
	return &NativeOptions{
		configFlags: genericclioptions.NewConfigFlags(true),
		IOStreams:   streams,
	}
}

// NewCmdEnvSecret creates a new cobra command for the kubectl-envsecret plugin.
//
// Example usage:
// streams := genericiooptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
// cmd := NewCmdEnvSecret(streams)
// cmd.Execute()
func NewCmdEnvSecret(streams genericiooptions.IOStreams) *cobra.Command {
	o := NewNativeOptions(streams)

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

	o.configFlags.AddFlags(cmd.Flags())

	return cmd
}

// rootCmd is the base command for the kubectl-envsecret plugin.
var rootCmd *cobra.Command = NewCmdEnvSecret(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kubectl-envsecret.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	var flags *pflag.FlagSet = pflag.NewFlagSet("kubectl-envsecret", pflag.ExitOnError)
	pflag.CommandLine = flags
}
