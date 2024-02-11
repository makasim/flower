package engine

type TaskID string

type Task struct {
	ID  TaskID
	Rev int64

	ProcessID  ProcessID
	ProcessRev int64
}
