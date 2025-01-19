package main

import (
	"dev-server/cmd/server"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "dev",
		Short: "A command for dev tools",
	}

	rootCmd.AddCommand(server.ServerCmd)

	rootCmd.Execute()
}
