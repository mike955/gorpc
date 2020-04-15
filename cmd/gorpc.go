package cmd

import (
	"context"
	"gorpc/internal/server/grpc"
	v1 "gorpc/internal/service/v1"
)

// func NewGorpcCommand() {
// 	//cmd := &cobra.Command{
// 	//
// 	//}
// 	//return cmd
// }

func Run() error {
	ctx := context.Background()

	v1API := v1.NewTestServer()
	return grpc.RunServer(ctx, v1API)
}
