package go_todoist

import "encoding/json"

type Filter struct {
	Color     int    `json:"color"`
	ID        int    `json:"id"`
	IsDeleted int    `json:"is_deleted"`
	ItemOrder int    `json:"item_order"`
	Name      string `json:"name"`
	Query     string `json:"query"`
	UserID    int    `json:"user_id"`
}

type FilterManager struct {
	Filters []Filter
}

func (fm *FilterManager) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &fm.Filters)
	if err != nil {
		return err
	}

	return nil
}
