package go_todoist

import "encoding/json"

type User struct {
	AutoReminder      int           `json:"auto_reminder"`
	AvatarBig         string        `json:"avatar_big"`
	AvatarMedium      string        `json:"avatar_medium"`
	AvatarS640        string        `json:"avatar_s640"`
	AvatarSmall       string        `json:"avatar_small"`
	Beta              int           `json:"beta"`
	BusinessAccountID interface{}   `json:"business_account_id"`
	CompletedCount    int           `json:"completed_count"`
	CompletedToday    int           `json:"completed_today"`
	DateFormat        int           `json:"date_format"`
	DefaultReminder   interface{}   `json:"default_reminder"`
	Email             string        `json:"email"`
	Features          struct{}      `json:"features"`
	FullName          string        `json:"full_name"`
	GuideMode         bool          `json:"guide_mode"`
	HasPushReminders  bool          `json:"has_push_reminders"`
	ID                int           `json:"id"`
	ImageID           interface{}   `json:"image_id"`
	InboxProject      int           `json:"inbox_project"`
	IsBizAdmin        bool          `json:"is_biz_admin"`
	IsDummy           int           `json:"is_dummy"`
	IsPremium         bool          `json:"is_premium"`
	JoinDate          string        `json:"join_date"`
	Karma             float32           `json:"karma"`
	KarmaTrend        string        `json:"karma_trend"`
	MobileHost        interface{}   `json:"mobile_host"`
	MobileNumber      interface{}   `json:"mobile_number"`
	NextWeek          int           `json:"next_week"`
	PremiumUntil      interface{}   `json:"premium_until"`
	Restriction       int           `json:"restriction"`
	SortOrder         int           `json:"sort_order"`
	StartDay          int           `json:"start_day"`
	StartPage         string        `json:"start_page"`
	Theme             int           `json:"theme"`
	TimeFormat        int           `json:"time_format"`
	Timezone          string        `json:"timezone"`
	Token             string        `json:"token"`
	TzOffset          []interface{} `json:"tz_offset"`
}

type UserManager struct {
	User User
}

func (um *UserManager) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &um.User)
	if err != nil {
		return err
	}

	return nil
}