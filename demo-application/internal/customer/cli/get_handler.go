package cli

import (
	"customer_app/internal/customer/logic"
	"fmt"
)

type GetHandler struct {
	customerValidator CustomerValidator
	customerService   logic.CustomerService
}

func NewGetHandler(
	customerValidator CustomerValidator,
	customerService logic.CustomerService,
) *GetHandler {
	return &GetHandler{
		customerValidator: customerValidator,
		customerService:   customerService,
	}
}

func (g *GetHandler) Handle(args []string) error {
	// 1. parse arguments to get customer ID
	id, err := g.customerValidator.ValidateGetArgs(args)
	if err != nil {
		return fmt.Errorf("validate args: %w", err)
	}

	// 2. pass ID to logic to fetch customer
	customer, err := g.customerService.Get(id)
	if err != nil {
		return fmt.Errorf("fetch user: %w", err)
	}

	// 3. write customer data to terminal output (stdout)
	fmt.Printf("Customer %v:\n", customer.ID)
	fmt.Printf("FirstName: %s\n", customer.FirstName)
	if customer.MiddleName != nil {
		fmt.Printf("MiddleName: %s\n", *customer.MiddleName)
	}
	fmt.Printf("LastName: %s\n", customer.LastName)
	fmt.Printf("Age: %v\n", customer.Age)

	return nil
}
