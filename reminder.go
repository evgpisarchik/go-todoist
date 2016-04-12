package go_todoist

import "encoding/json"

type Reminder struct {
	DateLang   string `json:"date_lang"`
	DateString string `json:"date_string"`
	DueDate    string `json:"due_date"`
	DueDateUtc string `json:"due_date_utc"`
	ID         int    `json:"id"`
	IsDeleted  int    `json:"is_deleted"`
	ItemID     int    `json:"item_id"`
	MmOffset   int    `json:"mm_offset"`
	NotifyUID  int    `json:"notify_uid"`
	Service    string `json:"service"`
	Type       string `json:"type"`
}

type ReminderManager struct {
	Reminders []Reminder
}

func (rm *ReminderManager) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &rm.Reminders)
	if err != nil {
		return err
	}

	return nil
}