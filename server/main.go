package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	card "luhn-service/server/card"
	pb "luhn-service/server/proto"
)

type grpcServer struct {
	pb.UnimplementedLuhnServer
}

func (server *grpcServer) Validate(ctx context.Context, request *pb.LuhnServiceRequest) (*pb.LuhnServiceResponse, error) {
	isValid, err := card.ValidateCardNumber(request.Card.Number)

	if err != nil {
		log.Fatalf("failed to validate %v", err)
	}

	fmt.Println(isValid, request.Card.Number)

	return &pb.LuhnServiceResponse{Valid: isValid}, nil
}

func getEnvVariable(variableName string) string {
	path, err := filepath.Abs("../.env")

	if err != nil {
		log.Fatalf("Env file not found")
	}

	loadErr := godotenv.Load(path)

	if loadErr != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(variableName)
}

func main() {
	listenedServer, err := net.Listen(getEnvVariable("PROTOCOL"), getEnvVariable("PORT"))

	if err != nil {
		log.Fatalf("failed to listen: %v, %v, %v", err, getEnvVariable("PROTOCOL"), getEnvVariable("PORT"))
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
