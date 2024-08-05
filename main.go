package main

import (
	"engine/proto"
	"engine/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8090"
)

func main() {

	grpcEngineConnection(port)

}
func grpcEngineConnection(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
	grpcServer := grpc.NewServer()

	proto.RegisterEngineRequestServer(grpcServer, &service.FetchUser{})
	log.Printf("Server started at : %v", lis.Addr())

	err1 := grpcServer.Serve(lis)
	if err1 != nil {
		log.Fatalf("Failed to start: %v", err1)
	}
}
