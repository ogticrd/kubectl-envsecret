package cmd_test

import (
	"bytes"
	"testing"

	"github.com/ogticrd/kubectl-envsecret/cmd"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"k8s.io/cli-runtime/pkg/genericiooptions"
)

func TestNewRootCmdOptions(t *testing.T) {
	inBuf := new(bytes.Buffer)
	outBuf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)
	streams := genericiooptions.IOStreams{In: inBuf, Out: outBuf, ErrOut: errBuf}

	// options := cmd.NewRootCmdOptions(streams)
	options := cmd.NewRootCmdOptions(streams)

	assert.NotNil(t, options)
	assert.Equal(t, streams, options.IOStreams)
}

func TestNewCmdEnvSecret(t *testing.T) {
	inBuf := new(bytes.Buffer)
	outBuf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)
	streams := genericiooptions.IOStreams{In: inBuf, Out: outBuf, ErrOut: errBuf}

	rootCmd := cmd.NewCmdEnvSecret(streams)

	assert.NotNil(t, rootCmd)
	assert.Equal(t, rootCmd.Use, cmd.Use)
	assert.Equal(t, rootCmd.Short, cmd.ShortDescription)
	assert.Equal(t, rootCmd.Annotations[cobra.CommandDisplayNameAnnotation], cmd.DisplayName)
	assert.Equal(t, rootCmd.Long, cmd.LongDescription)

	// Test if the persistent flags include config flags
	assert.NotNil(t, rootCmd.PersistentFlags().Lookup("kubeconfig"))
}

func TestCmdExecute(t *testing.T) {
	inBuf := new(bytes.Buffer)
	outBuf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)
	streams := genericiooptions.IOStreams{In: inBuf, Out: outBuf, ErrOut: errBuf}

	rootCmd := cmd.NewCmdEnvSecret(streams)
	rootCmd.SetArgs([]string{}) // Simulate running the command with --help
	err := rootCmd.Execute()
	if err != nil {
		t.Error("error running root command")
	}

	assert.Nil(t, err)
	assert.Contains(t, outBuf.String(), cmd.LongDescription)
}
