package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	addr = ":9001"
)

func main() {
	grpcServer := grpc.NewServer()
	NewGRPCServiceHandler(grpcServer)

	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to create a connection: %s", addr)
	}
	defer conn.Close()
	log.Println("server is listening on port", addr)
	//grpc server

	if err := grpcServer.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
