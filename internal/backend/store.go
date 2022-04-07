package backend

import "petstore"

type StoreService interface {
	PlaceOrder(*petstore.Order) error
	GetOrder(id string) (*petstore.Order, error)
	DeleteOrder(orderID string) error
}
