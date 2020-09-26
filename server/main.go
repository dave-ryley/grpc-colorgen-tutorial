package main

import (
	"context"
	"fmt"
	pb "github.com/dave-ryley/grpc-colorgen-tutorial/server"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
)

const (
	port		= ":50051"
	colorBytes	= 3
)

type server struct {}

func (s *server) GetRandomColor(ctx context.Context, curr *pb.CurrentColor) (*pb.NewColor, error) {
	hex := "#" + randomHex()
	log.Printf("Client's current color: [#%v] sending [%v]", curr.Color, hex)
	return &pb.NewColor{Color: hex}, nil
}

func randomHex() string {
	bytes := make([]byte, colorBytes)
	if _, err := rand.Read(bytes); err != nil {
		log.Panicln("Error generating random hex value", err)
	}
	return fmt.Sprintf("%x", bytes)
}

func main() {
	if lis, err := net.Listen("tcp", port); err != nil {
		log.Fatalf("Failed to listen on port [%s]: %v", port, err)
	}
	s := grpc.NewServer()
	pb.RegisterColorGeneratorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
