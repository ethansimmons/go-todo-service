// ./cmd/client/main.go
package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	pb "simmons/todo_service/proto/item"
)

var itemServiceAddr string

func main() {
	// Set up a connection to the item server.
	fmt.Println("Connecting to item service via", itemServiceAddr)
	conn, err := grpc.NewClient(itemServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to item service: %v", err)
	}

	defer conn.Close()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	if err = pb.RegisterItemsHandler(context.Background(), mux, conn); err != nil {
		log.Fatalf("failed to register the item server: %v", err)
	}

	addr := "0.0.0.0:8080"
	fmt.Println("API gateway server is running on " + addr)
	if err = http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("gateway server closed abruptly: ", err)
	}
}
