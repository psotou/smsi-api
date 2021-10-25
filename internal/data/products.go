package data

import "time"

type ProductEnvelope struct {
	Product Product `json:"product"`
}
type Product struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SKU         string `json:"sku"`
	CategoryID  int32  `json:"category_id"`
	InventoryID int32  `json:"inventory_id"`
	Price       uint32 `json:"price"`
	// the hyphen (-) directive is used because we never want these fields
	// to appear in the JSON output
	CreatedAt  time.Time `json:"-"`
	ModifiedAt time.Time `json:"-"`
	DeletedAt  time.Time `json:"-"`
}
