// ./cmd/server/main.go
package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"simmons/todo_service/internal"
	pb "simmons/todo_service/proto/item"
)

func main() {
	log.Printf("grpc-ping: starting server...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "443"
		log.Printf("Defaulting to port %s", port)
	}

	// create a TCP listener on the specified port
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a gRPC server instance
	server := grpc.NewServer()

	// create a item service instance with a reference to the db
	db := internal.NewDB()
	itemService := internal.NewItemService(db)

	// register the item service with the grpc server
	pb.RegisterItemsServer(server, &itemService)

	// register reflection service on grpc server
	reflection.Register(server)

	// start listening to requests
	log.Printf("server listening at %v", listener.Addr())
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
