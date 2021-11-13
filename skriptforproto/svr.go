package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
	pb "skriptforproto/gen/proto/go/proto/v1"
)

func main() {
	if err := Main(); err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
	}
}

type server struct {
	pb.UnimplementedTestServiceServer
}

func (s *server) Test(ctx context.Context, r *pb.TestRequest) (*pb.TestResponse, error) {
	fmt.Println(r)
	return &pb.TestResponse{B: "Hello " + r.A}, nil
}

var _ pb.TestServiceServer = (*server)(nil)

func Main() error {
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterTestServiceServer(s, &server{})
	return s.Serve(ln)
}
