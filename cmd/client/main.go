// ./cmd/client/main.go
package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"net/http"
	"simmons/todo_service/protogen/golang/item"
)

var itemServiceAddr string

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got / request")
	_, err := io.WriteString(w, "This is the root of my service.\n")
	if err != nil {
		log.Fatalf("error writing response: %v", err)
	}
}

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
	if err = items.RegisterItemsHandler(context.Background(), mux, conn); err != nil {
		log.Fatalf("failed to register the item server: %v", err)
	}

	// start listening to requests from the gateway server
	http.HandleFunc("/", getRoot)
	
	addr := "0.0.0.0:8080"
	fmt.Println("API gateway server is running on " + addr)
	if err = http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("gateway server closed abruptly: ", err)
	}
}
