package go_todoist

import "encoding/json"

type Item struct {
	AssignedByUID  int         `json:"assigned_by_uid"`
	Checked        int         `json:"checked"`
	Children       interface{} `json:"children"`
	Collapsed      int         `json:"collapsed"`
	Content        string      `json:"content"`
	DateAdded      string      `json:"date_added"`
	DateLang       string      `json:"date_lang"`
	DateString     string      `json:"date_string"`
	DayOrder       int         `json:"day_order"`
	DueDate        interface{} `json:"due_date"`
	DueDateUtc     interface{} `json:"due_date_utc"`
	ID             int         `json:"id"`
	InHistory      int         `json:"in_history"`
	Indent         int         `json:"indent"`
	IsArchived     int         `json:"is_archived"`
	IsDeleted      int         `json:"is_deleted"`
	ItemOrder      int         `json:"item_order"`
	Labels         []int       `json:"labels"`
	Priority       int         `json:"priority"`
	ProjectID      int         `json:"project_id"`
	ResponsibleUID interface{} `json:"responsible_uid"`
	SyncID         interface{} `json:"sync_id"`
	UserID         int         `json:"user_id"`
}

type ItemManager struct {
	Items []Item
}

func (im *ItemManager) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &im.Items)
	if err != nil {
		return err
	}

	return nil
}