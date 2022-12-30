package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type CustomerValidator interface {
	// ValidateGetArgs parses the Customer ID from a set of args
	ValidateGetArgs(args []string) (int, error)

	ValidateCreateArgs(args []string) (*CreateCustomerJSON, error)
}

func NewCustomerValidatorImpl() *CustomerValidatorImpl {
	return &CustomerValidatorImpl{}
}

type CustomerValidatorImpl struct{}

func (c *CustomerValidatorImpl) ValidateGetArgs(args []string) (int, error) {
	if len(args) != 1 {
		return 0, errors.New("get customer argument list incorrect length")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, fmt.Errorf("parse get user arg as int: %w", err)
	}

	return id, nil
}

func (c *CustomerValidatorImpl) ValidateCreateArgs(args []string) (*CreateCustomerJSON, error) {
	if len(args) != 1 {
		return nil, errors.New("get customer argument list incorrect length")
	}
	filename := args[0]

	var err error
	var result CreateCustomerJSON

	file, err := loadFile(err, filename)
	if err == nil {
		defer file.Close()
	}
	err = parseJSON(err, file, &result)
	err = notNil(err, result.FirstName, "first_name is required")
	err = notNil(err, result.LastName, "last_name is required")
	err = notNil(err, result.Age, "age is required")
	err = stringMax(err, result.FirstName, 10, "first_name has invalid length")
	err = stringMax(err, result.MiddleName, 10, "middle_name has invalid length")
	err = stringMax(err, result.LastName, 10, "last_name has invalid length")

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func stringMax(err error, value *string, max int, msg string) error {
	if err != nil || value == nil {
		return err
	}

	if len(*value) > max {
		return errors.New(msg)
	}

	return nil
}

func notNil[T any](err error, value *T, msg string) error {
	if err != nil {
		return err
	}

	if value == nil {
		return errors.New(msg)
	}

	return nil
}

func loadFile(err error, path string) (io.ReadCloser, error) {
	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if err == os.ErrNotExist {
		return nil, fmt.Errorf("input file does not exist: %s", path)
	}
	if err != nil {
		return nil, fmt.Errorf("error opening input file at: %s", path)
	}

	return file, nil
}

func parseJSON(err error, jsonReader io.Reader, result interface{}) error {
	if err != nil {
		return err
	}

	err = json.NewDecoder(jsonReader).Decode(result)
	if err != nil {
		return fmt.Errorf("failed to parse input JSON: %w", err)
	}

	return nil
}
