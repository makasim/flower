package inmemorydriver

import "github.com/makasim/flower/engine"

type Driver interface {
	Commit(cmds ...*engine.Command) error
}
