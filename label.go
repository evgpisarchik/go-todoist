package go_todoist

import "encoding/json"

type Label struct {
	Color     int    `json:"color"`
	ID        int    `json:"id"`
	IsDeleted int    `json:"is_deleted"`
	ItemOrder int    `json:"item_order"`
	Name      string `json:"name"`
	UID       int    `json:"uid"`
}

type LabelManager struct  {
	Labels []Label
}

func (lm *LabelManager) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &lm.Labels)
	if err != nil {
		return err
	}

	return nil
}