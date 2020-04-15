package grpc

import (
	"context"
	"google.golang.org/grpc/reflection"
	v1 "gorpc/api/proto/v1"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

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
	reflection.Register(server)
	v1.RegisterTestServiceServer(server, v1API)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-c
		log.Println("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("shutting down grpc server ...")
			server.GracefulStop()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}()

	// start gPRC server
	log.Println("starting gRPC server ...")
	return server.Serve(listen)
}
