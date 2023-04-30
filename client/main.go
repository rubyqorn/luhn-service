package main

import (
	"context"
	"flag"
	"log"

	pb "luhn-service/client/proto"
	utils "luhn-service/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Card struct {
	Number int64
}

func main() {
	creditCardNumber := flag.Int64("creditCardNumber", 0, "int64")
	flag.Parse()

	resource := utils.GetEnvVariable("HOST") + utils.GetEnvVariable("PORT")

	connection, err := grpc.Dial(resource, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer connection.Close()

	client := pb.NewLuhnClient(connection)
	card := &pb.Card{Number: *creditCardNumber}
	grpcServerResponse, err := client.Validate(context.Background(), &pb.LuhnServiceRequest{Card: card})

	if err != nil {
		log.Fatalf("failed to validate: %v", err)
	}

	log.Printf("validation result: %v", grpcServerResponse.Valid)
}
