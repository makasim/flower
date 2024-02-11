package engine

type ProcessID string
type NodeID string
type BehaviorID string
type TransitionID string

type Process struct {
	ID  ProcessID
	Rev int64

	Nodes       []Node
	Transitions []Transition
}

type Node struct {
	ID  NodeID
	Rev int64

	BehaviorID     BehaviorID
	BehaviorConfig any
}

type Transition struct {
	ID     TransitionID
	FromID NodeID
	ToID   NodeID
}

type Behavior interface {
	Process(t *Task) (TransitionID, error)
}
