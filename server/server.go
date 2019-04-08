package server

import (
	"fmt"
	"github.com/evenfound/even-p2pchat-go/listener"
	"github.com/evenfound/even-p2pchat-go/server/api"
	"github.com/evenfound/even-p2pchat-go/server/handlers"
	"google.golang.org/grpc"
	"net"
)

//go:generate protoc --proto_path=proto --go_out=plugins=grpc:./api message.proto
//go:generate protoc --proto_path=proto --go_out=plugins=grpc:./api subscribe.proto

const (
	DefaultServerPort = ":8091"
)

type Server struct {
	grpcServer *grpc.Server
}

func (server *Server) Start(errorHandler chan<- error, eventListener listener.OnListener) {
	var lis, err = net.Listen("tcp", DefaultServerPort)

	if err != nil {
		errorHandler <- err
		return
	}

	/* @TODO remove this line */
	fmt.Println("Server started : PORT %v", DefaultServerPort)

	var (
		// New gRPC server
		gRPCServer = grpc.NewServer()

		// registering server handlers
		messageHandler   = handlers.NewMessageHandler(errorHandler)
		subscribeHandler = handlers.NewSubscribeHandler(errorHandler, eventListener)
	)

	api.RegisterChatServiceServer(gRPCServer, messageHandler)
	api.RegisterSubscribeServiceServer(gRPCServer, subscribeHandler)

	server.grpcServer = gRPCServer

	if err := gRPCServer.Serve(lis); err != nil {
		errorHandler <- err
	}
}

func (server *Server) Stop() {
	server.grpcServer.GracefulStop()
}
