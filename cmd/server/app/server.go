package app

import (
	"context"
	apiV1Pb "github.com/Yessentemir256/grpc-project/pkg/api/v1"
	"log"
	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	apiV1Pb.UnimplementedMessageServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) FindMessageByID(
	ctx context.Context,
	request *apiV1Pb.MessageRequest,
) (*apiV1Pb.MessageResponse, error) {
	log.Print(request)

	return &apiV1Pb.MessageResponse{
		Id:      1,
		Content: "Hello world",
		Created: timestamppb.Now(),
	}, nil

	// или для ошибки:
	// return nil, status.Errorf(
	// codes.NotFound,
	// "message with id %d not found", request.GetId(),
	// )
}
