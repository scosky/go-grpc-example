package main

import (
	"net"

	"github.com/scosky/grpc-service/services"
	"google.golang.org/grpc"
)

func main() {
	rpcService := grpc.NewServer()
	services.RegisterUserServiceServer(rpcService, new(services.UserService))

	lis, _ := net.Listen("tcp", ":8081")
	rpcService.Serve(lis)
}
