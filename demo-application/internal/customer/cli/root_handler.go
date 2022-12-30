package cli

import (
	"errors"
	"fmt"
)

type RootHandler struct {
	getCmd    ErrCmdHandler
	createCmd ErrCmdHandler
}

func NewRootHandler(
	getCmd ErrCmdHandler,
	createCmd ErrCmdHandler,
) CmdHandler {
	root := &RootHandler{
		getCmd:    getCmd,
		createCmd: createCmd,
	}

	return WithErrorOutput(root)
}

func (r *RootHandler) Handle(args []string) error {
	if len(args) < 1 {
		return errors.New("subcommand is required")
	}

	switch args[0] {
	case "get":
		return r.getCmd.Handle(args[1:])
	case "create":
		return r.createCmd.Handle(args[1:])
	}

	return fmt.Errorf("invalid subcommand: %s", args[0])
}
