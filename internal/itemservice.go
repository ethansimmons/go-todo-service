// ./internal/itemservice.go
package internal

import (
	"context"
	"log"
	"simmons/todo_service/proto/item"
)

// ItemService should implement the ItemsServer interface generated from grpc.
//
// UnimplementedItemsServer must be embedded to have forwarded compatible implementations.
type ItemService struct {
	db *DB
	item.UnimplementedItemsServer
}

// NewItemService creates a new Item
func NewItemService(db *DB) ItemService {
	return ItemService{db: db}
}

// AddItem implements the AddItem method of the grpc ItemService interface to add a new item
func (o *ItemService) AddItem(_ context.Context, req *item.PayloadWithSingleItem) (*item.Empty, error) {
	log.Printf("Received an add-item request")
	err := o.db.AddItem(req.GetItem())
	return &item.Empty{}, err
}

// GetItems implements the GetItems method of the grpc ItemService interface to add a new item
func (o *ItemService) GetItems(_ context.Context, _ *item.Empty) (*item.PayloadWithItems, error) {
	log.Printf("Received an get-items request")
	elems, err := o.db.GetItems()

	response := &item.PayloadWithItems{}

	response.Items = elems

	return response, err
}
