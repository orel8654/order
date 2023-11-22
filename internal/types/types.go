package types

import "time"

type Product struct {
	UUID        string
	Name        string
	Description string
	Type        int
	Weight      int32
	Price       float64
	Created_at  time.Time
}

type OrderItem struct {
	Count       int32
	ProductUUID string
}

type Menu struct {
	Salads    []Product
	Garnishes []Product
	Meats     []Product
	Soups     []Product
	Drinks    []Product
	Desserts  []Product
}

type CreateOrderUUID struct {
	UUID  string
	Order Menu
}
