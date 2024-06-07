package main

import (
	"os"

	"github.com/ogticrd/kubectl-envsecret/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func main() {
	var flags *pflag.FlagSet = pflag.NewFlagSet("kubectl-envsecret", pflag.ExitOnError)
	pflag.CommandLine = flags

	// rootCmd is the base command for the kubectl-envsecret plugin.
	var rootCmd *cobra.Command = cmd.NewCmdEnvSecret(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
