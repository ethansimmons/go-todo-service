// ./internal/db.go
package internal

import (
	"fmt"
	"simmons/todo_service/protogen/golang/item"
)

type DB struct {
	collection []*items.Item
}

// NewDB creates a new array to mimic the behaviour of a in-memory database
func NewDB() *DB {
	return &DB{
		collection: make([]*items.Item, 0),
	}
}

// AddItem adds a new item to the DB collection. Returns an error on duplicate ids
func (d *DB) AddItem(item *items.Item) error {
	for _, i := range d.collection {
		if i.ItemId == item.ItemId {
			return fmt.Errorf("duplicate item id: %d", item.GetItemId())
		}
	}
	d.collection = append(d.collection, item)
	return nil
}

func (d *DB) GetItems() ([]*items.Item, error) {
	return d.collection, nil
}
