package customer

import (
	"fmt"
	"os"
)

// What will the program do?
// it's a CLI
// Supports these operations / subcommands:
// get		--		./customers get ${CUSTOMER_ID}
// create	--		./customers create ${PATH_TO_JSON_FILE_WITH_CUSTOMER_DATA}

func RunCLIApplication() {
	app, err := newCustomerApp()
	if err != nil {
		fmt.Println("error initializing application!")
		os.Exit(1)
	}

	app.RootHandler.Handle(os.Args[1:])

	closeCustomerApp(app)
}
