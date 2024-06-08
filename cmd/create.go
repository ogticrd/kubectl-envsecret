package cmd

import (
	"github.com/ogticrd/kubectl-envsecret/internal/k8s"
	"github.com/ogticrd/kubectl-envsecret/internal/parser"
	"github.com/ogticrd/kubectl-envsecret/internal/utils"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/rest"
)

// CreateOptions contains the options for the create command.
type CreateOptions struct {
	genericclioptions.IOStreams
	configFlags  *genericclioptions.ConfigFlags
	restConfig   *rest.Config
	namespace    string
	secretName   string
	envFilePaths []string
}

// NewCreateOptions initializes CreateOptions with the provided IO streams.
//
// Example usage:
// streams := genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
// options := NewCreateOptions(streams)
func NewCreateOptions(streams genericclioptions.IOStreams) *CreateOptions {
	return &CreateOptions{
		configFlags:  genericclioptions.NewConfigFlags(true),
		IOStreams:    streams,
		envFilePaths: []string{".env"},
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
		Use:   "create [secret name] [flags]",
		Short: "Create a Kubernetes secret from a .env file with multiline support.",
		Long: `The create command allows you to generate a Kubernetes secret from a .env file, including support for multiline environment variables. 

  This command reads the specified .env file, processes its contents, and creates a Kubernetes secret that can be applied to your cluster. This is particularly useful for managing sensitive configuration data with complex, multiline values in a streamlined and efficient manner.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := o.Complete(cmd, args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			if err := o.Run(); err != nil {
				return err
			}
			return nil
		},
	}

	createCmd.Flags().StringSliceVar(&o.envFilePaths, "from-env-file", o.envFilePaths, "Specify the path to a file to read key=val pairs to create a secret.")
	createCmd.MarkFlagFilename("from-env-file")

	return createCmd
}

func (o *CreateOptions) Complete(cmd *cobra.Command, args []string) error {
	o.secretName = args[0]

	var err error

	envFilePaths, err := cmd.Flags().GetStringSlice("from-env-file")
	if err != nil {
		return err
	}
	o.envFilePaths = utils.RemoveDuplicatedStringE(envFilePaths)

	o.restConfig, err = o.configFlags.ToRESTConfig()
	if err != nil {
		return err
	}

	ns, err := cmd.Flags().GetString("namespace")
	if err != nil {
		return err
	}

	if len(ns) == 0 {
		o.namespace = ""
	} else {
		o.namespace = ns
	}

	return nil
}

func (o *CreateOptions) Validate() error {
	// Validate that paths exists
	return utils.ValidatePaths(o.envFilePaths)
}

func (o *CreateOptions) Run() error {
	var err error

	client, err := k8s.NewK8sClientFromConfig(k8s.NewK8sConfig(o.restConfig, o.namespace))
	if err != nil {
		return err
	}

	parsedFile, err := parser.Load(o.envFilePaths...)
	if err != nil {
		return err
	}

	if err := client.CreateSecret(o.secretName, parsedFile); err != nil {
		return err
	}

	return nil
}
