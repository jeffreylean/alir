package cmd

import (
	"github.com/jeffreylean/alir/config"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	var file string

	var rootCmd = &cobra.Command{
		Use:   "alir <command> [flags]",
		Short: "Ingester",
		Long:  "Low latency fan out ingester",
	}

	c := new(config.Config)

	if file != "" {
		c.ConfigFile = file
	}

	rootCmd.AddCommand(serverCmd(c))

	rootCmd.Flags().StringVarP(&file, "file", "f", "", "Custom config file location")
	return rootCmd
}
