package go_todoist

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type ID string // id may be tempId (string) or real id (int)

func (id *ID) UnmarshalJSON(b []byte) error {
	*id = ID(string(b))
	return nil
}

func (id *ID) MarshalJSON() ([]byte, error) {
	// if id is tempId (string) return empty id
	i, err := strconv.Atoi(string(*id))
	if err != nil {
		return []byte(`""`), nil
	} else {
		b, err := json.Marshal(map[string]int{
			"id": i,
		})
		if err != nil {
			return nil, err
		}
		return b, nil
	}
}

type Project struct {
	ArchivedDate      interface{} `json:"archived_date,omitempty"`
	ArchivedTimestamp int         `json:"archived_timestamp,omitempty"`
	Collapsed         int         `json:"collapsed,omitempty"`
	Color             int         `json:"color,omitempty"`
	ID                ID          `json:"id,omitempty"`
	Indent            int         `json:"indent,omitempty"`
	IsArchived        int         `json:"is_archived,omitempty"`
	IsDeleted         int         `json:"is_deleted,omitempty"`
	ItemOrder         int         `json:"item_order,omitempty"`
	Name              string      `json:"name,omitempty"`
	Shared            bool        `json:"shared,omitempty"`
	UserID            int         `json:"user_id,omitempty"`
}

func (p Project) String() string {
	return fmt.Sprintf("%s(%v)", p.Name, p.ID)
}

type ProjectManager struct {
	Api *TodoistAPI

	Projects []*Project
}

func (pm *ProjectManager) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &pm.Projects)
	if err != nil {
		return err
	}

	return nil
}

func (pm *ProjectManager) Add(project *Project) {
	command := NewCommand("project_add", project)
	project.ID = ID(command.TempID)
	pm.Projects = append(pm.Projects, project)
	pm.Api.Commands.Push(command)
}

func (pm *ProjectManager) Update(project *Project) {
	command := NewCommand("project_update", project)
	pm.Api.Commands.Push(command)
}

func (pm *ProjectManager) Delete(project *Project) {
	arg := make(map[string][]interface{})
	arg["ids"] = append(arg["ids"], project.ID)

	command := NewCommand("project_delete", arg)
	pm.Api.Commands.Push(command)
}

func (pm *ProjectManager) Archive(project *Project) {
	command := NewCommand("project_archive", project)
	pm.Api.Commands.Push(command)
}

func (pm *ProjectManager) Unarchive(project *Project) {
	command := NewCommand("project_unarchive", project)
	pm.Api.Commands.Push(command)
}

func (pm *ProjectManager) Find(id string) *Project {
	for i := 0; i < len(pm.Projects); i++ {
		if pm.Projects[i].ID == ID(id) {
			return pm.Projects[i]
		}
	}
	return nil
}

func (pm *ProjectManager) ReplaceTempId(tempId, realId string) error {
	return nil
}

func (pm *ProjectManager) DeleteAll() {
	for i := 0; i < len(pm.Projects); i++ {
		pm.Delete(pm.Projects[i])
	}
}
