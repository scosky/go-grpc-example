package main

import (
	context "context"
	"fmt"
	"log"

	"github.com/scosky/grpc-client/services"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	userClient := services.NewUserServiceClient(conn)
	userResponse, err := userClient.GetUserInfo(context.Background(), &services.UserRequest{UserId: 23})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("userResponse:", userResponse)
}
