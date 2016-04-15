package go_todoist

import (
	"github.com/satori/go.uuid"
)

type CommandQueue []*Command

func (q *CommandQueue) Push(n *Command) {
    *q = append(*q, n)
}

func (q *CommandQueue) Pop() (n *Command) {
    n = (*q)[0]
    *q = (*q)[1:]
    return
}

func (q *CommandQueue) Clear() {
	*q = (*q)[:0]
}

func (q *CommandQueue) Len() int {
    return len(*q)
}

func (q *CommandQueue) IsEmpty() bool {
	return len(*q) == 0
}

type Command struct {
	Type string `json:"type"`
	Args interface{} `json:"args"`
	UUID string `json:"uuid"`
	TempID string `json:"temp_id"`
}

func NewCommand(tp string, args interface{}) *Command {
	return &Command{
		Type: tp,
		Args: args,
		UUID: generateUUID(),
		TempID: generateUUID(),
	}
}

func generateUUID() string {
	u := uuid.NewV4()
	return u.String()
}