package cli

type CmdHandler interface {
	Handle(args []string)
}
