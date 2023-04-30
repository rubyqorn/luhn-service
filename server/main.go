package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	card "luhn-service/server/card"
	pb "luhn-service/server/proto"
	utils "luhn-service/utils"
)

type grpcServer struct {
	pb.UnimplementedLuhnServer
}

func (server *grpcServer) Validate(ctx context.Context, request *pb.LuhnServiceRequest) (*pb.LuhnServiceResponse, error) {
	isValid, err := card.ValidateCardNumber(request.Card.Number)

	if err != nil {
		log.Fatalf("failed to validate %v", err)
	}

	return &pb.LuhnServiceResponse{Valid: isValid}, nil
}

func main() {
	listenedServer, err := net.Listen(utils.GetEnvVariable("PROTOCOL"), utils.GetEnvVariable("PORT"))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	newGrpcServer := grpc.NewServer()
	luhnServiceGrpcServer := &grpcServer{}
	reflection.Register(newGrpcServer)
	pb.RegisterLuhnServer(newGrpcServer, luhnServiceGrpcServer)

	log.Printf("server linstening at: %v", listenedServer.Addr())

	if err := newGrpcServer.Serve(listenedServer); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
