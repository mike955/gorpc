package server

import (
	"context"
	"github.com/spf13/cobra"
	"gorpc/internal/server/grpc"
	v1 "gorpc/internal/service/v1"
)

func NewCmdServer() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "server",
		Short:  `start grpc server`,
		Long:   ``,
		PreRun: preRun,
		Run:    run,
	}
	return cmd
}

func preRun(c *cobra.Command, args []string) {

}

func run(c *cobra.Command, args []string) {
	ctx := context.Background()
	v1API := v1.NewTestServer()
	if err := grpc.RunServer(ctx, v1API); err != nil {
		panic(err)
	}
}
