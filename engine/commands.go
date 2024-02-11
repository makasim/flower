package engine

type Command interface {
	Execute() error
}

type StartCommand struct {
}

func (cmd *StartCommand) Execute() error {
	return nil
}

type EndCommand struct {
}

func (cmd *EndCommand) Execute() error {
	return nil
}
