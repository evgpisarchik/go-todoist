package go_todoist

import (
	"github.com/satori/go.uuid"
)

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
	}
}

func generateUUID() string {
	u := uuid.NewV4()
	return u.String()
}