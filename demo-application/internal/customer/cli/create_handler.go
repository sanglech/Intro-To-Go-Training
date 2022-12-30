package cli

import (
	"customer_app/internal/customer/logic"
	"fmt"
)

type CreateHandler struct {
	customerValidator CustomerValidator
	customerService   logic.CustomerService
}

func NewCreateHandler(
	customerValidator CustomerValidator,
	customerService logic.CustomerService,
) *CreateHandler {
	return &CreateHandler{
		customerValidator: customerValidator,
		customerService:   customerService,
	}
}

func (c *CreateHandler) Handle(args []string) error {
	// 1. Parse arguments and load customer data from file
	inputData, err := c.customerValidator.ValidateCreateArgs(args)
	if err != nil {
		return fmt.Errorf("validate create: %w", err)
	}

	// 1.A. map input data to logic format
	customer := &logic.Customer{
		FirstName:  *inputData.FirstName,
		MiddleName: inputData.MiddleName,
		LastName:   *inputData.LastName,
		Age:        *inputData.Age,
	}

	// 2. Pass customer data to logic to save customer
	id, err := c.customerService.Create(customer)
	if err != nil {
		return fmt.Errorf("create customer: %w", err)
	}

	// 3. Print new customer ID
	fmt.Printf("New customer created with ID: %v\n", id)

	return nil
}
