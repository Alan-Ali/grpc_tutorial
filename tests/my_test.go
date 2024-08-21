package tests

import (
	"fmt"
	"gRPC_tutorial/protobufs"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
	"time"
)

func TestMyclient(t *testing.T) {

	// CREATING OUR CLIENT WITH THE HOST THAT WE HAVE
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Failed to connect: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("An error occured: %v", err)
		}
	}(conn)

	// THIS IS ADDING THE CLIENT THAT WE CREATED TO THE PROTOBUFFS
	hs := protobufs.NewHelloServiceClient(conn)

	// THIS IS THE CONTEXT, WE MAKE A TIME FOR OUR TEST TIMEOUT
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	req := &protobufs.HelloApiRequest{
		Type: true,
	}

	api, err := hs.HelloApi(ctx, req)
	if err != nil {
		fmt.Printf("An error occured: %v", err)
	}

	if api != nil {
		t.Logf("HelloApi Response: %v", api.Message)
	} else {
		t.Error("Sorry, there was an error")
	}

}
