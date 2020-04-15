package cmd

import (
	"gorpc/cmd/client"
	"gorpc/cmd/server"

	"github.com/spf13/cobra"
)

func NewCmdGorpc() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "gorpc",
		Version: "0.0.1",
		Short: `gorpc:  easily bootstrap grpc project`,
		Long: `gorpc:  easily bootstrap grpc project`,
	}

	cmd.AddCommand(server.NewCmdServer())
	cmd.AddCommand(client.NewCmdClient())
	return cmd
}
