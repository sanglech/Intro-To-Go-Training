package cli

type ErrCmdHandler interface {
	Handle(args []string) error
}
