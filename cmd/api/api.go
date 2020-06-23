package main

import (
	"context"
	"fmt"
	"github.com/soichisumi-sandbox/grpc-opencensus-sample/app"
	"github.com/soichisumi-sandbox/grpc-opencensus-sample/app/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	grpcPort = ":3000"
)

type HealthCheckService struct{}

func (s *HealthCheckService) HealthCheck(ctx context.Context, in *proto.Empty) (*proto.HealthCheckResponse, error) {
	return &proto.HealthCheckResponse{Status: "OK"}, nil
}

func main() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %+v\n", err)
	}

	grpcServer := grpc.NewServer()
	server := app.NewServer()
	proto.RegisterServerServer(grpcServer, server)
	proto.RegisterHealthCheckServiceServer(grpcServer, new(HealthCheckService))
	reflection.Register(grpcServer)
	fmt.Printf("api is running on port: %s\n", grpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
