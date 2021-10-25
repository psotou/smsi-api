package data

import "time"

type Product struct {
	ID          int64
	Name        string
	Description string
	SKU         string
	CategoryID  int32
	InventoryID int32
	Price       uint32
	CreatedAt   time.Time
	ModifiedAt  time.Time
	DeletedAt   time.Time
}
