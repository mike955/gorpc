package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	v1 "gorpc/api/proto/v1"

	"google.golang.org/grpc"
)

const (
	PORT = "8005"
)

func RunServer(ctx context.Context, v1API v1.TestServiceServer) error {
	listen, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(err)
	}
	server := grpc.NewServer()
	v1.RegisterTestServiceServer(server, v1API)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down grpc server ...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	// start gPRC server
	log.Println("starting gRPC server ...")
	return server.Serve(listen)
}
