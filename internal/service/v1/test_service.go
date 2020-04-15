package v1

import (
	"context"
	v1 "gorpc/api/proto/v1"
	"log"
)

type testServer struct {
	bar string
}

// NewTestServer
func NewTestServer() v1.TestServiceServer {
	return &testServer{
		bar: "faa",
	}
}

func (s *testServer) Get(ctx context.Context, req *v1.GetRequest) (*v1.GetResponse, error) {
	log.Printf("received: %v", req.GetId())
	return &v1.GetResponse{Id: 12345, Name: "mike"}, nil
}
