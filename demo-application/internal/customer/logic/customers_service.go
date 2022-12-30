package logic

import (
	"errors"
	"fmt"
)

type CustomerService interface {
	Get(id int) (*Customer, error)

	// Create returns the unique int ID of the new customer
	Create(data *Customer) (int, error)
}

func NewCustomerServiceImpl(
	customerRepository CustomerRepository,
) *CustomerServiceImpl {
	return &CustomerServiceImpl{
		customerRepository: customerRepository,
	}
}

type CustomerServiceImpl struct {
	customerRepository CustomerRepository
}

func (c *CustomerServiceImpl) Get(id int) (*Customer, error) {
	customer, err := c.customerRepository.GetCustomer(id)
	if errors.Is(err, ErrNotFound) {
		return nil, fmt.Errorf("customer with ID %v not found", id)
	}
	if err != nil {
		return nil, fmt.Errorf("get customer from storage: %w", err)
	}

	return customer, nil
}

func (c *CustomerServiceImpl) Create(data *Customer) (int, error) {
	return c.customerRepository.CreateCustomer(data)
}
