package task

import (
	"encoding/json"
)

// TaskSpec is the specification for a task
type TaskSpec struct {
	Name     string   `json:"name"`
	Image    string   `json:"image"`
	Init     string   `json:"init"`
	InitArgs []string `json:"initArgs"`
}

// MarshalBinary marshals a TaskSpec
func (s *TaskSpec) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

// UnmarshalBinary unmarshals a TaskSpec
func (s *TaskSpec) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}

type TaskInfo struct {
	Id       string      `json:"id"`
	Metadata interface{} `json:"metadata"`
}
