// ./internal/db.go
package internal

import (
	"fmt"
	"simmons/todo_service/proto/item"
)

type DB struct {
	collection []*item.Item
}

// NewDB creates a new array to mimic the behaviour of a in-memory database
func NewDB() *DB {
	return &DB{
		collection: make([]*item.Item, 0),
	}
}

// AddItem adds a new item to the DB collection. Returns an error on duplicate ids
func (d *DB) AddItem(item *item.Item) error {
	for _, i := range d.collection {
		if i.ItemId == item.ItemId {
			return fmt.Errorf("duplicate item id: %d", item.GetItemId())
		}
	}
	d.collection = append(d.collection, item)
	return nil
}

func (d *DB) GetItems() ([]*item.Item, error) {
	return d.collection, nil
}
