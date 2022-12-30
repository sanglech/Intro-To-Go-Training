package customer

import (
	"customer_app/internal/customer/cli"
	"customer_app/internal/customer/dal"
	"customer_app/internal/customer/logic"
	"database/sql"
	"fmt"
)

type customerApp struct {
	db *sql.DB

	CustomerRepositoryMySQLImpl *dal.CustomerRepositoryMySQLImpl

	CustomerServiceImpl *logic.CustomerServiceImpl

	CustomerValidatorImpl *cli.CustomerValidatorImpl

	GetHandler    *cli.GetHandler
	CreateHandler *cli.CreateHandler
	RootHandler   cli.CmdHandler
}

func newCustomerApp() (*customerApp, error) {
	var result customerApp
	var err error

	result.db, err = dal.NewMySQLClient(&dal.MySQLClientConfig{
		User:      "docker",
		Password:  "docker",
		Address:   "127.0.0.1:3306",
		DB:        "introtogo",
		TimeoutMs: 5000,
	})
	if err != nil {
		return nil, fmt.Errorf("create MySQL connection: %w", err)
	}

	result.CustomerRepositoryMySQLImpl = dal.NewCustomerRepositoryMySQLImpl(result.db)

	result.CustomerServiceImpl = logic.NewCustomerServiceImpl(result.CustomerRepositoryMySQLImpl)

	result.CustomerValidatorImpl = cli.NewCustomerValidatorImpl()

	result.CreateHandler = cli.NewCreateHandler(result.CustomerValidatorImpl, result.CustomerServiceImpl)
	result.GetHandler = cli.NewGetHandler(result.CustomerValidatorImpl, result.CustomerServiceImpl)
	result.RootHandler = cli.NewRootHandler(result.GetHandler, result.CreateHandler)

	return &result, nil
}

func closeCustomerApp(app *customerApp) {
	app.db.Close()
}
