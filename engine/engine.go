package engine

type driver interface {
	Commit(cmds ...*Command) error
}

type Engine struct {
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) Execute(t *Task) *Process {
	return &Process{}
}
