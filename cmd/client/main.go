package main

import (
	"context"
	"google.golang.org/grpc"
	v1 "gorpc/api/proto/v1"
	"log"
	"time"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

func main()  {
	// get configuration
	//address := flag.String("server", "", "gRPC server in format host:port")
	//flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewTestServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := v1.GetRequest{
		Id: 12,
	}
	res, err := c.Get(ctx, &req)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res)

}