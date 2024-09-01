// ./internal/itemservice.go
package internal

import (
	"context"
	"log"
	"simmons/todo_service/protogen/golang/item"
)

// ItemService should implement the ItemsServer interface generated from grpc.
//
// UnimplementedItemsServer must be embedded to have forwarded compatible implementations.
type ItemService struct {
	db *DB
	items.UnimplementedItemsServer
}

// NewItemService creates a new Item
func NewItemService(db *DB) ItemService {
	return ItemService{db: db}
}

// AddItem implements the AddItem method of the grpc ItemService interface to add a new item
func (o *ItemService) AddItem(_ context.Context, req *items.PayloadWithSingleItem) (*items.Empty, error) {
	log.Printf("Received an add-item request")
	err := o.db.AddItem(req.GetItem())
	return &items.Empty{}, err
}

// GetItems implements the GetItems method of the grpc ItemService interface to add a new item
func (o *ItemService) GetItems(_ context.Context, _ *items.Empty) (*items.PayloadWithItems, error) {
	log.Printf("Received an get-items request")
	elems, err := o.db.GetItems()

	response := &items.PayloadWithItems{}

	response.Items = elems

	return response, err
}
