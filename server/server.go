package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/cemgurhan/exp-grpc-gateway/gen/go"
	"github.com/cemgurhan/exp-grpc-gateway/server/categories"

	"google.golang.org/grpc"
)

const port int = 9090

func RunServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCategoryServer(grpcServer, &categories.CategoryServer{})
	fmt.Printf("Running gRPC server at port %d\n", port)
	grpcServer.Serve(lis)
}

func main() {
	RunServer()
}
