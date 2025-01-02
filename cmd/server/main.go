package main

import (
	apiV1Pb "github.com/Yessentemir256/grpc-project/pkg/api/v1"
	"github.com/Yessentemir256/grpc-project/server/app"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

func main() {
	err := execute(
		net.JoinHostPort("go.yessentemir.dev", "9999"),
		"server.crt",
		"server-private.key",
	)
	if err != nil {
		log.Print(err)
	}
}

func execute(addr string, certFile string, privateKeyFile string) error {
	creds, err := credentials.NewServerTLSFromFile(certFile, privateKeyFile)
	if err != nil {
		return err
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Print(err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	server := app.NewServer()
	apiV1Pb.RegisterMessageServiceServer(grpcServer, server)
	return grpcServer.Serve(listener)
}
