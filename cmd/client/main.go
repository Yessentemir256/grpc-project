package main

import (
	"context"
	apiV1Pb "github.com/Yessentemir256/grpc-project/pkg/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"time"
)

func main() {
	err := execute(
		net.JoinHostPort("go.yessentemir.dev", "9999"),
		"server.crt",
	)
	if err != nil {
		log.Print(err)
	}
}

func execute(addr string, certFile string) error {
	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		log.Print(err)
	}

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Print(err)
	}

	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()

	client := apiV1Pb.NewMessageServiceClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	response, err := client.FindMessageByID(ctx, &apiV1Pb.MessageRequest{
		Id: 1,
	})
	if err != nil {
		log.Print(err)
	}
	log.Print(response.String())
	return nil
}
