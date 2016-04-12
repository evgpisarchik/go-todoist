package go_todoist

import (
	"encoding/json"
)

type Project struct {
	ArchivedDate      interface{} `json:"archived_date"`
	ArchivedTimestamp int         `json:"archived_timestamp"`
	Collapsed         int         `json:"collapsed"`
	Color             int         `json:"color"`
	ID                int         `json:"id"`
	Indent            int         `json:"indent"`
	IsArchived        int         `json:"is_archived"`
	IsDeleted         int         `json:"is_deleted"`
	ItemOrder         int         `json:"item_order"`
	Name              string      `json:"name"`
	Shared            bool        `json:"shared"`
	UserID            int         `json:"user_id"`
}

func (p Project) String() string {
	return p.Name
}

type ProjectManager struct {
	Projects []Project
}

func (pm *ProjectManager) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &pm.Projects)
	if err != nil {
		return err
	}

	return nil
}