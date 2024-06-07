package cmd

import (
	"fmt"

	"github.com/ogticrd/kubectl-envsecret/internal/utils"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// CreateOptions contains the options for the create command.
type CreateOptions struct {
	genericclioptions.IOStreams          // Input/output streams for the CLI.
	envFilePath                 []string // Paths to the .env files to be processed.
}

// NewCreateOptions initializes CreateOptions with the provided IO streams.
//
// Example usage:
// streams := genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
// options := NewCreateOptions(streams)
func NewCreateOptions(streams genericclioptions.IOStreams) *CreateOptions {
	return &CreateOptions{
		IOStreams:   streams,
		envFilePath: []string{"."},
	}
}

// NewCmdCreate creates a new cobra command for creating Kubernetes secrets from .env files.
//
// Example usage:
// streams := genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
// cmd := NewCmdCreate(streams)
// cmd.Execute()
func NewCmdCreate(streams genericclioptions.IOStreams) *cobra.Command {
	o := NewCreateOptions(streams)

	// createCmd represents the create command
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a Kubernetes secret from a .env file with multiline support.",
		Long: `The create command allows you to generate a Kubernetes secret from a .env file, including support for multiline environment variables. 

  This command reads the specified .env file, processes its contents, and creates a Kubernetes secret that can be applied to your cluster. This is particularly useful for managing sensitive configuration data with complex, multiline values in a streamlined and efficient manner.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ns, err := cmd.Flags().GetString("namespace")
			if err != nil {
				return err
			}

			// Avoid multiple flags with the same value
			o.envFilePath = utils.RemoveDuplicatedStringE(o.envFilePath)
			fmt.Println("create called with", o.envFilePath)

			fmt.Println("The namespace is", ns)

			return nil
		},
	}

	createCmd.Flags().StringSliceVar(&o.envFilePath, "from-env-file", o.envFilePath, "")

	return createCmd
}
