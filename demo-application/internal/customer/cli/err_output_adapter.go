package cli

import "fmt"

func WithErrorOutput(next ErrCmdHandler) CmdHandler {
	return &errorCmdHandlerAdapter{
		next: next,
	}
}

type errorCmdHandlerAdapter struct {
	next ErrCmdHandler
}

func (e *errorCmdHandlerAdapter) Handle(args []string) {
	err := e.next.Handle(args)
	if err != nil {
		e.printError(err)
	}
}

func (e *errorCmdHandlerAdapter) printError(err error) {
	fmt.Println("An error occurred!")
	fmt.Println("Error: " + err.Error())
	fmt.Println("Usage:")
	fmt.Println("./customers get ${CUSTOMER_ID}")
	fmt.Println("./customers create ${PATH_TO_JSON_FILE_WITH_CUSTOMER_DATA}")
}
