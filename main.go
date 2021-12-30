package main

import (
	"log"
	"net"

	pb "github.com/zxcv9203/grpc_study/ecommerce"
	"github.com/zxcv9203/grpc_study/service"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &service.Server{})
	log.Printf("Starting gRPC listener on port : " + port)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to Serve : %v", err)
	}
}
