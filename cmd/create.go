/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var envFilePath string = "./"

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Kubernetes secret from a .env file with multiline support.",
	Long: `The create command allows you to generate a Kubernetes secret from a .env file, including support for multiline environment variables. 

This command reads the specified .env file, processes its contents, and creates a Kubernetes secret that can be applied to your cluster. This is particularly useful for managing sensitive configuration data with complex, multiline values in a streamlined and efficient manner.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called with", envFilePath)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().StringVar(&envFilePath, "from-env-file", envFilePath, "")
}
