/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ogticrd/kubectl-envsecret/internal/utils"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

type CreateOptions struct {
	genericclioptions.IOStreams
	envFilePath []string
}

func NewCreateOptions(streams genericclioptions.IOStreams) *CreateOptions {
	return &CreateOptions{
		IOStreams:   streams,
		envFilePath: []string{"."},
	}
}

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
