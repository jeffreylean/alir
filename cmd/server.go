/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"

	"github.com/jeffreylean/alir/config"
	"github.com/jeffreylean/alir/internal/server"
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
			if err := c.Load(); err != nil {
				return err
			}

			s := server.Start(c)
			if err := s.Listen(context.Background()); err != nil {
				return err
			}

			return nil
		},
	}
	return cmd
}
