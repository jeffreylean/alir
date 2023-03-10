/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/jeffreylean/alir/config"
	"github.com/spf13/cobra"
)

func serverCmd(c *config.Config) *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "server <command>",
		Aliases: []string{"s"},
		Short:   "Server",
		Long:    "Server execution",
	}

	// Add command
	cmd.AddCommand(startCommand(c))

	return cmd

}

func startCommand(c *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "start",
		Aliases: []string{"s"},
		Short:   "Start the server",
		RunE: func(cmd *cobra.Command, args []string) error {
			c.Load()
			return nil
		},
	}
	return cmd
}
