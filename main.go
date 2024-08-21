package main

import (
	"context"
	"fmt"
	"gRPC_tutorial/protobufs"
	"google.golang.org/grpc"
	"net"
)

type HelloService struct {
	protobufs.UnimplementedHelloServiceServer
}

func (hs *HelloService) HelloApi(ctx context.Context, req *protobufs.HelloApiRequest) (*protobufs.HelloApiResponse, error) {

	type_ := req.Type

	if type_ {

		return &protobufs.HelloApiResponse{
			Message: "The type is true",
		}, nil

	} else {

		return &protobufs.HelloApiResponse{
			Message: "The type is false",
		}, nil
	}

}

func main() {

	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Print("There was an error while listening to the address")
	}

	grpcServer := grpc.NewServer()
	protobufs.RegisterHelloServiceServer(grpcServer, &HelloService{})

	err = grpcServer.Serve(server)
	if err != nil {
		fmt.Println("There was an error while Serving the gRPC server")
	}

}
