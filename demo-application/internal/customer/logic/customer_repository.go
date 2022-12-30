package logic

import "errors"

var ErrNotFound = errors.New("not found")

type CustomerRepository interface {
	CreateCustomer(c *Customer) (int, error)

	// GetCustomer returns ErrNotFound if the customer is not found
	GetCustomer(id int) (*Customer, error)
}
